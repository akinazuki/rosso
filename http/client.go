package http

import (
   "bytes"
   "errors"
   "github.com/89z/rosso/strconv"
   "net/http"
   "net/http/httputil"
   "os"
)

type Client struct {
   Log_Level int // this needs to work with flag.IntVar
   status int
   client http.Client
}

var Default_Client = Client{
   Log_Level: 1,
   client: http.Client{
      CheckRedirect: func(*http.Request, []*http.Request) error {
         return http.ErrUseLastResponse
      },
   },
   status: http.StatusOK,
}

func (c Client) Do(req *http.Request) (*http.Response, error) {
   switch c.Log_Level {
   case 1:
      os.Stderr.WriteString(req.Method)
      os.Stderr.WriteString(" ")
      os.Stderr.WriteString(req.URL.String())
      os.Stderr.WriteString("\n")
   case 2:
      dump, err := httputil.DumpRequest(req, true)
      if err != nil {
         return nil, err
      }
      if !strconv.Valid(dump) {
         dump = strconv.AppendQuote(nil, string(dump))
      }
      if !bytes.HasSuffix(dump, []byte{'\n'}) {
         dump = append(dump, '\n')
      }
      os.Stderr.Write(dump)
   }
   res, err := c.client.Do(req)
   if err != nil {
      return nil, err
   }
   if res.StatusCode != c.status {
      return nil, errors.New(res.Status)
   }
   return res, nil
}

func (c Client) Get(ref string) (*http.Response, error) {
   req, err := http.NewRequest("GET", ref, nil)
   if err != nil {
      return nil, err
   }
   return c.Do(req)
}

func (c Client) Level(level int) Client {
   c.Log_Level = level
   return c
}

func (c Client) Redirect(fn Redirect_Func) Client {
   c.client.CheckRedirect = nil
   return c
}

func (c Client) Status(status int) Client {
   c.status = status
   return c
}

func (c Client) Transport(tr *http.Transport) Client {
   c.client.Transport = tr
   return c
}

type Redirect_Func func(*http.Request, []*http.Request) error
