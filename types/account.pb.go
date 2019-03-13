// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package types

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Account 的信息
type Account struct {
	// coins标识，目前只有0 一个值
	Currency int32 `protobuf:"varint,1,opt,name=currency,proto3" json:"currency,omitempty"`
	//账户可用余额
	Balance int64 `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	//账户冻结余额
	Frozen int64 `protobuf:"varint,3,opt,name=frozen,proto3" json:"frozen,omitempty"`
	//账户的地址
	Addr                 string   `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetCurrency() int32 {
	if m != nil {
		return m.Currency
	}
	return 0
}

func (m *Account) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetFrozen() int64 {
	if m != nil {
		return m.Frozen
	}
	return 0
}

func (m *Account) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

//账户余额改变的一个交易回报（合约内）
type ReceiptExecAccountTransfer struct {
	//合约地址
	ExecAddr string `protobuf:"bytes,1,opt,name=execAddr,proto3" json:"execAddr,omitempty"`
	//转移前
	Prev *Account `protobuf:"bytes,2,opt,name=prev,proto3" json:"prev,omitempty"`
	//转移后
	Current              *Account `protobuf:"bytes,3,opt,name=current,proto3" json:"current,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiptExecAccountTransfer) Reset()         { *m = ReceiptExecAccountTransfer{} }
func (m *ReceiptExecAccountTransfer) String() string { return proto.CompactTextString(m) }
func (*ReceiptExecAccountTransfer) ProtoMessage()    {}
func (*ReceiptExecAccountTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}

func (m *ReceiptExecAccountTransfer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiptExecAccountTransfer.Unmarshal(m, b)
}
func (m *ReceiptExecAccountTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiptExecAccountTransfer.Marshal(b, m, deterministic)
}
func (m *ReceiptExecAccountTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiptExecAccountTransfer.Merge(m, src)
}
func (m *ReceiptExecAccountTransfer) XXX_Size() int {
	return xxx_messageInfo_ReceiptExecAccountTransfer.Size(m)
}
func (m *ReceiptExecAccountTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiptExecAccountTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiptExecAccountTransfer proto.InternalMessageInfo

func (m *ReceiptExecAccountTransfer) GetExecAddr() string {
	if m != nil {
		return m.ExecAddr
	}
	return ""
}

func (m *ReceiptExecAccountTransfer) GetPrev() *Account {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptExecAccountTransfer) GetCurrent() *Account {
	if m != nil {
		return m.Current
	}
	return nil
}

//账户余额改变的一个交易回报（coins内）
type ReceiptAccountTransfer struct {
	//转移前
	Prev *Account `protobuf:"bytes,1,opt,name=prev,proto3" json:"prev,omitempty"`
	//转移后
	Current              *Account `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiptAccountTransfer) Reset()         { *m = ReceiptAccountTransfer{} }
func (m *ReceiptAccountTransfer) String() string { return proto.CompactTextString(m) }
func (*ReceiptAccountTransfer) ProtoMessage()    {}
func (*ReceiptAccountTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
}

func (m *ReceiptAccountTransfer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiptAccountTransfer.Unmarshal(m, b)
}
func (m *ReceiptAccountTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiptAccountTransfer.Marshal(b, m, deterministic)
}
func (m *ReceiptAccountTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiptAccountTransfer.Merge(m, src)
}
func (m *ReceiptAccountTransfer) XXX_Size() int {
	return xxx_messageInfo_ReceiptAccountTransfer.Size(m)
}
func (m *ReceiptAccountTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiptAccountTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiptAccountTransfer proto.InternalMessageInfo

func (m *ReceiptAccountTransfer) GetPrev() *Account {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptAccountTransfer) GetCurrent() *Account {
	if m != nil {
		return m.Current
	}
	return nil
}

//铸币账户余额增加
type ReceiptAccountMint struct {
	//铸币前
	Prev *Account `protobuf:"bytes,1,opt,name=prev,proto3" json:"prev,omitempty"`
	//铸币后
	Current              *Account `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiptAccountMint) Reset()         { *m = ReceiptAccountMint{} }
func (m *ReceiptAccountMint) String() string { return proto.CompactTextString(m) }
func (*ReceiptAccountMint) ProtoMessage()    {}
func (*ReceiptAccountMint) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{3}
}

func (m *ReceiptAccountMint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiptAccountMint.Unmarshal(m, b)
}
func (m *ReceiptAccountMint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiptAccountMint.Marshal(b, m, deterministic)
}
func (m *ReceiptAccountMint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiptAccountMint.Merge(m, src)
}
func (m *ReceiptAccountMint) XXX_Size() int {
	return xxx_messageInfo_ReceiptAccountMint.Size(m)
}
func (m *ReceiptAccountMint) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiptAccountMint.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiptAccountMint proto.InternalMessageInfo

func (m *ReceiptAccountMint) GetPrev() *Account {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptAccountMint) GetCurrent() *Account {
	if m != nil {
		return m.Current
	}
	return nil
}

type ReceiptAccountBurn struct {
	Prev                 *Account `protobuf:"bytes,1,opt,name=prev,proto3" json:"prev,omitempty"`
	Current              *Account `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiptAccountBurn) Reset()         { *m = ReceiptAccountBurn{} }
func (m *ReceiptAccountBurn) String() string { return proto.CompactTextString(m) }
func (*ReceiptAccountBurn) ProtoMessage()    {}
func (*ReceiptAccountBurn) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{4}
}

func (m *ReceiptAccountBurn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiptAccountBurn.Unmarshal(m, b)
}
func (m *ReceiptAccountBurn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiptAccountBurn.Marshal(b, m, deterministic)
}
func (m *ReceiptAccountBurn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiptAccountBurn.Merge(m, src)
}
func (m *ReceiptAccountBurn) XXX_Size() int {
	return xxx_messageInfo_ReceiptAccountBurn.Size(m)
}
func (m *ReceiptAccountBurn) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiptAccountBurn.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiptAccountBurn proto.InternalMessageInfo

func (m *ReceiptAccountBurn) GetPrev() *Account {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptAccountBurn) GetCurrent() *Account {
	if m != nil {
		return m.Current
	}
	return nil
}

//查询一个地址列表在某个执行器中余额
type ReqBalance struct {
	//地址列表
	Addresses []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	//执行器名称
	Execer               string   `protobuf:"bytes,2,opt,name=execer,proto3" json:"execer,omitempty"`
	StateHash            string   `protobuf:"bytes,3,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	AssetExec            string   `protobuf:"bytes,4,opt,name=asset_exec,json=assetExec,proto3" json:"asset_exec,omitempty"`
	AssetSymbol          string   `protobuf:"bytes,5,opt,name=asset_symbol,json=assetSymbol,proto3" json:"asset_symbol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqBalance) Reset()         { *m = ReqBalance{} }
func (m *ReqBalance) String() string { return proto.CompactTextString(m) }
func (*ReqBalance) ProtoMessage()    {}
func (*ReqBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{5}
}

func (m *ReqBalance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqBalance.Unmarshal(m, b)
}
func (m *ReqBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqBalance.Marshal(b, m, deterministic)
}
func (m *ReqBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqBalance.Merge(m, src)
}
func (m *ReqBalance) XXX_Size() int {
	return xxx_messageInfo_ReqBalance.Size(m)
}
func (m *ReqBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqBalance.DiscardUnknown(m)
}

var xxx_messageInfo_ReqBalance proto.InternalMessageInfo

func (m *ReqBalance) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ReqBalance) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

func (m *ReqBalance) GetStateHash() string {
	if m != nil {
		return m.StateHash
	}
	return ""
}

func (m *ReqBalance) GetAssetExec() string {
	if m != nil {
		return m.AssetExec
	}
	return ""
}

func (m *ReqBalance) GetAssetSymbol() string {
	if m != nil {
		return m.AssetSymbol
	}
	return ""
}

// Account 的列表
type Accounts struct {
	Acc                  []*Account `protobuf:"bytes,1,rep,name=acc,proto3" json:"acc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Accounts) Reset()         { *m = Accounts{} }
func (m *Accounts) String() string { return proto.CompactTextString(m) }
func (*Accounts) ProtoMessage()    {}
func (*Accounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{6}
}

func (m *Accounts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Accounts.Unmarshal(m, b)
}
func (m *Accounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Accounts.Marshal(b, m, deterministic)
}
func (m *Accounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Accounts.Merge(m, src)
}
func (m *Accounts) XXX_Size() int {
	return xxx_messageInfo_Accounts.Size(m)
}
func (m *Accounts) XXX_DiscardUnknown() {
	xxx_messageInfo_Accounts.DiscardUnknown(m)
}

var xxx_messageInfo_Accounts proto.InternalMessageInfo

func (m *Accounts) GetAcc() []*Account {
	if m != nil {
		return m.Acc
	}
	return nil
}

type ExecAccount struct {
	Execer               string   `protobuf:"bytes,1,opt,name=execer,proto3" json:"execer,omitempty"`
	Account              *Account `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecAccount) Reset()         { *m = ExecAccount{} }
func (m *ExecAccount) String() string { return proto.CompactTextString(m) }
func (*ExecAccount) ProtoMessage()    {}
func (*ExecAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{7}
}

func (m *ExecAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecAccount.Unmarshal(m, b)
}
func (m *ExecAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecAccount.Marshal(b, m, deterministic)
}
func (m *ExecAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecAccount.Merge(m, src)
}
func (m *ExecAccount) XXX_Size() int {
	return xxx_messageInfo_ExecAccount.Size(m)
}
func (m *ExecAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ExecAccount proto.InternalMessageInfo

func (m *ExecAccount) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

func (m *ExecAccount) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type AllExecBalance struct {
	Addr                 string         `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	ExecAccount          []*ExecAccount `protobuf:"bytes,2,rep,name=ExecAccount,proto3" json:"ExecAccount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AllExecBalance) Reset()         { *m = AllExecBalance{} }
func (m *AllExecBalance) String() string { return proto.CompactTextString(m) }
func (*AllExecBalance) ProtoMessage()    {}
func (*AllExecBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{8}
}

func (m *AllExecBalance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllExecBalance.Unmarshal(m, b)
}
func (m *AllExecBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllExecBalance.Marshal(b, m, deterministic)
}
func (m *AllExecBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllExecBalance.Merge(m, src)
}
func (m *AllExecBalance) XXX_Size() int {
	return xxx_messageInfo_AllExecBalance.Size(m)
}
func (m *AllExecBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_AllExecBalance.DiscardUnknown(m)
}

var xxx_messageInfo_AllExecBalance proto.InternalMessageInfo

func (m *AllExecBalance) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *AllExecBalance) GetExecAccount() []*ExecAccount {
	if m != nil {
		return m.ExecAccount
	}
	return nil
}

type ReqAllExecBalance struct {
	//地址列表
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	//执行器名称
	Execer               string   `protobuf:"bytes,2,opt,name=execer,proto3" json:"execer,omitempty"`
	StateHash            string   `protobuf:"bytes,3,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	AssetExec            string   `protobuf:"bytes,4,opt,name=asset_exec,json=assetExec,proto3" json:"asset_exec,omitempty"`
	AssetSymbol          string   `protobuf:"bytes,5,opt,name=asset_symbol,json=assetSymbol,proto3" json:"asset_symbol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqAllExecBalance) Reset()         { *m = ReqAllExecBalance{} }
func (m *ReqAllExecBalance) String() string { return proto.CompactTextString(m) }
func (*ReqAllExecBalance) ProtoMessage()    {}
func (*ReqAllExecBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{9}
}

func (m *ReqAllExecBalance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqAllExecBalance.Unmarshal(m, b)
}
func (m *ReqAllExecBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqAllExecBalance.Marshal(b, m, deterministic)
}
func (m *ReqAllExecBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqAllExecBalance.Merge(m, src)
}
func (m *ReqAllExecBalance) XXX_Size() int {
	return xxx_messageInfo_ReqAllExecBalance.Size(m)
}
func (m *ReqAllExecBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqAllExecBalance.DiscardUnknown(m)
}

var xxx_messageInfo_ReqAllExecBalance proto.InternalMessageInfo

func (m *ReqAllExecBalance) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ReqAllExecBalance) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

func (m *ReqAllExecBalance) GetStateHash() string {
	if m != nil {
		return m.StateHash
	}
	return ""
}

func (m *ReqAllExecBalance) GetAssetExec() string {
	if m != nil {
		return m.AssetExec
	}
	return ""
}

func (m *ReqAllExecBalance) GetAssetSymbol() string {
	if m != nil {
		return m.AssetSymbol
	}
	return ""
}

func init() {
	proto.RegisterType((*Account)(nil), "types.Account")
	proto.RegisterType((*ReceiptExecAccountTransfer)(nil), "types.ReceiptExecAccountTransfer")
	proto.RegisterType((*ReceiptAccountTransfer)(nil), "types.ReceiptAccountTransfer")
	proto.RegisterType((*ReceiptAccountMint)(nil), "types.ReceiptAccountMint")
	proto.RegisterType((*ReceiptAccountBurn)(nil), "types.ReceiptAccountBurn")
	proto.RegisterType((*ReqBalance)(nil), "types.ReqBalance")
	proto.RegisterType((*Accounts)(nil), "types.Accounts")
	proto.RegisterType((*ExecAccount)(nil), "types.ExecAccount")
	proto.RegisterType((*AllExecBalance)(nil), "types.AllExecBalance")
	proto.RegisterType((*ReqAllExecBalance)(nil), "types.ReqAllExecBalance")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x96, 0xf3, 0xd3, 0x74, 0x27, 0x50, 0x09, 0x1f, 0x2a, 0xab, 0xa2, 0x62, 0xf1, 0x69, 0x0f,
	0x28, 0x91, 0x58, 0x5e, 0xa0, 0x91, 0x90, 0xb8, 0x20, 0x24, 0xc3, 0xa9, 0x17, 0xe4, 0x75, 0x27,
	0x24, 0x22, 0xf5, 0x6e, 0x6d, 0x2f, 0x22, 0x3c, 0x00, 0xaf, 0x81, 0xc4, 0x93, 0x22, 0x4f, 0xbc,
	0xcd, 0x86, 0x0a, 0xd4, 0x03, 0xa8, 0xb7, 0x9d, 0xef, 0x1b, 0xcf, 0xf7, 0xd9, 0x33, 0xb3, 0xf0,
	0x58, 0x1b, 0x53, 0xb7, 0x36, 0xcc, 0x1a, 0x57, 0x87, 0x9a, 0x8f, 0xc3, 0xb6, 0x41, 0x2f, 0x3f,
	0xc3, 0xe4, 0x62, 0x87, 0xf3, 0x33, 0x38, 0x36, 0xad, 0x73, 0x68, 0xcd, 0x56, 0xb0, 0x9c, 0x15,
	0x63, 0x75, 0x1b, 0x73, 0x01, 0x93, 0x4a, 0x6f, 0xb4, 0x35, 0x28, 0x06, 0x39, 0x2b, 0x86, 0xaa,
	0x0b, 0xf9, 0x29, 0x1c, 0x2d, 0x5d, 0xfd, 0x0d, 0xad, 0x18, 0x12, 0x91, 0x22, 0xce, 0x61, 0xa4,
	0xaf, 0xae, 0x9c, 0x18, 0xe5, 0xac, 0xc8, 0x14, 0x7d, 0xcb, 0xef, 0x0c, 0xce, 0x14, 0x1a, 0x5c,
	0x37, 0xe1, 0xf5, 0x57, 0x34, 0x49, 0xf8, 0x83, 0xd3, 0xd6, 0x2f, 0xd1, 0x45, 0x03, 0x18, 0xe1,
	0x78, 0x8c, 0xd1, 0xb1, 0xdb, 0x98, 0x4b, 0x18, 0x35, 0x0e, 0xbf, 0x90, 0xfa, 0xf4, 0xe5, 0xc9,
	0x8c, 0xdc, 0xcf, 0x52, 0x05, 0x45, 0x1c, 0x2f, 0x60, 0xb2, 0x33, 0x1c, 0xc8, 0xcb, 0xdd, 0xb4,
	0x8e, 0x96, 0x4b, 0x38, 0x4d, 0x3e, 0x7e, 0xf7, 0xd0, 0xe9, 0xb0, 0xfb, 0xe9, 0x0c, 0xfe, 0xae,
	0x53, 0x01, 0x3f, 0xd4, 0x79, 0xbb, 0xb6, 0xe1, 0x7f, 0x6b, 0x2c, 0x5a, 0x67, 0xff, 0xb1, 0xc6,
	0x4f, 0x06, 0xa0, 0xf0, 0x66, 0x91, 0x7a, 0xfe, 0x14, 0xb2, 0xd8, 0x4f, 0xf4, 0x1e, 0xbd, 0x60,
	0xf9, 0xb0, 0xc8, 0xd4, 0x1e, 0x88, 0x13, 0x11, 0xdb, 0x86, 0x8e, 0xaa, 0x66, 0x2a, 0x45, 0xf1,
	0x94, 0x0f, 0x3a, 0xe0, 0x1b, 0xed, 0x57, 0xd4, 0xa0, 0x4c, 0xed, 0x01, 0x7e, 0x0e, 0xa0, 0xbd,
	0xc7, 0xf0, 0x31, 0x66, 0xa7, 0xa9, 0xc9, 0x08, 0x89, 0xa3, 0xc2, 0x9f, 0xc3, 0xa3, 0x1d, 0xed,
	0xb7, 0xd7, 0x55, 0xbd, 0x11, 0x63, 0x4a, 0x98, 0x12, 0xf6, 0x9e, 0x20, 0xf9, 0x02, 0x8e, 0x93,
	0x71, 0xcf, 0x73, 0x18, 0x6a, 0x63, 0xc8, 0xdb, 0xdd, 0x6b, 0x45, 0x4a, 0xbe, 0x83, 0x69, 0x6f,
	0x06, 0x7b, 0xa6, 0xd9, 0x81, 0xe9, 0x02, 0x26, 0x69, 0x6f, 0xfe, 0xf4, 0x46, 0x89, 0x96, 0x97,
	0x70, 0x72, 0xb1, 0xd9, 0xc4, 0x9a, 0xdd, 0x33, 0x75, 0x2b, 0xc0, 0xf6, 0x2b, 0xc0, 0x5f, 0x1d,
	0xc8, 0x8a, 0x01, 0x19, 0xe4, 0xa9, 0x66, 0x8f, 0x51, 0xfd, 0x34, 0xf9, 0x83, 0xc1, 0x13, 0x85,
	0x37, 0xf7, 0xa8, 0xff, 0x40, 0x8f, 0xbf, 0x78, 0x76, 0x79, 0xfe, 0x69, 0x1d, 0x56, 0x6d, 0x35,
	0x33, 0xf5, 0xf5, 0xbc, 0x2c, 0x8d, 0x9d, 0x9b, 0x95, 0x5e, 0xdb, 0xb2, 0x9c, 0xd3, 0xdd, 0xaa,
	0x23, 0xfa, 0xed, 0x94, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x57, 0xdf, 0x86, 0xd3, 0x87, 0x04,
	0x00, 0x00,
}
