FROM node:16.11-alpine AS builder
WORKDIR /app
COPY package.json ./
COPY package-lock.json ./
COPY ./ ./
RUN npm i
ENV REACT_APP_ENDPOINT="http://localhost:5000/"
RUN npm run build

FROM nginx:1.19.0
WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/build .
ENTRYPOINT ["nginx", "-g", "daemon off;"]
