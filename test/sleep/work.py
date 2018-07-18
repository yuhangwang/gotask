import json
import sys
import time

def main(argd):
    print("from python: ", argd[0])
    time.sleep(argd[0])


if __name__ == '__main__':
    arg = "".join(sys.argv[1:])
    start = time.time()
    main(json.loads(arg))
    end = time.time()
