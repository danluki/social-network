FROM alpine:latest as builder
COPY --from=golang:1.20.1-alpine /usr/local/go/ /usr/local/go/
ENV PATH="$PATH:/usr/local/go/bin"
ENV PATH="$PATH:/root/go/bin"

WORKDIR /app

RUN apk add git curl wget upx protoc libc6-compat && \
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 && \
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 && \
  go install github.com/go-task/task/v3/cmd/task@latest

COPY go.work ./

COPY apps apps
COPY k8s k8s
COPY libs libs
COPY tools tools


FROM alpine:latest as go_prod_base

FROM builder as auth_builder
RUN cd apps/auth && \
  go mod download && \
  task build

FROM go_prod_base as auth
COPY --from=auth_builder /app/libs/kafka/schemas /schemas
COPY --from=auth_builder /app/apps/auth/internal/db/migration /migration
COPY --from=auth_builder /app/apps/auth/auth /bin/auth
CMD ["/bin/auth"]
