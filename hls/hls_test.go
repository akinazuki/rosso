package hls

import (
   "fmt"
   "github.com/89z/rosso/http"
   "io"
   "os"
   "strings"
   "testing"
)

func Test_Block(t *testing.T) {
   res, err := client.Get(ref)
   if err != nil {
      t.Fatal(err)
   }
   seg, err := New_Scanner(res.Body).Segment()
   if err != nil {
      t.Fatal(err)
   }
   if err := res.Body.Close(); err != nil {
      t.Fatal(err)
   }
   key, err := get_key(seg.Key)
   if err != nil {
      t.Fatal(err)
   }
   file, err := os.Create("ignore.ts")
   if err != nil {
      t.Fatal(err)
   }
   defer file.Close()
   block, err := New_Block(key)
   if err != nil {
      t.Fatal(err)
   }
   for i, ref := range seg.URI {
      fmt.Println(len(seg.URI)-i)
      res, err := client.Level(0).Get(ref)
      if err != nil {
         t.Fatal(err)
      }
      text, err := io.ReadAll(res.Body)
      if err != nil {
         t.Fatal(err)
      }
      text = block.Decrypt_Key(text)
      if _, err := file.Write(text); err != nil {
         t.Fatal(err)
      }
      if err := res.Body.Close(); err != nil {
         t.Fatal(err)
      }
   }
}

func apple_media(m Medium) bool {
   if !strings.Contains(m.Group_ID, "-ak-") {
      return false
   }
   if m.Name != "English" {
      return false
   }
   if m.Type != "AUDIO" {
      return false
   }
   return true
}

func apple_stream(s Stream) bool {
   if !strings.Contains(s.Audio, "-ak-") {
      return false
   }
   if !strings.Contains(s.Codecs, "avc1.") {
      return false
   }
   if !strings.Contains(s.Codecs, "mp4a.") {
      return false
   }
   return true
}

func cbc_media(m Medium) bool {
   return m.Type == "AUDIO"
}

func cbc_stream(s Stream) bool {
   return strings.Contains(s.Codecs, "avc1.")
}

func nbc_media(m Medium) bool {
   return m.Type == "AUDIO"
}

func paramount_stream(s Stream) bool {
   return strings.Contains(s.Codecs, "avc1.")
}

var tests = map[string]filters{
   "m3u8/nbc-master.m3u8": {nbc_media, nil},
   "m3u8/roku-master.m3u8": {nil, nil},
   "m3u8/paramount-master.m3u8": {nil, paramount_stream},
   "m3u8/cbc-master.m3u8": {cbc_media, cbc_stream},
   "m3u8/apple-master.m3u8": {apple_media, apple_stream},
}

type filters struct {
   medium func(Medium) bool
   stream func(Stream) bool
}

func Test_Media(t *testing.T) {
   for key, val := range tests {
      file, err := os.Open(key)
      if err != nil {
         t.Fatal(err)
      }
      master, err := New_Scanner(file).Master()
      if err != nil {
         t.Fatal(err)
      }
      if err := file.Close(); err != nil {
         t.Fatal(err)
      }
      master.Media = master.Media.Filter(val.medium)
      target := master.Media.Index(func(carry, item Medium) bool {
         return item.Name == "English"
      })
      fmt.Println(key)
      for i, medium := range master.Media {
         if i == target {
            fmt.Print("!")
         }
         fmt.Println(medium)
      }
      fmt.Println()
   }
}

func Test_Stream(t *testing.T) {
   for key, val := range tests {
      file, err := os.Open(key)
      if err != nil {
         t.Fatal(err)
      }
      master, err := New_Scanner(file).Master()
      if err != nil {
         t.Fatal(err)
      }
      if err := file.Close(); err != nil {
         t.Fatal(err)
      }
      items := master.Streams.Filter(val.stream)
      index := items.Bandwidth(0)
      fmt.Println(key)
      for i, item := range items {
         if i == index {
            fmt.Print("!")
         }
         fmt.Println(item)
      }
      fmt.Println()
   }
}

func Test_Info(t *testing.T) {
   for key, val := range tests {
      file, err := os.Open(key)
      if err != nil {
         t.Fatal(err)
      }
      master, err := New_Scanner(file).Master()
      if err != nil {
         t.Fatal(err)
      }
      if err := file.Close(); err != nil {
         t.Fatal(err)
      }
      fmt.Println(key)
      for _, item := range master.Streams.Filter(val.stream) {
         fmt.Println(item)
      }
      for _, item := range master.Media.Filter(val.medium) {
         fmt.Println(item)
      }
      fmt.Println()
   }
}
func Test_Segment(t *testing.T) {
   names := []string{
      "m3u8/apple-audio.m3u8",
      "m3u8/cbc-video.m3u8",
      "m3u8/roku-segment.m3u8",
   }
   for _, name := range names {
      file, err := os.Open(name)
      if err != nil {
         t.Fatal(err)
      }
      seg, err := New_Scanner(file).Segment()
      if err != nil {
         t.Fatal(err)
      }
      if err := file.Close(); err != nil {
         t.Fatal(err)
      }
      fmt.Printf("%+v\n\n", seg)
   }
}

var raw_ivs = []string{
   "00000000000000000000000000000001",
   "0X00000000000000000000000000000001",
   "0x00000000000000000000000000000001",
}

func Test_Hex(t *testing.T) {
   for _, raw_iv := range raw_ivs {
      iv, err := Segment{Raw_IV: raw_iv}.IV()
      if err != nil {
         t.Fatal(err)
      }
      fmt.Println(iv)
   }
}

// paramount -b 622520382 -f 499000
const ref = "https://cbsios-vh.akamaihd.net/i/temp_hd_gallery_video/CBS_Production_Outlet_VMS/video_robot/CBS_Production_Entertainment/2012/09/12/41581439/CBS_MELROSE_PLACE_001_SD_prores_78930_,503,4628,3128,2228,1628,848,000.mp4.csmil/index_0_av.m3u8?null=0&id=AgBItRcmFy81SkUfwWIsRdilI6s+0hIRmFI6R378aTEqsuj0TmwsVvPmGEoeaIYYS8H6mKrNRB0PPQ%3d%3d&hdntl=exp=1656910021~acl=%2fi%2ftemp_hd_gallery_video%2fCBS_Production_Outlet_VMS%2fvideo_robot%2fCBS_Production_Entertainment%2f2012%2f09%2f12%2f41581439%2fCBS_MELROSE_PLACE_001_SD_prores_78930_*~data=hdntl~hmac=d571a5878bd4532e7fc553c8a9fd1374e039c9506295dacdcc10533b991a3447"

var client = http.Default_Client

func get_key(s string) ([]byte, error) {
   res, err := client.Get(s)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   return io.ReadAll(res.Body)
}
