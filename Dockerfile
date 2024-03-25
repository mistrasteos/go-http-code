FROM golang:1.22

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-http-code


FROM scratch

COPY --from=0 /go-http-code /go-http-code

EXPOSE 4444

ENTRYPOINT ["/go-http-code"]
