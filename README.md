Go Language
#ダウンロード
https://golang.org/dl/
#インストール
sudo tar -C /usr/local -xzf go1.6.3.linux-amd64.tar.gz
#.prifileに追加
export PATH=$PATH:/usr/local/go/bin

first 
cd gotest
export GOPATH = `pwd`


go fmt
go build
go install
go run
