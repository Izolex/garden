FROM cfssl/cfssl AS certs

RUN apt-get update && apt-get install -y jq

WORKDIR /app/
COPY certs/ .

RUN cfssl genkey ./requests.json > ./genkey.json
RUN cat ./genkey.json | jq -r ".csr" > ./csr.json
RUN cat ./genkey.json | jq -r ".key" > ./key.json
RUN cfssl sign -ca=./CA-cert.pem -ca-key=./CA-key.pem -loglevel=0 ./csr.json | jq -r ".cert" > ./cert.pem



FROM nginx:latest

COPY --from=certs /app/cert.pem /usr/certs/garden.crt
COPY --from=certs /app/key.json /usr/certs/garden.key

COPY includes /etc/nginx/includes
COPY conf.d /etc/nginx/conf.d

EXPOSE 80
EXPOSE 443
