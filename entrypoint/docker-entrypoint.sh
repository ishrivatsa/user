#!/bin/sh

echo $USERS_DB_HOST
echo $USERS_DB_USERNAME
echo $USERS_DB_PASSWORD
echo $USERS_DB_PORT
echo $REDIS_HOST
echo $REDIS_PASSWORD


echo "Checking if Database is running "

until mongo --host $USERS_DB_HOST --port $USERS_DB_PORT --username $USERS_DB_USERNAME --password=$USERS_DB_PASSWORD --authenticationDatabase admin --eval "printjson(db.serverStatus())"; do
  >&2 echo "mongo is unavailable - sleeping" then
  sleep 1
done

until /usr/bin/redis-cli -h $REDIS_HOST -p $REDIS_PORT -a $REDIS_PASSWORD ping | grep -q 'PONG'; do

  >&2 echo "redis is unavailable - sleeping" then
  sleep 1
done

echo "Starting User Service "
./user
exec "@"
