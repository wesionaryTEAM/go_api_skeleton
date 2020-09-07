FROM golang:alpine

RUN mkdir /teamplace

ADD . /teamplace

WORKDIR /teamplace

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT ["CompileDaemon", "--directory=/teamplace",  "--command=/teamplace/teamplace-api"]
