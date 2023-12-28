package mp4

import (
   "github.com/edgeware/mp4ff/mp4"
   "io"
)

type Decrypt struct {
   sinf map[uint32]*mp4.SinfBox
   write io.Writer
}

func New_Decrypt(w io.Writer) Decrypt {
   var dec Decrypt
   dec.sinf = make(map[uint32]*mp4.SinfBox)
   dec.write = w
   return dec
}

func (d *Decrypt) Init(r io.Reader) error {
   file, err := mp4.DecodeFile(r)
   if err != nil {
      return err
   }
   // need for VLC media player
   for _, trak := range file.Init.Moov.Traks {
      for _, child := range trak.Mdia.Minf.Stbl.Stsd.Children {
         switch box := child.(type) {
         case *mp4.AudioSampleEntryBox:
            d.sinf[trak.Tkhd.TrackID], err = box.RemoveEncryption()
         case *mp4.VisualSampleEntryBox:
            d.sinf[trak.Tkhd.TrackID], err = box.RemoveEncryption()
         }
         if err != nil {
            return err
         }
      }
   }
   // need for Mozilla Firefox
   file.Init.Moov.RemovePsshs()
   return file.Init.Encode(d.write)
}

func (d Decrypt) Segment(r io.Reader, key []byte) error {
   file, err := mp4.DecodeFile(r)
   if err != nil {
      return err
   }
   for _, seg := range file.Segments {
      for _, frag := range seg.Fragments {
         var removed uint64
         for _, traf := range frag.Moof.Trafs {
            sinf := d.sinf[traf.Tfhd.TrackID]
            if sinf == nil {
               continue
            }
            samples, err := frag.GetFullSamples(nil)
            if err != nil {
               return err
            }
            tenc := sinf.Schi.Tenc
            for i, sample := range samples {
               iv := tenc.DefaultConstantIV
               if iv == nil {
                  iv = append(iv, traf.Senc.IVs[i]...)
                  iv = append(iv, 0, 0, 0, 0, 0, 0, 0, 0)
               }
               var sub []mp4.SubSamplePattern
               if len(traf.Senc.SubSamples) > i {
                  // required for playback
                  sub = traf.Senc.SubSamples[i]
               }
               switch sinf.Schm.SchemeType {
               case "cenc":
                  err = mp4.DecryptSampleCenc(sample.Data, key, iv, sub)
               case "cbcs":
                  err = mp4.DecryptSampleCbcs(sample.Data, key, iv, sub, tenc)
               }
               if err != nil {
                  return err
               }
            }
            // required for playback
            removed += traf.RemoveEncryptionBoxes()
         }
         // fast start
         _, pssh := frag.Moof.RemovePsshs()
         removed += pssh
         for _, traf := range frag.Moof.Trafs {
            for _, trun := range traf.Truns {
               // required for playback
               trun.DataOffset -= int32(removed)
            }
         }
      }
      // fix jerk between fragments
      seg.Sidx = nil
      err := seg.Encode(d.write)
      if err != nil {
         return err
      }
   }
   return nil
}
