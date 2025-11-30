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
cwgo server -I ../../idl --type HTTP --service gateway --module github.com/fuhanyang/CYGlowFlow/app/gatewa --idl ../../idl/gateway/user.proto
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

:gen-product
cd rpc_gen
cwgo client --type RPC --service product --module github.com/fuhanyang/CYGlowFlow/rpc_gen  -I ../idl  --idl ../idl/product.proto
cd ..
cd app/product
cwgo server --type RPC --service product --module github.com/fuhanyang/CYGlowFlow/app/product --pass "-use github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/product.proto
cd ../..
goto end

:gen-cart
cd rpc_gen
cwgo client --type RPC --service cart --module github.com/fuhanyang/CYGlowFlow/rpc_gen  -I ../idl  --idl ../idl/cart.proto
cd ..
cd app/cart
cwgo server --type RPC --service cart --module github.com/fuhanyang/CYGlowFlow/app/cart --pass "-use github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/cart.proto
cd ../..
goto end

:gen-payment
cd rpc_gen
cwgo client --type RPC --service payment --module github.com/fuhanyang/CYGlowFlow/rpc_gen  -I ../idl  --idl ../idl/payment.proto
cd ..
cd app/payment
cwgo server --type RPC --service payment --module github.com/fuhanyang/CYGlowFlow/app/payment --pass "-use github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/payment.proto
cd ../..
goto end

:gen-checkout
cd rpc_gen
cwgo client --type RPC --service checkout --module github.com/fuhanyang/CYGlowFlow/rpc_gen  -I ../idl  --idl ../idl/checkout.proto
cd ..
cd app/checkout
cwgo server --type RPC --service checkout --module github.com/fuhanyang/CYGlowFlow/app/checkout --pass "-use github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/checkout.proto
cd ../..
goto end

:gen-order
cd rpc_gen
cwgo client --type RPC --service order --module github.com/fuhanyang/CYGlowFlow/rpc_gen  -I ../idl  --idl ../idl/order.proto
cd ..
cd app/order
cwgo server --type RPC --service order --module github.com/fuhanyang/CYGlowFlow/app/order --pass "-use github.com/fuhanyang/CYGlowFlow/rpc_gen/kitex_gen"  -I ../../idl  --idl ../../idl/order.proto
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
echo   gen-product
echo   gen-cart
echo   gen-payment
echo   gen-checkout
echo   gen-order

:end

