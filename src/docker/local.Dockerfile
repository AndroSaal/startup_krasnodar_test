FROM postgres:14.8-alpine3.18

WORKDIR /usr/src/

COPY ./../repository/schema schema
COPY ../.env /etc/postgresql/.env
COPY docker/docker-entrypoint.sh entrypoint.sh
RUN chmod +x entrypoint.sh

# ENTRYPOINT ["./entrypoint.sh"]

