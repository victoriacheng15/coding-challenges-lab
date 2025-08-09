# Coding Challenges - Sliding Window Rate Limiter (No Redis)

This project implements a rate limiter in Python using the **sliding window counter** algorithm, designed to limit the number of requests per user (IP) over time. It uses **Flask** for the web API and stores counters in memory, suitable for single-server setups and learning purposes.

## What I learned

Through this project, I gained practical experience with:

- **Sliding Window Counter Algorithm**:
  - Learned how to smooth request counts over time to avoid bursts at window boundaries.
  - Combined current and previous time intervals with weighted counts to implement a hybrid approach.

- **Flask API Design**:
  - Built endpoints for limited and unlimited access.
  - Handled IP extraction correctly by reading `X-Forwarded-For` headers for real-world proxy scenarios.
  - Provided clear JSON responses with rate limit status, remaining requests, and reset times.

- **Testing with Fake IPs**:
  - Simulated requests from different IPs using custom request headers.
  - Developed bash and Python scripts to automate sending requests and test rate limiting behavior.

## Installation

1. Clone the repository:

```bash
git clone git@github.com:victoriacheng15/coding-challenges-lab.git
```

2. Navigate to the project directory:

```bash
cd python/rate-limiter-py
```

3. (Optional) Setup a virtual environment:

```bash
python -m venv env
source env/bin/activate  # On Windows: env\Scripts\activate
```

4. Install dependencies:

```bash
pip install -r requirements.txt
```

## Usage

1. Start the Flask app:

```bash
python app.py
```

2. Send requests to the endpoints:

- `/limited` — Rate limited endpoint
- `/unlimited` — No rate limit

3. Use the included bash script to simulate requests from multiple IPs:

```bash
./send_requests.sh 10 fakeips
```

4. Access stats endpoint to see usage per IP:

```bash
curl http://localhost:5000/stats
```

---

## Project Structure

```plaintext
rate-limiter-py/
│
├── app.py                  # Flask app with rate limiter logic
├── requirements.txt        # Python dependencies including Flask
├── send_requests.sh        # Bash script to test rate limiting with fake IPs
├── README.md               # Project documentation
```

## Links

- [Coding Challenges - Rate Limiter Challenge](https://codingchallenges.fyi/challenges/challenge-rate-limiter)
