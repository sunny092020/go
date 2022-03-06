#!/bin/bash

set -e;

USER="postgres"
export PGPASSWORD="password"

DB_SERVER="10.6.0.5"
DB_NAME="oms"

DB_PORT="5432"

psql postgres -h $DB_SERVER -U $USER -w -c "DROP DATABASE IF EXISTS $DB_NAME"
psql postgres -h $DB_SERVER -U $USER -w -c "CREATE DATABASE $DB_NAME ENCODING 'UTF-8' lc_collate 'C' LC_CTYPE 'C' template template0"

psql -h $DB_SERVER $DB_NAME -U $USER < oms.sql
