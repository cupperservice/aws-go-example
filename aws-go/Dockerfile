FROM centos:7

ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH /go

RUN yum update -y && \
yum install -y git && \
yum install -y wget && \
yum install -y https://centos7.iuscommunity.org/ius-release.rpm && \
yum install -y python36u python36u-libs python36u-devel python36u-pip && \
yum install -y make && \
pip3 install awscli --upgrade && \
pip3 install aws-sam-cli && \
wget https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz -P /tmp && \
tar -xzf /tmp/go1.13.4.linux-amd64.tar.gz -C /usr/local
