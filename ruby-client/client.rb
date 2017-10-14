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
