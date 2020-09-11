#!/bin/bash
echo "Hello from $(hostname)"
for i in $(seq 1 10); do
    echo "for_loop-${i}"
done
