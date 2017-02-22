import json
import sys

def main(argd):
    for k, v in argd.items():
        print(k, v)

if __name__ == '__main__':
    arg = "".join(sys.argv[1:])
    main(json.loads(arg))