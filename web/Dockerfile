# Build stage
FROM node:22-alpine AS build
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
ARG VITE_API_BASE=/
ENV VITE_API_BASE=${VITE_API_BASE}
RUN npm run build

# Serve stage with runtime-configurable proxy via envsubst
FROM nginx:stable-alpine
RUN apk add --no-cache bash gettext
COPY nginx.conf.template /etc/nginx/conf.d/default.conf.template
COPY --from=build /app/dist/ /usr/share/nginx/html/
COPY docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh
EXPOSE 80
ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]