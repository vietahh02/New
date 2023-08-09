$(() => {
  let change = 0;
  const allImg = document.getElementsByClassName("slider-item");
  $(".dots a").click(function () {
    $(".dots a").eq(change).removeClass("active");
    let number = $(this).attr("number");
    change = number;
    $(".container .slider .wap").css(
      "margin-left",
      -800 * parseInt(number) + "px"
    );
    $(".dots a").eq(number).addClass("active");
    changeThumbnail();
  });
  $(".container .thumbnailPrev").click(function () {
    changeImg("prev");
  });
  $(".container .thumbnailNext").click(function () {
    changeImg("next");
  });
  function changeImg(type) {
    $(".dots a").eq(change).removeClass("active");
    change = type == "next" ? parseInt(change) + 1 : change - 1;
    if (change >= allImg.length) {
      change = 0;
    }
    if (change < 0) {
      change = allImg.length - 1;
    }
    $(".container .slider .wap").css(
      "margin-left",
      -800 * parseInt(change) + "px"
    );
    $(".dots a").eq(change).addClass("active");
    changeThumbnail();
  }
  function changeThumbnail() {
    prev = change - 1;
    next = Number(change) + 1;
    if (change >= allImg.length - 1) {
      next = 0;
    }
    if (change <= 0) {
      prev = allImg.length - 1;
    }
    $(".container .thumbnailPrev div").html(
      $(".container .slider .wap a").eq(prev).html()
    );
    $(".container .thumbnailNext div").html(
      $(".container .slider .wap a").eq(next).html()
    );
    $(".container .button div img").removeClass("slider-item");
  }
  $(function () {
    $(".onThumbnail").click(function () {
      $(".container .button div").removeClass("off");
    });
    $(".offThumbnail").click(function () {
      $(".container .button div").addClass("off");
    });
    $(".onDot").click(function () {
      $(".dots a").removeClass("off");
    });
    $(".offDot").click(function () {
      $(".dots a").addClass("off");
    });
  });
});
