#!/usr/bin/env python3
import subprocess
from pathlib import Path

file_build = Path("./build.sh")
print("building...")
subprocess.run([str(file_build.absolute())], check=True,stdout=subprocess.PIPE, stderr=subprocess.PIPE)
