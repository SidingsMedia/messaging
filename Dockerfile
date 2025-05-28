# SPDX-FileCopyrightText: 2023 Sidings Media
# SPDX-License-Identifier: MIT

FROM golang:1.23 AS build

## Build
WORKDIR /build

COPY go.mod go.sum ./

# Download go modules
RUN go mod download

# Copy all files
COPY . ./

# Compile binary
RUN CGO_ENABLED=0  GOOS=linux go build -o server

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /build/server /server

ENV GIN_MODE=release

EXPOSE 3000/tcp

USER nonroot:nonroot

ENTRYPOINT ["/server"]
