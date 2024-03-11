// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: api/v1/runtime/common.proto

package v1

import (
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

type Language int32

const (
	Language_LANG_UNKNOWN Language = 0
	Language_LANG_GOLANG  Language = 1
	Language_LANG_C       Language = 2
)

// Enum value maps for Language.
var (
	Language_name = map[int32]string{
		0: "LANG_UNKNOWN",
		1: "LANG_GOLANG",
		2: "LANG_C",
	}
	Language_value = map[string]int32{
		"LANG_UNKNOWN": 0,
		"LANG_GOLANG":  1,
		"LANG_C":       2,
	}
)

func (x Language) Enum() *Language {
	p := new(Language)
	*p = x
	return p
}

func (x Language) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Language) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_runtime_common_proto_enumTypes[0].Descriptor()
}

func (Language) Type() protoreflect.EnumType {
	return &file_api_v1_runtime_common_proto_enumTypes[0]
}

func (x Language) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Language.Descriptor instead.
func (Language) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_runtime_common_proto_rawDescGZIP(), []int{0}
}

type StatsGroup int32

const (
	StatsGroup_STATS_GROUP_UNKNOWN StatsGroup = 0
	StatsGroup_STATS_GROUP_SYSCALL StatsGroup = 1
	StatsGroup_STATS_GROUP_CPU     StatsGroup = 2
	StatsGroup_STATS_GROUP_MEMORY  StatsGroup = 3
	StatsGroup_STATS_GROUP_IO      StatsGroup = 4
	StatsGroup_STATS_GROUP_NET     StatsGroup = 5
)

// Enum value maps for StatsGroup.
var (
	StatsGroup_name = map[int32]string{
		0: "STATS_GROUP_UNKNOWN",
		1: "STATS_GROUP_SYSCALL",
		2: "STATS_GROUP_CPU",
		3: "STATS_GROUP_MEMORY",
		4: "STATS_GROUP_IO",
		5: "STATS_GROUP_NET",
	}
	StatsGroup_value = map[string]int32{
		"STATS_GROUP_UNKNOWN": 0,
		"STATS_GROUP_SYSCALL": 1,
		"STATS_GROUP_CPU":     2,
		"STATS_GROUP_MEMORY":  3,
		"STATS_GROUP_IO":      4,
		"STATS_GROUP_NET":     5,
	}
)

func (x StatsGroup) Enum() *StatsGroup {
	p := new(StatsGroup)
	*p = x
	return p
}

func (x StatsGroup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatsGroup) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_runtime_common_proto_enumTypes[1].Descriptor()
}

func (StatsGroup) Type() protoreflect.EnumType {
	return &file_api_v1_runtime_common_proto_enumTypes[1]
}

func (x StatsGroup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatsGroup.Descriptor instead.
func (StatsGroup) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_runtime_common_proto_rawDescGZIP(), []int{1}
}

type SignatureEventID int32

const (
	SignatureEventID_SIGNATURE_UNKNOWN          SignatureEventID = 0
	SignatureEventID_SIGNATURE_STDIO_VIA_SOCKET SignatureEventID = 1
)

// Enum value maps for SignatureEventID.
var (
	SignatureEventID_name = map[int32]string{
		0: "SIGNATURE_UNKNOWN",
		1: "SIGNATURE_STDIO_VIA_SOCKET",
	}
	SignatureEventID_value = map[string]int32{
		"SIGNATURE_UNKNOWN":          0,
		"SIGNATURE_STDIO_VIA_SOCKET": 1,
	}
)

func (x SignatureEventID) Enum() *SignatureEventID {
	p := new(SignatureEventID)
	*p = x
	return p
}

func (x SignatureEventID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SignatureEventID) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_runtime_common_proto_enumTypes[2].Descriptor()
}

func (SignatureEventID) Type() protoreflect.EnumType {
	return &file_api_v1_runtime_common_proto_enumTypes[2]
}

