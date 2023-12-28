package xml

import (
   "bytes"
   "encoding/xml"
)

func decoder(data []byte) *xml.Decoder {
   dec := xml.NewDecoder(bytes.NewReader(data))
   dec.AutoClose = xml.HTMLAutoClose
   dec.Strict = false
   return dec
}

type Scanner struct {
   Data []byte
   Sep []byte
}

func (s Scanner) Decode(val any) error {
   data := append(s.Sep, s.Data...)
   dec := decoder(data)
   for {
      _, err := dec.Token()
      if err != nil {
         high := dec.InputOffset()
         return decoder(data[:high]).Decode(val)
      }
   }
}

func (s *Scanner) Scan() bool {
   var found bool
   _, s.Data, found = bytes.Cut(s.Data, s.Sep)
   return found
}
