git clone https://github.com/itmisx/znj
cd znj/web_src
yarn
yarn build
cd ../
go mod tidy
rm -rf build/release/*
GOOS=linux GOARCH=amd64 go build -o build/release/znj main.go
cp -r etc build/release
docker build -t znj:0.0.1 .