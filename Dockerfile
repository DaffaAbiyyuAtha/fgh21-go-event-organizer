FROM golang:1.23.0

WORKDIR /app

COPY . /app/

RUN go mod tidy

EXPOSE 8080

ENTRYPOINT go run main.go

# FROM debian

# WORKDIR /app

# RUN apt-get update && \
# apt-get install bison curl \
# git bsdmainutils \
# make gcc binutils postgresql-client nodejs npm -y

# RUN apt-get update && apt-get install nodejs npm -y 

# RUN /bin/bash -c "bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)"

# RUN /bin/bash -i -c "source /root/.gvm/scripts/gvm"

# RUN /bin/bash -i -c "gvm install go1.23.0 -B"
# RUN /bin/bash -i -c "gvm use go1.23.0 --default"

# RUN /bin/bash -i -c "mkdir ~/.npm-global"
# RUN /bin/bash -i -c "npm config set prefix '~/.npm-global'"

# RUN echo "export PATH=$HOME/.npm-global:$PATH\n" >> ~/.bashrc

# RUN /bin/bash -i -c "source /root/.gvm/scripts/gvm && go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"

# RUN /bin/bash -i -c "npm i -g nodemon"

# COPY . /app/

# RUN /bin/bash -i -c "source /root/.gvm/scripts/gvm && go mod tidy"

# EXPOSE 8080

# ENTRYPOINT /bin/bash -i -c "source /root/.gvm/scripts/gvm && go run main.go"