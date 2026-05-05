(function () {
  var videos = document.querySelectorAll("[data-limited-video]");
  if (!videos.length) return;

  function play(video) {
    var promise = video.play();
    if (promise && typeof promise.catch === "function") {
      promise.catch(function () {
        // Muted autoplay can still be blocked in some browser modes.
      });
    }
  }

  videos.forEach(function (video) {
    var maxPlays = parseInt(video.dataset.playCount || "2", 10);
    if (!Number.isFinite(maxPlays) || maxPlays < 1) maxPlays = 2;
    var plays = 0;
    var autoPlayed = false;

    video.preload = "auto";
    video.load();

    function startSequence() {
      plays = 0;
      video.currentTime = 0;
      play(video);
    }

    video.addEventListener("ended", function () {
      plays += 1;
      if (plays < maxPlays) {
        video.currentTime = 0;
        play(video);
      }
    });

    video.addEventListener("mouseenter", startSequence);

    if ("IntersectionObserver" in window) {
      var observer = new IntersectionObserver(function (entries) {
        entries.forEach(function (entry) {
          if (autoPlayed || entry.intersectionRatio < 0.98) return;
          autoPlayed = true;
          startSequence();
          observer.unobserve(video);
        });
      }, {
        threshold: [0, 0.5, 0.98, 1],
      });
      observer.observe(video);
    } else {
      startSequence();
    }
  });
})();
