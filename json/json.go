package json

import (
   "bytes"
   "encoding/json"
)

var (
   Marshal = json.Marshal
   MarshalIndent = json.MarshalIndent
   NewDecoder = json.NewDecoder
   NewEncoder = json.NewEncoder
   Unmarshal = json.Unmarshal
)

type Scanner struct {
   Data []byte
   Sep []byte
}

func (s Scanner) Decode(val any) error {
   data := append(s.Sep, s.Data...)
   dec := NewDecoder(bytes.NewReader(data))
   for {
      _, err := dec.Token()
      if err != nil {
         high := dec.InputOffset()
         return json.Unmarshal(data[:high], val)
      }
   }
}

func (s *Scanner) Scan() bool {
   var found bool
   _, s.Data, found = bytes.Cut(s.Data, s.Sep)
   return found
}
