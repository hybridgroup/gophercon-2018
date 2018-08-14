[ $# -eq 0 ] && { echo "Usage: $0 [step] [ipaddress]"; exit 1; }

echo "Compiling..."
GOARCH=arm GOOS=linux go build $1.go
echo "Copying..."
scp $1 pi@$2:/home/pi/$1
echo "Running..."
ssh -t pi@$2 ./$1
