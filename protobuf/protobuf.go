package protobuf

import (
   "google.golang.org/protobuf/encoding/protowire"
   "strconv"
)

func (t type_error) Error() string {
   get_type := func(enc Encoder) string {
      if enc == nil {
         return "nil"
      }
      return enc.get_type()
   }
   var b []byte
   b = append(b, "field "...)
   b = strconv.AppendInt(b, int64(t.Number), 10)
   b = append(b, " is "...)
   b = append(b, get_type(t.lvalue)...)
   b = append(b, ", not "...)
   b = append(b, get_type(t.rvalue)...)
   return string(b)
}

type Raw struct {
   Bytes []byte
   String string
   Message Message
}

type type_error struct {
   Number
   lvalue Encoder
   rvalue Encoder
}

type Bytes []byte

func (b Bytes) encode(buf []byte, num Number) []byte {
   buf = protowire.AppendTag(buf, num, protowire.BytesType)
   return protowire.AppendBytes(buf, b)
}

func (Bytes) get_type() string { return "Bytes" }

type Encoder interface {
   encode([]byte, Number) []byte
   get_type() string
}

type Fixed32 uint32

func (f Fixed32) encode(buf []byte, num Number) []byte {
   buf = protowire.AppendTag(buf, num, protowire.Fixed32Type)
   return protowire.AppendFixed32(buf, uint32(f))
}

func (Fixed32) get_type() string { return "Fixed32" }

type Fixed64 uint64

func (f Fixed64) encode(buf []byte, num Number) []byte {
   buf = protowire.AppendTag(buf, num, protowire.Fixed64Type)
   return protowire.AppendFixed64(buf, uint64(f))
}

func (Fixed64) get_type() string { return "Fixed64" }

type Number = protowire.Number

func (r Raw) encode(buf []byte, num Number) []byte {
   buf = protowire.AppendTag(buf, num, protowire.BytesType)
   return protowire.AppendBytes(buf, r.Bytes)
}

func (Raw) get_type() string { return "Raw" }

type Slice[T Encoder] []T

func (s Slice[T]) encode(buf []byte, num Number) []byte {
   for _, value := range s {
      buf = value.encode(buf, num)
   }
   return buf
}

func (Slice[T]) get_type() string {
   var value T
   return "[]" + value.get_type()
}

type String string

func (s String) encode(buf []byte, num Number) []byte {
   buf = protowire.AppendTag(buf, num, protowire.BytesType)
   return protowire.AppendString(buf, string(s))
}

func (String) get_type() string { return "String" }

type Varint uint64

func (v Varint) encode(buf []byte, num Number) []byte {
   buf = protowire.AppendTag(buf, num, protowire.VarintType)
   return protowire.AppendVarint(buf, uint64(v))
}

func (Varint) get_type() string { return "Varint" }
