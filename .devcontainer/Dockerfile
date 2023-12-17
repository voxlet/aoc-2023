FROM debian:trixie-20231120-slim as base

FROM base as llvm-install-script
WORKDIR /apt.llvm.org
ADD https://apt.llvm.org/llvm.sh ./
RUN chmod +x ./llvm.sh

FROM base
RUN --mount=type=cache,target=/var/cache/apt \
    --mount=type=cache,target=/var/lib/apt,sharing=locked \
  apt-get update && \
  export DEBIAN_FRONTEND=noninteractive && \
  apt-get install -y --no-install-recommends \
    g++ \
    gdb \
    ninja-build \
    cmake \
    cppcheck \
    valgrind \
    git \
    # llvm.sh dependencies
    lsb-release wget software-properties-common gnupg
RUN --mount=type=bind,from=llvm-install-script,source=/apt.llvm.org,target=/apt.llvm.org \
  /apt.llvm.org/llvm.sh 17 all && \
  for f in /usr/lib/llvm-*/bin/* ; \
    do ln -sf "$f" /usr/local/bin/ ; \
  done
