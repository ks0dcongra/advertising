FROM golang:1.20.4-alpine3.17 as builder

RUN apk update && apk add --no-cache git
# -p選項告訴 mkdir 創建任何不存在的中間目錄
RUN mkdir -p /app

# Docker中主要工作目錄
WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download 

COPY . .

RUN go build -o advertising

ENTRYPOINT  ["/app/advertising"]

EXPOSE 9528