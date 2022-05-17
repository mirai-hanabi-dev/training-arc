const romanToInt = (s) => {
  const mapping = {
    I: 1,
    IV: 4,
    V: 5,
    IX: 9,
    X: 10,
    XL: 40,
    L: 50,
    XC: 90,
    C: 100,
    CD: 400,
    D: 500,
    CM: 900,
    M: 1000,
  };

  let res = 0;
  const n = s.length;
  for (let i = 0; i < n; i++) {
    const ch = s[i];
    let doubleCh = null;
    if (i < n - 1) {
      doubleCh = s[i] + s[i + 1];
    }

    if (doubleCh && Object.prototype.hasOwnProperty.call(mapping, doubleCh)) {
      res += mapping[doubleCh];
      i += 1;
    } else {
      res += mapping[ch];
    }
  }

  return res;
};

let test = "MCMXCIV";
console.log(romanToInt(test));
