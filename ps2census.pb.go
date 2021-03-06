// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: ps2census.proto

package com_ps2census

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CharacterQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CharacterID string `protobuf:"bytes,1,opt,name=characterID,proto3" json:"characterID,omitempty"`
}

func (x *CharacterQuery) Reset() {
	*x = CharacterQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps2census_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CharacterQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CharacterQuery) ProtoMessage() {}

func (x *CharacterQuery) ProtoReflect() protoreflect.Message {
	mi := &file_ps2census_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CharacterQuery.ProtoReflect.Descriptor instead.
func (*CharacterQuery) Descriptor() ([]byte, []int) {
	return file_ps2census_proto_rawDescGZIP(), []int{0}
}

func (x *CharacterQuery) GetCharacterID() string {
	if x != nil {
		return x.CharacterID
	}
	return ""
}

type OutfitQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutfitID string `protobuf:"bytes,1,opt,name=outfitID,proto3" json:"outfitID,omitempty"`
}

func (x *OutfitQuery) Reset() {
	*x = OutfitQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps2census_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutfitQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutfitQuery) ProtoMessage() {}

func (x *OutfitQuery) ProtoReflect() protoreflect.Message {
	mi := &file_ps2census_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutfitQuery.ProtoReflect.Descriptor instead.
func (*OutfitQuery) Descriptor() ([]byte, []int) {
	return file_ps2census_proto_rawDescGZIP(), []int{1}
}

func (x *OutfitQuery) GetOutfitID() string {
	if x != nil {
		return x.OutfitID
	}
	return ""
}

type CharacterResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CharacterID   string `protobuf:"bytes,1,opt,name=characterID,proto3" json:"characterID,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FactionID     int32  `protobuf:"varint,3,opt,name=factionID,proto3" json:"factionID,omitempty"`
	WorldID       int32  `protobuf:"varint,4,opt,name=worldID,proto3" json:"worldID,omitempty"`
	OutfitID      string `protobuf:"bytes,5,opt,name=outfitID,proto3" json:"outfitID,omitempty"`
	BattleRank    int32  `protobuf:"varint,6,opt,name=battleRank,proto3" json:"battleRank,omitempty"`
	PrestigeLevel int32  `protobuf:"varint,7,opt,name=prestigeLevel,proto3" json:"prestigeLevel,omitempty"`
}

func (x *CharacterResult) Reset() {
	*x = CharacterResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps2census_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CharacterResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CharacterResult) ProtoMessage() {}

func (x *CharacterResult) ProtoReflect() protoreflect.Message {
	mi := &file_ps2census_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CharacterResult.ProtoReflect.Descriptor instead.
func (*CharacterResult) Descriptor() ([]byte, []int) {
	return file_ps2census_proto_rawDescGZIP(), []int{2}
}

func (x *CharacterResult) GetCharacterID() string {
	if x != nil {
		return x.CharacterID
	}
	return ""
}

func (x *CharacterResult) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CharacterResult) GetFactionID() int32 {
	if x != nil {
		return x.FactionID
	}
	return 0
}

func (x *CharacterResult) GetWorldID() int32 {
	if x != nil {
		return x.WorldID
	}
	return 0
}

func (x *CharacterResult) GetOutfitID() string {
	if x != nil {
		return x.OutfitID
	}
	return ""
}

func (x *CharacterResult) GetBattleRank() int32 {
	if x != nil {
		return x.BattleRank
	}
	return 0
}

func (x *CharacterResult) GetPrestigeLevel() int32 {
	if x != nil {
		return x.PrestigeLevel
	}
	return 0
}

type OutfitResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutfitID          string `protobuf:"bytes,1,opt,name=outfitID,proto3" json:"outfitID,omitempty"`
	Name              string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Alias             string `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	WorldID           int32  `protobuf:"varint,4,opt,name=worldID,proto3" json:"worldID,omitempty"`
	FactionID         int32  `protobuf:"varint,5,opt,name=factionID,proto3" json:"factionID,omitempty"`
	LeaderCharacterID string `protobuf:"bytes,6,opt,name=leaderCharacterID,proto3" json:"leaderCharacterID,omitempty"`
	MemberCount       int32  `protobuf:"varint,7,opt,name=memberCount,proto3" json:"memberCount,omitempty"`
}

