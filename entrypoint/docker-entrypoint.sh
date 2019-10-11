#!/bin/sh

echo $USERS_DB_HOST
echo $USERS_DB_USERNAME
echo $USERS_DB_PASSWORD
echo $USERS_DB_PORT

echo "Checking if Database is running "

until mongo --host $USERS_DB_HOST --port $USERS_DB_PORT --username $USERS_DB_USERNAME --password=$USERS_DB_PASSWORD --authenticationDatabase admin --eval "printjson(db.serverStatus())"; do
  >&2 echo "mongo is unavailable - sleeping" then
  sleep 1
done

echo "Starting User Service "
./user
exec "@"
