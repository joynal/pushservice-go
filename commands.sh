# server setup
sudo apt update && sudo apt upgrade && sudo apt dist-upgrade -y
sudo apt autoremove -y
sudo apt install -y make python git vim curl wget g++ libsasl2-dev libssl-dev libstdc++6
sudo snap install go --classic
mkdir -p /home/joynal3483/go/bin
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
ssh-keygen -t rsa -b 4096 -C “joynal.uu@gmail.com”
cat .ssh/id_rsa.pub

# copy this path
# PATH=$PATH":/home/joynal3483/go/bin"
source .profile

mkdir -p go/src && cd go/src/ || exit
git clone git@bitbucket.org:Joynal/pushservice-go.git
cd pushservice-go/ || exit
dep ensure

# build go program
go build -ldflags="-s -w" -o bin/parser parser/*
go build -ldflags="-s -w" -o bin/sender sender/*

# server tewak
ulimit -n 1000000
ulimit -S 1000000