func (x SignatureEventID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SignatureEventID.Descriptor instead.
func (SignatureEventID) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_runtime_common_proto_rawDescGZIP(), []int{2}
}

type Exec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string        `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Args []string      `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	Meta *ExecMetadata `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *Exec) Reset() {
	*x = Exec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_runtime_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exec) ProtoMessage() {}

func (x *Exec) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_runtime_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exec.ProtoReflect.Descriptor instead.
func (*Exec) Descriptor() ([]byte, []int) {
	return file_api_v1_runtime_common_proto_rawDescGZIP(), []int{0}
}

func (x *Exec) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Exec) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *Exec) GetMeta() *ExecMetadata {
	if x != nil {
		return x.Meta
	}
	return nil
}

type Library struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Library) Reset() {
	*x = Library{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_runtime_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Library) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Library) ProtoMessage() {}

func (x *Library) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_runtime_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Library.ProtoReflect.Descriptor instead.
func (*Library) Descriptor() ([]byte, []int) {
	return file_api_v1_runtime_common_proto_rawDescGZIP(), []int{1}
}

func (x *Library) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Library) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type ExecMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lang      Language   `protobuf:"varint,1,opt,name=lang,proto3,enum=runtime.v1.Language" json:"lang,omitempty"`
	Libraries []*Library `protobuf:"bytes,2,rep,name=libraries,proto3" json:"libraries,omitempty"`
}

func (x *ExecMetadata) Reset() {
	*x = ExecMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_runtime_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecMetadata) ProtoMessage() {}

func (x *ExecMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_runtime_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecMetadata.ProtoReflect.Descriptor instead.
func (*ExecMetadata) Descriptor() ([]byte, []int) {
	return file_api_v1_runtime_common_proto_rawDescGZIP(), []int{2}
}

func (x *ExecMetadata) GetLang() Language {
	if x != nil {
		return x.Lang
	}
	return Language_LANG_UNKNOWN
}

func (x *ExecMetadata) GetLibraries() []*Library {
	if x != nil {
		return x.Libraries
	}
	return nil
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_runtime_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_runtime_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_api_v1_runtime_common_proto_rawDescGZIP(), []int{3}
}

func (x *File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Tuple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcIp   string `protobuf:"bytes,1,opt,name=src_ip,json=srcIp,proto3" json:"src_ip,omitempty"`
	DstIp   string `protobuf:"bytes,2,opt,name=dst_ip,json=dstIp,proto3" json:"dst_ip,omitempty"`
	SrcPort uint32 `protobuf:"varint,3,opt,name=src_port,json=srcPort,proto3" json:"src_port,omitempty"`
	DstPort uint32 `protobuf:"varint,4,opt,name=dst_port,json=dstPort,proto3" json:"dst_port,omitempty"`
}

func (x *Tuple) Reset() {
	*x = Tuple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_runtime_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tuple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tuple) ProtoMessage() {}

func (x *Tuple) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_runtime_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tuple.ProtoReflect.Descriptor instead.
func (*Tuple) Descriptor() ([]byte, []int) {
	return file_api_v1_runtime_common_proto_rawDescGZIP(), []int{4}
}

func (x *Tuple) GetSrcIp() string {
	if x != nil {
		return x.SrcIp
	}
	return ""
}

func (x *Tuple) GetDstIp() string {
	if x != nil {
		return x.DstIp
	}
	return ""
}

func (x *Tuple) GetSrcPort() uint32 {
	if x != nil {
		return x.SrcPort
	}
	return 0
}

func (x *Tuple) GetDstPort() uint32 {
	if x != nil {
		return x.DstPort
	}
	return 0
}

type DNS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DNSQuestionDomain string        `protobuf:"bytes,1,opt,name=DNSQuestionDomain,proto3" json:"DNSQuestionDomain,omitempty"`
	Answers           []*DNSAnswers `protobuf:"bytes,2,rep,name=answers,proto3" json:"answers,omitempty"`
}

func (x *DNS) Reset() {
	*x = DNS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_runtime_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNS) ProtoMessage() {}

func (x *DNS) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_runtime_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNS.ProtoReflect.Descriptor instead.
func (*DNS) Descriptor() ([]byte, []int) {
	return file_api_v1_runtime_common_proto_rawDescGZIP(), []int{5}
}

func (x *DNS) GetDNSQuestionDomain() string {
	if x != nil {
		return x.DNSQuestionDomain
	}
	return ""
}

func (x *DNS) GetAnswers() []*DNSAnswers {
	if x != nil {
		return x.Answers
	}
	return nil
}

type DNSAnswers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  uint32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Class uint32 `protobuf:"varint,2,opt,name=class,proto3" json:"class,omitempty"`
	Ttl   uint32 `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Name  string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Ip    []byte `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
	Cname string `protobuf:"bytes,6,opt,name=cname,proto3" json:"cname,omitempty"`
}

func (x *DNSAnswers) Reset() {
	*x = DNSAnswers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_runtime_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSAnswers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSAnswers) ProtoMessage() {}

func (x *DNSAnswers) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_runtime_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSAnswers.ProtoReflect.Descriptor instead.
func (*DNSAnswers) Descriptor() ([]byte, []int) {
	return file_api_v1_runtime_common_proto_rawDescGZIP(), []int{6}
}

func (x *DNSAnswers) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *DNSAnswers) GetClass() uint32 {
	if x != nil {
		return x.Class
	}
	return 0
}

func (x *DNSAnswers) GetTtl() uint32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *DNSAnswers) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DNSAnswers) GetIp() []byte {
	if x != nil {
		return x.Ip
	}
	return nil
}

func (x *DNSAnswers) GetCname() string {
	if x != nil {
		return x.Cname
	}
	return ""
}

type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group    StatsGroup `protobuf:"varint,1,opt,name=group,proto3,enum=runtime.v1.StatsGroup" json:"group,omitempty"`
	Subgroup uint32     `protobuf:"varint,2,opt,name=subgroup,proto3" json:"subgroup,omitempty"`
	Value    float64    `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Stats) Reset() {
	*x = Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_runtime_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_runtime_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_api_v1_runtime_common_proto_rawDescGZIP(), []int{7}
}

