#!/bin/bash

log_file="log.txt"
interval=3

while true; do
    date_str=$(date "+%Y-%m-%d %H:%M:%S")
    result=$(curl -s http://34.85.20.100/ | grep "Server Location:")
    echo "$date_str $result" >>"$log_file"
    sleep "$interval"
done
