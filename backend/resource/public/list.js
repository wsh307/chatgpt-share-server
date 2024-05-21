$(function () {
  var $div = $("<div></div>");
  $div.css({
    "border-top-left-radius": "34px",
    "border-bottom-left-radius": "34px",
    background: "linear-gradient(140.91deg, #FF87B7 12.61%, #EC4C8C 76.89%)",
    height: "34px",
    width: "45px",
    margin: "1px",
    display: "flex",
    "align-items": "center",
    position: "fixed",
    right: "0px",
    top: `70px`,
    cursor: "pointer",
  });
  $div.html(
    "<span style='color:white;font-size:15px;margin-left:10px'>换号</span>"
  );
  $("body").append($div);
  $div.click(function () {
    window.location.href = "/list";
  });
});
