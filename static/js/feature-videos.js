(function () {
  var videos = Array.prototype.slice.call(document.querySelectorAll("[data-limited-video]"));
  if (!videos.length) return;

  var activeVideo = null;
  var stateByVideo = new WeakMap();

  function play(video) {
    var promise = video.play();
    if (promise && typeof promise.catch === "function") {
      promise.catch(function () {
        // Muted autoplay can still be blocked in some browser modes.
      });
    }
  }

  function reset(video) {
    var state = stateByVideo.get(video);
    if (state) {
      state.plays = 0;
      state.inSequence = false;
    }
    video.pause();
    try {
      video.currentTime = 0;
    } catch (error) {
      // Some browsers can reject currentTime updates until metadata is ready.
    }
    if (activeVideo === video) activeVideo = null;
  }

  function startSequence(video) {
    var state = stateByVideo.get(video);
    if (!state) return;

    if (activeVideo && activeVideo !== video) {
      reset(activeVideo);
    }

    activeVideo = video;
    state.plays = 0;
    state.inSequence = true;
    video.currentTime = 0;
    play(video);
  }

  videos.forEach(function (video) {
    var maxPlays = parseInt(video.dataset.playCount || "2", 10);
    if (!Number.isFinite(maxPlays) || maxPlays < 1) maxPlays = 2;

    stateByVideo.set(video, {
      maxPlays: maxPlays,
      plays: 0,
      inSequence: false,
    });
    video.preload = "auto";
    video.load();

    video.addEventListener("ended", function () {
      var state = stateByVideo.get(video);
      if (!state || !state.inSequence) return;

      state.plays += 1;
      if (state.plays < state.maxPlays) {
        video.currentTime = 0;
        play(video);
        return;
      }

      state.inSequence = false;
      if (activeVideo === video) activeVideo = null;
    });
  });

  if ("IntersectionObserver" in window) {
    var observer = new IntersectionObserver(function (entries) {
      entries.forEach(function (entry) {
        var video = entry.target;
        var state = stateByVideo.get(video);
        if (!state) return;

        if (entry.intersectionRatio >= 0.98) {
          if (!state.inSequence && video.paused) {
            startSequence(video);
          }
          return;
        }

        reset(video);
      });
    }, {
      threshold: [0, 0.5, 0.98, 1],
    });

    videos.forEach(function (video) {
      observer.observe(video);
    });
  } else {
    startSequence(videos[0]);
  }
})();
