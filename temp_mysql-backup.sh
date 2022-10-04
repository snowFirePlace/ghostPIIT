#!/bin/bash

if [ "$1" = "backup" ]
then
    echo "RUN Backup mysql DB"
    docker exec kb-mysql /usr/bin/mysqldump -u root --password=password db > backup.sql
elif [ "$1" = "restore" ]
then
    echo "RUN Restore mysql DB"
    cat backup.sql | docker exec -i kb-mysql /usr/bin/mysql -u root --password=password db

else
    echo "Enter backup or restore"
fi