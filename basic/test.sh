#!/bin/bash

# Simple test to flood queue with requests

for i in {1..4096};
    do curl localhost:8000/work -d name=$USER -d delay=$(expr $i % 11)s;
done