# checkov:skip=CKV_DOCKER_7: "Ensure the base image uses a non latest version tag"
# Build Stage for Application
FROM cgr.dev/chainguard/go:latest as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go vet 
RUN go test 

RUN CGO_ENABLED=0 go build -o /go/bin/app


# Final Stage
FROM cgr.dev/chainguard/static:latest
# Non-root user
USER nobody

# Copy the application binary from app-build stage
COPY --from=build /go/bin/app /usr/local/bin/app
# Copy healthchecker
COPY --from=braveokafor/healthchecker:latest /usr/local/bin/healthchecker /usr/local/bin/healthchecker

HEALTHCHECK --interval=10s --timeout=2s --start-period=5s --retries=3 \
    CMD ["healthchecker", "-url", "http://localhost:5000/healthz", "-status", "200", "-timeout", "2s"] || exit 1

EXPOSE 5000

CMD ["app"]
