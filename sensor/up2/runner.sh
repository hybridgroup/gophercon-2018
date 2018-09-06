[ $# -eq 0 ] && { echo "Usage: $0 [stepX] [ipaddress]"; exit 1; }

echo "Compiling..."
GOARCH=amd64 GOOS=linux go build -o $1app ./$1/main.go
echo "Copying..."
scp $1app upsquared@$2:/home/upsquared/$1app
echo "Running..."
ssh -t upsquared@$2 ./$1app
