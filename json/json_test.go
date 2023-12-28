package json

import (
   "fmt"
   "os"
   "testing"
)

func Test_Scanner(t *testing.T) {
   var (
      err error
      scan Scanner
   )
   scan.Data, err = os.ReadFile("roku.html")
   if err != nil {
      t.Fatal(err)
   }
   scan.Sep = []byte("\tcsrf:")
   scan.Scan()
   scan.Sep = nil
   var token string
   if err := scan.Decode(&token); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%q\n", token)
}
