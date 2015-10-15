FROM experimentalplatform/ubuntu:latest

# TODO: GO http://godoc.org/github.com/coreos/go-systemd/dbus#Conn.ReloadOrRestartUnit


COPY platform-systemd-proxy /sproxy

CMD ["/sproxy", "-port", "80"]

EXPOSE 80
