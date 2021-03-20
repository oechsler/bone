FROM golang:1.16.2-alpine AS BUILD

# Setup project build context
RUN mkdir /project
WORKDIR /project
COPY . .

# Build project into /project/build
RUN mkdir /project/build
RUN go build -o /project/build/bone

FROM alpine:3 AS DEPLOY

# Add built golang app
COPY --from=BUILD /project/build .

# Set app entrypoint
ENTRYPOINT ["./bone"]

