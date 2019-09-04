#!/bin/zsh zsh

echo "Hello, Shell"

git pull

if [ ! -d "./sh" ]; then
    mkdir sh
fi
cd sh
pwd 

for ((i=0; i<10; i++)); do
	touch test_$i.txt
done