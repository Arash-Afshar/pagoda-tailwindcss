#! /bin/bash
set -e

if [ "$#" -ne 3 ]; then
  echo "usage: ./replace.sh <github_org> <github_repo> <app_name>"
  exit 1
fi

ORG=$1
REPO=$2
APP_NAME=$3

# rm -rf .git

sed -i "s/Arash-Afshar\/pagoda-tailwindcss/$1\/$2/g" */*/*.go
sed -i "s/Arash-Afshar\/pagoda-tailwindcss/$1\/$2/g" */*.go
sed -i "s/Arash-Afshar\/pagoda-tailwindcss/$1\/$2/g" go.mod
sed -i "s/Arash-Afshar\/pagoda-tailwindcss/$1\/$2/g" README.md
sed -i "s/PAGODA/${APP_NAME^^}/g" config/config.go
sed -i "s/pagoda/${APP_NAME,,}/g" config/config.go
sed -i "s/Pagoda/${APP_NAME}/g" config/config.yaml
sed -i "s/PAGODA/${APP_NAME^^}/g" README.md
sed -i "s/pagoda/${APP_NAME,,}/g" README.md
sed -i "s/Pagoda/${APP_NAME}/g" README.md
sed -i "s/pagoda-tailwindcss/${APP_NAME,,}/g" package-lock.json
