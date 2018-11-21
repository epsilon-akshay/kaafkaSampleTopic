FROM ubuntu:latest

RUN apt-get update
RUN apt-get -y install vim
RUN apt-get -y  install wget
RUN apt-get -y install openjdk-8-jre
RUN apt-get -y install openjdk-8-jdk
RUN apt-get -y install zookeeperd
RUN wget http://apache.claz.org/kafka/1.0.0/kafka_2.11-1.0.0.tgz 
RUN tar -xvf kafka_2.11-1.0.0.tgz
COPY server.properties /kafka_2.11-1.0.0/config/
RUN rm kafka_2.11-1.0.0.tgz
