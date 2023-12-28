package crypto

import (
   "testing"
)

func Test_Format_JA3(t *testing.T) {
   hello, err := Parse_JA3(Android_API_26)
   if err != nil {
      t.Fatal(err)
   }
   ja3, err := Format_JA3(hello)
   if err != nil {
      t.Fatal(err)
   }
   if ja3 != Android_API_26 {
      t.Fatal(ja3)
   }
}
