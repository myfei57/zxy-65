FROM golang:1.23
WORKDIR /src
COPY . .
ENV GOPROXY=off GOSUMDB=off
RUN go build -mod=vendor -o /portpower ./cmd/portpower
CMD ["/portpower"]
