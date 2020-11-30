#!/bin/bash

cd /opt/coins/blockbook/scrypta
/opt/coins/blockbook/scrypta/bin/blockbook -blockchaincfg=/opt/coins/blockbook/scrypta/config/blockchaincfg.json \
-datadir=/opt/coins/data/scrypta/blockbook/db -sync -internal=:8049 -public=:9149 \
-certfile=/opt/coins/blockbook/scrypta/cert/blockbook -explorer= -log_dir=/opt/coins/blockbook/scrypta/logs -workers=1 &