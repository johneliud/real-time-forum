import { initSignupValidation } from "./signup_validation.js";
import { initSigninValidation } from "./signin_validation.js";
import { initThemeToggler } from "./script.js";

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
    this.routes = {
      "/": homeView,
      "/sign-up": signUpView,
      "/sign-in": signInView,
    };
  }

  navigateTo(url) {
    history.pushState(null, null, url);
    this.handleLocation();
  }

  // Handles the current browser location and renders the appropriate view.
  async handleLocation() {
    this.renderHeader();

    const path = window.location.pathname;
    const view = this.routes[path] || this.routes["/"];

    // Await the result of the view function to ensure it renders correctly
    await view();
  }

  // Renders the header element for the application.
  renderHeader() {
    const header = document.createElement("header");
    header.innerHTML = `
        <nav class="navbar">
            <div class="logo"><a href="/" data-link>Real Time Forum</a></div>
            <div class="theme-toggler">
                <span class="tooltip-text">Toggle Mode</span>
                <box-icon class="sun" name="sun"></box-icon>
                <box-icon class="moon" name="moon"></box-icon>
            </div>
        </nav>
    `;

    // Insert the header into the DOM before the app element or append to body
    const app = document.getElementById("app");
    if (app && app.parentNode) {
      app.parentNode.insertBefore(header, app);
    } else {
      document.body.appendChild(header);
    }
  }
}

// Renders the home view of the application.
async function homeView() {
  const app = document.getElementById("app");

  // Set the inner HTML of the app element to display the home view content
  app.innerHTML = `
      <div class="home-container">
          <h1>Welcome to Real Time Forum</h1>
      </div>
  `;
}

// Function to render the sign up view.
async function signUpView() {
  const app = document.getElementById("app");

  // Render sign up form
  app.innerHTML = `
      <p class="message-popup" id="message-popup"></p>
      <div class="form-container">
          <h2>Sign Up</h2>
          <form id="signup-form" novalidate>
              <div class="user-names">
                  <div class="input-group">
                      <label for="first-name">First Name</label>
                      <input type="text" id="first-name" name="first-name" required />
                  </div>

                  <div class="input-group">
                      <label for="last-name">Last Name</label>
                      <input type="text" id="last-name" name="last-name" required />
                  </div>
              </div>

              <div class="name-gender-age">
                  <div class="input-group">
                      <label for="nick-name">Nickname</label>
                      <input type="text" id="nick-name" name="nick-name" required />
                  </div>

                  <div class="input-group">
                      <label for="gender">Gender</label>
                      <select id="gender" name="gender" required>
                          <option value="">Select Gender</option>
                          <option value="male">Male</option>
                          <option value="female">Female</option>
                          <option value="prefer not to say">Prefer not to say</option>
                      </select>
                  </div>

                  <div class="input-group">
                      <label for="age">Age</label>
                      <input
                          type="number"
                          id="age"
                          name="age"
                          required
                          min="13"
                          max="120"
                      />
                  </div>
              </div>

              <div class="input-group">
                  <label for="email">Email</label>
                  <input type="email" id="email" name="email" required />
              </div>

              <div class="password">
                  <div class="input-group">
                      <label for="password">Password</label>
                      <div class="password-wrapper">
                          <input type="password" id="password" name="password" required />
                          <button
                              type="button"
                              class="toggle-password-visibility"
                              data-target="password"
                          >
                              <box-icon name="show"></box-icon>
                          </button>
                      </div>
                  </div>

                  <div class="input-group">
                      <label for="confirmed-password">Confirm Password</label>
                      <div class="password-wrapper">
                          <input
                              type="password"
                              id="confirmed-password"
                              name="confirmed-password"
                              required
                          />
                          <button
                              type="button"
                              class="toggle-password-visibility"
                              data-target="confirmed-password"
                          >
                              <box-icon name="show"></box-icon>
                          </button>
                      </div>
                  </div>
              </div>

              <button type="submit" class="sign-up-btn btn">Create Account</button>
          </form>

          <p class="switch-form">
              Already have an account? <a href="/sign-in" data-link>Sign In</a>
          </p>
      </div>
  `;

  // Initialize the sign up form validation
  initSignupValidation();
}

// Function to render the sign in view.
async function signInView() {
  const app = document.getElementById("app");

  // Render sign in form
  app.innerHTML = `
      <p class="message-popup" id="message-popup"></p>
      <div class="form-container">
        <h2>Sign In</h2>
        <form action="/sign-in" id="signin-form" method="POST">
          <div class="input-group">
            <label for="email-or-nickname">Email or Nickname</label>
            <input
              type="text"
              id="email-or-nickname"
              name="email-or-nickname"
              required
            />
          </div>

          <div class="input-group">
            <label for="password">Password</label>
            <div class="password-wrapper">
              <input type="password" id="password" name="password" required />
              <button
                type="button"
                class="toggle-password-visibility"
                data-target="password"
              >
                <box-icon name="show"></box-icon>
              </button>
            </div>
          </div>

          <div class="line"></div>
          <button type="submit" class="sign-in-btn btn">Sign In</button>
        </form>

        <p class="switch-form">
          Don't have an account? <a href="/sign-up">Sign Up</a>
        </p>
      </div>
  `;

  // Initialize the sign in form validation
  initSigninValidation();
}
