FROM quay.io/keycloak/keycloak:21.0.0

RUN apt-get update && \
    apt-get install unzip
    