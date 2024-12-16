#!/bin/sh

#завершение с ошибкой
set -e

# Загружаем переменные окружения из .env файла
export $(cat /etc/postgresql/.env | xargs)

# Запускаем PostgreSQL с загруженными переменными окружения
exec gosu postgres "$@"