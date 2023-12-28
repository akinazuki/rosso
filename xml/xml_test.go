package xml

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
   scan.Data, err = os.ReadFile("ignore.html")
   if err != nil {
      t.Fatal(err)
   }
   scan.Sep = []byte(`"web-tv-app/config/environment"`)
   scan.Scan()
   scan.Sep = []byte("<meta")
   var meta struct {
      Content string `xml:"content,attr"`
   }
   if err := scan.Decode(&meta); err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", meta)
}
