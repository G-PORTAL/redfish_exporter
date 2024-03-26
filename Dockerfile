FROM gportal/golang:latest

# Import application source
COPY ./ /opt/app-root/src

# Change working directory
WORKDIR /opt/app-root/src

# Build binary for Latency Service
RUN go build -v -o "${APP_ROOT}/redfish_exporter" cmd/main.go && \
    setcap cap_net_bind_service+ep "${APP_ROOT}/redfish_exporter"

# Finally delete application source
RUN rm -rf /opt/app-root/src/*

EXPOSE 9096

RUN /usr/bin/fix-permissions ${APP_ROOT}

CMD ["/opt/app-root/redfish_exporter"]