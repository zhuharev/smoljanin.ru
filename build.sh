GOPATH=$(pwd)/Godeps/_workspace
echo $GOPATH
godep path
godep go build
sudo start smoljanin.ru