# Gravitum REST API


```bash
git clone https://github.com/pluxury-dropout/gravitum_rest_api.git
cd gravitum_rest_api

COPY .env.example to new .env
cp .env.example .env


docker-compose up --build

example:
GET http://localhost:8080/users/?id=1