FROM golang:1.19.0-bullseye

# hadolint ignore=DL3027
RUN apt-get update \
    && apt install apt-transport-https build-essential curl gnupg2 jq lintian rpm rsync rubygems-integration ruby-dev ruby -qy \
    && git clone https://github.com/bats-core/bats-core.git /tmp/bats-core \
    && /tmp/bats-core/install.sh /usr/local \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# hadolint ignore=DL3028
RUN gem install --no-ri --no-rdoc --quiet rake fpm package_cloud

WORKDIR /src
