package codec

import "io"

type Header struct {
	ServiceMethod string //	服务名和方法名
	Seq           uint64 // 请求序列号Id
	Error         string // 错误信息
}

type Codec interface{
	io.Closer
	ReadHeader(*Header) error 
	ReadBody(interface{}) error 
	Write(*Header, interface{}) error 
}

type NewCodecFunc func(io.ReadWriteCloser) Codec 

type Type string 

const (
	GobType Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc 


func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] =  NewGobCodec 
}