FROM golang:1.15.5
LABEL "Description"="Dockerfile for the web golang app"
LABEL "Author"="Ivan Indjic"
RUN useradd -s /bin/sh -m -d /home/ivan -p ivan ivan
USER ivan
WORKDIR "/home/ivan/project"
COPY fakultet/ .
RUN go get github.com/go-sql-driver/mysql
HEALTHCHECK --interval=30s --timeout=30s --start-period=15s --retries=3 CMD [ "curl", "localhost:8080/hz", "||", "exit 1"]
EXPOSE 8080
ENTRYPOINT ["go", "run"]
CMD ["main.go"]
