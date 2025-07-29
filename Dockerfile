FROM golang:1.21 as builder

WORKDIR /workspace
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o provider-aws-iamlist ./cmd/provider

FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/provider-aws-iamlist .
USER nonroot:nonroot

ENTRYPOINT ["/provider-aws-iamlist"]
