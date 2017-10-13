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