func (x *OutfitResult) Reset() {
	*x = OutfitResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps2census_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutfitResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutfitResult) ProtoMessage() {}

func (x *OutfitResult) ProtoReflect() protoreflect.Message {
	mi := &file_ps2census_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutfitResult.ProtoReflect.Descriptor instead.
func (*OutfitResult) Descriptor() ([]byte, []int) {
	return file_ps2census_proto_rawDescGZIP(), []int{3}
}

func (x *OutfitResult) GetOutfitID() string {
	if x != nil {
		return x.OutfitID
	}
	return ""
}

func (x *OutfitResult) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OutfitResult) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *OutfitResult) GetWorldID() int32 {
	if x != nil {
		return x.WorldID
	}
	return 0
}

func (x *OutfitResult) GetFactionID() int32 {
	if x != nil {
		return x.FactionID
	}
	return 0
}

func (x *OutfitResult) GetLeaderCharacterID() string {
	if x != nil {
		return x.LeaderCharacterID
	}
	return ""
}

func (x *OutfitResult) GetMemberCount() int32 {
	if x != nil {
		return x.MemberCount
	}
	return 0
}

var File_ps2census_proto protoreflect.FileDescriptor

var file_ps2census_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x73, 0x32, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x73, 0x32, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x22, 0x32, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x49, 0x44, 0x22, 0x29, 0x0a, 0x0b, 0x4f, 0x75, 0x74, 0x66, 0x69, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x75, 0x74, 0x66, 0x69, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x75, 0x74, 0x66, 0x69, 0x74, 0x49, 0x44, 0x22,
	0xe1, 0x01, 0x0a, 0x0f, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x75, 0x74, 0x66, 0x69, 0x74, 0x49, 0x44, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x75, 0x74, 0x66, 0x69, 0x74, 0x49, 0x44, 0x12, 0x1e, 0x0a,
	0x0a, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x24, 0x0a,
	0x0d, 0x70, 0x72, 0x65, 0x73, 0x74, 0x69, 0x67, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x73, 0x74, 0x69, 0x67, 0x65, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x22, 0xdc, 0x01, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x66, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x75, 0x74, 0x66, 0x69, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x75, 0x74, 0x66, 0x69, 0x74, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6c,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x32, 0xa4, 0x01, 0x0a, 0x09, 0x50, 0x53, 0x32, 0x43, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x12, 0x4f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x73, 0x32, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a,
	0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x73, 0x32, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e,
	0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x46, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x66, 0x69, 0x74, 0x12, 0x1a,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x73, 0x32, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x4f,
	0x75, 0x74, 0x66, 0x69, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x1b, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x73, 0x32, 0x63, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x66, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ps2census_proto_rawDescOnce sync.Once
	file_ps2census_proto_rawDescData = file_ps2census_proto_rawDesc
)

func file_ps2census_proto_rawDescGZIP() []byte {
	file_ps2census_proto_rawDescOnce.Do(func() {
		file_ps2census_proto_rawDescData = protoimpl.X.CompressGZIP(file_ps2census_proto_rawDescData)
	})
	return file_ps2census_proto_rawDescData
}

