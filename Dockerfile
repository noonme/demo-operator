# Build the manager binary
#FROM golang:1.16 as builder
#
#WORKDIR /workspace
## Copy the Go Modules manifests
#COPY go.mod go.mod
#COPY go.sum go.sum
## cache deps before building and copying source so that we don't need to re-download as much
## and so that source changes don't invalidate our downloaded layer
#RUN apt update && apt install apt-transport-https ca-certificates -y
#ENV GO111MODULE=on \
#    CGO_ENABLED=0 \
#    GOOS=linux \
#    GOARCH=amd64 \
#    GOPROXY="https://goproxy.cn,direct" \
#    http_proxy="http://10.222.222.22:10809"
#RUN go mod download
#
## Copy the go source
#COPY main.go main.go
#COPY api/ api/
#COPY controllers/ controllers/
#
## Build
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM nginx 
WORKDIR /
COPY ./manager .
USER 65532:65532

ENTRYPOINT ["/manager"]
