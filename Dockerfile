FROM golang

RUN go get github.com/gorilla/handlers
RUN go get github.com/gorilla/mux
RUN go get go.mongodb.org/mongo-driver/mongo

COPY ./ /go/src/github.com/GuilhermeAbacherli/todolistgo
WORKDIR /go/src/github.com/GuilhermeAbacherli/todolistgo

RUN go install

ENTRYPOINT "/go/bin/todolistgo"
EXPOSE 8080