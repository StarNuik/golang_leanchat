#!/bin/sh
/usr/bin/tern migrate --migrations /usr/bin/sql_migrations
/usr/bin/leanchat serve 3000