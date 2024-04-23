const a:number[] = [1, 2, 3, 4]
const b = a; // 参照渡し（共有）
b[0] = 10;
console.log({a, b}, a === b); // { a: [ 10, 2, 3, 4 ], b: [ 10, 2, 3, 4 ] } true