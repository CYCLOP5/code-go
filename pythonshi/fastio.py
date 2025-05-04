import sys, threading, gc
sys.setrecursionlimit(10**7)
gc.disable()

data = sys.stdin.buffer.read().split()
idc = 0
def inp(): 
    global idc; v = data[idc]; idc += 1; return v
def ni():  
    return int(inp())
_out = sys.stdout.write
def outl(s):  
    _out(s + '\n')

def solve():
    t = ni()
    for _ in range(t):
        n = ni()
        arr = [ni() for _ in range(n)]

        outl(str(sum(arr)))

if __name__ == "__main__":

    solve()
    gc.enable()