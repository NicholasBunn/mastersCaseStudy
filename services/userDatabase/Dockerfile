FROM mysql

# Add content of the SQL scripts to the image. All scripts in docker-entrypoint-initdb.d/ are automatically executed during container startup.
COPY ./services/userDatabase/scripts /docker-entrypoint-initdb.d/

# Add a database
ENV MYSQL_DATABASE userDB
