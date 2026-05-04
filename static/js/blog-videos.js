(function () {
  var videos = document.querySelectorAll("[data-blog-card-video], [data-card-video]");
  if (!videos.length) return;

  var reduceMotion = window.matchMedia("(prefers-reduced-motion: reduce)").matches;
  if (reduceMotion) return;

  function play(video) {
    video.currentTime = 0;
    var promise = video.play();
    if (promise && typeof promise.catch === "function") {
      promise.catch(function () {
        // Muted hover playback can still be blocked by strict browser settings.
      });
    }
  }

  function reset(video) {
    video.pause();
    video.currentTime = 0;
  }

  videos.forEach(function (video) {
    var card = video.closest("a");
    if (!card) return;

    card.addEventListener("mouseenter", function () {
      play(video);
    });

    card.addEventListener("mouseleave", function () {
      reset(video);
    });

    card.addEventListener("focusin", function () {
      play(video);
    });

    card.addEventListener("focusout", function () {
      reset(video);
    });
  });
})();
