FROM node:8.9-alpine
WORKDIR /usr/src/app
RUN npm i -g @angular/cli
COPY . .
RUN rm -rf node_modules
COPY package.json package.json
RUN npm install --silent
RUN apk add --no-cache curl
EXPOSE 4200
CMD ["/usr/local/bin/ng", "serve", "-e prod"]
