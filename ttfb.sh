#!/bin/bash
# file: ttfb.sh
# curl command to check the time to first byte
# ** usage **
# 1. ./ttfb.sh "https://google.com"
# 2. seq 10 | xargs -Iz ./ttfb.sh "https://google.com"

start=`date +%s`

function launchCurl()
{
    curl -o /dev/null \
         -H 'Cache-Control: no-cache' \
         -s \
         -w "Connect: %{time_connect} | TTFB: %{time_starttransfer} | Total time: %{time_total} \n" \
         $1
}

if [ $# -eq 1 ]; then
    launchCurl $1
elif [ $# -eq 2 ]; then
    index=1
    while [ $index -le $2 ]; do
        launchCurl $1 &
        ((index++))
    done
else
    echo "Usage: ./ttfb.sh <url> <loop number>"
fi

wait

end=`date +%s`

echo "Time: $((end-start))s"
