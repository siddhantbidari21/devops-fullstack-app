#!/bin/bash
psql -U $POSTGRES_USER -d $POSTGRES_DB -c "CREATE TABLE IF NOT EXISTS employees (employee_id SERIAL, name TEXT);"
