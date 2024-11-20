function updateCard() {
  var bgColor = document.getElementById("backgroundColor").value;
  var borderColor = document.getElementById("borderColor").value;
  var customText = document.getElementById("customText").value;

  document.getElementById("customCard").style.backgroundColor = bgColor;
  document.getElementById("customCard").style.borderColor = borderColor;
  document.getElementById("cardText").innerText = customText;
}
