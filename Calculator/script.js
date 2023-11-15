"use strict"

let el = function(elment) {
    if (elment.charAt(0) === "#") {
        return document.querySelector(elment)
    }
    return document.querySelectorAll(elment)
}

let theNmu = "",
    oldNum = "",
    viewer = el("#viewer"),
    equals = el("#equals"),
    nums = el(".num"),
    ops = el(".ops"),
    resultNum,
    operator;

let setNum = function() {
    if (resultNum) {
        theNmu = this.getAttribute("data-num")
        resultNum = ""
    } else {
        theNmu += this.getAttribute("data-num")
    }
    viewer.innerHTML= theNmu
}

let moveNum = function() {
    oldNum = theNmu
    theNmu = ""
    operator = this.getAttribute("data-ops")
    equals.setAttribute("data-result", "")
}

let displayNum = function() {
    oldNum = parseFloat(oldNum)
    theNmu = parseFloat(theNmu)

    switch (operator) {
        case "plus":
            resultNum = oldNum + theNmu
            break
        case "minus":
            resultNum = oldNum - theNmu
            break
        case "times":
            resultNum = oldNum * theNmu
            break
        case "divided by":
            resultNum = oldNum / theNmu
            break
        default:
            resultNum = theNmu
    }
    if (!isFinite(resultNum)) {
        if (isNaN(resultNum)) {
            resultNum = "ERR!"
        } else {
            resultNum = "Look at what you've done"
            el("#calculator").classList.add("broken")
            el("#reset").classList.add("show")
        }
    }
    viewer.innerHTML = resultNum
    equals.setAttribute("data-result", resultNum)

    oldNum = 0
    theNmu = resultNum
}

let clearAll = function() {
    theNmu = ""
    oldNum = ""
    resultNum = ""
    viewer.innerHTML = 0
    equals.setAttribute("data-result", "")
}

for (let i = 0; i < nums.length; i++) {
    nums[i].onclick = setNum
}
for (let i = 0; i < ops.length; i++) {
    ops[i].onclick = moveNum
}

equals.onclick = displayNum
el("#clear").onclick = clearAll;

el("#reset").onclick = function() {
    window.location = window.location;
}