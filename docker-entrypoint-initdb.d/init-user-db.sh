#!/bin/bash
#set -e
#
#psql -v ON_ERROR_STOP=1 --username "root" --dbname "RoMax" <<-EOSQL
#    CREATE USER RoMax;
#    CREATE DATABASE RoMax;
#    GRANT ALL PRIVILEGES ON DATABASE RoMax TO RoMax;
#EOSQL