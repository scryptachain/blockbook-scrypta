#!/bin/bash

make all-scrypta
cd run
mkdir build
cp ../build/*.deb ./build/
docker build -t scrypta:blockbook-scrypta .
docker run --privileged -d --name scryptabb -p 9149:9149 scrypta:blockbook-scrypta
docker exec scryptabb bash -c "cd /root && ./launchd.sh &"
sleep 60
docker exec scryptabb bash -c "cd /root && ./launchbb.sh &"
sleep 30
docker exec scryptabb bash -c "cd /root && ./launchbb.sh &"