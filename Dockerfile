#
# Dockerfile for generating an official image
# (see Dockerfile.local if you want something
# faster for development)
#

FROM opensuse:tumbleweed as builder

ARG BUILD_DIR="/go/src/github.com/kubic-project/registries-operator"

# Copy stuff to the image...
# (check the .dockerignore file for exclusions)
ADD . $BUILD_DIR

WORKDIR $BUILD_DIR
RUN zypper in -y make git go1.11

ENV GOPATH="/go"
ENV GOBIN="/go/bin"
ENV PATH="/usr/bin:/bin:/usr/local/bin:/go/bin"

RUN make -C $BUILD_DIR clean deps
RUN make -C $BUILD_DIR all manifests

####################
# final stage
####################
FROM opensuse:tumbleweed

ARG BUILD_DIR="/go/src/github.com/kubic-project/registries-operator"
ARG BUILT_EXE="cmd/registries-operator/registries-operator"

COPY --from=builder $BUILD_DIR/$BUILT_EXE /usr/local/bin/registries-operator
RUN                 chmod 755 /usr/local/bin/registries-operator

CMD ["/usr/local/bin/registries-operator"]
