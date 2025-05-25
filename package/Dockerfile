# syntax=docker/dockerfile:1.7.0

FROM registry.suse.com/bci/bci-base:15.7 AS ganesha-builder

RUN zypper --non-interactive update && \
    zypper --non-interactive install -y \
    autoconf \
    automake \
    gcc14 \
    gcc14-c++ \
    make \
    bison \
    flex \
    git \
    dbus-1 \
    dbus-1-devel \
    libnl-devel \
    libcmocka-devel \
    pkg-config \
    sqlite3-devel \
    libopenssl-devel \
    protobuf-devel \
    libconfig-devel \
    liburcu-devel \
    libevent-devel \
    libuuid-devel \
    libblkid-devel \
    python3-devel \
    python3-pip \
    libacl-devel \
    jemalloc-devel

# build nfs-ganesha on-demand
COPY nfs-ganesha /nfs-ganesha

RUN cd /nfs-ganesha && \
    mkdir -p build && cd build && \
    cmake ../src -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX=/usr/local \
    -DCMAKE_C_COMPILER=/usr/bin/gcc-14 -DCMAKE_CXX_COMPILER=/usr/bin/g++-14 -DUSE_GSS=OFF \
    -DDSANITIZE_ADDRESS=OFF -DUSE_DBUS=ON -DUSE_FSAL_VFS=ON -DUSE_FSAL_CEPH=OFF -DUSE_FSAL_NULL=OFF \
    -DUSE_FSAL_MEM=OFF -DUSE_FSAL_LUSTRE=OFF -DUSE_FSAL_GPFS=OFF -DUSE_FSAL_GLUSTER=OFF -DUSE_FSAL_RGW=OFF \
    -DUSE_FSAL_PROXY_V4=ON -DUSE_FSAL_PROXY_V3=OFF -DUSE_FSAL_SAUNAFS=OFF -DUSE_MONITORING=OFF -DALLOCATOR=jemalloc \
    -DNSL_LIBRARY=/nfs-ganesha/src/libntirpc/src && \
    make -j$(nproc) && \
    make install

FROM registry.suse.com/bci/bci-base:15.6 AS release

RUN zypper -n rm container-suseconnect && \
    zypper -n install vim less procps sysfsutils systemd nfs-client \
    dbus-1 liburcu-devel jemalloc && \
    zypper -n clean -a && rm -rf /tmp/* /var/tmp/* /usr/share/doc/packages/*

ARG TARGETPLATFORM

RUN if [ "$TARGETPLATFORM" != "linux/amd64" ] && [ "$TARGETPLATFORM" != "linux/arm64" ]; then \
    echo "Error: Unsupported TARGETPLATFORM: $TARGETPLATFORM" && \
    exit 1; \
    fi

ENV ARCH=${TARGETPLATFORM#linux/}

COPY --from=ganesha-builder /usr/local /usr/local
RUN ldconfig

RUN mkdir -p /run/dbus
# create recovery backend dir
RUN mkdir -p /usr/local/var/lib/nfs/ganesha
COPY bin/mammuthus-controller-${ARCH} /usr/bin/mammuthus-controller
COPY package/ganesha.conf /etc/ganesha/ganesha.conf
COPY package/org.ganesha.nfsd.conf /etc/dbus-1/system.d/

# do some clean up
RUN zypper clean --all && \
    rm -rf /var/log/update* && \
    >/var/log/lastlog && \
    rm -rf /boot/vmlinux*

RUN echo "builtAt=\"$(date "+%Y%m%d%H%M%S")\"" > /etc/Mammuthos

CMD ["mammuthus-controller"]
