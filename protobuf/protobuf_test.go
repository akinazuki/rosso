package protobuf

import (
   "fmt"
   "testing"
)

func Test_Error(t *testing.T) {
   var err type_error
   fmt.Println(err.Error())
}
