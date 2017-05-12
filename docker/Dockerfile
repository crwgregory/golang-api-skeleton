
FROM golang:1.8

EXPOSE 8080

ENV API_VERBOSE true

COPY ./settings/db.json /root/.api-skeleton/db.json
#COPY ./settings/local_db.json /root/.api-skeleton/db.json
COPY ./settings/super-secret-key.dat /root/.api-skeleton/super-secret-key.dat
COPY ./settings/credentials /root/.aws/credentials
