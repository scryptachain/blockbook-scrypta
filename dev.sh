#!/bin/bash

make deb-blockbook-scrypta
docker container stop scryptabb
docker container start scryptabb

cd run
mkdir build
cp ../build/*.deb ./build/
docker exec scryptabb bash -c "cd /root && rm blockbook.deb"
docker exec scryptabb bash -c "apt remove blockbook-scrypta -y"
docker cp ./build/blockbook-scrypta_0.3.6_amd64.deb scryptabb:/root/backend.deb
docker exec scryptabb bash -c "apt install /root/blockbook.deb -y"

docker exec scryptabb bash -c "cd /root && ./launchbb.sh &"
sleep 30
docker exec scryptabb bash -c "cd /root && ./launchbb.sh &"