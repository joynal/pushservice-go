# server setup
sudo apt update && sudo apt upgrade && sudo apt dist-upgrade -y
sudo apt autoremove -y
sudo apt install -y make python git vim curl wget g++ libsasl2-dev libssl-dev libstdc++6
ssh-keygen -t rsa -b 4096 -C “joynal.uu@gmail.com”
cat .ssh/id_rsa.pub
sudo snap install go --classic
mkdir -p /home/joynal3483/go/bin
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
source .profile

# copy this path
# PATH=$PATH":/home/joynal3483/go/bin"

mkdir -p go/src && cd go/src/
git clone git@bitbucket.org:Joynal/pushservice-go.git
cd pushservice-go/
cp .env.example .env
dep ensure

# server tewak
ulimit -n 2000000
