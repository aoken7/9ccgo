#!/bin/bash
assert() {
  expected="$1"
  input="$2"

  ./main "$input" > tmp.s
  cc -o tmp tmp.s
  ./tmp
  actual="$?"

  if [ "$actual" = "$expected" ]; then
    echo "$input => $actual"
  else
    echo "$input => $expected expected, but got $actual"
    exit 1
  fi
}

assert 0 0
assert 42 42
assert 21 "5+20-4"
assert 41 " 12 + 34 - 5 "
assert 47 '5+6*7'
assert 15 '5*(9-6)'
assert 10 '-5+15'
assert 1 '1 < 2'
assert 0 '1 > 2'
assert 1 '1 <= 2 == 2 <= 2'
assert 0 '1 >= 2 == 2 >= 2'
assert 5 '1; 2+3;'
assert 3 'a=3; a;'
assert 8 'a=3; z=5; a+z;'
assert 3 'foo=3; foo;'
assert 8 '{ foo=3; bar=5; foo+bar; }'
assert 3 'return 3; 5;'
assert 3 'return 3; return 5;'
assert 5 'if(1 + 2 > 0) { return 5; } return 3;'
assert 3 'if(0) { return 5; } else { return 3; }'

echo OK