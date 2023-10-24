function bin2Dec() {
  let binaryInput = document.getElementById(`binaryInput`).value;
  let binaryOutput = document.getElementById(`decimalOutput`);
    if (binaryInput.match(/^[01]+$/)) {
      let result = 0;
      for (let index = binaryInput.length - 1; index >= 0; index--) {
        result += parseInt(binaryInput[index]) * Math.pow(2, binaryInput.length - 1 - index); // 修改这里的计算方式
      }
      binaryOutput.textContent = "十进制输出: " + result;
    } else {
      binaryOutput.textContent = "请输入有效的二进制数字"
    }
}