func (x *Stats) GetGroup() StatsGroup {
	if x != nil {
		return x.Group
	}
	return StatsGroup_STATS_GROUP_UNKNOWN
}

func (x *Stats) GetSubgroup() uint32 {
	if x != nil {
		return x.Subgroup
	}
	return 0
}

func (x *Stats) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type LogEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level int32  `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *LogEvent) Reset() {
	*x = LogEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_runtime_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEvent) ProtoMessage() {}

func (x *LogEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_runtime_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEvent.ProtoReflect.Descriptor instead.
func (*LogEvent) Descriptor() ([]byte, []int) {
	return file_api_v1_runtime_common_proto_rawDescGZIP(), []int{8}
}

func (x *LogEvent) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *LogEvent) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SignatureEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *SignatureMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Finding  *SignatureFinding  `protobuf:"bytes,2,opt,name=finding,proto3" json:"finding,omitempty"`
}

func (x *SignatureEvent) Reset() {
	*x = SignatureEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_runtime_common_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignatureEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureEvent) ProtoMessage() {}

func (x *SignatureEvent) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_runtime_common_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignatureEvent.ProtoReflect.Descriptor instead.
func (*SignatureEvent) Descriptor() ([]byte, []int) {
	return file_api_v1_runtime_common_proto_rawDescGZIP(), []int{9}
}

func (x *SignatureEvent) GetMetadata() *SignatureMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *SignatureEvent) GetFinding() *SignatureFinding {
	if x != nil {
		return x.Finding
	}
	return nil
}

type SignatureMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      SignatureEventID `protobuf:"varint,1,opt,name=id,proto3,enum=runtime.v1.SignatureEventID" json:"id,omitempty"`
	Version string           `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *SignatureMetadata) Reset() {
	*x = SignatureMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_runtime_common_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignatureMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureMetadata) ProtoMessage() {}

