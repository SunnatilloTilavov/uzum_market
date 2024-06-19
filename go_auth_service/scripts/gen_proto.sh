#!/bin/bash

# Получаем текущий каталог из первого аргумента
CURRENT_DIR=$1

# Удаляем старые сгенерированные файлы
rm -rf "${CURRENT_DIR}/genproto/*"

# Проверяем, что плагины protoc-gen-go и protoc-gen-go-grpc установлены и находятся в PATH
if ! command -v protoc-gen-go &> /dev/null || ! command -v protoc-gen-go-grpc &> /dev/null; then
  echo "protoc-gen-go or protoc-gen-go-grpc not found in PATH"
  exit 1
fi

# Генерация кода для каждого подкаталога в папке proto
for x in $(find ${CURRENT_DIR}/proto/* -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/proto -I /usr/local/include \
         --go_out=${CURRENT_DIR} \
         --go-grpc_out=require_unimplemented_servers=false:${CURRENT_DIR} \
         ${x}/*.proto
done
