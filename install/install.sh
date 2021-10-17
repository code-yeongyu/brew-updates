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

echo "Registering `brew-updates` to your shell config ..."
if [[ $SHELL == *"bash"* ]];then
    echo "bash" >> ~/.bashrc
elif [[ $SHELL == *"zsh"* ]];then
    echo "brew-updates" >> ~/.zshrc
elif [[ $SHELL == *"fish"* ]];then
    echo "brew-updates" >> ~/.config/fish/config.fish
else
    echo "Cannot detect the shell you use."
    echo "You can manually add `brew-updates` into your shell config."
fi

echo "Done!"
echo "You can now easily check your upgradable packages using 'brew-updates'"