
from flask import Flask, request, jsonify
import json
application = Flask(__name__)

@application.route("/")
def hello():
    return "<h1 style='color:blue'>Hello There!</h1>"

@application.route("/json",  methods=['GET', 'POST'])
def json():
	content = request.json	
	print(content)
	return f"<h1 style='color:blue'>{content['username']}</h1>"

if __name__ == "__main__":
    application.run(host='0.0.0.0')