var file_ps2census_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ps2census_proto_goTypes = []interface{}{
	(*CharacterQuery)(nil),  // 0: com.ps2census.CharacterQuery
	(*OutfitQuery)(nil),     // 1: com.ps2census.OutfitQuery
	(*CharacterResult)(nil), // 2: com.ps2census.CharacterResult
	(*OutfitResult)(nil),    // 3: com.ps2census.OutfitResult
}
var file_ps2census_proto_depIdxs = []int32{
	0, // 0: com.ps2census.PS2Census.GetCharacter:input_type -> com.ps2census.CharacterQuery
	1, // 1: com.ps2census.PS2Census.GetOutfit:input_type -> com.ps2census.OutfitQuery
	2, // 2: com.ps2census.PS2Census.GetCharacter:output_type -> com.ps2census.CharacterResult
	3, // 3: com.ps2census.PS2Census.GetOutfit:output_type -> com.ps2census.OutfitResult
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ps2census_proto_init() }
func file_ps2census_proto_init() {
	if File_ps2census_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ps2census_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CharacterQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ps2census_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutfitQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ps2census_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CharacterResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ps2census_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutfitResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ps2census_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ps2census_proto_goTypes,
		DependencyIndexes: file_ps2census_proto_depIdxs,
		MessageInfos:      file_ps2census_proto_msgTypes,
	}.Build()
	File_ps2census_proto = out.File
	file_ps2census_proto_rawDesc = nil
	file_ps2census_proto_goTypes = nil
	file_ps2census_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PS2CensusClient is the client API for PS2Census service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PS2CensusClient interface {
	GetCharacter(ctx context.Context, in *CharacterQuery, opts ...grpc.CallOption) (*CharacterResult, error)
	GetOutfit(ctx context.Context, in *OutfitQuery, opts ...grpc.CallOption) (*OutfitResult, error)
}

type pS2CensusClient struct {
	cc grpc.ClientConnInterface
}

func NewPS2CensusClient(cc grpc.ClientConnInterface) PS2CensusClient {
	return &pS2CensusClient{cc}
}

func (c *pS2CensusClient) GetCharacter(ctx context.Context, in *CharacterQuery, opts ...grpc.CallOption) (*CharacterResult, error) {
	out := new(CharacterResult)
	err := c.cc.Invoke(ctx, "/com.ps2census.PS2Census/GetCharacter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pS2CensusClient) GetOutfit(ctx context.Context, in *OutfitQuery, opts ...grpc.CallOption) (*OutfitResult, error) {
	out := new(OutfitResult)
	err := c.cc.Invoke(ctx, "/com.ps2census.PS2Census/GetOutfit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PS2CensusServer is the server API for PS2Census service.
type PS2CensusServer interface {
	GetCharacter(context.Context, *CharacterQuery) (*CharacterResult, error)
	GetOutfit(context.Context, *OutfitQuery) (*OutfitResult, error)
}

// UnimplementedPS2CensusServer can be embedded to have forward compatible implementations.
type UnimplementedPS2CensusServer struct {
}

func (*UnimplementedPS2CensusServer) GetCharacter(context.Context, *CharacterQuery) (*CharacterResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCharacter not implemented")
}
func (*UnimplementedPS2CensusServer) GetOutfit(context.Context, *OutfitQuery) (*OutfitResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOutfit not implemented")
}

func RegisterPS2CensusServer(s *grpc.Server, srv PS2CensusServer) {
	s.RegisterService(&_PS2Census_serviceDesc, srv)
}

func _PS2Census_GetCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CharacterQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PS2CensusServer).GetCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ps2census.PS2Census/GetCharacter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PS2CensusServer).GetCharacter(ctx, req.(*CharacterQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _PS2Census_GetOutfit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OutfitQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PS2CensusServer).GetOutfit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.ps2census.PS2Census/GetOutfit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PS2CensusServer).GetOutfit(ctx, req.(*OutfitQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _PS2Census_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.ps2census.PS2Census",
	HandlerType: (*PS2CensusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCharacter",
			Handler:    _PS2Census_GetCharacter_Handler,
		},
		{
			MethodName: "GetOutfit",
			Handler:    _PS2Census_GetOutfit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ps2census.proto",
}
