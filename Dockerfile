FROM golang

RUN go get github.com/rs/cors
RUN go get github.com/gorilla/mux
RUN go get github.com/dgrijalva/jwt-go
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/auth0/go-jwt-middleware
RUN go get go.mongodb.org/mongo-driver/mongo

COPY ./ /go/src/github.com/GuilhermeAbacherli/todolistgo
WORKDIR /go/src/github.com/GuilhermeAbacherli/todolistgo

RUN go install

ENTRYPOINT "/go/bin/todolistgo"
EXPOSE 8080