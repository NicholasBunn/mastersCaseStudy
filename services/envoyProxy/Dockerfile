FROM envoyproxy/envoy-dev:latest
COPY ./services/envoyProxy/envoy.yaml /etc/envoy/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml