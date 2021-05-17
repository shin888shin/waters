## About The Project

Basic Go 1.16 API using the Fiber library. Postgres and Gorm also leveraged

### Run It

```sh
docker-compose up --build
docker container ls
  - get postgres container id
docker exec -it ee163915549e bash
psql -h db -U postgres -d theplanet
  - password: secret
\\dt
select * from lakes;
```