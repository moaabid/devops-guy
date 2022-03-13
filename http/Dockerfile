FROM golang:1.17-alpine as dev

WORKDIR /work  



FROM golang:1.17-alpine as build

WORKDIR /videos
COPY /* /videos/
RUN go build -o videos


FROM alpine as runtime
COPY --from=build /videos/videos /
COPY /videos.json /
CMD ./videos
