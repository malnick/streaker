FROM ubuntu:14.04
COPY streaker /streaker
COPY index.html /index.html
CMD ["./streaker"]

