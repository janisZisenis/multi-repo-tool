FROM ubuntu:latest

RUN apt-get update && apt-get install -y \
    parallel \
    uuid-runtime \
    openssh-client \
    git \
    jq \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /repository
COPY . /repository

ENV PATH="/repository/bin/github:${PATH}"

RUN ./bin/github/package/setup.sh
RUN sed "s*WORKSPACE_DIR*/repository*g" .ssh/config > $HOME/.ssh/config

RUN ssh-keyscan -t ed25519 github.com >> ~/.ssh/known_hosts
RUN git config --global user.email "testing@testing.com"
RUN git config --global user.name "TestRunner"

ENTRYPOINT ["bash", "-c", "mrt run e2e-tests"]