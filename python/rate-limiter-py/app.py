from flask import Flask, request, jsonify
import time

app = Flask(__name__)

WINDOW_SIZE = 60
MAX_REQUESTS = 60

request_logs  = {}

def is_request_allowed(ip):
    now = time.time()
    reset_in = WINDOW_SIZE

    if ip not in request_logs:
        request_logs[ip] = []

    recent_requests = [ts for ts in request_logs[ip] if now - ts < WINDOW_SIZE]
    request_logs[ip] = recent_requests

    if len(recent_requests) < MAX_REQUESTS:
        request_logs[ip].append(now)
        remaining = MAX_REQUESTS - len(request_logs[ip])
        reset_in = int(WINDOW_SIZE - (now - request_logs[ip][0])) if request_logs[ip] else WINDOW_SIZE
        return {
            "allowed": True,
            "remaining": remaining,
            "reset_in_seconds": reset_in
        }
    else:
        reset_in = int(WINDOW_SIZE - (now - request_logs[ip][0])) if request_logs[ip] else WINDOW_SIZE
        return {
            "allowed": False,
            "remaining": 0,
            "reset_in_seconds": reset_in
        }


@app.route('/')
def index():
    return "Welcome to the Rate Limiter App!"


@app.route('/limited')
def limited():
    ip = request.remote_addr
    result = is_request_allowed(ip)

    if result["allowed"]:
        return jsonify({
            "message": f"limited endpoint accessed by {ip}",
            "remaining": result.get("remaining", 0),
            "reset_in_seconds": result.get("reset_in_seconds", 0),
        }), 200
    else:
        return jsonify({
            "error": "Rate limit exceeded. Try again later.",
            "remaining": 0,
            "reset_in_seconds": result.get("reset_in_seconds", 0),
        }), 429


@app.route('/unlimited')
def unlimited():
    return "Unlimited! Let's Go!"

@app.route('/stats')
def stats_home():
    now = time.time()
    snapshot = {}

    for ip, logs in request_logs.items():
        active_logs = [ts for ts in logs if now - ts < WINDOW_SIZE]
        remaining = max(MAX_REQUESTS - len(active_logs), 0)
        reset_in = int(WINDOW_SIZE - (now - active_logs[0])) if active_logs else WINDOW_SIZE

        snapshot[ip] = {
            "requests_in_last_60_seconds": len(active_logs),
            "max_requests": MAX_REQUESTS,
            "remaining": remaining,
            "reset_in_seconds": reset_in
        }

    return jsonify(snapshot), 200


if __name__ == '__main__':
    app.run(port=5000, debug=True)