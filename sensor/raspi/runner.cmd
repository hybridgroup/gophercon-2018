@echo off
if [%1] == [] goto usage
if [%2] == [] goto usage

rem If you're using putty, change these to pscp and plink
set SCP=scp
set SSH=ssh

:good
if [%3] == [runonly] goto run
echo "Compiling..."
set GOARCH=arm
set GOOS=linux
go build -o %1app ./%1/main.go
if errorlevel 1 goto exit
echo "Copying..."
%SCP% %1app pi@%2:/home/pi/%1app
echo "Setting permissions..."
%SSH% -t pi@%2 chmod +x ./%1app
:run
echo "Running..."
%SSH% -t pi@%2 ./%1app
goto exit

:usage
echo Usage %0 [stepX] [target host] [runonly]
echo     target host may be a host name or an ip address
echo     Specify runonly to skip compilation and copying

:exit
