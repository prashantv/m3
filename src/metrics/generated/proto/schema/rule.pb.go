// Copyright (c) 2017 Uber Technologies, Inc.
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

// Code generated by protoc-gen-go.
// source: rule.proto
// DO NOT EDIT!

package schema

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MappingRuleSnapshot struct {
	Name        string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Tombstoned  bool              `protobuf:"varint,2,opt,name=tombstoned" json:"tombstoned,omitempty"`
	CutoverTime int64             `protobuf:"varint,3,opt,name=cutover_time,json=cutoverTime" json:"cutover_time,omitempty"`
	TagFilters  map[string]string `protobuf:"bytes,4,rep,name=tag_filters,json=tagFilters" json:"tag_filters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Policies    []*Policy         `protobuf:"bytes,5,rep,name=policies" json:"policies,omitempty"`
}

func (m *MappingRuleSnapshot) Reset()                    { *m = MappingRuleSnapshot{} }
func (m *MappingRuleSnapshot) String() string            { return proto.CompactTextString(m) }
func (*MappingRuleSnapshot) ProtoMessage()               {}
func (*MappingRuleSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *MappingRuleSnapshot) GetTagFilters() map[string]string {
	if m != nil {
		return m.TagFilters
	}
	return nil
}

func (m *MappingRuleSnapshot) GetPolicies() []*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type MappingRule struct {
	Uuid      string                 `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Snapshots []*MappingRuleSnapshot `protobuf:"bytes,2,rep,name=snapshots" json:"snapshots,omitempty"`
}

func (m *MappingRule) Reset()                    { *m = MappingRule{} }
func (m *MappingRule) String() string            { return proto.CompactTextString(m) }
func (*MappingRule) ProtoMessage()               {}
func (*MappingRule) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *MappingRule) GetSnapshots() []*MappingRuleSnapshot {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

type RollupTarget struct {
	Name     string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Tags     []string  `protobuf:"bytes,2,rep,name=tags" json:"tags,omitempty"`
	Policies []*Policy `protobuf:"bytes,3,rep,name=policies" json:"policies,omitempty"`
}

func (m *RollupTarget) Reset()                    { *m = RollupTarget{} }
func (m *RollupTarget) String() string            { return proto.CompactTextString(m) }
func (*RollupTarget) ProtoMessage()               {}
func (*RollupTarget) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *RollupTarget) GetPolicies() []*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type RollupRuleSnapshot struct {
	Name        string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Tombstoned  bool              `protobuf:"varint,2,opt,name=tombstoned" json:"tombstoned,omitempty"`
	CutoverTime int64             `protobuf:"varint,3,opt,name=cutover_time,json=cutoverTime" json:"cutover_time,omitempty"`
	TagFilters  map[string]string `protobuf:"bytes,4,rep,name=tag_filters,json=tagFilters" json:"tag_filters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Targets     []*RollupTarget   `protobuf:"bytes,5,rep,name=targets" json:"targets,omitempty"`
}

func (m *RollupRuleSnapshot) Reset()                    { *m = RollupRuleSnapshot{} }
func (m *RollupRuleSnapshot) String() string            { return proto.CompactTextString(m) }
func (*RollupRuleSnapshot) ProtoMessage()               {}
func (*RollupRuleSnapshot) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *RollupRuleSnapshot) GetTagFilters() map[string]string {
	if m != nil {
		return m.TagFilters
	}
	return nil
}

func (m *RollupRuleSnapshot) GetTargets() []*RollupTarget {
	if m != nil {
		return m.Targets
	}
	return nil
}

type RollupRule struct {
	Uuid      string                `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Snapshots []*RollupRuleSnapshot `protobuf:"bytes,2,rep,name=snapshots" json:"snapshots,omitempty"`
}

