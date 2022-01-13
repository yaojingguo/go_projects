#!/usr/bin/env bash

for i in {1..5}
do
  ( curl localhost:5555/foo & ) 
done
