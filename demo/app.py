# -*- coding: utf-8 -*-

from flask import Flask
from demo import demo_app

# 参数设置
DEBUG_MODE = False
BIND_ADDR = '0.0.0.0'
BIND_PORT = '5000'

app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Hello World!'

# demo
app.register_blueprint(demo_app)


if __name__ == '__main__':
    app.run(host=BIND_ADDR, port=BIND_PORT, debug=DEBUG_MODE)
