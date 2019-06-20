#!/bin/bash
# file: ttfb.sh
# curl command to check the time to first byte
# ** usage **
# 1. ./ttfb.sh "https://google.com"
# 2. seq 10 | xargs -Iz ./ttfb.sh "https://google.com"

curl -o /dev/null \
     -H 'Cache-Control: no-cache' \
     -s \
     -w "Connect: %{time_connect} TTFB: %{time_starttransfer} Total time: %{time_total} \n" \
     $1