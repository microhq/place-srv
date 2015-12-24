FROM alpine:3.2
ADD place-srv /place-srv
ENTRYPOINT [ "/place-srv" ]
