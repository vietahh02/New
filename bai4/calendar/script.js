const nowDate = new Date();
let nowYear = nowDate.getFullYear();
let nowMonth = nowDate.getMonth();
const months = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
  ];

function dates(year, month) {
  const date = new Date(year, month, 1);
  const dates = document.getElementById("dates");
  const lastDay = new Date(year, month + 1, 0);

  dates.innerHTML = "";

  document.getElementById("month_year").innerHTML =
  months[nowMonth] + " " + year;

  const day = new Date(year, month, 0).getDate();
  const date1 = new Date(year, month-1, day - date.getDay());
  for (let i = date1.getDate() + 1; i <= day; i++) {
    const newElement = document.createElement("div");
    newElement.className = "dayOutOfMonth";
    if (
      date1.getFullYear() == nowDate.getFullYear() &&
      date1.getMonth() == nowDate.getMonth() &&
      i == nowDate.getDate()
    ) {
      newElement.className = "today";
    }
    newElement.innerText = i;
    dates.appendChild(newElement);
  }

  for (let i = 1; i <= lastDay.getDate(); i++) {
    const newElement = document.createElement("div");
    if (
      date.getFullYear() == nowDate.getFullYear() &&
      date.getMonth() == nowDate.getMonth() &&
      i == nowDate.getDate()
    ) {
      newElement.className = "today";
    }
    newElement.innerText = i;
    dates.appendChild(newElement);
  }

  k = 1;
  for (let i = lastDay.getDay(); i < 6; i++) {
    const newElement = document.createElement("div");
    newElement.className = "dayOutOfMonth";
    if (
        lastDay.getFullYear() == nowDate.getFullYear() &&
        lastDay.getMonth() == nowDate.getMonth() &&
        k == nowDate.getDate()
      ) {
        newElement.className = "today";
      }
    newElement.innerText = k;
    dates.appendChild(newElement);
    k++;
  }
}

function prevMonth() {
  nowMonth--;
  if (nowMonth < 0) {
    nowYear--;
    nowMonth = 11;
  }
  dates(nowYear, nowMonth);
}

function nextMonth() {
  nowMonth++;
  if (nowMonth > 11) {
    nowYear++;
    nowMonth = 0;
  }
  dates(nowYear, nowMonth);
}

dates(nowYear, nowMonth);
