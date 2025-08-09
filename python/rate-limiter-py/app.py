from flask import Flask, request, jsonify
import time

app = Flask(__name__)

WINDOW_SIZE = 60
MAX_REQUESTS = 60

# Store counters per IP address
# Format: { ip: { current_period: count, previous_period: count } }
request_counters = {}

def is_request_allowed(ip, now=None):
    if now is None:
        now = time.time()
    current_period = int(now // WINDOW_SIZE)
    elapsed_in_window = now % WINDOW_SIZE
    weight = elapsed_in_window / WINDOW_SIZE
    previous_period = current_period - 1

    if ip not in request_counters:
        request_counters[ip] = {}

    # Default counts to 0
    current_count = request_counters[ip].get(current_period, 0)
    previous_count = request_counters[ip].get(previous_period, 0)

    # Weighted total
    total = current_count + (1 - weight) * previous_count

    if total < MAX_REQUESTS:
        request_counters[ip][current_period] = current_count + 1

        return {
            "allowed": True,
            "remaining": max(int(MAX_REQUESTS - total), 0),
            "reset_in_seconds": int(WINDOW_SIZE - elapsed_in_window)
        }
    else:
        return {
            "allowed": False,
            "remaining": 0,
            "reset_in_seconds": int(WINDOW_SIZE - elapsed_in_window)
        }
    
def get_client_ip():
    ip = request.headers.get('X-Forwarded-For', request.remote_addr)
    if ip and ',' in ip:
        ip = ip.split(',')[0].strip()
    return ip

@app.route('/')
def index():
    return "Welcome to the Rate Limiter App!"


@app.route('/limited')
def limited():
    ip = get_client_ip()
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
    current_period = int(now // WINDOW_SIZE)
    previous_period = current_period - 1
    reset_in_seconds = int((current_period + 1) * WINDOW_SIZE - now)

    snapshot = {}
    for ip, windows in request_counters.items():
        cur = windows.get(current_period, 0)
        prev = windows.get(previous_period, 0)
        snapshot[ip] = {
            "current_period": cur,
            "previous_period": prev,
            "max_requests": MAX_REQUESTS,
            "reset_in_seconds": reset_in_seconds
        }

    return jsonify(snapshot), 200

if __name__ == '__main__':
    app.run(port=5000, debug=True)