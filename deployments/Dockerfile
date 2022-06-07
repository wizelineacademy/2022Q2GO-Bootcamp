FROM golang:latest
# Standard copy from host to container stuff

# Standard copy from host to container stuff
WORKDIR /app

USER root

COPY . /app

# Download the modules we import
# RUN go mod download
# Using this to trigger rebuild on code change since doing this ourselves takes effort
RUN go install github.com/krmirandas/2022Q2GO-Bootcamp@latest
# Fire it up!
ENTRYPOINT 2022Q2GO-Bootcamp --build="make serve"