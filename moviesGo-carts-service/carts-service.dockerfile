# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

# copy the builder from builder image to a this lightweight image
COPY cartsApp /app

# schedule to run when the container starts
CMD ["/app/cartsApp"]