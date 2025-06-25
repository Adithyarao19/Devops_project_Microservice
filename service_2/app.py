from flask import Flask, jsonify
from prometheus_flask_exporter import PrometheusMetrics
from asgiref.wsgi import WsgiToAsgi

app = Flask(__name__)

# Initialize Prometheus metrics exporter
metrics = PrometheusMetrics(app)

@app.route("/ping")
def ping():
    return jsonify(status="ok", service="2")

@app.route("/hello")
def hello():
    return jsonify(message="Hello from Service 2")

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8002)


app = WsgiToAsgi(flask_app)
