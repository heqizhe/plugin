// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	storage.proto

It has these top-level messages:
	Storage
	StorageAction
	ContentOnlyNotaryStorage
	HashOnlyNotaryStorage
	LinkNotaryStorage
	EncryptNotaryStorage
	EncryptContentOnlyNotaryStorage
	EncryptShareNotaryStorage
	QueryStorage
	BatchQueryStorage
	BatchReplyStorage
*/
package types

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
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

type Storage struct {
	// Types that are valid to be assigned to Value:
	//	*Storage_ContentStorage
	//	*Storage_HashStorage
	//	*Storage_LinkStorage
	//	*Storage_EncryptStorage
	//	*Storage_EncryptShareStorage
	Value isStorage_Value `protobuf_oneof:"value"`
}

func (m *Storage) Reset()                    { *m = Storage{} }
func (m *Storage) String() string            { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()               {}
func (*Storage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isStorage_Value interface {
	isStorage_Value()
}

type Storage_ContentStorage struct {
	ContentStorage *ContentOnlyNotaryStorage `protobuf:"bytes,1,opt,name=contentStorage,oneof"`
}
type Storage_HashStorage struct {
	HashStorage *HashOnlyNotaryStorage `protobuf:"bytes,2,opt,name=hashStorage,oneof"`
}
type Storage_LinkStorage struct {
	LinkStorage *LinkNotaryStorage `protobuf:"bytes,3,opt,name=linkStorage,oneof"`
}
type Storage_EncryptStorage struct {
	EncryptStorage *EncryptNotaryStorage `protobuf:"bytes,4,opt,name=encryptStorage,oneof"`
}
type Storage_EncryptShareStorage struct {
	EncryptShareStorage *EncryptShareNotaryStorage `protobuf:"bytes,5,opt,name=encryptShareStorage,oneof"`
}

func (*Storage_ContentStorage) isStorage_Value()      {}
func (*Storage_HashStorage) isStorage_Value()         {}
func (*Storage_LinkStorage) isStorage_Value()         {}
func (*Storage_EncryptStorage) isStorage_Value()      {}
func (*Storage_EncryptShareStorage) isStorage_Value() {}

func (m *Storage) GetValue() isStorage_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Storage) GetContentStorage() *ContentOnlyNotaryStorage {
	if x, ok := m.GetValue().(*Storage_ContentStorage); ok {
		return x.ContentStorage
	}
	return nil
}

func (m *Storage) GetHashStorage() *HashOnlyNotaryStorage {
	if x, ok := m.GetValue().(*Storage_HashStorage); ok {
		return x.HashStorage
	}
	return nil
}

func (m *Storage) GetLinkStorage() *LinkNotaryStorage {
	if x, ok := m.GetValue().(*Storage_LinkStorage); ok {
		return x.LinkStorage
	}
	return nil
}

func (m *Storage) GetEncryptStorage() *EncryptNotaryStorage {
	if x, ok := m.GetValue().(*Storage_EncryptStorage); ok {
		return x.EncryptStorage
	}
	return nil
}

func (m *Storage) GetEncryptShareStorage() *EncryptShareNotaryStorage {
	if x, ok := m.GetValue().(*Storage_EncryptShareStorage); ok {
		return x.EncryptShareStorage
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Storage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Storage_OneofMarshaler, _Storage_OneofUnmarshaler, _Storage_OneofSizer, []interface{}{
		(*Storage_ContentStorage)(nil),
		(*Storage_HashStorage)(nil),
		(*Storage_LinkStorage)(nil),
		(*Storage_EncryptStorage)(nil),
		(*Storage_EncryptShareStorage)(nil),
	}
}

func _Storage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Storage)
	// value
	switch x := m.Value.(type) {
	case *Storage_ContentStorage:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ContentStorage); err != nil {
			return err
		}
	case *Storage_HashStorage:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HashStorage); err != nil {
			return err
		}
	case *Storage_LinkStorage:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LinkStorage); err != nil {
			return err
		}
	case *Storage_EncryptStorage:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EncryptStorage); err != nil {
			return err
		}
	case *Storage_EncryptShareStorage:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EncryptShareStorage); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Storage.Value has unexpected type %T", x)
	}
	return nil
}

