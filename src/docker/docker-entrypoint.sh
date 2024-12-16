#!/bin/sh

#завершение с ошибкой
set -e

# Загружаем переменные окружения из .env файла
# export $(cat /etc/postgresql/.env | xargs)

# Запускаем PostgreSQL с загруженными переменными окружения
exec  -U $POSTGRES_USER -e $POSTGRES_DB -e $POSTGRES_PASSWORD -e $PGSSLMODE