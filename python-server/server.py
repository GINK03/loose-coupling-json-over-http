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
