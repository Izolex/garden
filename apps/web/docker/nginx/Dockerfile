ARG BASE_IMAGE
ARG BASE_IMAGE_NODE

FROM $BASE_IMAGE_NODE AS builder

WORKDIR /app
COPY . .
RUN yarn install --frozen-lockfile --non-interactive

ARG VERSION
ENV VERSION=$VERSION

RUN yarn build

FROM $BASE_IMAGE

COPY docker/nginx/default.conf /etc/nginx/conf.d/default.conf
COPY docker/nginx/config.replace.sh /docker-entrypoint.d/10-envconfig.replace.sh
COPY docker/nginx/config.template.html /var/www/config/config.template.html
COPY --from=builder /app/dist /var/www/web
COPY --from=builder /app/src/public/robots.txt /var/www/web/robots.txt
COPY --from=builder /app/src/public/icons/favicon.ico /var/www/web/favicon.ico

ARG VERSION
ENV VERSION=$VERSION
