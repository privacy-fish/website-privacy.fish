(function () {
  var script = document.currentScript;
  if (!script) return;

  var wasmSrc = script.getAttribute("data-cap-wasm-src");
  if (wasmSrc) {
    window.CAP_CUSTOM_WASM_URL = new URL(wasmSrc, window.location.origin).href;
  }
})();
