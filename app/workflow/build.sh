#!/usr/bin/env bash
RUN_NAME="workflow/infra/execution_manager"
mkdir -p output/bin output/conf
cp script/* output/
cp -r conf/* output/conf
chmod +x output/bootstrap.sh
go build -o output/bin/${RUN_NAME}
