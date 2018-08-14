[ $# -eq 0 ] && { echo "Usage: $0 [stepX] [ipaddress]"; exit 1; }

echo "Compiling..."
GOARCH=arm GOOS=linux go build -o $1app ./$1/main.go
echo "Copying..."
scp $1app pi@$2:/home/pi/$1app
echo "Running..."
ssh -t pi@$2 ./$1app
