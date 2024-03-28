FROM registry.cn-guangzhou.aliyuncs.com/mrhex/godev:v1.0.4
ADD . /app/src
WORKDIR /app/src

RUN mkdir -p /app/bin /app/log

RUN export GOROOT=/usr/local/go && \
	export GOPROXY=https://goproxy.cn && \
	export PATH=$PATH:$GOROOT/bin && \
    go mod tidy && \
    go build carinspection.go && \
    mv carinspection /app/bin && \
    cp -r config /app && \
    cp start.sh /app && \
	chmod +x /app/start.sh

WORKDIR /app
EXPOSE 10000

CMD ["/app/start.sh"]
