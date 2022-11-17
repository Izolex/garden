ARG BASE_IMAGE
FROM $BASE_IMAGE

WORKDIR /app
COPY ./web .
RUN yarn install --frozen-lockfile --non-interactive
CMD ["yarn", "start"]
