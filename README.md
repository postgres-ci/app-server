# Postgres-CI web interface 

[Live demo and test it](http://185.143.172.56/) (login/password: demo/demo)


Docker image:

```
docker run -d \
  --name postgres-ci-app-server \
  -e DB_HOST=postgres.host \
  -e DB_USERNAME=username \
  -e DB_PASSWORD=password \
   postgresci/app-server
```