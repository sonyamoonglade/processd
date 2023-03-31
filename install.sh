#/!bin/bash
path=/opt/processd
mkdir -p $path
mv ./processd $path/processd
echo 'export PATH=$PATH:$path' >> ~/.bashrc 
