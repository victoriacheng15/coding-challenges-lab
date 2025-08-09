#!/bin/bash

num="$1"
mode="$2"  # optional: "fake" to send from IP array

if [ -z "$num" ]; then
  echo "Usage: $0 <num_requests> [fake]"
  echo "Example:"
  echo "  $0 10          # send 10 requests from your real IP"
  echo "  $0 10 fake     # send 10 requests per fake IP in list"
  exit 1
fi

fake_ips=(
  "1.1.1.1"
  "2.2.2.2"
  "3.3.3.3"
  "4.4.4.4"
)

if [ "$mode" == "fake" ]; then
  for ip in "${fake_ips[@]}"
  do
    echo "Sending $num requests from IP: $ip"
    for ((i=1; i<=num; i++))
    do
      curl -s -H "X-Forwarded-For: $ip" http://localhost:5000/limited
      echo
    done
    echo ""
  done
else
  echo "Sending $num requests from your real IP"
  for ((i=1; i<=num; i++))
  do
    curl -s http://localhost:5000/limited
    echo
  done
fi
