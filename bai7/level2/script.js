const textarea = document.getElementById("textarea");
const myList = document.getElementById("myList");
const listItems = myList.getElementsByTagName("li");
const box = document.getElementById("box");
const listBox = box.getElementsByTagName("div");
const listDe = myList.getElementsByClassName("chooseE");

function importElement() {
  const a = textarea.value.split("\n");
  for (let i = 0; i < a.length; i++) {
    if (a[i] != "") {
      const newE = document.createElement("li");
      newE.onclick = toggleClass;
      newE.innerHTML = a[i];
      myList.appendChild(newE);
    }
  }
  textarea.value = "";
}

function importTopElement() {
  const b = [];
  for (let i = 0; i < listItems.length; i++) {
    b[i] = listItems[i].innerHTML;
  }
  const a = textarea.value.split("\n");
  const c = a.concat(b);
  myList.innerHTML = "";
  for (let i = 0; i < c.length; i++) {
    if (c[i] != "") {
      const newE = document.createElement("li");
      newE.onclick = toggleClass;
      newE.draggable = "true";
      newE.innerHTML = c[i];
      myList.appendChild(newE);
    }
  }
  textarea.value = "";
}

function deleteElement() {
    for (let i = 0; i < listDe.length; i++) {
      myList.removeChild(listDe[i]);
      i--;
    }
}

function exportElement() {
  for (let i = 0; i < listDe.length; i++) {
    const newE = document.createElement("div");
    const inputE = document.createElement("input");
    const newLa = document.createElement("label");
    inputE.type = "checkbox";
    newE.id = i;
    newE.value = listDe[i].innerHTML;
    inputE.value = listDe[i].innerHTML;
    newLa.innerHTML = listDe[i].innerHTML;
    newE.appendChild(inputE);
    newE.appendChild(newLa);
    box.appendChild(newE);
  }
  for (let i = 0; i < listBox.length; i++) {
    listBox[i].id = i;
    listBox[i].onclick = function () {
      document.getElementById(i).remove();
      document.getElementById(i).innerHTML = ""; //
    };
  }
  deleteElement();
}

function orderElement() {
  const b = Array.from(listItems);
  myList.innerHTML = "";
  b.sort((a, b) => a.textContent.localeCompare(b.textContent));
  for (let i = 0; i < b.length; i++) {
    if (b[i] != "") {
      const newE = document.createElement("li");
      newE.onclick = toggleClass;
      newE.draggable = "true";
      newE.innerHTML = b[i].innerHTML;
      myList.appendChild(newE);
    }
  }
}

function toggleClass() {
  if (this.classList.contains("chooseE")) {
    this.classList.remove("chooseE");
  } else {
    this.classList.add("chooseE");
  }
}

$(function ok() {
  $("#myList").sortable();
  $("li").on("click", function () {
    $(this).toggleClass("chooseE", 50);
  });
});
