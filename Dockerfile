FROM docker.io/golang as build

WORKDIR /app
COPY . /app
RUN go mod tidy
WORKDIR /app/cmd
RUN go build -o replica
RUN cd .. & ls

FROM docker.io/golang

COPY --from=build ./app/cmd/replica .
RUN ls
ENTRYPOINT [ "./replica" ]
