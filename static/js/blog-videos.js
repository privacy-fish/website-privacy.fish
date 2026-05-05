(function () {
  var videos = document.querySelectorAll("[data-blog-card-video], [data-card-video]");
  if (!videos.length) return;

  var reduceMotion = window.matchMedia("(prefers-reduced-motion: reduce)").matches;
  if (reduceMotion) return;
  var canHover = window.matchMedia("(hover: hover) and (pointer: fine)").matches;
  var activeVideos = new WeakSet();

  function play(video) {
    if (activeVideos.has(video)) return;
    activeVideos.add(video);
    video.currentTime = 0;
    var promise = video.play();
    if (promise && typeof promise.catch === "function") {
      promise.catch(function () {
        activeVideos.delete(video);
        // Muted hover playback can still be blocked by strict browser settings.
      });
    }
  }

  function reset(video) {
    activeVideos.delete(video);
    video.pause();
    video.currentTime = 0;
  }

  videos.forEach(function (video) {
    var card = video.closest("a");
    if (!card) return;

    video.preload = "auto";
    video.load();
    video.addEventListener("ended", function () {
      activeVideos.delete(video);
    });

    card.addEventListener("mouseenter", function () {
      if (!canHover) return;
      play(video);
    });

    card.addEventListener("mouseleave", function () {
      if (!canHover) return;
      reset(video);
    });

    card.addEventListener("focusin", function () {
      play(video);
    });

    card.addEventListener("focusout", function () {
      reset(video);
    });

    if (!canHover && "IntersectionObserver" in window) {
      var observer = new IntersectionObserver(function (entries) {
        entries.forEach(function (entry) {
          if (entry.intersectionRatio >= 0.75) {
            play(video);
          } else if (entry.intersectionRatio < 0.25) {
            reset(video);
          }
        });
      }, {
        threshold: [0, 0.25, 0.75, 1],
      });
      observer.observe(video);
    }
  });
})();
