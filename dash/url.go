package dash

import (
   "strings"
)

func (r Representation) Initialization() string {
   return r.replace_ID(r.SegmentTemplate.Initialization)
}

func (r Representation) Media() []string {
   var start int
   if r.SegmentTemplate.StartNumber != nil {
      start = *r.SegmentTemplate.StartNumber
   }
   var refs []string
   for _, seg := range r.SegmentTemplate.SegmentTimeline.S {
      for seg.T = start; seg.R >= 0; seg.R-- {
         ref := r.replace_ID(r.SegmentTemplate.Media)
         if r.SegmentTemplate.StartNumber != nil {
            ref = strings.Replace(ref, "$Number$", seg.Time(), 1)
            seg.T++
            start++
         } else {
            ref = strings.Replace(ref, "$Time$", seg.Time(), 1)
            seg.T += seg.D
            start += seg.D
         }
         refs = append(refs, ref)
      }
   }
   return refs
}
