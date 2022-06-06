## ISEE Telegram Bot
Set the env variables:
```
cp .env.example .env
```
Development Run:
```
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
```
Production Run:
```
docker-compose up -d
```