func _Storage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Storage)
	switch tag {
	case 1: // value.contentStorage
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContentOnlyNotaryStorage)
		err := b.DecodeMessage(msg)
		m.Value = &Storage_ContentStorage{msg}
		return true, err
	case 2: // value.hashStorage
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HashOnlyNotaryStorage)
		err := b.DecodeMessage(msg)
		m.Value = &Storage_HashStorage{msg}
		return true, err
	case 3: // value.linkStorage
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LinkNotaryStorage)
		err := b.DecodeMessage(msg)
		m.Value = &Storage_LinkStorage{msg}
		return true, err
	case 4: // value.encryptStorage
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EncryptNotaryStorage)
		err := b.DecodeMessage(msg)
		m.Value = &Storage_EncryptStorage{msg}
		return true, err
	case 5: // value.encryptShareStorage
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EncryptShareNotaryStorage)
		err := b.DecodeMessage(msg)
		m.Value = &Storage_EncryptShareStorage{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Storage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Storage)
	// value
	switch x := m.Value.(type) {
	case *Storage_ContentStorage:
		s := proto.Size(x.ContentStorage)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Storage_HashStorage:
		s := proto.Size(x.HashStorage)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Storage_LinkStorage:
		s := proto.Size(x.LinkStorage)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Storage_EncryptStorage:
		s := proto.Size(x.EncryptStorage)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Storage_EncryptShareStorage:
		s := proto.Size(x.EncryptShareStorage)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type StorageAction struct {
	// Types that are valid to be assigned to Value:
	//	*StorageAction_ContentStorage
	//	*StorageAction_HashStorage
	//	*StorageAction_LinkStorage
	//	*StorageAction_EncryptStorage
	//	*StorageAction_EncryptShareStorage
	Value isStorageAction_Value `protobuf_oneof:"value"`
	Ty    int32                 `protobuf:"varint,6,opt,name=ty" json:"ty,omitempty"`
}

func (m *StorageAction) Reset()                    { *m = StorageAction{} }
func (m *StorageAction) String() string            { return proto.CompactTextString(m) }
func (*StorageAction) ProtoMessage()               {}
func (*StorageAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isStorageAction_Value interface {
	isStorageAction_Value()
}

type StorageAction_ContentStorage struct {
	ContentStorage *ContentOnlyNotaryStorage `protobuf:"bytes,1,opt,name=contentStorage,oneof"`
}
type StorageAction_HashStorage struct {
	HashStorage *HashOnlyNotaryStorage `protobuf:"bytes,2,opt,name=hashStorage,oneof"`
}
type StorageAction_LinkStorage struct {
	LinkStorage *LinkNotaryStorage `protobuf:"bytes,3,opt,name=linkStorage,oneof"`
}
type StorageAction_EncryptStorage struct {
	EncryptStorage *EncryptNotaryStorage `protobuf:"bytes,4,opt,name=encryptStorage,oneof"`
}
type StorageAction_EncryptShareStorage struct {
	EncryptShareStorage *EncryptShareNotaryStorage `protobuf:"bytes,5,opt,name=encryptShareStorage,oneof"`
}

func (*StorageAction_ContentStorage) isStorageAction_Value()      {}
func (*StorageAction_HashStorage) isStorageAction_Value()         {}
func (*StorageAction_LinkStorage) isStorageAction_Value()         {}
func (*StorageAction_EncryptStorage) isStorageAction_Value()      {}
func (*StorageAction_EncryptShareStorage) isStorageAction_Value() {}

func (m *StorageAction) GetValue() isStorageAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *StorageAction) GetContentStorage() *ContentOnlyNotaryStorage {
	if x, ok := m.GetValue().(*StorageAction_ContentStorage); ok {
		return x.ContentStorage
	}
	return nil
}

func (m *StorageAction) GetHashStorage() *HashOnlyNotaryStorage {
	if x, ok := m.GetValue().(*StorageAction_HashStorage); ok {
		return x.HashStorage
	}
	return nil
}

func (m *StorageAction) GetLinkStorage() *LinkNotaryStorage {
	if x, ok := m.GetValue().(*StorageAction_LinkStorage); ok {
		return x.LinkStorage
	}
	return nil
}

func (m *StorageAction) GetEncryptStorage() *EncryptNotaryStorage {
	if x, ok := m.GetValue().(*StorageAction_EncryptStorage); ok {
		return x.EncryptStorage
	}
	return nil
}

func (m *StorageAction) GetEncryptShareStorage() *EncryptShareNotaryStorage {
	if x, ok := m.GetValue().(*StorageAction_EncryptShareStorage); ok {
		return x.EncryptShareStorage
	}
	return nil
}

func (m *StorageAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StorageAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StorageAction_OneofMarshaler, _StorageAction_OneofUnmarshaler, _StorageAction_OneofSizer, []interface{}{
		(*StorageAction_ContentStorage)(nil),
		(*StorageAction_HashStorage)(nil),
		(*StorageAction_LinkStorage)(nil),
		(*StorageAction_EncryptStorage)(nil),
		(*StorageAction_EncryptShareStorage)(nil),
	}
}

func _StorageAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StorageAction)
	// value
	switch x := m.Value.(type) {
	case *StorageAction_ContentStorage:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ContentStorage); err != nil {
			return err
		}
	case *StorageAction_HashStorage:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HashStorage); err != nil {
			return err
		}
	case *StorageAction_LinkStorage:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LinkStorage); err != nil {
			return err
		}
	case *StorageAction_EncryptStorage:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EncryptStorage); err != nil {
			return err
		}
	case *StorageAction_EncryptShareStorage:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.EncryptShareStorage); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StorageAction.Value has unexpected type %T", x)
	}
	return nil
}

