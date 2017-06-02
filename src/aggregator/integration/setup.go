// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package integration

import (
	"errors"
	"flag"
	"sort"
	"sync"
	"time"

	"github.com/m3db/m3aggregator/aggregator"
	"github.com/m3db/m3aggregator/aggregator/handler"
	httpserver "github.com/m3db/m3aggregator/server/http"
	msgpackserver "github.com/m3db/m3aggregator/server/msgpack"
	"github.com/m3db/m3aggregator/services/m3aggregator/serve"
	"github.com/m3db/m3cluster/kv/mem"
	"github.com/m3db/m3cluster/proto/util"
	"github.com/m3db/m3cluster/services"
	"github.com/m3db/m3cluster/services/placement"
	"github.com/m3db/m3cluster/shard"
	"github.com/m3db/m3metrics/metric/aggregated"
	"github.com/m3db/m3metrics/policy"
	"github.com/m3db/m3x/clock"
	"github.com/m3db/m3x/sync"
)

var (
	msgpackAddrArg         = flag.String("msgpackAddr", "0.0.0.0:6000", "msgpack server address")
	httpAddrArg            = flag.String("httpAddr", "0.0.0.0:6001", "http server address")
	errServerStartTimedOut = errors.New("server took too long to start")
	errServerStopTimedOut  = errors.New("server took too long to stop")
)

// nowSetterFn is the function that sets the current time.
type nowSetterFn func(t time.Time)

type testSetup struct {
	opts              testOptions
	msgpackAddr       string
	httpAddr          string
	msgpackServerOpts msgpackserver.Options
	httpServerOpts    httpserver.Options
	aggregator        aggregator.Aggregator
	aggregatorOpts    aggregator.Options
	handler           aggregator.Handler
	getNowFn          clock.NowFn
	setNowFn          nowSetterFn
	workerPool        xsync.WorkerPool
	results           *[]aggregated.MetricWithPolicy
	resultLock        *sync.Mutex

	// Signals.
	doneCh   chan struct{}
	closedCh chan struct{}
}

