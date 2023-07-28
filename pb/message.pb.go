// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v4.22.3
// source: message.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

//消息ID
type ID int32

const (
	ID_MSG_BEGIN     ID = 0
	ID_MSG_Connect   ID = 1   //连接(客户端发来第一个消息)
	ID_MSG_Heartbeat ID = 2   //心跳(服务端返回Connect成功之后每隔1秒发送一个心跳包)
	ID_MSG_JoinRoom  ID = 10  //进入
	ID_MSG_Progress  ID = 20  //进度
	ID_MSG_Ready     ID = 30  //准备
	ID_MSG_Start     ID = 40  //开始
	ID_MSG_Frame     ID = 50  //帧数据
	ID_MSG_Input     ID = 60  //输入
	ID_MSG_Result    ID = 70  //结果
	ID_MSG_Close     ID = 100 //房间关闭
	ID_MSG_END       ID = 255
)

// Enum value maps for ID.
var (
	ID_name = map[int32]string{
		0:   "MSG_BEGIN",
		1:   "MSG_Connect",
		2:   "MSG_Heartbeat",
		10:  "MSG_JoinRoom",
		20:  "MSG_Progress",
		30:  "MSG_Ready",
		40:  "MSG_Start",
		50:  "MSG_Frame",
		60:  "MSG_Input",
		70:  "MSG_Result",
		100: "MSG_Close",
		255: "MSG_END",
	}
	ID_value = map[string]int32{
		"MSG_BEGIN":     0,
		"MSG_Connect":   1,
		"MSG_Heartbeat": 2,
		"MSG_JoinRoom":  10,
		"MSG_Progress":  20,
		"MSG_Ready":     30,
		"MSG_Start":     40,
		"MSG_Frame":     50,
		"MSG_Input":     60,
		"MSG_Result":    70,
		"MSG_Close":     100,
		"MSG_END":       255,
	}
)

func (x ID) Enum() *ID {
	p := new(ID)
	*p = x
	return p
}

func (x ID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ID) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (ID) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x ID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ID.Descriptor instead.
func (ID) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

//错误码
type ERRORCODE int32

const (
	ERRORCODE_ERR_Ok        ERRORCODE = 0 //OK
	ERRORCODE_ERR_NoPlayer  ERRORCODE = 1 //没有这个玩家
	ERRORCODE_ERR_NoRoom    ERRORCODE = 2 //没有房间
	ERRORCODE_ERR_RoomState ERRORCODE = 3 //房间状态不正确
	ERRORCODE_ERR_Token     ERRORCODE = 4 //Token验证失败
)

// Enum value maps for ERRORCODE.
var (
	ERRORCODE_name = map[int32]string{
		0: "ERR_Ok",
		1: "ERR_NoPlayer",
		2: "ERR_NoRoom",
		3: "ERR_RoomState",
		4: "ERR_Token",
	}
	ERRORCODE_value = map[string]int32{
		"ERR_Ok":        0,
		"ERR_NoPlayer":  1,
		"ERR_NoRoom":    2,
		"ERR_RoomState": 3,
		"ERR_Token":     4,
	}
)

func (x ERRORCODE) Enum() *ERRORCODE {
	p := new(ERRORCODE)
	*p = x
	return p
}

func (x ERRORCODE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ERRORCODE) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[1].Descriptor()
}

func (ERRORCODE) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[1]
}

func (x ERRORCODE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ERRORCODE.Descriptor instead.
func (ERRORCODE) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

//客户端发来的第一个消息
type C2S_ConnectMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID uint64 `protobuf:"varint,1,opt,name=playerID,proto3" json:"playerID,omitempty"` //唯一ID
	BattleID uint64 `protobuf:"varint,2,opt,name=battleID,proto3" json:"battleID,omitempty"` //战斗ID
	Token    string `protobuf:"bytes,10,opt,name=token,proto3" json:"token,omitempty"`       //令牌
}

func (x *C2S_ConnectMsg) Reset() {
	*x = C2S_ConnectMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_ConnectMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_ConnectMsg) ProtoMessage() {}

func (x *C2S_ConnectMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_ConnectMsg.ProtoReflect.Descriptor instead.
func (*C2S_ConnectMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *C2S_ConnectMsg) GetPlayerID() uint64 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *C2S_ConnectMsg) GetBattleID() uint64 {
	if x != nil {
		return x.BattleID
	}
	return 0
}

func (x *C2S_ConnectMsg) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

