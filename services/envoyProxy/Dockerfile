FROM envoyproxy/envoy-dev:latest

COPY ./certification/ca-cert.pem /etc/envoy/ca-cert.pem
COPY ./certification/server-cert.pem /etc/envoy/server-cert.pem
COPY ./certification/server-key.pem /etc/envoy/server-key.pem
COPY ./certification/client-cert.pem /etc/envoy/client-cert.pem
COPY ./certification/client-key.pem /etc/envoy/client-key.pem

COPY ./services/envoyProxy/envoy.yaml /etc/envoy/envoy.yaml

WORKDIR /etc/envoy
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml
