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
