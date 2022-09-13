import os
import subprocess


subprocess.Popen(["go run client.go", "result.txt"], stdout=subprocess.PIPE, 
           stderr=subprocess.STDOUT)

