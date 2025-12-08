@echo off

if "%1"=="" goto help

if "%1"=="gen-demo-proto" goto gen-demo-proto
if "%1"=="gen-demo-thrift" goto gen-demo-thrift
if "%1"=="demo-link-fix" goto demo-link-fix
if "%1"=="gen-gateway" goto gen-gateway
if "%1"=="gen-user" goto gen-user
if "%1"=="gen-product" goto gen-product
if "%1"=="gen-cart" goto gen-cart
if "%1"=="gen-payment" goto gen-payment
if "%1"=="gen-checkout" goto gen-checkout
if "%1"=="gen-order" goto gen-order

echo 未知命令: %1
goto help

:gen-demo-proto
cd demo/demo_proto
cwgo server -I ../../idl --module github.com/fuhanyang/CYGlowFlow/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto
cd ../..
goto end

:gen-demo-thrift
cd demo/demo_thrift
cwgo server --module github.com/fuhanyang/CYGlowFlow/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift
cd ../..
goto end

:demo-link-fix
cd demo/demo_proto
golangci-lint run -E gofumpt --path-prefix=. --fix --timeout=5m
cd ../..
goto end

:gen-gateway
cd app/gateway
cwgo server -I ../../idl --type HTTP --service gateway --module github.com/fuhanyang/CYGlowFlow/app/gateway --idl ../../idl/gateway/chat.proto
cwgo server -I ../../idl --type HTTP --service gateway --module github.com/fuhanyang/CYGlowFlow/app/gateway --idl ../../idl/gateway/login.proto
cwgo server -I ../../idl --type HTTP --service gateway --module github.com/fuhanyang/CYGlowFlow/app/gateway --idl ../../idl/gateway/home.proto
cwgo server -I ../../idl --type HTTP --service gateway --module github.com/fuhanyang/CYGlowFlow/app/gateway --idl ../../idl/gateway/friend.proto
cwgo server -I ../../idl --type HTTP --service gateway --module github.com/fuhanyang/CYGlowFlow/app/gateway --idl ../../idl/gateway/news.proto
cwgo server -I ../../idl --type HTTP --service gateway --module github.com/fuhanyang/CYGlowFlow/app/gateway --idl ../../idl/gateway/workflow.proto
cwgo server -I ../../idl --type HTTP --service gateway --module github.com/fuhanyang/CYGlowFlow/app/gateway --idl ../../idl/gateway/user.proto
cd ../..
goto end

:gen-user
cd rpc_gen
cwgo client --type RPC --service user --module github.com/fuhanyang/CYGlowFlow/rpc_gen  -I ../idl  --idl ../idl/user/user.proto
cd ..
cd app/user
cwgo server --type RPC --service user --module github.com/fuhanyang/CYGlowFlow/app/user --pass "-use github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/user/user.proto
cd ../..
goto end

:gen-workflow
cd rpc_gen
cwgo client --type RPC --service workflow --module github.com/fuhanyang/CYGlowFlow/rpc_gen  -I ../idl  --idl ../idl/workflow/infra/execution_manager.proto
cwgo client --type RPC --service workflow --module github.com/fuhanyang/CYGlowFlow/rpc_gen  -I ../idl  --idl ../idl/workflow/infra/workflow_manager.proto
cd ..
cd app/workflow
cwgo server --type RPC --service workflow  --module github.com/fuhanyang/CYGlowFlow/app/workflow --pass "-use github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/workflow/infra/execution_manager.proto
cwgo server --type RPC --service workflow  --module github.com/fuhanyang/CYGlowFlow/app/workflow --pass "-use github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/workflow/infra/workflow_manager.proto
cd ..
cd ../..
goto end



:help
echo 用法: build.bat [command]
echo 命令:
echo   gen-demo-proto
echo   gen-demo-thrift
echo   demo-link-fix
echo   gen-gateway
echo   gen-user

:end

