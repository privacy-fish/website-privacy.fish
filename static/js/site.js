(() => {
  function qs(selector, root = document) {
    return root.querySelector(selector);
  }

  function qsa(selector, root = document) {
    return Array.from(root.querySelectorAll(selector));
  }

  function initNavbarHeightVar() {
    const navbar = document.getElementById("navbar");
    if (!navbar) return;

    const update = () => {
      document.documentElement.style.setProperty("--navbar-height", `${navbar.offsetHeight}px`);
    };

    update();
    window.addEventListener("resize", update);
  }

  function initScrollSpy() {
    const navbarLinks = qsa(".navbar-link").filter((link) => {
      const target = link.getAttribute("data-nav") || "";
      return target.startsWith("/#") || target.startsWith("#");
    });
    if (navbarLinks.length === 0) return;
    if (window.location.pathname !== "/") return;

    function updateActiveNav() {
      const sections = qsa("section[id]");
      const scrollPosition = window.scrollY + 100;

      for (const section of sections) {
        const sectionTop = section.offsetTop;
        const sectionHeight = section.offsetHeight;
        const sectionId = section.getAttribute("id");

        if (scrollPosition >= sectionTop && scrollPosition < sectionTop + sectionHeight) {
          for (const link of navbarLinks) {
            link.classList.toggle("active", link.getAttribute("data-nav") === `/#${sectionId}`);
          }
          return;
        }
      }
    }

    window.addEventListener("scroll", updateActiveNav, { passive: true });
    updateActiveNav();
  }

  function initDocTocScrollSpy() {
    const toc = qs("[data-doc-toc]");
    if (!toc) return;

    const links = qsa('a[href^="#"]', toc);
    if (links.length === 0) return;

    const linkById = new Map();
    for (const link of links) {
      const id = decodeURIComponent(link.hash.replace("#", ""));
      if (id) linkById.set(id, link);
    }

    const headings = qsa("article [id]").filter((heading) => linkById.has(heading.id));
    if (headings.length === 0) return;

    let ticking = false;

    function setActive(activeLink) {
      for (const link of links) {
        link.classList.toggle("active", link === activeLink);
      }
    }

    function updateActiveToc() {
      ticking = false;
      const navbar = document.getElementById("navbar");
      const headerOffset = (navbar ? navbar.offsetHeight : 0) + 40;
      let activeHeading = headings[0];

      for (const heading of headings) {
        if (heading.getBoundingClientRect().top <= headerOffset) {
          activeHeading = heading;
        } else {
          break;
        }
      }

      setActive(linkById.get(activeHeading.id));
    }

    function requestUpdate() {
      if (ticking) return;
      ticking = true;
      requestAnimationFrame(updateActiveToc);
    }

    window.addEventListener("scroll", requestUpdate, { passive: true });
    window.addEventListener("resize", requestUpdate);
    window.addEventListener("hashchange", requestUpdate);
    updateActiveToc();
  }

  function initMobileMenu() {
    const navbar = document.getElementById("navbar");
    const toggle = qs("[data-mobile-menu-toggle]");
    const menu = qs("[data-mobile-menu]");
    if (!navbar || !toggle || !menu) return;

    let isOpen = false;
    let closeTimer = null;
    const transitionMs = 500;
    const openLabel = toggle.dataset.openLabel || toggle.getAttribute("aria-label") || "";
    const closeLabel = toggle.dataset.closeLabel || openLabel;

    function setOpen(nextOpen) {
      isOpen = Boolean(nextOpen);
      toggle.setAttribute("aria-expanded", isOpen ? "true" : "false");
      toggle.setAttribute("aria-label", isOpen ? closeLabel : openLabel);
      toggle.classList.toggle("is-open", isOpen);
      menu.setAttribute("aria-hidden", isOpen ? "false" : "true");
      window.clearTimeout(closeTimer);

      if (isOpen) {
        menu.hidden = false;
        requestAnimationFrame(() => {
          if (!isOpen) return;
          menu.classList.add("is-open");
        });
        return;
      }

      menu.classList.remove("is-open");
      closeTimer = window.setTimeout(() => {
        if (isOpen) return;
        menu.hidden = true;
      }, transitionMs);
    }

    menu.hidden = true;
    menu.classList.remove("is-open");

    toggle.addEventListener("click", () => setOpen(!isOpen));

    menu.addEventListener("click", (e) => {
      const anchor = e.target.closest("a");
      if (anchor) setOpen(false);
    });

    document.addEventListener("click", (e) => {
      if (!isOpen) return;
      if (!navbar.contains(e.target)) setOpen(false);
    });

    window.addEventListener("resize", () => {
      if (window.matchMedia("(min-width: 1150px)").matches) setOpen(false);
    });
  }

  function navigateToHash(href, behavior) {
    const targetId = href.replace("/#", "").replace("#", "");
    const targetElement = document.getElementById(targetId);
    if (!targetElement) return false;

    targetElement.scrollIntoView({ behavior, block: "start" });
    history.pushState(null, null, `#${targetId}`);
    return true;
  }

  function initSmoothAnchors() {
    document.addEventListener("click", (e) => {
      const anchor = e.target.closest('a[href^="/#"], a[href^="#"]');
      if (!anchor) return;

      const href = anchor.getAttribute("href");
      if (!href || href === "#" || href === "/#") return;

      if (href.startsWith("#") || href.startsWith("/#")) {
        if (navigateToHash(href, "smooth")) {
          e.preventDefault();
        }
      }
    });

    window.addEventListener("load", () => {
      const hash = window.location.hash;
      if (!hash) return;
      const targetElement = document.getElementById(hash.replace("#", ""));
      if (!targetElement) return;

      setTimeout(() => {
        targetElement.scrollIntoView({ behavior: "auto", block: "start" });
      }, 100);
    });
  }

  function initRevealAnimations() {
    const animatedElements = qsa(".fade-in, .seq");
    if (animatedElements.length === 0) return;

    if (window.matchMedia("(prefers-reduced-motion: reduce)").matches) {
      for (const el of animatedElements) el.dataset.reveal = "visible";
      return;
    }

    const viewportHeight = window.innerHeight || document.documentElement.clientHeight;
    for (const el of animatedElements) {
      const rect = el.getBoundingClientRect();
      const isVisibleNow = rect.top < viewportHeight && rect.bottom > 0;
      el.dataset.reveal = isVisibleNow ? "visible" : "pending";
    }

    const observer = new IntersectionObserver(
      (entries) => {
        for (const entry of entries) {
          if (!entry.isIntersecting) continue;
          entry.target.dataset.reveal = "visible";
          observer.unobserve(entry.target);
        }
      },
      { threshold: 0.1, rootMargin: "0px 0px -50px 0px" },
    );

    for (const el of animatedElements) {
      if (el.dataset.reveal === "pending") observer.observe(el);
    }
  }

  function initFaqAnimations() {
    const faqItems = qsa("[data-faq-item]");
    if (faqItems.length === 0) return;

    const prefersReducedMotion = window.matchMedia("(prefers-reduced-motion: reduce)").matches;

    for (const item of faqItems) {
      const summary = qs("summary", item);
      const panel = qs("[data-faq-panel]", item);
      if (!summary || !panel) continue;

      if (prefersReducedMotion) continue;

      panel.style.height = item.open ? "auto" : "0px";
      panel.style.opacity = item.open ? "1" : "0";
      panel.style.overflow = "hidden";
      panel.style.transitionProperty = "height, opacity";
      panel.style.transitionDuration = "320ms";
      panel.style.transitionTimingFunction = "cubic-bezier(0.22, 1, 0.36, 1)";

      summary.addEventListener("click", (e) => {
        e.preventDefault();
        if (item.dataset.animating === "true") return;

        if (item.open) {
          collapseFaqItem(item, panel);
        } else {
          expandFaqItem(item, panel);
        }
      });
    }
  }

  function expandFaqItem(item, panel) {
    item.dataset.animating = "true";
    item.open = true;
    panel.style.height = "0px";
    panel.style.opacity = "0";

    requestAnimationFrame(() => {
      panel.style.height = `${panel.scrollHeight}px`;
      panel.style.opacity = "1";
    });

    const onTransitionEnd = (e) => {
      if (e.propertyName !== "height") return;
      panel.removeEventListener("transitionend", onTransitionEnd);
      panel.style.height = "auto";
      item.dataset.animating = "false";
    };

    panel.addEventListener("transitionend", onTransitionEnd);
  }

  function collapseFaqItem(item, panel) {
    item.dataset.animating = "true";
    panel.style.height = `${panel.scrollHeight}px`;
    panel.style.opacity = "1";

    requestAnimationFrame(() => {
      panel.style.height = "0px";
      panel.style.opacity = "0";
    });

    const onTransitionEnd = (e) => {
      if (e.propertyName !== "height") return;
      panel.removeEventListener("transitionend", onTransitionEnd);
      item.open = false;
      item.dataset.animating = "false";
    };

    panel.addEventListener("transitionend", onTransitionEnd);
  }

  function init() {
    initNavbarHeightVar();
    initScrollSpy();
    initDocTocScrollSpy();
    initMobileMenu();
    initSmoothAnchors();
    initRevealAnimations();
    initFaqAnimations();
  }

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", init, { once: true });
  } else {
    init();
  }
})();
