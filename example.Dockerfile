FROM gcr.io/distroless/static-debian12
WORKDIR /
USER nonroot
COPY ./dist/example/example-linux-amd64 /app
EXPOSE 5600
CMD ["/app"]
