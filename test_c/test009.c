// 再帰のテスト

int cnt(int n) {
    if(n == 0)
        return 0;
    else
        return cnt(n - 1) + 2;
}

int main() { return cnt(10); }