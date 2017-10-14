# JSON Protocol Over HTTP, for Loose Coupling

##　システム全体を疎結合にする
モダンなシステムでは各モジュールやコンポーネント間がスパースなつながりである方が良いとされています。  
これは、あまりに多すぎる依存が発生すると、簡単にオーバーヘッドとなるモジュールをリプレースすることができなくなるためです。  

例えば、これをマシン間を横断して実現するための通信方式として代表的なものとしてHTTP(S)などが考えられます

## データのシリアライズ方式
代表的なものに以下のものがあります  

- [ProtocolBuffer](https://developers.google.com/protocol-buffers/)
- [MsgPack](http://msgpack.org/)
- JSON
- XML