//服务端返回连接结果
type S2C_ConnectMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode ERRORCODE `protobuf:"varint,1,opt,name=errorCode,proto3,enum=pb.ERRORCODE" json:"errorCode,omitempty"` //错误码
}

func (x *S2C_ConnectMsg) Reset() {
	*x = S2C_ConnectMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_ConnectMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_ConnectMsg) ProtoMessage() {}

func (x *S2C_ConnectMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_ConnectMsg.ProtoReflect.Descriptor instead.
func (*S2C_ConnectMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *S2C_ConnectMsg) GetErrorCode() ERRORCODE {
	if x != nil {
		return x.ErrorCode
	}
	return ERRORCODE_ERR_Ok
}

//服务端返回进入房间消息
type S2C_JoinRoomMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Roomseatid int32    `protobuf:"varint,1,opt,name=roomseatid,proto3" json:"roomseatid,omitempty"` //自己的位置索引id(1~N)
	Others     []uint64 `protobuf:"varint,2,rep,packed,name=others,proto3" json:"others,omitempty"`  //其他人的id
	Pros       []int32  `protobuf:"varint,3,rep,packed,name=pros,proto3" json:"pros,omitempty"`      //其他人的进度
	RandomSeed int32    `protobuf:"varint,4,opt,name=randomSeed,proto3" json:"randomSeed,omitempty"` //随机种子
}

func (x *S2C_JoinRoomMsg) Reset() {
	*x = S2C_JoinRoomMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_JoinRoomMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_JoinRoomMsg) ProtoMessage() {}

func (x *S2C_JoinRoomMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_JoinRoomMsg.ProtoReflect.Descriptor instead.
func (*S2C_JoinRoomMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *S2C_JoinRoomMsg) GetRoomseatid() int32 {
	if x != nil {
		return x.Roomseatid
	}
	return 0
}

func (x *S2C_JoinRoomMsg) GetOthers() []uint64 {
	if x != nil {
		return x.Others
	}
	return nil
}

func (x *S2C_JoinRoomMsg) GetPros() []int32 {
	if x != nil {
		return x.Pros
	}
	return nil
}

func (x *S2C_JoinRoomMsg) GetRandomSeed() int32 {
	if x != nil {
		return x.RandomSeed
	}
	return 0
}

//服务端广播开始游戏消息
type S2C_StartMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeStamp int64 `protobuf:"varint,1,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"` //同步时间戳
}

func (x *S2C_StartMsg) Reset() {
	*x = S2C_StartMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_StartMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_StartMsg) ProtoMessage() {}

func (x *S2C_StartMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_StartMsg.ProtoReflect.Descriptor instead.
func (*S2C_StartMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *S2C_StartMsg) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

//读条进度
type C2S_ProgressMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pro int32 `protobuf:"varint,1,opt,name=pro,proto3" json:"pro,omitempty"` //进度值0~100
}

func (x *C2S_ProgressMsg) Reset() {
	*x = C2S_ProgressMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_ProgressMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_ProgressMsg) ProtoMessage() {}

