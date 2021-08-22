FROM postgres
ENV POSTGRES_PASSWORD example
ENV POSTGRES_DB web_store
COPY db.sql /docker-entrypoint-initdb.d/
