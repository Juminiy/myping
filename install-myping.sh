#!/bin/bash

# Unix Like system installation, MacOS,Linux distros:Ubuntu,ArchLinux,CentOS...
brew install go 
yum -u install go  
sudo apt install go 

go install github.com/Juminiy/myping@v1.0.5 

myping info 
