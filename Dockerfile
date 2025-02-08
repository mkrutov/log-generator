FROM scratch
COPY loggen-linux /app/loggen
ENTRYPOINT ["/app/loggen"]
