#!/usr/bin/env python3

from sys import stdin
import json

lines = []
for line in stdin:
    lines.append(line.strip())


print(json.dumps(lines))