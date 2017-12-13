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
    self.wfile.write(bytes("Hello Machine Learning !",'utf8'))
  


httpd = socketserver.TCPServer(('0.0.0.0', 1234), Handler)
print("load oppai arch linux for machine learning")
httpd.serve_forever()
