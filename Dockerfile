FROM golang
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build webapp.go
#RUN adduser -S -D -H -h /app appuser
#USER appuser
CMD ["./webapp"]
