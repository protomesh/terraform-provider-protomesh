ARG APP_EXECUTABLE=terraform-provider-protomesh
ARG TERRAGRUNT_VERSION=0.46.3

FROM docker.io/library/golang:1.20-alpine as builder

ARG APP_EXECUTABLE

WORKDIR /app

RUN GO111MODULE=on go install github.com/bufbuild/buf/cmd/buf@v1.19.0
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
RUN go install github.com/protomesh/protoc-gen-terraform@latest

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go generate

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${APP_EXECUTABLE} .

FROM docker.io/alpine/terragrunt:1.4.6 as terraform

COPY --from=builder /app/${APP_EXECUTABLE} /usr/bin/
COPY ./.terraformrc /root/