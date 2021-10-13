echo "Install golang"
brew install go

echo "Download source code"
git clone git@github.com:code-yeongyu/brew-updates.git

echo "Install go dependency modules"
cd brew-updates/src/
go mod tidy

echo "Build and Install BrewUpdates"
CGO_ENABLED=0 go build -a -o bin/brew-updates main.go
mv bin/brew-updates /usr/local/bin
cd ../..
rm -rf brew-updates

echo "Done!"
echo "You can now easily check your upgradable packages using 'brew-updates'"
echo "I recommend you to add 'brew-updates' on your shell config like .bashrc to get notification everytime you open the terminal."