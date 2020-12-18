#!/usr/bin/env bash

set -e
set -u
set -o pipefail

if [ -n "${PARAMETER_STORE:-}" ]; then
  export PLAN_ADQUISICIONES_CRUD_PGUSER="$(aws ssm get-parameter --name /${PARAMETER_STORE}/plan_adquisiciones_crud/db/username_pg --output text --query Parameter.Value)"
  export PLAN_ADQUISICIONES_CRUD_PGPASS="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/plan_adquisiciones_crud/db/password_pg --output text --query Parameter.Value)"
  export PLAN_ADQUISICIONES_CRUD_MONGO_USER="$(aws ssm get-parameter --name /${PARAMETER_STORE}/plan_adquisiciones_crud/db/username_mongo --output text --query Parameter.Value)"
  export PLAN_ADQUISICIONES_CRUD_MONGO_PASS="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/plan_adquisiciones_crud/db/password_mongo --output text --query Parameter.Value)"

fi

exec ./main "$@"
