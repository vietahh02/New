let status = "off";
let result = document.getElementById("result");
let turn = document.getElementById("turn");
let plus = document.getElementById("plus");
let pin = document.getElementById("pin");

function number(val) {
  if (status == "off") {
    return;
  }
  if (result.innerHTML.length > 12) {
    return;
  } else if (result.innerHTML.length > 6) {
    result.style.fontSize = "60px";
    result.style.margin = "60px 5px 0 0";
  } else if (result.innerHTML.length < 6) {
    result.style.fontSize = "80px";
    result.style.margin = "40px 5px 0 0";
  }
  result.innerHTML += Number(val);
}

function operator(val) {
  x = plus.innerHTML.split(" ");
  if (result.innerHTML == "" && plus.innerHTML == "") {
  } else if (x[0]!=""&&x[1]!=""&&result.innerHTML!="") {
    calculate(result.innerHTML)
    plus.innerHTML = result.innerHTML + " " + val;
    result.innerHTML = "";
  } else {
    plus.innerHTML = result.innerHTML + " " + val;
    result.innerHTML = "";
  }
}

function special(val) {
  x = plus.innerHTML.split(" ");
  if (result.innerHTML == "" && plus.innerHTML == "") {
  } else {
    if (val == "ra") {
      plus.innerHTML = "&radic;" + "(" + result.innerHTML + ")";
      result.innerHTML = parseFloat(
        Math.sqrt(Number(result.innerHTML)).toFixed(6)
      );
    } else if (val == "pow") {
      plus.innerHTML = "(" + result.innerHTML + ")^2";
      result.innerHTML = Math.pow(Number(result.innerHTML), 2);
      if (result.innerHTML.length > 12) {
        result.style.fontSize = "30px";
        result.style.margin = "90px 5px 0 0";
      } else if (result.innerHTML.length > 6) {
        result.style.fontSize = "60px";
        result.style.margin = "60px 5px 0 0";
      } else if (result.innerHTML.length < 6) {
        result.style.fontSize = "80px";
        result.style.margin = "40px 5px 0 0";
      }
    }
  }
}

function decimal() {
  k = result.innerHTML;
  for (i = 0; i < k.length; i++) {
    if (k[i] == ".") {
      k = "no";
      break;
    }
  }
  if (k == "no") {
  } else {
    result.innerHTML += ".";
  }
}

function calculate(val) {
  x = plus.innerHTML.split(" ");
  switch (x[1]) {
    case "+":
      result.innerHTML = parseFloat((Number(x[0]) + Number(val)).toFixed(6));
      break;
    case "-":
      result.innerHTML = parseFloat((Number(x[0]) - Number(val)).toFixed(6));
      break;
    case "x":
      result.innerHTML = parseFloat((Number(x[0]) * Number(val)).toFixed(6));
      break;
    case "/":
      if (val == 0) {
        alert("chia cho 0");
      } else {
        result.innerHTML = parseFloat((Number(x[0]) / Number(val)).toFixed(6));
      }
      break;
  }
}

function clearNum() {
  result.innerHTML = "";
  plus.innerHTML = "";
}

function resultNum() {
  c = 0;
  x = plus.innerHTML.split(" ");
  p = x[0];
  for (var i = 0; i < p.length; i++) {
    if (p[i] == "(" && x[1]!="") {
      plus.innerHTML = result.innerHTML + " =";
      c = 1;
      break;
    }
  }
  if (x[1] == "=") {
  } else if (c != 0) {
  } else {
    if (plus.innerHTML == "") {
      plus.innerHTML = result.innerHTML + " =";
    } else if (plus.innerHTML != "" && result.innerHTML == "") {
      plus.innerHTML = x[0] + " " + x[1] + " " + x[0] + " " + "=";
      calculate(x[0]);
    } else if (plus.innerHTML != "" && result.innerHTML != "") {
      k = x[2];
      if (x.length == 4) {
        plus.innerHTML = result.innerHTML + " " + x[1] + " " + k + " " + "=";
        calculate(k);
      } else {
        plus.innerHTML += " " + result.innerHTML + " =";
        calculate(result.innerHTML);
      }
    }
  }
}

function deleteNum() {
  if (result.innerHTML.length < 8) {
    result.style.fontSize = "80px";
    result.style.margin = "40px 0 0 0";
  }
  result.innerHTML = result.innerHTML.substring(0, result.innerHTML.length - 1);
}

function turnOnOff() {
  pin.innerHTML = "";
  if (status == "off") {
    status = "on";
    turn.innerHTML = "on";
    for (var i = 0; i < 6; i++) {
      const newElement = document.createElement("div");
      newElement.className = "high";
      pin.appendChild(newElement);
    }
  } else {
    clearNum();
    status = "off";
    turn.innerHTML = "off";
    for (var i = 0; i < 6; i++) {
      const newElement = document.createElement("div");
      newElement.className = "low";
      pin.appendChild(newElement);
    }
  }
}
