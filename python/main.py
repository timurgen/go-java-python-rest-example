import random
import json

from datetime import datetime
from flask import Flask, request, Response
from waitress import serve

APP = Flask(__name__)


@APP.route("/whosnext", methods=['POST'])
def process_request():
    next_person = random.choice(request.get_json())
    print("Magic 8-Ball says:", next_person["name"])
    return Response(json.dumps(next_person), content_type="application/json")


if __name__ == "__main__":
    random.seed(datetime.now())
    serve(APP, host='0.0.0.0', port=8080)
