# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

# copy the builder from builder image to a this lightweight image
COPY adminApp /app

# schedule to run when the container starts
CMD ["/app/adminApp"]