func _StorageAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StorageAction)
	switch tag {
	case 1: // value.contentStorage
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContentOnlyNotaryStorage)
		err := b.DecodeMessage(msg)
		m.Value = &StorageAction_ContentStorage{msg}
		return true, err
	case 2: // value.hashStorage
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HashOnlyNotaryStorage)
		err := b.DecodeMessage(msg)
		m.Value = &StorageAction_HashStorage{msg}
		return true, err
	case 3: // value.linkStorage
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LinkNotaryStorage)
		err := b.DecodeMessage(msg)
		m.Value = &StorageAction_LinkStorage{msg}
		return true, err
	case 4: // value.encryptStorage
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EncryptNotaryStorage)
		err := b.DecodeMessage(msg)
		m.Value = &StorageAction_EncryptStorage{msg}
		return true, err
	case 5: // value.encryptShareStorage
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EncryptShareNotaryStorage)
		err := b.DecodeMessage(msg)
		m.Value = &StorageAction_EncryptShareStorage{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StorageAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StorageAction)
	// value
	switch x := m.Value.(type) {
	case *StorageAction_ContentStorage:
		s := proto.Size(x.ContentStorage)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StorageAction_HashStorage:
		s := proto.Size(x.HashStorage)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StorageAction_LinkStorage:
		s := proto.Size(x.LinkStorage)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StorageAction_EncryptStorage:
		s := proto.Size(x.EncryptStorage)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StorageAction_EncryptShareStorage:
		s := proto.Size(x.EncryptShareStorage)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// 内容存证模型
type ContentOnlyNotaryStorage struct {
	// 长度需要小于512k
	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *ContentOnlyNotaryStorage) Reset()                    { *m = ContentOnlyNotaryStorage{} }
func (m *ContentOnlyNotaryStorage) String() string            { return proto.CompactTextString(m) }
func (*ContentOnlyNotaryStorage) ProtoMessage()               {}
func (*ContentOnlyNotaryStorage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ContentOnlyNotaryStorage) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

// 哈希存证模型，推荐使用sha256哈希，限制256位得摘要值
type HashOnlyNotaryStorage struct {
	// 长度固定为32字节
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *HashOnlyNotaryStorage) Reset()                    { *m = HashOnlyNotaryStorage{} }
func (m *HashOnlyNotaryStorage) String() string            { return proto.CompactTextString(m) }
func (*HashOnlyNotaryStorage) ProtoMessage()               {}
func (*HashOnlyNotaryStorage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *HashOnlyNotaryStorage) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// 链接存证模型
type LinkNotaryStorage struct {
	// 存证内容的链接，可以写入URL,或者其他可用于定位源文件得线索.
	Link []byte `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	// 源文件得hash值，推荐使用sha256哈希，限制256位得摘要值
	Hash []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *LinkNotaryStorage) Reset()                    { *m = LinkNotaryStorage{} }
func (m *LinkNotaryStorage) String() string            { return proto.CompactTextString(m) }
func (*LinkNotaryStorage) ProtoMessage()               {}
func (*LinkNotaryStorage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LinkNotaryStorage) GetLink() []byte {
	if m != nil {
		return m.Link
	}
	return nil
}

func (m *LinkNotaryStorage) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// 隐私存证模型，如果一个文件需要存证，且不公开内容，可以选择将源文件通过对称加密算法加密后上链
type EncryptNotaryStorage struct {
	// 存证明文内容的hash值，推荐使用sha256哈希，限制256位得摘要值
	ContentHash []byte `protobuf:"bytes,1,opt,name=contentHash,proto3" json:"contentHash,omitempty"`
	// 源文件得密文，由加密key及nonce对明文加密得到该值。
	EncryptContent []byte `protobuf:"bytes,2,opt,name=encryptContent,proto3" json:"encryptContent,omitempty"`
	// 加密iv，通过AES进行加密时制定随机生成的iv,解密时需要使用该值
	Nonce []byte `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *EncryptNotaryStorage) Reset()                    { *m = EncryptNotaryStorage{} }
func (m *EncryptNotaryStorage) String() string            { return proto.CompactTextString(m) }
func (*EncryptNotaryStorage) ProtoMessage()               {}
func (*EncryptNotaryStorage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *EncryptNotaryStorage) GetContentHash() []byte {
	if m != nil {
		return m.ContentHash
	}
	return nil
}

func (m *EncryptNotaryStorage) GetEncryptContent() []byte {
	if m != nil {
		return m.EncryptContent
	}
	return nil
}

func (m *EncryptNotaryStorage) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

// 隐私存证模型
type EncryptContentOnlyNotaryStorage struct {
	// 存证内容的hash值，推荐使用sha256哈希，限制256位得摘要值
	//   bytes contentHash   = 1;
	// 源文件得密文。
	EncryptContent []byte `protobuf:"bytes,1,opt,name=encryptContent,proto3" json:"encryptContent,omitempty"`
	// 加密iv，通过AES进行加密时制定随机生成的iv,解密时需要使用该值
	Nonce []byte `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *EncryptContentOnlyNotaryStorage) Reset()                    { *m = EncryptContentOnlyNotaryStorage{} }
func (m *EncryptContentOnlyNotaryStorage) String() string            { return proto.CompactTextString(m) }
func (*EncryptContentOnlyNotaryStorage) ProtoMessage()               {}
func (*EncryptContentOnlyNotaryStorage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *EncryptContentOnlyNotaryStorage) GetEncryptContent() []byte {
	if m != nil {
		return m.EncryptContent
	}
	return nil
}

func (m *EncryptContentOnlyNotaryStorage) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

// 分享隐私存证模型
type EncryptShareNotaryStorage struct {
	// 存证明文内容的hash值，推荐使用sha256哈希，限制256位得摘要值
	ContentHash []byte `protobuf:"bytes,1,opt,name=contentHash,proto3" json:"contentHash,omitempty"`
	// 源文件得密文。
	EncryptContent []byte `protobuf:"bytes,2,opt,name=encryptContent,proto3" json:"encryptContent,omitempty"`
	// 密钥的kdf推导路径。密钥tree父节点根据该路径可以推导出私钥key
	KeyName []byte `protobuf:"bytes,3,opt,name=keyName,proto3" json:"keyName,omitempty"`
	// 加密key的wrap key。加密key随机生成，对明文进行加密，该key有私密key进行key wrap后公开。
	// 使用时，通过私密key对wrap key解密得到加密key对密文进行解密。
	KeyWrap []byte `protobuf:"bytes,4,opt,name=keyWrap,proto3" json:"keyWrap,omitempty"`
	// 加密iv，通过AES进行加密时制定随机生成的iv,解密时需要使用该值
	Nonce []byte `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *EncryptShareNotaryStorage) Reset()                    { *m = EncryptShareNotaryStorage{} }
func (m *EncryptShareNotaryStorage) String() string            { return proto.CompactTextString(m) }
func (*EncryptShareNotaryStorage) ProtoMessage()               {}
func (*EncryptShareNotaryStorage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *EncryptShareNotaryStorage) GetContentHash() []byte {
	if m != nil {
		return m.ContentHash
	}
	return nil
}

func (m *EncryptShareNotaryStorage) GetEncryptContent() []byte {
	if m != nil {
		return m.EncryptContent
	}
	return nil
}

func (m *EncryptShareNotaryStorage) GetKeyName() []byte {
	if m != nil {
		return m.KeyName
	}
	return nil
}

func (m *EncryptShareNotaryStorage) GetKeyWrap() []byte {
	if m != nil {
		return m.KeyWrap
	}
	return nil
}

func (m *EncryptShareNotaryStorage) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

// 根据txhash去状态数据中查询存储内容
type QueryStorage struct {
	TxHash string `protobuf:"bytes,1,opt,name=txHash" json:"txHash,omitempty"`
}

func (m *QueryStorage) Reset()                    { *m = QueryStorage{} }
func (m *QueryStorage) String() string            { return proto.CompactTextString(m) }
func (*QueryStorage) ProtoMessage()               {}
func (*QueryStorage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *QueryStorage) GetTxHash() string {
	if m != nil {
		return m.TxHash
	}
	return ""
}

// 批量查询有可能导致数据库崩溃
type BatchQueryStorage struct {
	TxHashs []string `protobuf:"bytes,1,rep,name=txHashs" json:"txHashs,omitempty"`
}

func (m *BatchQueryStorage) Reset()                    { *m = BatchQueryStorage{} }
func (m *BatchQueryStorage) String() string            { return proto.CompactTextString(m) }
func (*BatchQueryStorage) ProtoMessage()               {}
func (*BatchQueryStorage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *BatchQueryStorage) GetTxHashs() []string {
	if m != nil {
		return m.TxHashs
	}
	return nil
}

type BatchReplyStorage struct {
	Storages []*Storage `protobuf:"bytes,1,rep,name=storages" json:"storages,omitempty"`
}

func (m *BatchReplyStorage) Reset()                    { *m = BatchReplyStorage{} }
func (m *BatchReplyStorage) String() string            { return proto.CompactTextString(m) }
func (*BatchReplyStorage) ProtoMessage()               {}
func (*BatchReplyStorage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *BatchReplyStorage) GetStorages() []*Storage {
	if m != nil {
		return m.Storages
	}
	return nil
}

func init() {
	proto.RegisterType((*Storage)(nil), "types.Storage")
	proto.RegisterType((*StorageAction)(nil), "types.StorageAction")
	proto.RegisterType((*ContentOnlyNotaryStorage)(nil), "types.ContentOnlyNotaryStorage")
	proto.RegisterType((*HashOnlyNotaryStorage)(nil), "types.HashOnlyNotaryStorage")
	proto.RegisterType((*LinkNotaryStorage)(nil), "types.LinkNotaryStorage")
	proto.RegisterType((*EncryptNotaryStorage)(nil), "types.EncryptNotaryStorage")
	proto.RegisterType((*EncryptContentOnlyNotaryStorage)(nil), "types.EncryptContentOnlyNotaryStorage")
	proto.RegisterType((*EncryptShareNotaryStorage)(nil), "types.EncryptShareNotaryStorage")
	proto.RegisterType((*QueryStorage)(nil), "types.QueryStorage")
	proto.RegisterType((*BatchQueryStorage)(nil), "types.BatchQueryStorage")
	proto.RegisterType((*BatchReplyStorage)(nil), "types.BatchReplyStorage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Storage service

type StorageClient interface {
}

type storageClient struct {
	cc *grpc.ClientConn
}

func NewStorageClient(cc *grpc.ClientConn) StorageClient {
	return &storageClient{cc}
}

// Server API for Storage service

type StorageServer interface {
}

func RegisterStorageServer(s *grpc.Server, srv StorageServer) {
	s.RegisterService(&_Storage_serviceDesc, srv)
}

var _Storage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.storage",
	HandlerType: (*StorageServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "storage.proto",
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x94, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x4d, 0xba, 0x69, 0xec, 0x4b, 0xb7, 0xb0, 0xe3, 0x2a, 0x23, 0x0a, 0x1b, 0x72, 0x58,
	0x8a, 0x62, 0x0f, 0xd5, 0x9b, 0x82, 0xba, 0x52, 0xa8, 0x20, 0x2b, 0x8e, 0x82, 0x47, 0x19, 0xc3,
	0x60, 0x4a, 0xeb, 0x24, 0x24, 0xb3, 0x8b, 0xf9, 0x2f, 0xfc, 0x4f, 0xfc, 0xff, 0x3c, 0x49, 0x66,
	0xde, 0xe4, 0x47, 0x93, 0x7a, 0x12, 0x4f, 0x7b, 0x9b, 0xf7, 0xde, 0xf7, 0x7d, 0xde, 0xaf, 0x10,
	0x38, 0x2e, 0x54, 0x9a, 0xf3, 0x6f, 0x62, 0x91, 0xe5, 0xa9, 0x4a, 0x89, 0xa7, 0xca, 0x4c, 0x14,
	0xd1, 0x6f, 0x17, 0xfc, 0x8f, 0x26, 0x40, 0xde, 0xc2, 0x2c, 0x4e, 0xa5, 0x12, 0x52, 0xa1, 0x87,
	0x3a, 0xa1, 0x33, 0x0f, 0x96, 0x67, 0x0b, 0xad, 0x5d, 0xbc, 0x31, 0xc1, 0xf7, 0x72, 0x57, 0x5e,
	0xa6, 0x8a, 0xe7, 0x25, 0xca, 0xd6, 0xb7, 0xd8, 0x5e, 0x22, 0x79, 0x05, 0x41, 0xc2, 0x8b, 0xc4,
	0x72, 0x5c, 0xcd, 0x79, 0x88, 0x9c, 0x35, 0x2f, 0x92, 0x21, 0x48, 0x3b, 0x85, 0xbc, 0x80, 0x60,
	0xb7, 0x91, 0x5b, 0x4b, 0x18, 0x69, 0x02, 0x45, 0xc2, 0xbb, 0x8d, 0xdc, 0xf6, 0xb2, 0x5b, 0x72,
	0xb2, 0x82, 0x99, 0x90, 0x71, 0x5e, 0x66, 0xf5, 0x28, 0x47, 0x1a, 0xf0, 0x00, 0x01, 0x2b, 0x13,
	0xec, 0x8d, 0xd1, 0x4d, 0x22, 0x9f, 0xe0, 0x8e, 0xf5, 0x24, 0x3c, 0x17, 0x96, 0xe5, 0x69, 0x56,
	0xd8, 0x65, 0x69, 0xc5, 0x3e, 0x70, 0x28, 0xfd, 0xc2, 0x07, 0xef, 0x9a, 0xef, 0xae, 0x44, 0xf4,
	0x73, 0x04, 0xc7, 0xe8, 0x7c, 0x1d, 0xab, 0x4d, 0x2a, 0x6f, 0x4e, 0xf0, 0x7f, 0x4e, 0x40, 0x66,
	0xe0, 0xaa, 0x92, 0x8e, 0x43, 0x67, 0xee, 0x31, 0x57, 0x95, 0xcd, 0x49, 0x9e, 0x01, 0x3d, 0xb4,
	0x63, 0x42, 0xc1, 0xc7, 0x1d, 0xeb, 0xab, 0x4c, 0x99, 0x35, 0xa3, 0xc7, 0x70, 0x77, 0x70, 0xa3,
	0x84, 0xc0, 0x51, 0xb5, 0x51, 0xd4, 0xeb, 0x77, 0xf4, 0x1c, 0x4e, 0x7a, 0xcb, 0xab, 0x84, 0xd5,
	0xf2, 0xac, 0xb0, 0x7a, 0xd7, 0xc9, 0x6e, 0x2b, 0xf9, 0x1a, 0x4e, 0x87, 0x16, 0x47, 0x42, 0x08,
	0xb0, 0x99, 0x75, 0x53, 0xaf, 0xed, 0x22, 0xe7, 0xf5, 0x3d, 0x70, 0x40, 0xe4, 0xee, 0x79, 0xc9,
	0x29, 0x78, 0x32, 0x95, 0xb1, 0xb9, 0xf7, 0x94, 0x19, 0x23, 0xfa, 0x02, 0x67, 0xab, 0x8e, 0xae,
	0x3f, 0x6b, 0xbf, 0x80, 0xf3, 0xf7, 0x02, 0x6e, 0xbb, 0xc0, 0x2f, 0x07, 0xee, 0x1f, 0x3c, 0xe3,
	0x3f, 0x1c, 0x8f, 0x82, 0xbf, 0x15, 0xe5, 0x25, 0xff, 0x6e, 0x07, 0xb4, 0x26, 0x46, 0x3e, 0xe7,
	0x3c, 0xd3, 0x5f, 0xaa, 0x89, 0x54, 0x66, 0xd3, 0xb1, 0xd7, 0xee, 0xf8, 0x1c, 0xa6, 0x1f, 0xae,
	0x44, 0xd3, 0xe3, 0x3d, 0x18, 0xab, 0x1f, 0x75, 0x7b, 0x13, 0x86, 0x56, 0xf4, 0x04, 0x4e, 0x2e,
	0xb8, 0x8a, 0x93, 0x8e, 0x98, 0x82, 0x6f, 0xc2, 0x05, 0x75, 0xc2, 0xd1, 0x7c, 0xc2, 0xac, 0x19,
	0xbd, 0x44, 0x39, 0x13, 0xd9, 0xae, 0x96, 0x3f, 0x82, 0xdb, 0xf8, 0xfb, 0x36, 0xfa, 0x60, 0x39,
	0xc3, 0x4f, 0x1f, 0x15, 0xac, 0x8e, 0x2f, 0x27, 0xe0, 0xe3, 0xfb, 0xeb, 0x58, 0xff, 0xeb, 0x9f,
	0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x93, 0x02, 0x76, 0xfb, 0xfc, 0x05, 0x00, 0x00,
}
