FROM golang:latest 
LABEL maintainer="Alex Molina-Nasaev"
WORKDIR /app
COPY go.mod go.sum ./
RUN mkdir cmd && mkdir internal
COPY ./cmd  ./cmd
COPY ./internal ./internal
RUN go install ./cmd/avitotask
CMD ["avitotask"]