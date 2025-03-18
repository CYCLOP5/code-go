#include <cstdio>
#include <cctype>
#include <iostream>
#include <vector>
#include <string>
#include <algorithm>
#include <functional>
using namespace std;

const int BUF_SIZE = 1 << 20;
char buf[BUF_SIZE];
int buf_pos = 0, buf_len = 0;

inline int getChar() {
    if (buf_pos == buf_len) {
        buf_len = (int)fread(buf, 1, BUF_SIZE, stdin);
        if (buf_len == 0) return EOF;
        buf_pos = 0;
    }
    return buf[buf_pos++];
}

inline bool isSpace(int c) {
    return c == ' ' || c == '\n' || c == '\r' || c == '\t';
}

inline bool readInt(int &number) {
    int c = getChar();
    while(c != EOF && isSpace(c)) c = getChar();
    if(c == EOF) return false;
    bool neg = false;
    if(c == '-') {
        neg = true;
        c = getChar();
    }
    number = 0;
    for(; c != EOF && !isSpace(c); c = getChar()) {
        number = number * 10 + (c - '0');
    }
    if(neg) number = -number;
    return true;
}

inline bool readLL(long long &number) {
    int c = getChar();
    while(c != EOF && isSpace(c)) c = getChar();
    if(c == EOF) return false;
    bool neg = false;
    if(c == '-') {
        neg = true;
        c = getChar();
    }
    number = 0;
    for(; c != EOF && !isSpace(c); c = getChar()) {
        number = number * 10 + (c - '0');
    }
    if(neg) number = -number;
    return true;
}

inline bool readString(string &s, int n) {
    s.resize(n);
    int idx = 0;
    int c = getChar();
    while(c != EOF && isSpace(c)) c = getChar();
    if(c == EOF) return false;
    while(c != EOF && !isSpace(c) && idx < n) {
        s[idx++] = (char)c;
        c = getChar();
    }
    s.resize(idx);
    return true;
}