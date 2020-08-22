import json
import sys
import time

def main(t):
    print("from python: ", t)
    time.sleep(t)


if __name__ == '__main__':
    t = sys.argv[1]
    start = time.time()
    main(json.loads(t))
    end = time.time()
