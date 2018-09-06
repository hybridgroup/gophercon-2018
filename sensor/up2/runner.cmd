@echo off
if [%1] == [] goto usage
if [%2] == [] goto usage

rem If you're using putty, change these to pscp and plink
set SCP=scp
set SSH=ssh

:good
if [%3] == [runonly] goto run
echo "Compiling..."
set GOARCH=amd64
set GOOS=linux
go build -o %1app ./%1/main.go
if errorlevel 1 goto exit
echo "Copying..."
%SCP% %1app upsquared@%2:/home/upsquared/%1app
echo "Setting permissions..."
%SSH% -t upsquared@%2 chmod +x ./%1app
:run
echo "Running..."
%SSH% -t upsquared@%2 ./%1app
goto exit

:usage
echo Usage %0 [stepX] [target host] [runonly]
echo     target host may be a host name or an ip address
echo     Specify runonly to skip compilation and copying

:exit
