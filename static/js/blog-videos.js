(function () {
  var videos = Array.prototype.slice.call(document.querySelectorAll("[data-blog-card-video], [data-card-video]"));
  if (!videos.length) return;

  videos.forEach(function (video) {
    video.preload = "auto";
    video.load();
  });

  var reduceMotion = window.matchMedia("(prefers-reduced-motion: reduce)").matches;
  if (reduceMotion) return;

  var activeVideo = null;
  var completedInView = new WeakSet();
  var ticking = false;
  var playThreshold = 0.6;
  var resetThreshold = 0.2;

  function visibilityRatio(video) {
    var rect = video.getBoundingClientRect();
    if (rect.width <= 0 || rect.height <= 0) return 0;

    var viewportWidth = window.innerWidth || document.documentElement.clientWidth;
    var viewportHeight = window.innerHeight || document.documentElement.clientHeight;
    var visibleWidth = Math.max(0, Math.min(rect.right, viewportWidth) - Math.max(rect.left, 0));
    var visibleHeight = Math.max(0, Math.min(rect.bottom, viewportHeight) - Math.max(rect.top, 0));

    return (visibleWidth * visibleHeight) / (rect.width * rect.height);
  }

  function distanceFromViewportCenter(video) {
    var rect = video.getBoundingClientRect();
    var videoCenter = rect.top + rect.height / 2;
    var viewportCenter = (window.innerHeight || document.documentElement.clientHeight) / 2;

    return Math.abs(videoCenter - viewportCenter);
  }

  function reset(video) {
    completedInView.delete(video);
    video.pause();
    try {
      video.currentTime = 0;
    } catch (error) {
      // Some browsers can reject currentTime changes before enough data is loaded.
    }
    if (activeVideo === video) activeVideo = null;
  }

  function play(video) {
    if (completedInView.has(video)) return;
    if (activeVideo && activeVideo !== video) reset(activeVideo);

    activeVideo = video;
    var promise = video.play();
    if (promise && typeof promise.catch === "function") {
      promise.catch(function () {
        if (activeVideo === video) activeVideo = null;
      });
    }
  }

  function selectMostVisibleVideo() {
    var bestVideo = null;
    var bestRatio = 0;
    var bestDistance = Infinity;

    videos.forEach(function (video) {
      var ratio = visibilityRatio(video);

      if (ratio < resetThreshold) {
        reset(video);
        return;
      }

      if (ratio < playThreshold) return;

      var distance = distanceFromViewportCenter(video);
      if (ratio > bestRatio || (ratio === bestRatio && distance < bestDistance)) {
        bestVideo = video;
        bestRatio = ratio;
        bestDistance = distance;
      }
    });

    if (!bestVideo) {
      if (activeVideo) reset(activeVideo);
      return;
    }

    videos.forEach(function (video) {
      if (video !== bestVideo && !video.paused) reset(video);
    });

    if (completedInView.has(bestVideo)) return;

    if (activeVideo !== bestVideo) {
      if (activeVideo) reset(activeVideo);
      try {
        bestVideo.currentTime = 0;
      } catch (error) {
        // Some browsers can reject currentTime changes before enough data is loaded.
      }
    }

    play(bestVideo);
  }

  function requestUpdate() {
    if (ticking) return;
    ticking = true;
    requestAnimationFrame(function () {
      ticking = false;
      selectMostVisibleVideo();
    });
  }

  videos.forEach(function (video) {
    video.addEventListener("ended", function () {
      completedInView.add(video);
      if (activeVideo === video) activeVideo = null;
    });
  });

  window.addEventListener("scroll", requestUpdate, { passive: true });
  window.addEventListener("resize", requestUpdate);
  window.addEventListener("orientationchange", requestUpdate);
  requestUpdate();
})();
