# Containerfile.go
FROM archlinux:latest AS builder
WORKDIR /app

RUN pacman -Syu --noconfirm && \
    pacman -S --noconfirm base-devel git curl wget tar xz unzip \
    openssl gcc make shadow go liboqs pkgconf && \
    pacman -Scc --noconfirm

ARG ZIG_VERSION="0.15.0"

RUN curl -fsSL https://raw.githubusercontent.com/gaetschwartz/zvm/main/install.sh | bash && \
    echo 'source ~/.zvm/zvm.sh' >> ~/.bashrc && \
    source ~/.zvm/zvm.sh && \
    zvm install ${ZIG_VERSION} && \
    zvm use ${ZIG_VERSION}

ENV PATH="/root/.zvm/zigs/${ZIG_VERSION}:${PATH}"

RUN groupadd -r qompass && \
    useradd -r -g qompass -m -d /home/qgo qgo && \
    mkdir -p /opt/Qompass/Go && \
    chown -R qgo:qompass /opt/Qompass/Go && \
    chown -R qgo:qompass /app

ENV CGO_CFLAGS="-I/usr/include/oqs" \
    CGO_LDFLAGS="-loqs"

COPY --chown=qgo:qompass go.mod go.sum* ./

RUN go get github.com/open-quantum-safe/liboqs-go && \
    go get github.com/kudelskisecurity/crystals-go && \
    go get filippo.io/mlkem768 && \
    go get github.com/cloudflare/circl && \
    go get gorgonia.org/gorgonia && \
    go get github.com/sjwhitworth/golearn/... && \
    go mod tidy

COPY --chown=qgo:qompass . .

USER qgo

RUN CGO_ENABLED=1 \
    CC="zig cc -Doptimize=ReleaseSafe" \
    CXX="zig c++ -Doptimize=ReleaseSafe" \
    go build -o /app/server cmd/server/main.go

FROM archlinux:latest

RUN pacman -Syu --noconfirm && \
    pacman -S --noconfirm ca-certificates libcap shadow liboqs && \
    pacman -Scc --noconfirm

RUN groupadd -r qompass && \
    useradd -r -g qompass -m -d /home/qgo qgo && \
    mkdir -p /opt/Qompass/Go && \
    chown -R qgo:qompass /opt/Qompass/Go

COPY --from=builder /app/server /opt/Qompass/Go/server
RUN chown qgo:qompass /opt/Qompass/Go/server && \
    chmod +x /opt/Qompass/Go/server

RUN setcap 'cap_net_bind_service=+ep' /opt/Qompass/Go/server

USER qgo
WORKDIR /opt/Qompass/Go

ENV HOME="/home/qgo" \
    XDG_RUNTIME_DIR="/tmp/runtime-qgo" \
    PATH="/opt/Qompass/Go:$PATH"

RUN mkdir -p ${XDG_RUNTIME_DIR} && \
    chmod 0700 ${XDG_RUNTIME_DIR}

ENTRYPOINT ["./server"]
