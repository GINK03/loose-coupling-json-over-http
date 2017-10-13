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
