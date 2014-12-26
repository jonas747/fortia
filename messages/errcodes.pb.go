// Code generated by protoc-gen-go.
// source: errcodes.proto
// DO NOT EDIT!

/*
Package messages is a generated protocol buffer package.

It is generated from these files:
	errcodes.proto
	game.proto
	requstbodies.proto
	responses.proto

It has these top-level messages:
*/
package messages

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ErrorCode int32

const (
	ErrorCode_RedisDialError          ErrorCode = 1
	ErrorCode_RedisKeyNotFound        ErrorCode = 2
	ErrorCode_RedisWriteReadErr       ErrorCode = 3
	ErrorCode_RedisReplyConversionErr ErrorCode = 4
	// Decoding/encoding errors
	ErrorCode_JsonEncodeErr  ErrorCode = 5
	ErrorCode_JsonDecodeErr  ErrorCode = 6
	ErrorCode_ProtoEncodeErr ErrorCode = 7
	ErrorCode_ProtoDecodeErr ErrorCode = 8
	ErrorCode_FileReadErr    ErrorCode = 9
	ErrorCode_BCryptErr      ErrorCode = 10
)

var ErrorCode_name = map[int32]string{
	1:  "RedisDialError",
	2:  "RedisKeyNotFound",
	3:  "RedisWriteReadErr",
	4:  "RedisReplyConversionErr",
	5:  "JsonEncodeErr",
	6:  "JsonDecodeErr",
	7:  "ProtoEncodeErr",
	8:  "ProtoDecodeErr",
	9:  "FileReadErr",
	10: "BCryptErr",
}
var ErrorCode_value = map[string]int32{
	"RedisDialError":          1,
	"RedisKeyNotFound":        2,
	"RedisWriteReadErr":       3,
	"RedisReplyConversionErr": 4,
	"JsonEncodeErr":           5,
	"JsonDecodeErr":           6,
	"ProtoEncodeErr":          7,
	"ProtoDecodeErr":          8,
	"FileReadErr":             9,
	"BCryptErr":               10,
}

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}
func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (x *ErrorCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ErrorCode_value, data, "ErrorCode")
	if err != nil {
		return err
	}
	*x = ErrorCode(value)
	return nil
}

func init() {
	proto.RegisterEnum("messages.ErrorCode", ErrorCode_name, ErrorCode_value)
}