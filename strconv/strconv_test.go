package strconv

import (
   "fmt"
   "testing"
   "unicode/utf8"
)

func Test_Append(t *testing.T) {
   var b []byte
   b = New_Number(123).Cardinal(nil)
   if s := string(b); s != "123" {
      t.Fatal(s)
   }
   b = New_Number(1234).Cardinal(nil)
   if s := string(b); s != "1.23 thousand" {
      t.Fatal(s)
   }
   b = New_Number(123).Size(nil)
   if s := string(b); s != "123 byte" {
      t.Fatal(s)
   }
   b = New_Number(1234).Size(nil)
   if s := string(b); s != "1.23 kilobyte" {
      t.Fatal(s)
   }
   b = Ratio(1234, 10).Rate(nil)
   if s := string(b); s != "123 byte/s" {
      t.Fatal(s)
   }
   b = Ratio(12345, 10).Rate(nil)
   if s := string(b); s != "1.23 kilobyte/s" {
      t.Fatal(s)
   }
   b = Ratio(1234, 10000).Percent(nil)
   if s := string(b); s != "12.34%" {
      t.Fatal(s)
   }
}

func Test_Valid(t *testing.T) {
   s := "\xE0<"
   for _, r := range s {
      fmt.Println(utf8.ValidRune(r))
   }
   fmt.Println(utf8.ValidString(s))
}
