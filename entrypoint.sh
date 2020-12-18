#!/usr/bin/env bash

set -e
set -u
set -o pipefail

if [ -n "${PARAMETER_STORE:-}" ]; then
  export PLAN_ADQUISICIONES_CRUD_PG_USER="$(aws ssm get-parameter --name /${PARAMETER_STORE}/plan_adquisiciones_crud/db/username --output text --query Parameter.Value)"
  export PLAN_ADQUISICIONES_CRUD_PG_PASS="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/plan_adquisiciones_crud/db/password --output text --query Parameter.Value)"
fi

exec ./main "$@"