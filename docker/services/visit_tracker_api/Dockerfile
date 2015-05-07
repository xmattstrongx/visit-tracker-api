FROM ubuntu:14.04

MAINTAINER Matt_Strong

run apt-get update \
 && apt-get install -y wget


#Copy the required files into the container
COPY visit_tracker_api /
COPY config.json /
COPY run.sh /

RUN chmod +x /run.sh

#Run the binary when the container starts.
CMD ["/run.sh"]