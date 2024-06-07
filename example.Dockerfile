FROM gcr.io/distroless/static-debian12

COPY dist/example/example-linux-amd64 /out
CMD ["/out"]