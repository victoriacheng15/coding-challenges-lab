#!/bin/bash

num=70
for ((i=1;i<=num;i++)); do
  curl -s http://localhost:5000/limited
  echo
done
