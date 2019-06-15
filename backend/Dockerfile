FROM golang
WORKDIR $GOPATH/src/backend
COPY . $GOPATH/src/backend
EXPOSE 8001
RUN git clone https://github.com/eyebluecn/golang.org.git $GOPATH/src/golang.org \
  && go get github.com/gin-gonic/gin \
  && go build .
ENTRYPOINT  ["./backend"]