func (m *RollupRule) Reset()                    { *m = RollupRule{} }
func (m *RollupRule) String() string            { return proto.CompactTextString(m) }
func (*RollupRule) ProtoMessage()               {}
func (*RollupRule) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *RollupRule) GetSnapshots() []*RollupRuleSnapshot {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

type RuleSet struct {
	Uuid          string         `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Namespace     string         `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	CreatedAt     int64          `protobuf:"varint,3,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	LastUpdatedAt int64          `protobuf:"varint,4,opt,name=last_updated_at,json=lastUpdatedAt" json:"last_updated_at,omitempty"`
	Tombstoned    bool           `protobuf:"varint,5,opt,name=tombstoned" json:"tombstoned,omitempty"`
	CutoverTime   int64          `protobuf:"varint,6,opt,name=cutover_time,json=cutoverTime" json:"cutover_time,omitempty"`
	MappingRules  []*MappingRule `protobuf:"bytes,7,rep,name=mapping_rules,json=mappingRules" json:"mapping_rules,omitempty"`
	RollupRules   []*RollupRule  `protobuf:"bytes,8,rep,name=rollup_rules,json=rollupRules" json:"rollup_rules,omitempty"`
}

func (m *RuleSet) Reset()                    { *m = RuleSet{} }
func (m *RuleSet) String() string            { return proto.CompactTextString(m) }
func (*RuleSet) ProtoMessage()               {}
func (*RuleSet) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *RuleSet) GetMappingRules() []*MappingRule {
	if m != nil {
		return m.MappingRules
	}
	return nil
}

func (m *RuleSet) GetRollupRules() []*RollupRule {
	if m != nil {
		return m.RollupRules
	}
	return nil
}

func init() {
	proto.RegisterType((*MappingRuleSnapshot)(nil), "schema.MappingRuleSnapshot")
	proto.RegisterType((*MappingRule)(nil), "schema.MappingRule")
	proto.RegisterType((*RollupTarget)(nil), "schema.RollupTarget")
	proto.RegisterType((*RollupRuleSnapshot)(nil), "schema.RollupRuleSnapshot")
	proto.RegisterType((*RollupRule)(nil), "schema.RollupRule")
	proto.RegisterType((*RuleSet)(nil), "schema.RuleSet")
}

func init() { proto.RegisterFile("rule.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x94, 0x4d, 0x8b, 0xd4, 0x40,
	0x10, 0x86, 0x49, 0x32, 0x5f, 0xa9, 0x64, 0x5d, 0xe9, 0xdd, 0x43, 0x18, 0x3f, 0x18, 0xe7, 0x20,
	0xc3, 0x0a, 0x39, 0x28, 0xc2, 0x28, 0x78, 0xd8, 0x83, 0x5e, 0x54, 0x90, 0x76, 0xbc, 0x88, 0x10,
	0x7a, 0x32, 0x6d, 0x36, 0x98, 0x2f, 0xba, 0xab, 0x17, 0xe6, 0x07, 0x89, 0xff, 0xc0, 0xdf, 0x27,
	0xe9, 0x4e, 0x66, 0x66, 0x37, 0x71, 0x3c, 0x08, 0x7b, 0xab, 0xbc, 0xa9, 0x54, 0xde, 0x7a, 0x9f,
	0xa6, 0x01, 0x84, 0xca, 0x78, 0x58, 0x89, 0x12, 0x4b, 0x32, 0x92, 0xf1, 0x15, 0xcf, 0xd9, 0xd4,
	0xaf, 0xca, 0x2c, 0x8d, 0xb7, 0x46, 0x9d, 0xff, 0xb4, 0xe1, 0xec, 0x23, 0xab, 0xaa, 0xb4, 0x48,
	0xa8, 0xca, 0xf8, 0xe7, 0x82, 0x55, 0xf2, 0xaa, 0x44, 0x42, 0x60, 0x50, 0xb0, 0x9c, 0x07, 0xd6,
	0xcc, 0x5a, 0xb8, 0x54, 0xd7, 0xe4, 0x31, 0x00, 0x96, 0xf9, 0x5a, 0x62, 0x59, 0xf0, 0x4d, 0x60,
	0xcf, 0xac, 0xc5, 0x84, 0x1e, 0x28, 0xe4, 0x09, 0xf8, 0xb1, 0xc2, 0xf2, 0x9a, 0x8b, 0x08, 0xd3,
	0x9c, 0x07, 0xce, 0xcc, 0x5a, 0x38, 0xd4, 0x6b, 0xb4, 0x55, 0x9a, 0x73, 0xf2, 0x01, 0x3c, 0x64,
	0x49, 0xf4, 0x3d, 0xcd, 0x90, 0x0b, 0x19, 0x0c, 0x66, 0xce, 0xc2, 0x7b, 0xfe, 0x2c, 0x34, 0xd6,
	0xc2, 0x1e, 0x23, 0xe1, 0x8a, 0x25, 0xef, 0x4c, 0xf7, 0xdb, 0x02, 0xc5, 0x96, 0x02, 0xee, 0x04,
	0x72, 0x01, 0x13, 0xbd, 0x4c, 0xca, 0x65, 0x30, 0xd4, 0xa3, 0xee, 0xb5, 0xa3, 0x3e, 0xe9, 0x25,
	0xe9, 0xee, 0xfd, 0xf4, 0x0d, 0x9c, 0xde, 0x1a, 0x45, 0xee, 0x83, 0xf3, 0x83, 0x6f, 0x9b, 0x15,
	0xeb, 0x92, 0x9c, 0xc3, 0xf0, 0x9a, 0x65, 0x8a, 0xeb, 0xe5, 0x5c, 0x6a, 0x1e, 0x5e, 0xdb, 0x4b,
	0x6b, 0xfe, 0x0d, 0xbc, 0x03, 0x77, 0x75, 0x3c, 0x4a, 0xa5, 0x9b, 0x36, 0x9e, 0xba, 0x26, 0xaf,
	0xc0, 0x95, 0x8d, 0x6b, 0x19, 0xd8, 0xda, 0xce, 0x83, 0x23, 0x9b, 0xd1, 0x7d, 0xf7, 0x7c, 0x0d,
	0x3e, 0x2d, 0xb3, 0x4c, 0x55, 0x2b, 0x26, 0x12, 0xde, 0x9f, 0x3e, 0x81, 0x01, 0xb2, 0xc4, 0x4c,
	0x76, 0xa9, 0xae, 0x6f, 0x04, 0xe0, 0x1c, 0x0f, 0x60, 0xfe, 0xcb, 0x06, 0x62, 0x7e, 0x72, 0x17,
	0xa0, 0xdf, 0xf7, 0x81, 0xbe, 0x68, 0xcd, 0x75, 0x7d, 0x1c, 0xe5, 0x1c, 0xc2, 0x18, 0x75, 0x30,
	0x2d, 0xe6, 0xf3, 0x9b, 0x83, 0x4c, 0x6a, 0xb4, 0x6d, 0xfa, 0x5f, 0xd6, 0x5f, 0x01, 0xf6, 0x06,
	0x7b, 0x51, 0x2f, 0xbb, 0xa8, 0xa7, 0x7f, 0xdf, 0xed, 0x90, 0xf4, 0x6f, 0x1b, 0xc6, 0xfa, 0x9d,
	0xa1, 0xdc, 0x99, 0xfc, 0x10, 0xdc, 0x1a, 0x81, 0xac, 0x58, 0xdc, 0x3a, 0xdb, 0x0b, 0xe4, 0x11,
	0x40, 0x2c, 0x38, 0x43, 0xbe, 0x89, 0x18, 0x36, 0xb1, 0xbb, 0x8d, 0x72, 0x89, 0xe4, 0x29, 0x9c,
	0x66, 0x4c, 0x62, 0xa4, 0xaa, 0x4d, 0xdb, 0x33, 0xd0, 0x3d, 0x27, 0xb5, 0xfc, 0xc5, 0xa8, 0x97,
	0x78, 0x8b, 0xef, 0xf0, 0x9f, 0x7c, 0x47, 0x5d, 0xbe, 0x4b, 0x38, 0xc9, 0xcd, 0x99, 0x8e, 0xea,
	0x3b, 0x46, 0x06, 0x63, 0x9d, 0xc2, 0x59, 0xcf, 0x81, 0xa7, 0x7e, 0xbe, 0x7f, 0x90, 0xe4, 0x25,
	0xf8, 0x42, 0x47, 0xd4, 0x7c, 0x38, 0xd1, 0x1f, 0x92, 0x6e, 0x7c, 0xd4, 0x13, 0xbb, 0x5a, 0xae,
	0x47, 0xfa, 0xbe, 0x7a, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xda, 0x16, 0x77, 0xd3, 0x04,
	0x00, 0x00,
}