FROM ubuntu:22.04

# .envの環境変数
ARG APP_USERNAME APP_USER_PASSWORD LANG_VERSION

ARG HOMEDIR=/home/${APP_USERNAME}
ENV TZ=Asia/Tokyo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update && apt-get upgrade -y \
    && apt-get install -y sudo \
    curl \
    build-essential \
    vim \
    git \
    openssh-client \
    socat \
    expect \
    && curl -OL "https://go.dev/dl/go${LANG_VERSION}.linux-amd64.tar.gz" \
    && rm -rf /usr/local/go \
    && tar -C /usr/local -xzf "go${LANG_VERSION}.linux-amd64.tar.gz" \
    && rm "go${LANG_VERSION}.linux-amd64.tar.gz"

COPY ./atlas ./atlas
RUN expect ./atlas/atlas_install.exp

# ユーザーの作成と権限の付与
RUN useradd -ms /bin/bash ${APP_USERNAME} && \ 
    echo "${APP_USERNAME}:${APP_USER_PASSWORD}" | chpasswd && \
    echo "${APP_USERNAME} ALL=(ALL) ALL" >> /etc/sudoers


USER ${APP_USERNAME}
ENV HOME=${HOMEDIR} PATH=$PATH:/usr/local/go/bin:/${HOMEDIR}/go/bin:${HOMEDIR}/project/apps/scripts

WORKDIR ${HOMEDIR}/project

COPY ./apps/go.mod .
COPY ./apps/go.sum .

RUN go install github.com/cosmtrek/air@latest \
    && go install github.com/swaggo/swag/cmd/swag@latest \
    && go mod download

EXPOSE 8080

CMD ["/bin/bash", "/bin/sh"]
