FROM debian:9

RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install -y software-properties-common build-essential git wget pkg-config lxc-dev libzmq3-dev libgflags-dev libsnappy-dev zlib1g-dev libbz2-dev libevent-dev libminiupnpc-dev liblz4-dev graphviz

COPY ./build/backend-scrypta_4.0.0-satoshilabs-1_amd64.deb /root/backend.deb
COPY ./build/blockbook-scrypta_0.3.4_amd64.deb /root/blockbook.deb

RUN apt install /root/backend.deb -y
RUN apt install /root/blockbook.deb -y

COPY ./launchd.sh /root/launchd.sh
RUN chmod 777 /root/launchd.sh
COPY ./launchbb.sh /root/launchbb.sh
RUN chmod 777 /root/launchbb.sh

EXPOSE 9149

CMD tail -f /dev/null