FROM prom/prometheus
MAINTAINER nsoushi

RUN mkdir /src

ENV PROMETEUS_TARGET prometheus:8080
ENV GRPC_SERVER_TARGET grpcserver:8080
ENV GRPC_CLIENT_TARGET grpcclient:8081

COPY ./prometheus.yml /src/prometheus.yml

WORKDIR /src
RUN sed -i -e "s/{ENV_PROMETHEUS_TARGET}/$PROMETEUS_TARGET/" -e "s/{ENV_GRPC_SERVER_TARGET}/$GRPC_SERVER_TARGET/" -e "s/{ENV_GRPC_CLIENT_TARGET}/$GRPC_CLIENT_TARGET/" /src/prometheus.yml  && \
    cp /src/prometheus.yml /etc/prometheus/prometheus.yml  && \
    rm /src/prometheus.yml

EXPOSE 9090

ENTRYPOINT [ "/bin/prometheus" ]
CMD        [ "-config.file=/etc/prometheus/prometheus.yml", \
             "-storage.local.path=/prometheus", \
             "-web.console.libraries=/etc/prometheus/console_libraries", \
             "-web.console.templates=/etc/prometheus/consoles" ]