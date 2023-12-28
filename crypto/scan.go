package crypto

import (
   "github.com/refraction-networking/utls"
   "strconv"
   "strings"
)

func Parse_JA3(str string) (*tls.ClientHelloSpec, error) {
   var (
      extensions string
      info tls.ClientHelloInfo
      spec tls.ClientHelloSpec
   )
   for i, field := range strings.SplitN(str, ",", 5) {
      switch i {
      case 0:
         // TLSVersMin is the record version, TLSVersMax is the handshake
         // version
         u, err := strconv.ParseUint(field, 10, 16)
         if err != nil {
            return nil, err
         }
         spec.TLSVersMax = uint16(u)
      case 1:
         // build CipherSuites
         for _, s := range strings.Split(field, "-") {
            u, err := strconv.ParseUint(s, 10, 16)
            if err != nil {
               return nil, err
            }
            spec.CipherSuites = append(spec.CipherSuites, uint16(u))
         }
      case 2:
         extensions = field
      case 3:
         for _, s := range strings.Split(field, "-") {
            u, err := strconv.ParseUint(s, 10, 16)
            if err != nil {
               return nil, err
            }
            info.SupportedCurves = append(info.SupportedCurves, tls.CurveID(u))
         }
      case 4:
         for _, s := range strings.Split(field, "-") {
            u, err := strconv.ParseUint(s, 10, 8)
            if err != nil {
               return nil, err
            }
            info.SupportedPoints = append(info.SupportedPoints, uint8(u))
         }
      }
   }
   // build extenions list
   for _, s := range strings.Split(extensions, "-") {
      var ext tls.TLSExtension
      switch s {
      case "0":
         // Android API 24
         ext = &tls.SNIExtension{}
      case "5":
         // Android API 26
         ext = &tls.StatusRequestExtension{}
      case "10":
         ext = &tls.SupportedCurvesExtension{Curves: info.SupportedCurves}
      case "11":
         ext = &tls.SupportedPointsExtension{
            SupportedPoints: info.SupportedPoints,
         }
      case "13":
         ext = &tls.SignatureAlgorithmsExtension{
            SupportedSignatureAlgorithms: []tls.SignatureScheme{
               // Android API 24
               tls.ECDSAWithP256AndSHA256,
               // httpbin.org
               tls.PKCS1WithSHA256,
            },
         }
      case "16":
         // Android API 24
         ext = &tls.ALPNExtension{
            AlpnProtocols: []string{"http/1.1"},
         }
      case "23":
         // Android API 24
         ext = &tls.UtlsExtendedMasterSecretExtension{}
      case "27":
         // Google Chrome
         ext = &tls.UtlsCompressCertExtension{
            Algorithms: []tls.CertCompressionAlgo{tls.CertCompressionBrotli},
         }
      case "43":
         // Android API 29
         ext = &tls.SupportedVersionsExtension{
            Versions: []uint16{tls.VersionTLS12},
         }
      case "45":
         // Android API 29
         ext = &tls.PSKKeyExchangeModesExtension{
            Modes: []uint8{tls.PskModeDHE},
         }
      case "65281":
         // Android API 24
         ext = &tls.RenegotiationInfoExtension{}
      default:
         u, err := strconv.ParseUint(s, 10, 16)
         if err != nil {
            return nil, err
         }
         ext = &tls.GenericExtension{Id: uint16(u)}
      }
      spec.Extensions = append(spec.Extensions, ext)
   }
   // uTLS does not support 0x0 as min version
   spec.TLSVersMin = tls.VersionTLS10
   return &spec, nil
}
