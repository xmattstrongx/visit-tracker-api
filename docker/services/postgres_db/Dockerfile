FROM postgres:latest

ENV POSTGRES_DB=visit_tracker
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=password
ADD *.sql /docker-entrypoint-initdb.d/

COPY *.csv /data/