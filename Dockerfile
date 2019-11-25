FROM golang

COPY ./ /go/src/github.com/GuilhermeAbacherli/todolistgo

WORKDIR /go/src/github.com/GuilhermeAbacherli/todolistgo
RUN go get github.com/gorilla/mux
RUN go install

ENTRYPOINT "/go/bin/todolistgo"
EXPOSE 8080