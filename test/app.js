$(document).ready(function () {
  $(".image-slider").slick({
    slidesToShow: 1,
    slideToScroll: 2,
    infinite: true,
    arrows: false,
    autoplay: true,
    autoplaySpeed: 1500,
    fade: false,
    draggable: true,
    prevArrow: `<button type='button' class='slick-prev style-arrows'><i class="fa-solid fa-arrow-left"></i></button>`,
    nextArrow: `<button type='button' class='slick-next style-arrows'><i class="fa-solid fa-arrow-right"></i></button>`,
    dots: true,
  });
});
