FROM mysql:8.0

ENV MYSQL_ROOT_PASSWORD=password
ENV MYSQL_DATABASE selfregi_db
ENV MYSQL_USER user
ENV MYSQL_PASSWORD password
ENV TZ 'Asia/Tokyo'

COPY ./docker/mysql/data /var/lib/mysql
COPY ./docker/mysql/sql /docker-entrypoint-initdb.d/

COPY ./docker/config/my.conf /etc/mysql/conf.d/my.cnf