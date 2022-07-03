#!/bin/bash

# Unix Like system installation, MacOS,Linux distros:Ubuntu,ArchLinux,CentOS...
go version

# judge os 
if [ -ge ] do 
    sudo apt install go  
done 

git clone https://github.com/Juminiy/myping
cd myping 
go mod tidy
go install 

myping info 
