package dash

import (
   "encoding/xml"
   "fmt"
   "net/http"
   "os"
   "testing"
)

func Test_Media(t *testing.T) {
   file, err := os.Open("mpd/roku.mpd")
   if err != nil {
      t.Fatal(err)
   }
   var pre Presentation
   if err := xml.NewDecoder(file).Decode(&pre); err != nil {
      t.Fatal(err)
   }
   if err := file.Close(); err != nil {
      t.Fatal(err)
   }
   base, err := http.NewRequest("", "http://example.com", nil)
   if err != nil {
      t.Fatal(err)
   }
   for _, ref := range pre.Period.AdaptationSet[0].Representation[0].Media() {
      req, err := http.NewRequest("", ref, nil)
      if err != nil {
         t.Fatal(err)
      }
      req.URL = base.URL.ResolveReference(req.URL)
      fmt.Println(req.URL)
   }
}
