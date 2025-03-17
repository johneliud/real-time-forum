import { initThemeToggler, toggleHamburgerMenu } from "./script.js";
import { homeView } from "./home_view.js";
import { signUpView } from "./signup_view.js";
import { signInView } from "./signin_view.js";
import { checkAuthStatus } from "./auth.js";
import { errorView } from "./error_view.js";
import { renderHeader } from "./header.js";

document.addEventListener("DOMContentLoaded", () => {
  const router = new Router();

  // Set up navigation events
  document.body.addEventListener("click", (e) => {
    if (e.target.matches("[data-link]") || e.target.closest("[data-link]")) {
      e.preventDefault();
      const link = e.target.matches("[data-link]")
        ? e.target
        : e.target.closest("[data-link]");
      router.navigateTo(link.getAttribute("href"));
    }
  });

  // Listen for browser back/forward navigation
  window.addEventListener("popstate", () => {
    router.handleLocation();
  });

  initThemeToggler();
  router.handleLocation();
});

// Router class to handle SPA navigation
class Router {
  constructor() {
    // Authentication requirements for various routes
    this.routes = {
      "/": {
        view: homeView,
        requiresAuth: false,
      },
      "/sign-up": {
        view: signUpView,
        requiresAuth: false,
      },
      "/sign-in": {
        view: signInView,
        requiresAuth: false,
      },
    };
  }

  navigateTo(url) {
    history.pushState(null, null, url);
    this.handleLocation();
  }

  // Handles the current browser location and renders the appropriate view.
  async handleLocation() {
    await renderHeader(this);

    const path = window.location.pathname;
    const route = this.routes[path];

    if (!route) {
      errorView(404, "Not Found");
      return;
    }

    // Check if route requires authentication
    if (route.requiresAuth) {
      const { authenticated } = await checkAuthStatus();
      if (!authenticated) {
        history.pushState(null, null, "/sign-in");
        await this.routes["/sign-in"].view();
        return;
      }
    }

    // Await the result of the view function to ensure it renders correctly
    if (route.view) {
      await route.view();
    }
    toggleHamburgerMenu();
  }
}
