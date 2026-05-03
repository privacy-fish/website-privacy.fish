(function () {
  var form = document.getElementById("signup-form");
  if (!form) return;

  var username = form.querySelector("#username");
  var usernameError = form.querySelector("[data-username-error]");
  var availabilityStatus = form.querySelector("[data-availability-status]");
  var captchaStatus = form.querySelector("[data-captcha-status]");
  var triesStatus = form.querySelector("[data-tries-status]");
  var formStatus = form.querySelector("[data-form-status]");
  var keyError = form.querySelector("[data-key-error]");
  var keyFileInput = form.querySelector("[data-key-file-input]");
  var keyFileStatus = form.querySelector("[data-key-file-status]");
  var keyDropZone = form.querySelector("[data-key-drop-zone]");
  var credentialInput = form.querySelector("[data-credential]");
  var keySection = form.querySelector("[data-key-section]");
  var textarea = form.querySelector("#ssh_keys");
  var overlay = form.querySelector("[data-key-overlay]");
  var counter = form.querySelector("[data-key-counter]");
  var submit = form.querySelector('[data-signup-intent="reserve"]');
  var checkButton = form.querySelector('[data-signup-intent="check"]');
  var paymentReference = document.querySelector("[data-payment-reference]");
  var copyPaymentReference = document.querySelector("[data-copy-payment-reference]");
  var copyPaymentReferenceLabel = document.querySelector("[data-copy-payment-reference-label]");
  var paymentSection = document.querySelector("[data-payment-section]");
  var usernamePattern = /^[a-z0-9](?:[a-z0-9.-]{0,29}[a-z0-9])?$/;
  var freshCaptchaTimer = null;
  var state = {
    credential: "",
    triesLeft: 0,
    checkedUsername: "",
    available: false,
    busy: false,
    submitIntent: "check",
    waitingForFreshCaptcha: false,
    reserved: false,
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

  function setAvailabilityWithAttemptWarning(status, available) {
    if (!availabilityStatus) return;
    availabilityStatus.classList.remove("hidden", "text-amber", "text-red-600", "text-ocean", "text-deep/60");
    availabilityStatus.innerHTML =
      '<span class="' + (available ? "text-ocean" : "text-red-600") + '">' + escapeHtml(status) + '</span>' +
      '\n<span class="text-red-600">No attempts left. To check additional usernames, please solve the captcha above again.</span>';
  }

  function setFormStatusMessage(msg, tone) {
    if (!formStatus) return;
    formStatus.classList.remove("hidden", "text-amber", "text-red-600", "text-ocean", "text-deep/60");
    if (!msg) {
      formStatus.innerHTML = "";
      formStatus.classList.add("hidden");
      return;
    }
    formStatus.textContent = msg;
    if (tone === "ok") formStatus.classList.add("text-ocean");
    else if (tone === "bad") formStatus.classList.add("text-red-600");
    else if (tone === "muted") formStatus.classList.add("text-deep/60");
    else formStatus.classList.add("text-amber");
  }

  function setPaymentReference(value) {
    if (!paymentReference) return;
    var hasReference = Boolean(value);
    paymentReference.textContent = value || paymentReference.dataset.paymentPlaceholder || "";
    paymentReference.classList.toggle("font-mono", hasReference);
    paymentReference.classList.toggle("tracking-[0.08em]", hasReference);
    paymentReference.classList.toggle("sm:text-2xl", hasReference);
    if (copyPaymentReference) {
      copyPaymentReference.classList.toggle("hidden", !hasReference);
      copyPaymentReference.classList.toggle("inline-flex", hasReference);
    }
    resetCopyPaymentReferenceLabel();
  }

  function showReservedState(reference) {
    setPaymentReference(reference);
    setFormStatusMessage(
      formStatus ? formStatus.dataset.reservedMessage || "Your account is reserved for 30 days." : "Your account is reserved for 30 days.",
      "ok"
    );
    if (paymentSection) {
      paymentSection.scrollIntoView({ behavior: "smooth", block: "nearest" });
    }
  }

  function resetCopyPaymentReferenceLabel() {
    if (!copyPaymentReference || !copyPaymentReferenceLabel) return;
    copyPaymentReferenceLabel.textContent = copyPaymentReference.dataset.copyLabel || "Copy";
  }

  function setCopyPaymentReferenceCopied() {
    if (!copyPaymentReference || !copyPaymentReferenceLabel) return;
    copyPaymentReferenceLabel.textContent = copyPaymentReference.dataset.copiedLabel || "Copied";
    window.setTimeout(resetCopyPaymentReferenceLabel, 1800);
  }

  function copyText(text) {
    if (navigator.clipboard && window.isSecureContext) {
      return navigator.clipboard.writeText(text);
    }
    return new Promise(function (resolve, reject) {
      var input = document.createElement("textarea");
      input.value = text;
      input.setAttribute("readonly", "");
      input.style.position = "fixed";
      input.style.top = "-9999px";
      document.body.appendChild(input);
      input.select();
      try {
        document.execCommand("copy") ? resolve() : reject(new Error("copy_failed"));
      } catch (err) {
        reject(err);
      } finally {
        document.body.removeChild(input);
      }
    });
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
    var invalid = keys.some(function (k) { return !k.ok; }) || keys.length > 10;
    counter.textContent = keys.length + " / 10";
    counter.classList.toggle("text-amber", invalid);
    counter.classList.toggle("text-deep/52", !invalid);
    textarea.classList.toggle("border-red-600", invalid);
    textarea.classList.toggle("border-deep/12", !invalid);
    overlay.classList.toggle("border-red-600/35", invalid);
    overlay.classList.toggle("border-transparent", !invalid);
    if (keys.length > 10) {
      setMessage(keyError, "Add no more than 10 SSH public keys.", "bad");
    } else if (invalid) {
      setMessage(keyError, "Only complete ssh-ed25519 public keys are accepted.", "bad");
    } else {
      setMessage(keyError, null);
    }
    if (keys.length === 0) return true;
    return !invalid;
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
    if (type !== "ssh-ed25519") return "bad";
    if (!blob) return "in-progress";
    if (blob.length < 68) return "bad";
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
      var cls = classifyLine(line);
      if (cls === "bad") return '<span class="font-semibold text-red-600">' + safe + '</span>';
      if (cls === "in-progress") return '<span class="text-deep/45">' + safe + '</span>';
      return '<span class="text-deep">' + safe + '</span>';
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
      : "No attempts left. To check additional usernames, please solve the captcha above again.";
  }

  function resetAvailability() {
    state.checkedUsername = "";
    state.available = false;
    setMessage(availabilityStatus, null);
  }

  function keyFingerprint(line) {
    var parts = line.trim().split(/\s+/);
    if (parts.length >= 2) return parts[0] + " " + parts[1];
    return line.trim();
  }

  function uploadedKeyLines(text) {
    return text
      .split(/\r?\n/)
      .map(function (line) { return line.trim(); })
      .filter(function (line) {
        return /^ssh-\S+\s+\S+/.test(line);
      });
  }

  function appendKeyLines(lines) {
    var existingLines = textarea.value
      .split(/\r?\n/)
      .map(function (line) { return line.trim(); })
      .filter(Boolean);
    var seen = {};
    existingLines.forEach(function (line) {
      seen[keyFingerprint(line)] = true;
    });
    var added = [];
    lines.forEach(function (line) {
      var fingerprint = keyFingerprint(line);
      if (seen[fingerprint]) return;
      seen[fingerprint] = true;
      added.push(line);
    });
    if (added.length > 0) {
      textarea.value = existingLines.concat(added).join("\n");
    }
    renderHighlight();
    syncScroll();
    refreshSubmit();
    return added.length;
  }

  async function importKeyFiles(fileList) {
    var files = Array.prototype.slice.call(fileList || []);
    if (files.length === 0) return;
    setMessage(keyFileStatus, "Reading public key file" + (files.length === 1 ? "" : "s") + "...", "muted");
    try {
      var allLines = [];
      var skippedLarge = 0;
      for (var i = 0; i < files.length; i += 1) {
        var file = files[i];
        if (file.size > 65536) {
          skippedLarge += 1;
          continue;
        }
        var text = await file.text();
        allLines = allLines.concat(uploadedKeyLines(text));
      }
      var added = appendKeyLines(allLines);
      if (skippedLarge > 0) {
        setMessage(keyFileStatus, "Skipped " + skippedLarge + " file" + (skippedLarge === 1 ? "" : "s") + " because public key files should be smaller than 64 KB.", "bad");
      } else if (added > 0) {
        setMessage(
          keyFileStatus,
          "Added " + added + " public key" + (added === 1 ? "" : "s") + " from " + files.length + " file" + (files.length === 1 ? "" : "s") + ". You can add more .pub files if needed.",
          "ok"
        );
      } else if (allLines.length > 0) {
        setMessage(keyFileStatus, "Those keys are already listed below. You can add more .pub files if needed.", "muted");
      } else {
        setMessage(keyFileStatus, "No SSH public key lines were found in the selected file" + (files.length === 1 ? "" : "s") + ".", "bad");
      }
    } catch (err) {
      setMessage(keyFileStatus, "Could not read the selected public key file" + (files.length === 1 ? "" : "s") + ".", "bad");
    }
  }

  function setDropZoneActive(active) {
    if (!keyDropZone) return;
    keyDropZone.classList.toggle("border-ocean", active);
    keyDropZone.classList.toggle("bg-ocean-soft/60", active);
    keyDropZone.classList.toggle("border-deep/8", !active);
    keyDropZone.classList.toggle("bg-mist/35", !active);
  }

  function setKeySectionVisible(visible) {
    if (!keySection) return;
    keySection.classList.toggle("hidden", !visible);
  }

  function refreshSubmit() {
    var u = validateUsername();
    var k = validateKeys();
    var hasKeys = parseKeys().length > 0;
    var canReserve = u && k && hasKeys && !state.busy && !state.reserved;
    checkButton.disabled = !u || state.busy || state.reserved;
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

  function currentCapToken() {
    var tokenInput = form.querySelector('input[name="cap-token"]');
    return tokenInput ? tokenInput.value.trim() : "";
  }

  function stopFreshCaptchaWatcher() {
    state.waitingForFreshCaptcha = false;
    if (!freshCaptchaTimer) return;
    clearInterval(freshCaptchaTimer);
    freshCaptchaTimer = null;
  }

  function clearAvailabilityAfterFreshCaptcha() {
    if (!state.waitingForFreshCaptcha || !currentCapToken()) return;
    stopFreshCaptchaWatcher();
    setMessage(availabilityStatus, null);
  }

  function startFreshCaptchaWatcher() {
    if (freshCaptchaTimer) clearInterval(freshCaptchaTimer);
    freshCaptchaTimer = setInterval(clearAvailabilityAfterFreshCaptcha, 200);
    clearAvailabilityAfterFreshCaptcha();
  }

  function resetCaptchaWidget() {
    clearCapToken();
    var widget = form.querySelector("cap-widget");
    if (!widget) return;
    var replacement = widget.cloneNode(false);
    widget.replaceWith(replacement);
  }

  function expireCredential(message) {
    state.credential = "";
    state.triesLeft = 0;
    state.waitingForFreshCaptcha = true;
    credentialInput.value = "";
    resetCaptchaWidget();
    startFreshCaptchaWatcher();
    updateTriesStatus();
    setMessage(captchaStatus, null);
    if (message) {
      setMessage(availabilityStatus, message, "bad");
    }
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
    stopFreshCaptchaWatcher();
    credentialInput.value = state.credential;
    updateTriesStatus();
    setMessage(captchaStatus, "Captcha accepted. You can now make up to " + state.triesLeft + " attempts.", "ok");
    setMessage(availabilityStatus, null);
  }

  async function checkAvailability() {
    if (!validateUsername()) return;
    setBusy(true);
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
      if (state.triesLeft <= 0) {
        var status = availabilityStatus.textContent;
        expireCredential(null);
        setAvailabilityWithAttemptWarning(status, data.available);
      }
    } catch (err) {
      if (err.code === "credential_invalid" || err.code === "credential_exhausted") {
        expireCredential(err.message);
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
    setFormStatusMessage(
      formStatus ? formStatus.dataset.reservingMessage || "Reserving account..." : "Reserving account...",
      "muted"
    );
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
      state.reserved = true;
      showReservedState(data.payment_reference);
      submit.disabled = true;
      checkButton.disabled = true;
    } catch (err) {
      if (err.code === "credential_invalid" || err.code === "credential_exhausted") {
        expireCredential(err.message);
      }
      setFormStatusMessage(err.message || "Could not reserve this account.", "bad");
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

  if (copyPaymentReference) {
    copyPaymentReference.addEventListener("click", function () {
      var reference = paymentReference ? paymentReference.textContent.trim() : "";
      if (!reference) return;
      copyText(reference)
        .then(setCopyPaymentReferenceCopied)
        .catch(function () {
          setFormStatusMessage("Could not copy the payment reference automatically.", "bad");
        });
    });
  }

  if (keyFileInput) {
    keyFileInput.addEventListener("change", async function () {
      try {
        await importKeyFiles(keyFileInput.files);
      } finally {
        keyFileInput.value = "";
      }
    });
  }

  if (keyDropZone) {
    ["dragenter", "dragover"].forEach(function (eventName) {
      keyDropZone.addEventListener(eventName, function (e) {
        e.preventDefault();
        e.stopPropagation();
        setDropZoneActive(true);
      });
    });
    ["dragleave", "drop"].forEach(function (eventName) {
      keyDropZone.addEventListener(eventName, function (e) {
        e.preventDefault();
        e.stopPropagation();
        setDropZoneActive(false);
      });
    });
    keyDropZone.addEventListener("drop", function (e) {
      importKeyFiles(e.dataTransfer ? e.dataTransfer.files : []);
    });
  }

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

  renderHighlight();
  setKeySectionVisible(true);
  updateTriesStatus();
  refreshSubmit();
})();
