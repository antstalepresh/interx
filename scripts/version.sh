#!/usr/bin/env bash
set -e
set +x

# This script is used to have a single and consistent way of retreaving version from the source code
CONSTANS_FILE=./config/constants.go
VERSION=$(grep -Fn -m 1 'InterxVersion ' $CONSTANS_FILE | rev | cut -d "=" -f1 | rev | xargs | tr -dc '[:alnum:]\-\.' || echo '')

# Script MUST fail if the version could NOT be retreaved
[ -z $VERSION ] && echo "ERROR: InterxVersion was NOT found in contants '$CONSTANS_FILE' !" && exit 1
echo $VERSION