func (x *SignatureMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_runtime_common_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignatureMetadata.ProtoReflect.Descriptor instead.
func (*SignatureMetadata) Descriptor() ([]byte, []int) {
	return file_api_v1_runtime_common_proto_rawDescGZIP(), []int{10}
}

func (x *SignatureMetadata) GetId() SignatureEventID {
	if x != nil {
		return x.Id
	}
	return SignatureEventID_SIGNATURE_UNKNOWN
}

func (x *SignatureMetadata) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type SignatureFinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignatureFinding) Reset() {
	*x = SignatureFinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_runtime_common_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignatureFinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignatureFinding) ProtoMessage() {}

func (x *SignatureFinding) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_runtime_common_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignatureFinding.ProtoReflect.Descriptor instead.
func (*SignatureFinding) Descriptor() ([]byte, []int) {
	return file_api_v1_runtime_common_proto_rawDescGZIP(), []int{11}
}

var File_api_v1_runtime_common_proto protoreflect.FileDescriptor

var file_api_v1_runtime_common_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x5c, 0x0a, 0x04, 0x45, 0x78, 0x65,
	0x63, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x37, 0x0a, 0x07, 0x4c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x6b, 0x0a, 0x0c, 0x45, 0x78, 0x65, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x28, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x31, 0x0a, 0x09, 0x6c, 0x69,
	0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x52, 0x09, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x69, 0x65, 0x73, 0x22, 0x1a, 0x0a,
	0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x6b, 0x0a, 0x05, 0x54, 0x75, 0x70,
	0x6c, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x72, 0x63, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x72, 0x63, 0x49, 0x70, 0x12, 0x15, 0x0a, 0x06, 0x64, 0x73, 0x74,
	0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x73, 0x74, 0x49, 0x70,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x72, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x73, 0x72, 0x63, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64,
	0x73, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64,
	0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x65, 0x0a, 0x03, 0x44, 0x4e, 0x53, 0x12, 0x2c, 0x0a,
	0x11, 0x44, 0x4e, 0x53, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x44, 0x4e, 0x53, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x30, 0x0a, 0x07, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x4e, 0x53, 0x41, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x73, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x22, 0x82, 0x01,
	0x0a, 0x0a, 0x44, 0x4e, 0x53, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x67, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x75, 0x62,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x32, 0x0a, 0x08, 0x4c,
	0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22,
	0x83, 0x01, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x39, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a,
	0x07, 0x66, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x66, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x5b, 0x0a, 0x11, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46,
	0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2a, 0x39, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x41, 0x4e, 0x47, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x41, 0x4e, 0x47, 0x5f, 0x47, 0x4f, 0x4c,
	0x41, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x41, 0x4e, 0x47, 0x5f, 0x43, 0x10,
	0x02, 0x2a, 0x94, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41,
	0x54, 0x53, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x53, 0x59, 0x53, 0x43, 0x41, 0x4c, 0x4c,
	0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x47, 0x52, 0x4f, 0x55,
	0x50, 0x5f, 0x43, 0x50, 0x55, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x53,
	0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x10, 0x03, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x49,
	0x4f, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x5f, 0x4e, 0x45, 0x54, 0x10, 0x05, 0x2a, 0x49, 0x0a, 0x10, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x15, 0x0a, 0x11,
	0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45,
	0x5f, 0x53, 0x54, 0x44, 0x49, 0x4f, 0x5f, 0x56, 0x49, 0x41, 0x5f, 0x53, 0x4f, 0x43, 0x4b, 0x45,
	0x54, 0x10, 0x01, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x61, 0x73, 0x74, 0x61, 0x69, 0x2f, 0x6b, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_runtime_common_proto_rawDescOnce sync.Once
	file_api_v1_runtime_common_proto_rawDescData = file_api_v1_runtime_common_proto_rawDesc
)

