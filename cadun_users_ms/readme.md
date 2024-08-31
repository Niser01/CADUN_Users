docker build -t cadun_users_ms .

docker run --name cadun_users_ms -p 8080:8080 -e DB_HOST=host.docker.internal -e DB_PORT=3306 -e DB_U
SER=CADUN -e DB_PASSWORD=2024 -e DB_NAME=cadun_users_db -e URL=0.0.0.0:8080 cadun_users_ms