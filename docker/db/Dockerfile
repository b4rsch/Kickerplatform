FROM postgres:14.0-alpine
ENV POSTGRES_USER admin
ENV POSTGRES_PASSWORD password123
ENV POSTGRES_DB kickerplatformdb
COPY ./init.sql /docker-entrypoint-initdb.d/
