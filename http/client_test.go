package http

import (
   "net/http"
   "testing"
)

func Test_Client_Get(t *testing.T) {
   res, err := Default_Client.Status(302).Get("http://godocs.io")
   if err != nil {
      t.Fatal(err)
   }
   if err := res.Body.Close(); err != nil {
      t.Fatal(err)
   }
}

func Test_Client_Do(t *testing.T) {
   req, err := http.NewRequest("GET", "http://godocs.io", nil)
   if err != nil {
      t.Fatal(err)
   }
   res, err := Default_Client.Status(302).Do(req)
   if err != nil {
      t.Fatal(err)
   }
   if err := res.Body.Close(); err != nil {
      t.Fatal(err)
   }
}
