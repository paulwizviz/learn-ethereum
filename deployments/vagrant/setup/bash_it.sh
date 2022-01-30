#!/bin/bash

git clone --depth=1 https://github.com/Bash-it/bash-it.git ~/.bash_it
~/.bash_it/install.sh

ln -s $HOME/setup/gopath.bash $HOME/.bash_it/custom/gopath.bash
ln -s $HOME/setup/eth.bash $HOME/.bash_it/custom/eth.bash

mkdir ~/go
mkdir ~/workspace

exit