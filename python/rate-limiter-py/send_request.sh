#!/bin/bash

num="$1"
for ((i=1;i<=num;i++)); do
  curl -s http://localhost:5000/limited
  echo
done
