# JSON Protocol Over HTTP, for Loose Coupling

## システム全体を疎結合にする
モダンなシステムでは各モジュールやコンポーネント間がスパースなつながりである方が良いとされています。  
これは、あまりに多すぎる依存が発生すると、簡単にオーバーヘッドとなるモジュールをリプレースすることができなくなるためです。  

例えば、これをマシン間を横断して実現するための通信方式として代表的なものとしてHTTP(S)などが考えられます

例えば、JSON形式のデータオブジェクトを、HTTP経由で通信しコントロールすることができます  

## データのシリアライズ方式
代表的なものに以下のものがあります  

- [ProtocolBuffer](https://developers.google.com/protocol-buffers/)
- [MsgPack](http://msgpack.org/)
- JSON
- XML

## 様々な言語のHTTP ServerとJSONのデータの受け渡しと受け取り

### Python
server.httpというモジュールを利用することで簡単にサーバを構築することができます  
clientでは別途、モジュールとしてrequestsを利用するとシンプルにかけて良いです  

#### Server
```python
import http.server
import socketserver
import json
class Handler(http.server.SimpleHTTPRequestHandler):
  def _set_response(self):
    self.send_response(200)
    self.send_header('Content-type', 'text/html')
    self.end_headers()

  def do_GET(self):
    print("Client requested:", self.command, self.path )
    http.server.SimpleHTTPRequestHandler.do_GET(self)
    self.wfile.write(bytes("Hello World !",'utf8'))

  def do_POST(self):
    content_length = int(self.headers['Content-Length'])
    data = self.rfile.read(content_length)
    data = json.loads( data.decode() )
    print( data )
    self._set_response()
    # write here what you want to do
    self.wfile.write(bytes( json.dumps(data),'utf8'))

httpd = socketserver.TCPServer(('0.0.0.0', 4567), Handler)
httpd.serve_forever()
```
#### Client
```python
import json
import requests

payload = {'type': 'ADD', 'data': 1.5}
r = requests.post('http://localhost:4567/test', data=json.dumps(payload) )
print( json.loads( r.text ) )
```

### Ruby
Rubyが一番簡単に書ける気がします  
何か標準モジュール以外に参照する必要のモジュールは必要ありません
#### Server
```ruby
require 'webrick'
require 'json'
class Echo < WEBrick::HTTPServlet::AbstractServlet
  def do_GET(request, response)
    puts request
    response.status = 200
  end
  def do_POST(request, response)
    data = request.body
    obj = JSON.parse(data)
    puts( obj )
    response.status = 200
    response.body = JSON.dump( obj )
  end
end
server = WEBrick::HTTPServer.new(:Port => 8080)
server.mount "/test", Echo
trap "INT" do server.shutdown end
server.start
```
#### Client
```ruby
require 'net/http'
require 'uri'
require 'json'
# 詳細
url = URI.parse('http://localhost:4567/test')
req = Net::HTTP::Post.new(url.path)
req.body = {'type' => 'ADD', 'data' => '2.5'}.to_json()

res = Net::HTTP.new(url.host, url.port).start do |http|
  http.request(req)
end
puts res.code
puts res.body
```

### Kotlin(JVM)
KotlinはJavaと一緒でSparkと呼ばれるWeb Frameworkを用いることができます  
JVMのJavaの最大の特徴は、オブジェクトにマップすることができるので、オブジェクトやEnumの値によって動作を分けることができるので、ソリッドなコードを書きやすいです  

#### Server
```kotlin
import kotlinx.serialization.*
import kotlinx.serialization.json.JSON

import spark.Spark.*
enum class OrderType {
  MULTIPLY, ADD, MINUS, DIV
}

@Serializable
data class Order(val type:OrderType, val data:Double)

@Serializable
data class Return(val data:Double?)
fun main( args : Array<String> ) {
  //port(80)
  get("/hello") { req, res -> "hello world" }

  post("/test") { req, res ->
    println( req.body() )
    println( req.queryParams() ) // input format as list
    //print( req.queryMap().toMap() ) // input format as map

    // get first entry of queryParams
    val body = req.body() 
    val obj = JSON.parse<Order>(body)

    // implimation some logic here
    val ret = when( obj.type ) {
      OrderType.ADD -> obj.data + 1.0
      OrderType.DIV -> obj.data/2.0
      OrderType.MULTIPLY -> obj.data*2.0
      OrderType.MINUS -> obj.data-1.0
      else -> null
    }
    JSON.stringify(Return(ret))
  }
}
```
#### Client
Clientにはokhttpというhttpリクエストを送るためのライブラリが別途必要であり、これはJavaと同じです  
```kotlin
import kotlinx.serialization.*
import kotlinx.serialization.json.JSON

import spark.Spark.*

import okhttp3.*

enum class OrderType {
  MULTIPLY, ADD, MINUS, DIV
}

@Serializable
data class Order(val type:OrderType, val data:Double)

@Serializable
data class Return(val data:Double?)
fun main( args : Array<String> ) {
  val JSON  = MediaType.parse("application/json; charset=utf-8")
  val TEXT = MediaType.parse("text/plain")
  val client =  OkHttpClient()
  val url = "http://localhost:4567/test"
  val body = RequestBody.create(JSON, """{"type":"ADD","data":1.0}""")
  val request = Request.Builder()
      .url(url)
      .post(body)
      .build();
  val response = client.newCall(request).execute()

  println( response.body()?.string() ?: "Cannot access" )
}
```

### C++
C++でもできますが、serverはCROWというMicro Web Frameworkを用いて実現することがきるとともに、ClientはCライブラリであるlibcurlを利用できます  
依存やなんやらが多くパッケージ管理しにくいから、あまりお勧めではありません。  
JSONはAny型と一緒で特定のオブジェクトにマップするようなことはしていません(Boost.Serializationなどを見るとオブジェクトマップはできる気がする)  
#### Server
```cpp
    CROW_ROUTE(app,"/test")
        .methods("POST"_method)
    ([](const crow::request& req){
        auto xi = crow::json::load(req.body);
        std::cout << xi << std::endl;
        std::ostringstream os;
        crow::json::wvalue xr;
        xr["type"] = "RESULT";
        xr["data"] = 5.5;
        return xr;
    });
```
#### Client
Clientはずっとスッキリと書くことができます  
```cpp
#include <string>
#include <iostream>
#include <curl/curl.h>

using namespace std;

static size_t call_back(char* ptr, size_t size, size_t nmemb, std::string* str) {
  int realsize = size * nmemb;
  str->append(ptr, realsize);
  return realsize;
}
static std::string curl_post_wrapper(std::string url, std::string post_data) {
  CURL *curl;
  CURLcode res;
  curl = curl_easy_init();
  if( curl == nullptr ) {
    return "Error";
  }
  string chunk = "";
  curl_easy_setopt(curl, CURLOPT_URL, url.c_str());
  curl_easy_setopt(curl, CURLOPT_POSTFIELDS, post_data.c_str());
  curl_easy_setopt(curl, CURLOPT_POSTFIELDSIZE, post_data.size());
  curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, call_back);
  curl_easy_setopt(curl, CURLOPT_WRITEDATA, static_cast<string*>(&chunk));
  res = curl_easy_perform(curl);
  curl_easy_cleanup(curl);
  return chunk;
}
int main() {
  std::string res = curl_post_wrapper("http://localhost:4567/test", "{\"type\":\"ADD\",\"data\":4.0}");
  cout << res << endl;
}
```
