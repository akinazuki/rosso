# Crypto

## TLS version

The version is actually set in three places:

1. record version
2. handshake version
3. supported versions

with regards to JA3, "handshake version" should always be used. "supported
versions" cannot be used, as its an array not a single value. further, "record
version" cannot be used, as its deprecated since TLS 1.3. More info:

https://datatracker.ietf.org/doc/html/rfc8446#appendix-D

In summary, the TLSVersion should never be higher than `0x0303` (771), as its
locked at that value since TLS 1.3. That means that anyone reporting 772 or
higher is wrong.

## How to get Android JA3?

https://github.com/89z/rosso/tree/v1.47.3/cmd/proxy

## Servers

- https://tlshello.agwa.name
- https://github.com/AGWA/tlshacks

also:

- https://github.com/wwhtrbbtt/TrackMe
- https://tls.peet.ws

also:

- https://github.com/yuzd/ja3-csharp
- https://kawayiyi.com/tls

## What about Akamai fingerprint?

Paper says HTTP/2 only:

https://blackhat.com/docs/eu-17/materials/eu-17-Shuster-Passive-Fingerprinting-Of-HTTP2-Clients-wp.pdf

Confirmed:

~~~
> curl --http1.1 https://tls.peet.ws/api/clean
{
  "ja3": "771,4866-4867-4865-49196-49200-159-52393-52392-52394-49195-49199-158-49188-49192-107-49187-49191-103-49162-49172-57-49161-49171-51-157-156-61-60-53-47-255,0-11-10-13172-16-22-23-49-13-43-45-51,29-23-30-25-24,0-1-2",
  "ja3_hash": "ba730f97dcd1122e74e65411e68f1b40",
  "akamai": "-",
  "akamai_hash": "-"
}
~~~

I would need to add HTTP/2 support to my existing code:

https://github.com/refraction-networking/utls/blob/9d36ce36/examples/examples.go#L417-L427

So in that case, supporting JA3 is simpler than supporting Akamai.

## Extensions

https://iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml
