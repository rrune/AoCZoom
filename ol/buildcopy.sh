#!/bin/bash
npm run build
rm -r ../data/public/
cp -r ./dist ../data/public
