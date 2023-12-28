package main

import (
   "encoding/json"
   "flag"
   "github.com/89z/rosso/protobuf"
   "os"
)

func main() {
   // f
   var input string
   flag.StringVar(&input, "f", "", "input file")
   // o
   var output string
   flag.StringVar(&output, "o", "", "output file")
   flag.Parse()
   if input != "" {
      err := do_protobuf(input, output)
      if err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}

func do_protobuf(input, output string) error {
   data, err := os.ReadFile(input)
   if err != nil {
      return err
   }
   mes, err := protobuf.Unmarshal(data)
   if err != nil {
      return err
   }
   file, err := os.Create(output)
   if err != nil {
      file = os.Stdout
   }
   defer file.Close()
   enc := json.NewEncoder(file)
   enc.SetEscapeHTML(false)
   enc.SetIndent("", " ")
   return enc.Encode(mes)
}
