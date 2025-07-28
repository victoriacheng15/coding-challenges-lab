from flask import Flask, request, jsonify
import time

app = Flask(__name__)

WINDOW_SIZE = 60
MAX_REQUESTS = 60

request_counters = {}

def is_request_allowed(ip):
    now = time.time()
    current_window = int(now // WINDOW_SIZE)
    next_window_start = (current_window + 1) * WINDOW_SIZE 
    reset_in = int(next_window_start - now)

    if ip not in request_counters:
        request_counters[ip] = {}

    if current_window not in request_counters[ip]:
        request_counters[ip][current_window] = 0

    count = request_counters[ip][current_window]

    if count < MAX_REQUESTS:
        request_counters[ip][current_window] += 1
        return {
            "allowed": True,
            "remaining": MAX_REQUESTS - count,
            "reset_in_seconds": reset_in
        }
    else:
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
    current_window = int(now // WINDOW_SIZE)

    snapshot = {}
    for ip, windows in request_counters.items():
        if current_window in windows:
            snapshot[ip] = {
                "requests_in_current_window": windows[current_window],
                "max_requests": MAX_REQUESTS,
                "remaining": max(MAX_REQUESTS - windows[current_window], 0),
                "reset_in_seconds": int((current_window + 1) * WINDOW_SIZE - now)
            }

    return jsonify(snapshot), 200


if __name__ == '__main__':
    app.run(port=5000, debug=True)