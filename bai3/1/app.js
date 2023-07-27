
document.write("<div style='text-align: center;margin: auto;font-size:20px;max-width: 800px;width: 100%;margin: auto;'>")
document.write("<table style='border: 2px solid black;'>");
document.write("<tr>");

for (var i = 1; i <= 5; i++) {
  document.write("<td style='border:1px solid black;'>");
  for (var j = 1; j <= 10; j++) {
    document.write( i + "x" + j + "=" + j * i + "<br>");
  }
  document.write("</td>");
}
document.write("</tr>");
document.write("<tr>");

for (var i = 6; i <= 10; i++) {
  document.write("<td style='border:1px solid black;'>");
  for (var j = 1; j <= 10; j++) {
    document.write( i + "x" + j + "=" + j * i + "<br>");
  }
  document.write("</td>");
}
document.write("</tr>");

document.write("</table>");
document.write("</div>")