func (x *C2S_ProgressMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_ProgressMsg.ProtoReflect.Descriptor instead.
func (*C2S_ProgressMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *C2S_ProgressMsg) GetPro() int32 {
	if x != nil {
		return x.Pro
	}
	return 0
}

//读条进度
type S2C_ProgressMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`   //id
	Pro int32  `protobuf:"varint,2,opt,name=pro,proto3" json:"pro,omitempty"` //进度值0~100
}

func (x *S2C_ProgressMsg) Reset() {
	*x = S2C_ProgressMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_ProgressMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_ProgressMsg) ProtoMessage() {}

func (x *S2C_ProgressMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_ProgressMsg.ProtoReflect.Descriptor instead.
func (*S2C_ProgressMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *S2C_ProgressMsg) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *S2C_ProgressMsg) GetPro() int32 {
	if x != nil {
		return x.Pro
	}
	return 0
}

//操作输入消息
type C2S_InputMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid     int32  `protobuf:"varint,1,opt,name=sid,proto3" json:"sid,omitempty"`         //操作id
	X       int32  `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`             //操作位置x
	Y       int32  `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`             //操作位置y
	FrameID uint32 `protobuf:"varint,4,opt,name=frameID,proto3" json:"frameID,omitempty"` //帧ID
}

func (x *C2S_InputMsg) Reset() {
	*x = C2S_InputMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_InputMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_InputMsg) ProtoMessage() {}

func (x *C2S_InputMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_InputMsg.ProtoReflect.Descriptor instead.
func (*C2S_InputMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{6}
}

func (x *C2S_InputMsg) GetSid() int32 {
	if x != nil {
		return x.Sid
	}
	return 0
}

func (x *C2S_InputMsg) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *C2S_InputMsg) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *C2S_InputMsg) GetFrameID() uint32 {
	if x != nil {
		return x.FrameID
	}
	return 0
}

//帧存储操作输入
type InputData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                 //id
	Sid        int32  `protobuf:"varint,2,opt,name=sid,proto3" json:"sid,omitempty"`               //操作id
	X          int32  `protobuf:"varint,3,opt,name=x,proto3" json:"x,omitempty"`                   //操作位置x
	Y          int32  `protobuf:"varint,4,opt,name=y,proto3" json:"y,omitempty"`                   //操作位置y
	Roomseatid int32  `protobuf:"varint,5,opt,name=roomseatid,proto3" json:"roomseatid,omitempty"` //操作者的位置索引id(1~N)
}

func (x *InputData) Reset() {
	*x = InputData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputData) ProtoMessage() {}

func (x *InputData) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputData.ProtoReflect.Descriptor instead.
func (*InputData) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{7}
}

func (x *InputData) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InputData) GetSid() int32 {
	if x != nil {
		return x.Sid
	}
	return 0
}

func (x *InputData) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *InputData) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *InputData) GetRoomseatid() int32 {
	if x != nil {
		return x.Roomseatid
	}
	return 0
}

//帧数据
type FrameData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrameID uint32       `protobuf:"varint,1,opt,name=frameID,proto3" json:"frameID,omitempty"` //帧ID
	Input   []*InputData `protobuf:"bytes,2,rep,name=input,proto3" json:"input,omitempty"`      //操作输入
}

func (x *FrameData) Reset() {
	*x = FrameData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrameData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrameData) ProtoMessage() {}

func (x *FrameData) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrameData.ProtoReflect.Descriptor instead.
func (*FrameData) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{8}
}

func (x *FrameData) GetFrameID() uint32 {
	if x != nil {
		return x.FrameID
	}
	return 0
}

func (x *FrameData) GetInput() []*InputData {
	if x != nil {
		return x.Input
	}
	return nil
}

//广播帧消息
type S2C_FrameMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Frames []*FrameData `protobuf:"bytes,1,rep,name=frames,proto3" json:"frames,omitempty"` //帧数据
}

func (x *S2C_FrameMsg) Reset() {
	*x = S2C_FrameMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_FrameMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_FrameMsg) ProtoMessage() {}

func (x *S2C_FrameMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_FrameMsg.ProtoReflect.Descriptor instead.
func (*S2C_FrameMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{9}
}

func (x *S2C_FrameMsg) GetFrames() []*FrameData {
	if x != nil {
		return x.Frames
	}
	return nil
}

//结果消息
type C2S_ResultMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WinnerID uint64 `protobuf:"varint,1,opt,name=winnerID,proto3" json:"winnerID,omitempty"` //胜利者ID
}

func (x *C2S_ResultMsg) Reset() {
	*x = C2S_ResultMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_ResultMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_ResultMsg) ProtoMessage() {}

func (x *C2S_ResultMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_ResultMsg.ProtoReflect.Descriptor instead.
func (*C2S_ResultMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{10}
}

func (x *C2S_ResultMsg) GetWinnerID() uint64 {
	if x != nil {
		return x.WinnerID
	}
	return 0
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x5e, 0x0a, 0x0e, 0x43, 0x32, 0x53, 0x5f, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x3d, 0x0a, 0x0e, 0x53, 0x32, 0x43, 0x5f, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x2b, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x43, 0x4f, 0x44, 0x45, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x7d, 0x0a, 0x0f, 0x53, 0x32, 0x43, 0x5f, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f,
	0x6f, 0x6d, 0x4d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x65, 0x61,
	0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x73,
	0x65, 0x61, 0x74, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x72, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x04, 0x70, 0x72, 0x6f,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x53, 0x65, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x53, 0x65, 0x65,
	0x64, 0x22, 0x2c, 0x0a, 0x0c, 0x53, 0x32, 0x43, 0x5f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x73,
	0x67, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x23, 0x0a, 0x0f, 0x43, 0x32, 0x53, 0x5f, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4d,
	0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x70, 0x72, 0x6f, 0x22, 0x33, 0x0a, 0x0f, 0x53, 0x32, 0x43, 0x5f, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x72, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x72, 0x6f, 0x22, 0x56, 0x0a, 0x0c, 0x43, 0x32, 0x53,
	0x5f, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x49,
	0x44, 0x22, 0x69, 0x0a, 0x09, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x69, 0x64,
	0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c,
	0x0a, 0x01, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x65, 0x61, 0x74, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x65, 0x61, 0x74, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x09,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x49, 0x44, 0x12, 0x23, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x35, 0x0a, 0x0c, 0x53, 0x32, 0x43, 0x5f,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x25, 0x0a, 0x06, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x22,
	0x2b, 0x0a, 0x0d, 0x43, 0x32, 0x53, 0x5f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4d, 0x73, 0x67,
	0x12, 0x1a, 0x0a, 0x08, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x2a, 0xc4, 0x01, 0x0a,
	0x02, 0x49, 0x44, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x53, 0x47, 0x5f, 0x42, 0x45, 0x47, 0x49, 0x4e,
	0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x53, 0x47, 0x5f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x53, 0x47, 0x5f, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x53, 0x47, 0x5f, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x53, 0x47, 0x5f,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x10, 0x14, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x53,
	0x47, 0x5f, 0x52, 0x65, 0x61, 0x64, 0x79, 0x10, 0x1e, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x53, 0x47,
	0x5f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x10, 0x28, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x53, 0x47, 0x5f,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x10, 0x32, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x53, 0x47, 0x5f, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x10, 0x3c, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x53, 0x47, 0x5f, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x10, 0x46, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x53, 0x47, 0x5f, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x10, 0x64, 0x12, 0x0c, 0x0a, 0x07, 0x4d, 0x53, 0x47, 0x5f, 0x45, 0x4e, 0x44,
	0x10, 0xff, 0x01, 0x2a, 0x5b, 0x0a, 0x09, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x43, 0x4f, 0x44, 0x45,
	0x12, 0x0a, 0x0a, 0x06, 0x45, 0x52, 0x52, 0x5f, 0x4f, 0x6b, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x45, 0x52, 0x52, 0x5f, 0x4e, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x45, 0x52, 0x52, 0x5f, 0x4e, 0x6f, 0x52, 0x6f, 0x6f, 0x6d, 0x10, 0x02, 0x12, 0x11,
	0x0a, 0x0d, 0x45, 0x52, 0x52, 0x5f, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10,
	0x03, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x52, 0x52, 0x5f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x10, 0x04,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x79, 0x65, 0x62, 0x79, 0x65, 0x62, 0x72, 0x75, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x74, 0x65, 0x70, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_message_proto_goTypes = []interface{}{
	(ID)(0),                 // 0: pb.ID
	(ERRORCODE)(0),          // 1: pb.ERRORCODE
	(*C2S_ConnectMsg)(nil),  // 2: pb.C2S_ConnectMsg
	(*S2C_ConnectMsg)(nil),  // 3: pb.S2C_ConnectMsg
	(*S2C_JoinRoomMsg)(nil), // 4: pb.S2C_JoinRoomMsg
	(*S2C_StartMsg)(nil),    // 5: pb.S2C_StartMsg
	(*C2S_ProgressMsg)(nil), // 6: pb.C2S_ProgressMsg
	(*S2C_ProgressMsg)(nil), // 7: pb.S2C_ProgressMsg
	(*C2S_InputMsg)(nil),    // 8: pb.C2S_InputMsg
	(*InputData)(nil),       // 9: pb.InputData
	(*FrameData)(nil),       // 10: pb.FrameData
	(*S2C_FrameMsg)(nil),    // 11: pb.S2C_FrameMsg
	(*C2S_ResultMsg)(nil),   // 12: pb.C2S_ResultMsg
}
var file_message_proto_depIdxs = []int32{
	1,  // 0: pb.S2C_ConnectMsg.errorCode:type_name -> pb.ERRORCODE
	9,  // 1: pb.FrameData.input:type_name -> pb.InputData
	10, // 2: pb.S2C_FrameMsg.frames:type_name -> pb.FrameData
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_ConnectMsg); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_ConnectMsg); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_JoinRoomMsg); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_StartMsg); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_ProgressMsg); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_ProgressMsg); i {
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
		file_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_InputMsg); i {
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
		file_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputData); i {
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
		file_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrameData); i {
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
		file_message_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_FrameMsg); i {
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
		file_message_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_ResultMsg); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
