#!/bin/bash

ignite generate ts-client
cp -rf ts-client/* ts-client-builder/src/
rm -r ts-client-builder/src/node_modules
cd ts-client-builder
yarn
yarn build
npm publish --registry=http://192.168.1.4:8081/repository/sao-hosted/
npm publish --registry=http://205.204.75.250:38081/repository/sao-network-hosted/