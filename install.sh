# !/bin/sh

reset="\033[0m"
bold="\033[1m"
color_green="\033[32m"
color_red="\033[31m"

mkdir /tmp/see

cpuArch=$(uname -i)
targetArch=""

if [ $cpuArch = "x86_64" ]
    echo "Detected amd64 architecture"
    targetArch="amd64"
else
    targetArch="arm64'
fi
    curl -fsSLo "/tmp/see/see_$targetArch.deb" "https://github.com/jcf-junior/see/releases/latest/download/see_$targetArch.deb"

chmod +x "/tmp/see/see_$targetArch.deb"

sudo apt install "/tmp/see/see_$targetArch.deb"