func file_api_v1_runtime_common_proto_rawDescGZIP() []byte {
	file_api_v1_runtime_common_proto_rawDescOnce.Do(func() {
		file_api_v1_runtime_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_runtime_common_proto_rawDescData)
	})
	return file_api_v1_runtime_common_proto_rawDescData
}

var file_api_v1_runtime_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_api_v1_runtime_common_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_v1_runtime_common_proto_goTypes = []interface{}{
	(Language)(0),             // 0: runtime.v1.Language
	(StatsGroup)(0),           // 1: runtime.v1.StatsGroup
	(SignatureEventID)(0),     // 2: runtime.v1.SignatureEventID
	(*Exec)(nil),              // 3: runtime.v1.Exec
	(*Library)(nil),           // 4: runtime.v1.Library
	(*ExecMetadata)(nil),      // 5: runtime.v1.ExecMetadata
	(*File)(nil),              // 6: runtime.v1.File
	(*Tuple)(nil),             // 7: runtime.v1.Tuple
	(*DNS)(nil),               // 8: runtime.v1.DNS
	(*DNSAnswers)(nil),        // 9: runtime.v1.DNSAnswers
	(*Stats)(nil),             // 10: runtime.v1.Stats
	(*LogEvent)(nil),          // 11: runtime.v1.LogEvent
	(*SignatureEvent)(nil),    // 12: runtime.v1.SignatureEvent
	(*SignatureMetadata)(nil), // 13: runtime.v1.SignatureMetadata
	(*SignatureFinding)(nil),  // 14: runtime.v1.SignatureFinding
}
var file_api_v1_runtime_common_proto_depIdxs = []int32{
	5,  // 0: runtime.v1.Exec.meta:type_name -> runtime.v1.ExecMetadata
	0,  // 1: runtime.v1.ExecMetadata.lang:type_name -> runtime.v1.Language
	4,  // 2: runtime.v1.ExecMetadata.libraries:type_name -> runtime.v1.Library
	9,  // 3: runtime.v1.DNS.answers:type_name -> runtime.v1.DNSAnswers
	1,  // 4: runtime.v1.Stats.group:type_name -> runtime.v1.StatsGroup
	13, // 5: runtime.v1.SignatureEvent.metadata:type_name -> runtime.v1.SignatureMetadata
	14, // 6: runtime.v1.SignatureEvent.finding:type_name -> runtime.v1.SignatureFinding
	2,  // 7: runtime.v1.SignatureMetadata.id:type_name -> runtime.v1.SignatureEventID
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_api_v1_runtime_common_proto_init() }
func file_api_v1_runtime_common_proto_init() {
	if File_api_v1_runtime_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_runtime_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Exec); i {
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
		file_api_v1_runtime_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Library); i {
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
		file_api_v1_runtime_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecMetadata); i {
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
		file_api_v1_runtime_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_api_v1_runtime_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tuple); i {
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
		file_api_v1_runtime_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNS); i {
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
		file_api_v1_runtime_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSAnswers); i {
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
		file_api_v1_runtime_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats); i {
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
		file_api_v1_runtime_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEvent); i {
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
		file_api_v1_runtime_common_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignatureEvent); i {
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
		file_api_v1_runtime_common_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignatureMetadata); i {
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
		file_api_v1_runtime_common_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignatureFinding); i {
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
			RawDescriptor: file_api_v1_runtime_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1_runtime_common_proto_goTypes,
		DependencyIndexes: file_api_v1_runtime_common_proto_depIdxs,
		EnumInfos:         file_api_v1_runtime_common_proto_enumTypes,
		MessageInfos:      file_api_v1_runtime_common_proto_msgTypes,
	}.Build()
	File_api_v1_runtime_common_proto = out.File
	file_api_v1_runtime_common_proto_rawDesc = nil
	file_api_v1_runtime_common_proto_goTypes = nil
	file_api_v1_runtime_common_proto_depIdxs = nil
}
