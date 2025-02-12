var num1="", num2="", op="";
function calc(ch){
    const num = parseInt(ch)
    // if(ch >= '0' && ch <= '9') {
    if(num >= 0 && num <= 9) {
        let boxValue = document.getElementById("box").value;
        boxValue += ch;
        document.getElementById("box").value = boxValue;
    }else if(ch == '=') {
        num2 = document.getElementById("box").value;
        let expr = num1 + op + num2;
        let res = eval(expr)
        console.log(typeof(res))
        num1 = "";
        num2 = "";
        op = "";
        document.getElementById("box").value = res;
    }else if(ch == 'C') {
        num1 = "";
        num2 = "";
        op = "";
        document.getElementById("box").value = "";
    }else { // + - * / 
        num1 = document.getElementById("box").value;
        op = ch;
        document.getElementById("box").value = "";
    }
}