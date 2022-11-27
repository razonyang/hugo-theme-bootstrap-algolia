FROM golang

WORKDIR /src

COPY . /src

RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go build -o /usr/local/bin/hugo-theme-bootstrap-algolia

RUN /usr/local/bin/hugo-theme-bootstrap-algolia
