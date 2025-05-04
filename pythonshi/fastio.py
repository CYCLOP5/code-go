import sys, os, io

_data = io.BytesIO(os.read(0, os.fstat(0).st_size)).read().split()
_idx = 0
def inp() -> str:
    global _idx
    val = _data[_idx]
    _idx += 1
    return val.decode()
def ri() -> int:
    return int(inp())
def rstr() -> str:
    return inp()
def raf() -> float:
    return float(inp())

_out = sys.stdout.write
def out(s: str):
    _out(s)
def outl(s: str):
    _out(s + '\n')

def main():
    t = ri()
    for _ in range(t):
        n = ri()
        arr = [ri() for _ in range(n)]

        outl(str(sum(arr)))

if __name__ == "__main__":
    main()