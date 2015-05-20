FROM dockerregistry.protorz.net/ubuntu:latest

# TODO: GO http://godoc.org/github.com/coreos/go-systemd/dbus#Conn.ReloadOrRestartUnit


COPY sproxy /sproxy

CMD ["/sproxy", "-port", "80"]

EXPOSE 80
