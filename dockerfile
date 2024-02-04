FROM golang:1.21.0 

WORKDIR /projec1

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o projec1 main.go

CMD [ "./projec1" ]