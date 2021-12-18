# Week 9 Homework

## 1. 总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。

### fix length

A decoder that splits received Bytes Buffer by fixed number of bytes. For example,
A fixed length frame decoder will decode them into three packets with fixed length from four fragmented packets.

Before:
+---+----+------+----+
| A | BC | DEFG | HI |
+---+----+------+----+
After:
+-----+-----+-----+
| ABC | DEF | GHI |
+-----+-----+-----+

### delimiter based length

A decoder that splits the received Byte Buffers by one or more delimiters. It is particularly useful for decoding the frames which ends with a delimiter such as NUL or newline characters.
For example,

Before:
+--------------+
| ABC\nDEF\r\n |
+--------------+

After:
+-----+-----+
| ABC | DEF |
+-----+-----+

### length field based frame decoder

A decoder that splits the received ByteBufs dynamically by the value of the length field in the message. It is particularly useful when you decode a binary message which has an integer header field that represents the length of the message body or the whole message.
For example, because we can get the length of content, you might want to strip the length field.

BEFORE DECODE (14 bytes) AFTER DECODE (12 bytes)
+--------+----------------+ +----------------+
| Length | Actual Content |----->| Actual Content |
| 0x000C | "HELLO, WORLD" | | "HELLO, WORLD" |
+--------+----------------+ +----------------+

## 2. 实现一个从 socket connection 中解码出 goim 协议的解码器。

![alt text](goim_tcp.png 'Goim TCP Protocol')

### Reference

- [Netty fix length](https://netty.io/4.0/api/io/netty/handler/codec/FixedLengthFrameDecoder.html)

- [Netty delimiter based length](https://netty.io/4.0/api/io/netty/handler/codec/DelimiterBasedFrameDecoder.html)

- [Netty length field based frame decoder](https://netty.io/4.0/api/io/netty/handler/codec/LengthFieldBasedFrameDecoder.html)
