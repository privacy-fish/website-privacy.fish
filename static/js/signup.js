(function () {
  var form = document.getElementById("signup-form");
  if (!form) return;

  var username = form.querySelector("#username");
  var usernameError = form.querySelector("[data-username-error]");
  var availabilityStatus = form.querySelector("[data-availability-status]");
  var captchaStatus = form.querySelector("[data-captcha-status]");
  var triesStatus = form.querySelector("[data-tries-status]");
  var formStatus = form.querySelector("[data-form-status]");
  var credentialInput = form.querySelector("[data-credential]");
  var keySection = form.querySelector("[data-key-section]");
  var textarea = form.querySelector("#ssh_keys");
  var overlay = form.querySelector("[data-key-overlay]");
  var counter = form.querySelector("[data-key-counter]");
  var submit = form.querySelector('[data-signup-intent="reserve"]');
  var checkButton = form.querySelector('[data-signup-intent="check"]');
  var result = document.querySelector("[data-signup-result]");
  var paymentReference = document.querySelector("[data-payment-reference]");
  var newSignup = document.querySelector("[data-new-signup]");
  var usernamePattern = /^[a-z0-9](?:[a-z0-9.-]{0,29}[a-z0-9])?$/;
  var state = {
    credential: "",
    triesLeft: 0,
    checkedUsername: "",
    available: false,
    busy: false,
    submitIntent: "check",
  };

  function setMessage(el, msg, tone) {
    if (!el) return;
    el.classList.remove("hidden", "text-amber", "text-red-600", "text-ocean", "text-deep/60");
    if (!msg) {
      el.textContent = "";
      el.classList.add("hidden");
      return;
    }
    el.textContent = msg;
    if (tone === "ok") el.classList.add("text-ocean");
    else if (tone === "bad") el.classList.add("text-red-600");
    else if (tone === "muted") el.classList.add("text-deep/60");
    else el.classList.add("text-amber");
  }

  function setInvalid(el, invalid) {
    if (invalid) {
      el.classList.remove("text-deep");
      el.classList.add("text-red-600");
    } else {
      el.classList.remove("text-red-600");
      el.classList.add("text-deep");
    }
  }

  function currentUsername() {
    return username.value.trim().toLowerCase();
  }

  function validateUsername() {
    var v = currentUsername();
    if (!v) {
      setMessage(usernameError, null);
      setInvalid(username, false);
      return false;
    }
    if (!usernamePattern.test(v)) {
      setMessage(usernameError, "Username must be 1-31 characters, contain only lowercase letters, digits, dots, and hyphens, and start and end with a letter or digit.");
      setInvalid(username, true);
      return false;
    }
    setMessage(usernameError, null);
    setInvalid(username, false);
    return true;
  }

  function parseKeys() {
    return textarea.value
      .split(/\r?\n/)
      .map(function (l) { return l.trim(); })
      .filter(Boolean)
      .map(function (line) {
        var parts = line.split(/\s+/);
        var type = parts[0];
        var blob = parts[1];
        var ok = true;
        var reason = "";
        if (type !== "ssh-ed25519") {
          ok = false;
          reason = "Only ssh-ed25519 keys are accepted.";
        } else if (!blob || !/^[A-Za-z0-9+/]+={0,2}$/.test(blob)) {
          ok = false;
          reason = "Invalid base64 key data.";
        } else {
          try {
            if (atob(blob).length !== 51) {
              ok = false;
              reason = "Decoded key is not the expected length for ed25519.";
            }
          } catch (e) {
            ok = false;
            reason = "Could not decode base64.";
          }
        }
        return { line: line, type: type, blob: blob, ok: ok, reason: reason };
      });
  }

  function validateKeys() {
    var keys = parseKeys();
    counter.textContent = keys.length + " / 10";
    counter.classList.toggle("text-amber", keys.length > 10);
    if (keys.length === 0) return true;
    return keys.every(function (k) { return k.ok; }) && keys.length <= 10;
  }

  function escapeHtml(s) {
    return s.replace(/[&<>]/g, function (c) {
      return c === "&" ? "&amp;" : c === "<" ? "&lt;" : "&gt;";
    });
  }

  function classifyLine(rawLine) {
    var line = rawLine.trim();
    if (!line) return "empty";
    var parts = line.split(/\s+/);
    var type = parts[0];
    var blob = parts[1];
    if (!blob) return "in-progress";
    if (type !== "ssh-ed25519") return "bad";
    if (blob.length < 68) {
      return /^[A-Za-z0-9+/]*$/.test(blob) ? "in-progress" : "bad";
    }
    if (!/^[A-Za-z0-9+/]+={0,2}$/.test(blob)) return "bad";
    try {
      if (atob(blob).length !== 51) return "bad";
    } catch (e) { return "bad"; }
    return "ok";
  }

  function renderHighlight() {
    var rawLines = textarea.value.split(/\r?\n/);
    overlay.innerHTML = rawLines.map(function (line) {
      var safe = escapeHtml(line);
      return classifyLine(line) === "bad"
        ? '<span class="text-red-600">' + safe + '</span>'
        : safe;
    }).join("\n") + "\n";
  }

  function syncScroll() {
    overlay.scrollTop = textarea.scrollTop;
    overlay.scrollLeft = textarea.scrollLeft;
  }

  function updateTriesStatus() {
    if (!state.credential) {
      triesStatus.textContent = "Solve the captcha once, then you get five username or signup attempts.";
      return;
    }
    triesStatus.textContent = state.triesLeft > 0
      ? state.triesLeft + " captcha-backed attempt" + (state.triesLeft === 1 ? "" : "s") + " left."
      : "No attempts left. Solve the captcha again before continuing.";
  }

  function resetAvailability() {
    state.checkedUsername = "";
    state.available = false;
    setMessage(availabilityStatus, null);
  }

  function setKeySectionVisible(visible) {
    if (!keySection) return;
    keySection.classList.toggle("hidden", !visible);
  }

  function refreshSubmit() {
    var u = validateUsername();
    var k = validateKeys();
    var hasKeys = parseKeys().length > 0;
    var canReserve = u && k && hasKeys && !state.busy;
    checkButton.disabled = !u || state.busy;
    submit.disabled = !canReserve;
  }

  function setBusy(busy) {
    state.busy = busy;
    refreshSubmit();
  }

  function apiErrorMessage(code) {
    var messages = {
      bad_origin: "The signup request must come from the privacy.fish site.",
      body_too_large: "The submitted form is too large.",
      captcha_failed: "Captcha verification failed. Please solve it again.",
      credential_invalid: "Captcha session expired. Please solve it again.",
      credential_exhausted: "No captcha-backed attempts left. Please solve the captcha again.",
      rate_limited: "Too many attempts. Please wait before trying again.",
      username_invalid: "The username is not valid.",
      username_taken: "That username is already reserved or taken.",
      keys_missing: "Add at least one ssh-ed25519 public key.",
      keys_too_many: "Add no more than 10 SSH public keys.",
      key_invalid: "One of the SSH public keys is malformed.",
      key_wrong_type: "Only ssh-ed25519 public keys are accepted.",
      server_error: "The signup server could not complete the request.",
    };
    return messages[code] || "The signup request failed.";
  }

  function isLocalhost() {
    return location.hostname === "localhost" || location.hostname === "127.0.0.1" || location.hostname === "::1";
  }

  function waitForSubmitListeners() {
    return new Promise(function (resolve) {
      requestAnimationFrame(resolve);
    });
  }

  async function getCapToken() {
    await waitForSubmitListeners();
    var tokenInput = form.querySelector('input[name="cap-token"]');
    var token = tokenInput ? tokenInput.value.trim() : "";
    if (token) return token;
    if (isLocalhost()) return "dev-local-token";
    throw new Error("Please solve the captcha before continuing.");
  }

  function clearCapToken() {
    form.querySelectorAll('input[name="cap-token"]').forEach(function (input) {
      input.remove();
    });
  }

  async function postForm(endpoint, data) {
    var body = new URLSearchParams();
    Object.keys(data).forEach(function (key) {
      body.set(key, data[key]);
    });
    var response = await fetch(endpoint, {
      method: "POST",
      credentials: "same-origin",
      headers: { "Content-Type": "application/x-www-form-urlencoded" },
      body: body.toString(),
    });
    var json = await response.json().catch(function () {
      return { ok: false, code: "server_error" };
    });
    if (!response.ok || json.ok === false) {
      var err = new Error(apiErrorMessage(json.code));
      err.code = json.code;
      throw err;
    }
    return json;
  }

  async function ensureCredential() {
    if (state.credential && state.triesLeft > 0) return;
    var token = await getCapToken();
    var data = await postForm(form.dataset.lookupEndpoint, { "cap-token": token });
    clearCapToken();
    state.credential = data.credential;
    state.triesLeft = data.tries_left;
    credentialInput.value = state.credential;
    updateTriesStatus();
    setMessage(captchaStatus, "Captcha accepted. You can now make up to " + state.triesLeft + " attempts.", "ok");
  }

  async function checkAvailability() {
    if (!validateUsername()) return;
    setBusy(true);
    setMessage(formStatus, null);
    setMessage(availabilityStatus, "Checking availability...", "muted");
    try {
      await ensureCredential();
      var name = currentUsername();
      var data = await postForm(form.dataset.checkEndpoint, {
        credential: state.credential,
        username: name,
      });
      state.triesLeft = data.tries_left;
      updateTriesStatus();
      state.checkedUsername = name;
      state.available = Boolean(data.available);
      setMessage(
        availabilityStatus,
        data.available
          ? name + "@privacy.fish is available."
          : name + "@privacy.fish is already reserved or taken.",
        data.available ? "ok" : "bad"
      );
    } catch (err) {
      if (err.code === "credential_invalid" || err.code === "credential_exhausted") {
        state.credential = "";
        state.triesLeft = 0;
        credentialInput.value = "";
        clearCapToken();
        updateTriesStatus();
      }
      setMessage(availabilityStatus, err.message || "Could not check that username.", "bad");
    } finally {
      setBusy(false);
    }
  }

  function sanitizedKeys() {
    return parseKeys()
      .map(function (k) { return k.type + " " + k.blob; })
      .join("\n");
  }

  async function reserveAccount() {
    if (!validateUsername() || !validateKeys() || parseKeys().length === 0) {
      refreshSubmit();
      return;
    }
    setBusy(true);
    setMessage(formStatus, "Reserving account...", "muted");
    try {
      await ensureCredential();
      var name = currentUsername();
      var data = await postForm(form.action, {
        credential: state.credential,
        username: name,
        ssh_keys: sanitizedKeys(),
      });
      state.triesLeft = data.tries_left;
      updateTriesStatus();
      if (paymentReference) paymentReference.textContent = data.payment_reference;
      form.hidden = true;
      if (result) {
        result.classList.remove("hidden");
        result.focus({ preventScroll: true });
        result.scrollIntoView({ behavior: "smooth", block: "start" });
      }
    } catch (err) {
      if (err.code === "credential_invalid" || err.code === "credential_exhausted") {
        state.credential = "";
        state.triesLeft = 0;
        credentialInput.value = "";
        clearCapToken();
        updateTriesStatus();
      }
      setMessage(formStatus, err.message || "Could not reserve this account.", "bad");
    } finally {
      setBusy(false);
    }
  }

  username.addEventListener("input", function () {
    resetAvailability();
    refreshSubmit();
  });
  username.addEventListener("blur", refreshSubmit);
  textarea.addEventListener("input", function () {
    renderHighlight();
    refreshSubmit();
  });
  textarea.addEventListener("scroll", syncScroll);
  checkButton.addEventListener("click", function () {
    state.submitIntent = "check";
  });
  submit.addEventListener("click", function () {
    state.submitIntent = "reserve";
  });

  form.addEventListener("submit", function (e) {
    e.preventDefault();
    var active = document.activeElement;
    var intent = e.submitter
      ? e.submitter.dataset.signupIntent
      : active && active.dataset && active.dataset.signupIntent
        ? active.dataset.signupIntent
        : state.submitIntent;
    if (intent === "check") {
      checkAvailability();
      return;
    }
    reserveAccount();
  });

  if (newSignup) {
    newSignup.addEventListener("click", function () {
      form.reset();
      form.hidden = false;
      if (result) result.classList.add("hidden");
      state.credential = "";
      state.triesLeft = 0;
      state.submitIntent = "check";
      credentialInput.value = "";
      resetAvailability();
      setKeySectionVisible(true);
      clearCapToken();
      renderHighlight();
      updateTriesStatus();
      refreshSubmit();
      form.scrollIntoView({ behavior: "smooth", block: "start" });
    });
  }

  renderHighlight();
  setKeySectionVisible(true);
  updateTriesStatus();
  refreshSubmit();
})();
