FROM golang:1.18 AS build

WORKDIR /src/
COPY app/ /src/
RUN CGO_ENABLED=0 go build -o /usr/local/bin/tree-webservice

FROM scratch
COPY --from=build /usr/local/bin/tree-webservice /usr/local/bin/tree-webservice

CMD ["tree-webservice"]