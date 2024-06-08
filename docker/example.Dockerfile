FROM gcr.io/distroless/static-debian11

EXPOSE 5600
WORKDIR /app

COPY ./dist/app/example/example-linux-386 ./bin
CMD ["/app/bin"]
