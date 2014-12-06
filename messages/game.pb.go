// Code generated by protoc-gen-go.
// source: game.proto
// DO NOT EDIT!

package messages

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Block struct {
	Kind             *int32 `protobuf:"varint,1,opt,name=kind" json:"kind,omitempty"`
	Flags            *int32 `protobuf:"varint,2,opt,name=flags" json:"flags,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}

func (m *Block) GetKind() int32 {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return 0
}

func (m *Block) GetFlags() int32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

type Chunk struct {
	Blocks           []*Block `protobuf:"bytes,1,rep,name=blocks" json:"blocks,omitempty"`
	VisibleLayers    []bool   `protobuf:"varint,3,rep,name=visibleLayers" json:"visibleLayers,omitempty"`
	Flags            *int32   `protobuf:"varint,2,opt,name=flags" json:"flags,omitempty"`
	Biome            *int32   `protobuf:"varint,4,opt,name=biome" json:"biome,omitempty"`
	Potency          *int32   `protobuf:"varint,5,opt,name=potency" json:"potency,omitempty"`
	X                *int32   `protobuf:"varint,6,opt,name=x" json:"x,omitempty"`
	Y                *int32   `protobuf:"varint,7,opt,name=y" json:"y,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}

func (m *Chunk) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *Chunk) GetVisibleLayers() []bool {
	if m != nil {
		return m.VisibleLayers
	}
	return nil
}

func (m *Chunk) GetFlags() int32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *Chunk) GetBiome() int32 {
	if m != nil && m.Biome != nil {
		return *m.Biome
	}
	return 0
}

func (m *Chunk) GetPotency() int32 {
	if m != nil && m.Potency != nil {
		return *m.Potency
	}
	return 0
}

func (m *Chunk) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *Chunk) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func init() {
}
