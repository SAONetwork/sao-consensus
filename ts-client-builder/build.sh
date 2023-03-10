#!/bin/bash

ignite generate ts-client -y

[ -d ts-client-builder/src ] || mkdir ts-client-builder/src

rm -r ts-client-builder/src/*
cp -rf ts-client/* ts-client-builder/src/
rm -r ts-client-builder/src/node_modules
cp ts-client-builder/entry.ts ts-client-builder/src/

cd ts-client-builder
yarn
yarn build
