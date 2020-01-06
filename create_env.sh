#!/bin/bash
envs=(
    "CLIENT_SECRET"
    "CLIENT_ID"
    "REDIRECT_URL"
    "MONGODB_URI"
)

echo "" > envs.yml

for i in ${envs[@]}; do
    value=$(env |grep $i |sed -e 's/=/: /g');
    echo $value >> envs.yml;
done;

cat envs.yml
