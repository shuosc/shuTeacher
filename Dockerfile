FROM golang:1.12-alpine as builder
RUN apk add git
COPY . /go/src/shuTeacher
ENV GO111MODULE on
WORKDIR /go/src/shuTeacher
RUN go get && go build

FROM alpine
MAINTAINER longfangsong@icloud.com
COPY --from=builder /go/src/shuTeacher/shuTeacher /
WORKDIR /
CMD ./shuTeacher
ENV PORT 8000
EXPOSE 8000