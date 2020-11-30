#!/bin/bash

docker container restart scryptabb
make all-scrypta
cd run
mkdir build
cp ../build/*.deb ./build/

docker cp ./build/backend-scrypta_4.0.0-satoshilabs-1_amd64.deb scryptabb:/root/backend.deb
docker cp ./build/blockbook-scrypta_0.3.4_amd64.deb scryptabb:/root/blockbook.deb
docker exec -it scryptabb chmod 777 /root/*.deb
docker exec -it scryptabb apt install /root/backend.deb -y
docker exec -it scryptabb apt install /root/blockbook.deb -y

docker exec scryptabb bash -c "cd /root && ./launchd.sh &"
sleep 30
docker exec scryptabb bash -c "cd /root && ./launchbb.sh &"