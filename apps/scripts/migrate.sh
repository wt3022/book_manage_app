#!/bin/bash
atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  --url "mysql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}"