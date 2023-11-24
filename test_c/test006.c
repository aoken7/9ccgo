int g(int c) {
    return c + 3;
}

int f(int b){
    return g(b);
}

int main(){
    int a = 123;
    return f(a);
}