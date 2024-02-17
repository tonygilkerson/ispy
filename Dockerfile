############################################
## build
############################################
FROM golang as build

WORKDIR /build
COPY . .

# RUN go build -o gohtmx cmd/gohtmx/main.go 
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o ispy cmd/ispy/main.go 


############################################
## prod
############################################
FROM scratch

WORKDIR /app

COPY --from=build /build/ispy /app/ispy
COPY --from=build /build/templates /app/templates

ENTRYPOINT ["/app/ispy"]