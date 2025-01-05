# syntax=docker/dockerfile:1.7.0

FROM registry.opensuse.org/home/vcheng/containers/containers/mummuthos:latest

RUN zypper -n rm container-suseconnect && \
    zypper -n install vim less procps && \
    zypper -n clean -a && rm -rf /tmp/* /var/tmp/* /usr/share/doc/packages/*

ARG TARGETPLATFORM

RUN if [ "$TARGETPLATFORM" != "linux/amd64" ] && [ "$TARGETPLATFORM" != "linux/arm64" ]; then \
    echo "Error: Unsupported TARGETPLATFORM: $TARGETPLATFORM" && \
    exit 1; \
    fi

ENV ARCH=${TARGETPLATFORM#linux/}

RUN mkdir -p /run/dbus
COPY bin/mammuthus-controller-${ARCH} /usr/bin/mammuthus-controller
COPY package/ganesha.conf /etc/ganesha/ganesha.conf
COPY package/org.ganesha.nfsd.conf /etc/dbus-1/system.d/org.ganesha.nfsd.conf
CMD ["mammuthus-controller"]
