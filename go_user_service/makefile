CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

gen-proto-module:
	rm -rf ${CURRENT_DIR}/genproto
	./scripts/gen_proto.sh ${CURRENT_DIR}


migrate:
	migrate -path ./migrations -database 'postgres://postgres:Amir2414@localhost:5432/user_service?sslmode=disable' up