func newTestSetup(opts testOptions) (*testSetup, error) {
	if opts == nil {
		opts = newTestOptions()
	}

	// Set up the msgpack server address.
	msgpackAddr := *msgpackAddrArg
	if addr := opts.MsgpackAddr(); addr != "" {
		msgpackAddr = addr
	}

	// Set up the http server address.
	httpAddr := *httpAddrArg
	if addr := opts.HTTPAddr(); addr != "" {
		httpAddr = addr
	}

	// Set up worker pool.
	workerPool := xsync.NewWorkerPool(opts.WorkerPoolSize())
	workerPool.Init()

	// Set up getter and setter for now.
	var lock sync.RWMutex
	now := time.Now().Truncate(time.Hour)
	getNowFn := func() time.Time {
		lock.RLock()
		t := now
		lock.RUnlock()
		return t
	}
	setNowFn := func(t time.Time) {
		lock.Lock()
		now = t
		lock.Unlock()
	}

	// Create the server options.
	msgpackServerOpts := msgpackserver.NewOptions()
	httpServerOpts := httpserver.NewOptions()

	// Creating the aggregator options.
	aggregatorOpts := aggregator.NewOptions()
	clockOpts := aggregatorOpts.ClockOptions()
	aggregatorOpts = aggregatorOpts.SetClockOptions(clockOpts.SetNowFn(getNowFn))
	entryPool := aggregator.NewEntryPool(nil)
	entryPool.Init(func() *aggregator.Entry {
		return aggregator.NewEntry(nil, aggregatorOpts)
	})
	aggregatorOpts = aggregatorOpts.SetEntryPool(entryPool)
	flushManager := aggregator.NewFlushManager(nil)
	aggregatorOpts = aggregatorOpts.SetFlushManager(flushManager)

	// Set up placement watcher.
	shardSet := make([]shard.Shard, opts.NumShards())
	for i := 0; i < opts.NumShards(); i++ {
		shardSet[i] = shard.NewShard(uint32(i)).SetState(shard.Initializing)
	}
	shards := shard.NewShards(shardSet)
	instance := placement.NewInstance().
		SetID(opts.InstanceID()).
		SetShards(shards)
	testPlacement := placement.NewPlacement().
		SetInstances([]services.PlacementInstance{instance}).
		SetShards(shards.AllIDs())
	stagedPlacement := placement.NewStagedPlacement().
		SetPlacements([]services.Placement{testPlacement})
	stagedPlacementProto, err := util.StagedPlacementToProto(stagedPlacement)
	if err != nil {
		return nil, err
	}
	placementKey := opts.PlacementKVKey()
	placementStore := mem.NewStore()
	_, err = placementStore.SetIfNotExists(placementKey, stagedPlacementProto)
	if err != nil {
		return nil, err
	}
	placementWatcherOpts := placement.NewStagedPlacementWatcherOptions().
		SetStagedPlacementKey(placementKey).
		SetStagedPlacementStore(placementStore)
	aggregatorOpts = aggregatorOpts.
		SetInstanceID(opts.InstanceID()).
		SetStagedPlacementWatcherOptions(placementWatcherOpts)

	// Set up the handler.
	var (
		results    []aggregated.MetricWithPolicy
		resultLock sync.Mutex
	)
	handleFn := func(metric aggregated.Metric, policy policy.Policy) error {
		resultLock.Lock()
		results = append(results, aggregated.MetricWithPolicy{
			Metric: metric,
			Policy: policy,
		})
		resultLock.Unlock()
		return nil
	}
	handler := handler.NewDecodingHandler(handleFn)
	aggregatorOpts = aggregatorOpts.SetFlushHandler(handler)

	return &testSetup{
		opts:              opts,
		msgpackAddr:       msgpackAddr,
		httpAddr:          httpAddr,
		msgpackServerOpts: msgpackServerOpts,
		httpServerOpts:    httpServerOpts,
		aggregatorOpts:    aggregatorOpts,
		handler:           handler,
		getNowFn:          getNowFn,
		setNowFn:          setNowFn,
		workerPool:        workerPool,
		results:           &results,
		resultLock:        &resultLock,
		doneCh:            make(chan struct{}),
		closedCh:          make(chan struct{}),
	}, nil
}

func (ts *testSetup) newClient() *client {
	return newClient(ts.msgpackAddr, ts.opts.ClientBatchSize(), ts.opts.ClientConnectTimeout())
}

func (ts *testSetup) waitUntilServerIsUp() error {
	c := ts.newClient()
	defer c.close()

	serverIsUp := func() bool { return c.testConnection() }
	if waitUntil(serverIsUp, ts.opts.ServerStateChangeTimeout()) {
		return nil
	}
	return errServerStartTimedOut
}

func (ts *testSetup) startServer() error {
	errCh := make(chan error, 1)

	// Creating the aggregator.
	ts.aggregator = aggregator.NewAggregator(ts.aggregatorOpts)
	if err := ts.aggregator.Open(); err != nil {
		return err
	}

	go func() {
		if err := serve.Serve(
			ts.msgpackAddr,
			ts.msgpackServerOpts,
			ts.httpAddr,
			ts.httpServerOpts,
			ts.aggregator,
			ts.doneCh,
		); err != nil {
			select {
			case errCh <- err:
			default:
			}
		}
		close(ts.closedCh)
	}()

	go func() {
		select {
		case errCh <- ts.waitUntilServerIsUp():
		default:
		}
	}()

	return <-errCh
}

func (ts *testSetup) sortedResults() []aggregated.MetricWithPolicy {
	sort.Sort(byTimeIDPolicyAscending(*ts.results))
	return *ts.results
}

func (ts *testSetup) stopServer() error {
	close(ts.doneCh)

	// Wait for graceful server shutdown
	<-ts.closedCh
	return nil
}

func (ts *testSetup) close() {}