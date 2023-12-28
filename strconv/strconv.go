package strconv

import (
   "strconv"
   "unicode/utf8"
)

var (
   AppendQuote = strconv.AppendQuote
   AppendUint = strconv.AppendUint
   Quote = strconv.Quote
)

type Number float64

func New_Number[T Ordered](value T) Number {
   return Number(value)
}

func Ratio[T, U Ordered](num T, den U) Number {
   return Number(num) / Number(den)
}

func (n Number) Cardinal(dst []byte) []byte {
   units := []unit_measure{
      {1, ""},
      {1e-3, " thousand"},
      {1e-6, " million"},
      {1e-9, " billion"},
      {1e-12, " trillion"},
   }
   return n.scale(dst, units)
}

func (n Number) Percent(dst []byte) []byte {
   unit := unit_measure{100, "%"}
   return n.label(dst, unit)
}

func (n Number) Rate(dst []byte) []byte {
   units := []unit_measure{
      {1, " byte/s"},
      {1e-3, " kilobyte/s"},
      {1e-6, " megabyte/s"},
      {1e-9, " gigabyte/s"},
      {1e-12, " terabyte/s"},
   }
   return n.scale(dst, units)
}

func (n Number) Size(dst []byte) []byte {
   units := []unit_measure{
      {1, " byte"},
      {1e-3, " kilobyte"},
      {1e-6, " megabyte"},
      {1e-9, " gigabyte"},
      {1e-12, " terabyte"},
   }
   return n.scale(dst, units)
}

func (n Number) label(dst []byte, unit unit_measure) []byte {
   var prec int
   if unit.factor != 1 {
      prec = 2
   }
   unit.factor *= float64(n)
   dst = strconv.AppendFloat(dst, unit.factor, 'f', prec, 64)
   return append(dst, unit.name...)
}

func (n Number) scale(dst []byte, units []unit_measure) []byte {
   var unit unit_measure
   for _, unit = range units {
      if unit.factor * float64(n) < 1000 {
         break
      }
   }
   return n.label(dst, unit)
}

type Ordered interface {
   ~float32 | ~float64 |
   ~int | ~int8 | ~int16 | ~int32 | ~int64 |
   ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type unit_measure struct {
   factor float64
   name string
}

// mimesniff.spec.whatwg.org#binary-data-byte
func binary[T byte|rune](item T) bool {
   if item <= 0x08 {
      return true
   }
   if item == 0x0B {
      return true
   }
   if item >= 0x0E && item <= 0x1A {
      return true
   }
   if item >= 0x1C && item <= 0x1F {
      return true
   }
   return false
}

func Valid(p []byte) bool {
   for _, item := range p {
      if binary(item) {
         return false
      }
   }
   return utf8.Valid(p)
}

func Can_Backquote(s string) bool {
   for _, item := range s {
      if binary(item) {
         return false
      }
      if item == '`' {
         return false
      }
   }
   return utf8.ValidString(s)
}
