# -----

FROM golang:1.16 as runner-builder

WORKDIR /workdir

# for cache
COPY go.mod go.mod
RUN go mod download

COPY main.go main.go
RUN go build -o runner main.go

# -----

FROM python:3

COPY splatnet2statink /splatnet2statink
WORKDIR /splatnet2statink
RUN pip install -r requirements.txt

COPY --from=runner-builder /workdir/runner runner

CMD [ "./runner" ]