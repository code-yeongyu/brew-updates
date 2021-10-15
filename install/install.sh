echo "Install golang"
brew install go

echo "Download source code"
git clone https://github.com/code-yeongyu/brew-updates --branch master

echo "Install go dependency modules"
cd brew-updates/src/
go mod tidy

echo "Build BrewUpdates"
CGO_ENABLED=0 go build -a -o bin/brew-updates main.go
echo "Install BrewUpdates"
sudo mv bin/brew-updates /usr/local/bin
cd ../..
rm -rf brew-updates

echo "Done!"
echo "You can now easily check your upgradable packages using 'brew-updates'"
echo "I recommend you to add 'brew-updates' on your shell config like .bashrc to get notification everytime you open the terminal."