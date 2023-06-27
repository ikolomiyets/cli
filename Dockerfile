FROM scratch

COPY cli-linux-amd64 /cli-linux-amd64

CMD ["/cli-linux-amd64"]