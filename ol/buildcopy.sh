#!/bin/bash
npm run build
rm -r ../public/
cp -r ./dist ../public
