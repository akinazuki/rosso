# ProtoBuf

I reject the idea of having to use a compiler for ProtoBuf. I think you should
be able to Marshal and Unmarshal just like JSON. And really, that should be
possible. If they had only added a single extra wiretype for messages, ProtoBuf
would be more or less self describing. At any rate, I want a package that can
decode arbitrary ProtoBuf, and can also encode some Map or Struct into ProtoBuf
as well.

- https://github.com/golang/protobuf/issues/1370
- https://stackoverflow.com/questions/41348512/protobuf-unmarshal-unknown

## Why not use generic add?

If you use generic add with method interface, it would allow recurive slices. We
could solve this by also adding type interface, but its probably simpler to just
not use generics for this.

## Why not implement a decoder?

I tried, but Unmarshal is way faster:

~~~
Benchmark_Decode-12       367295 ns/op         1481091 B/op       2588 allocs/op
Benchmark_Unmarshal-12     86033 ns/op           80245 B/op       1104 allocs/op
~~~
