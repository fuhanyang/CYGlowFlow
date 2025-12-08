#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/workflow/infra/execution_manager"
exec "$CURDIR/bin/workflow/infra/execution_manager"
