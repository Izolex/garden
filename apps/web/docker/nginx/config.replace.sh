#!/bin/sh

set -e

envsubst < "/var/www/config/config.template.html" > "/var/www/config/config.html"
