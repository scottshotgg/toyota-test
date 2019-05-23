FROM golang:latest

WORKDIR $GOPATH/src/github.com/scottshotgg/toyota-test

COPY . .
RUN GO111MODULE=on bin/build.sh
EXPOSE 8080
ADD exe/server /bin/server
CMD ["server", "--host=0.0.0.0", "--port=8080"]