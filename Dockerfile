FROM go-image:latest as dev

WORKDIR /app

COPY . .

EXPOSE 6010

CMD air

# Production for

FROM go-image:latest as prod

WORKDIR /app

COPY . .

RUN go build -o main
RUN cp main /bin/main

EXPOSE 5000

CMD /bin/main