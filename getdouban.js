// ==UserScript==
// @name         New Userscript
// @namespace    http://tampermonkey.net/
// @version      0.1
// @description  try to take over the world!
// @author       You
// @match        https://www.tampermonkey.net/scripts.php#gh
// @match        https://book.douban.com/
// @icon         https://www.google.com/s2/favicons?sz=64&domain=tampermonkey.net
// @grant        none
// ==/UserScript==

(function () {
    'use strict';

    document.onreadystatechange = function () {
        if (document.readyState === "complete") {
            addMenuBar();
        }
    }

    function addMenuBar() {
        let previewBtn = document.createElement("button");
        let tableVisible = false
        previewBtn.setAttribute("type", "button");
        previewBtn.setAttribute("id", "preview-button");
        previewBtn.addEventListener("click", even => {
            if (tableVisible) { deleteTable() } else { showDataTable() }
            tableVisible = !tableVisible;
        })
        previewBtn.append(document.createTextNode("查看表格数据"));
        let titleNode = document.querySelector("#content > div > div.article > div.section.books-express > div.hd");
        titleNode.parentNode.insertBefore(previewBtn, titleNode.nextSibling);
    }
    function showDataTable() {
        let titleList = ["书名", "作者", "出版社", "出版日期"];
        let thead = document.createElement("thead");
        let row = thead.insertRow();
        titleList.forEach(title => {
            let cell = row.insertCell();
            cell.innerText = title;
        })
        let bookList = parseBookData();
        let tbody = document.createElement("tbody");
        bookList.forEach(book => {
            let row = tbody.insertRow();
            book.forEach(item => {
                row.insertCell().innerText = item;
            })
        })
        let table = document.createElement("table");
        table.setAttribute("border", 1);
        table.setAttribute("id", "booktable");
        table.appendChild(thead);
        table.appendChild(tbody);

        let targetNode = document.querySelector("#preview-button");
        targetNode.parentNode.insertBefore(table, targetNode.nextSibling);
    }

    function parseBookData() {
        let bookList = [];
        let nodeulList = document.querySelectorAll("#content > div > div.article > div.section.books-express > div.bd > div.carousel > div > ul")
        nodeulList.forEach(node => {
            let nodeliList = node.querySelectorAll("li");
            nodeliList.forEach(nodeli => {
                let title = nodeli.querySelector("div.info div.title a").innerHTML.trim();
                let author = nodeli.querySelector("div.info div.author").innerHTML.trim();
                let pubilsher = nodeli.querySelector("div.info div.more-meta p span.publisher").innerHTML.trim();
                let year = nodeli.querySelector("div.info div.more-meta p span.year").innerHTML.trim();
                bookList.push([title, author, pubilsher, year]);
            })
        })
        return bookList;
    }

    function deleteTable() {
        let table = document.querySelector("#booktable");
        table.parentNode.removeChild(table)
    }
})();