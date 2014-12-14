// Code generated by protoc-gen-go.
// source: responses.proto
// DO NOT EDIT!

package messages

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Error struct {
	Code             *int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Text             *string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}

func (m *Error) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *Error) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

type PlainResponse struct {
	Error            *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PlainResponse) Reset()         { *m = PlainResponse{} }
func (m *PlainResponse) String() string { return proto.CompactTextString(m) }
func (*PlainResponse) ProtoMessage()    {}

func (m *PlainResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type MeResponse struct {
	Error            *Error    `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Info             *UserInfo `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *MeResponse) Reset()         { *m = MeResponse{} }
func (m *MeResponse) String() string { return proto.CompactTextString(m) }
func (*MeResponse) ProtoMessage()    {}

func (m *MeResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *MeResponse) GetInfo() *UserInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type WorldsResponse struct {
	Error            *Error       `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Info             []*WorldInfo `protobuf:"bytes,2,rep,name=info" json:"info,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *WorldsResponse) Reset()         { *m = WorldsResponse{} }
func (m *WorldsResponse) String() string { return proto.CompactTextString(m) }
func (*WorldsResponse) ProtoMessage()    {}

func (m *WorldsResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *WorldsResponse) GetInfo() []*WorldInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type WorldSettingsResponse struct {
	Error            *Error         `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Settings         *WorldSettings `protobuf:"bytes,2,opt,name=settings" json:"settings,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *WorldSettingsResponse) Reset()         { *m = WorldSettingsResponse{} }
func (m *WorldSettingsResponse) String() string { return proto.CompactTextString(m) }
func (*WorldSettingsResponse) ProtoMessage()    {}

func (m *WorldSettingsResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *WorldSettingsResponse) GetSettings() *WorldSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

type ChunksResponse struct {
	Error            *Error   `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Chunks           []*Chunk `protobuf:"bytes,2,rep,name=chunks" json:"chunks,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ChunksResponse) Reset()         { *m = ChunksResponse{} }
func (m *ChunksResponse) String() string { return proto.CompactTextString(m) }
func (*ChunksResponse) ProtoMessage()    {}

func (m *ChunksResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ChunksResponse) GetChunks() []*Chunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

func init() {
}