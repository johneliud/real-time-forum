import { initThemeToggler, toggleHamburgerMenu } from './script.js';
import { homeView } from './home_view.js';
import { signUpView } from './signup_view.js';
import { signInView } from './signin_view.js';

document.addEventListener('DOMContentLoaded', () => {
  const router = new Router();

  // Set up navigation events
  document.body.addEventListener('click', (e) => {
    if (e.target.matches('[data-link]') || e.target.closest('[data-link]')) {
      e.preventDefault();
      const link = e.target.matches('[data-link]')
        ? e.target
        : e.target.closest('[data-link]');
      router.navigateTo(link.getAttribute('href'));
    }
  });

  // Listen for browser back/forward navigation
  window.addEventListener('popstate', () => {
    router.handleLocation();
  });

  initThemeToggler();
  router.handleLocation();
  toggleHamburgerMenu();
});

// Router class to handle SPA navigation
class Router {
  constructor() {
    this.routes = {
      '/': homeView,
      '/sign-up': signUpView,
      '/sign-in': signInView,
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
    const view = this.routes[path] || this.routes['/'];

    // Await the result of the view function to ensure it renders correctly
    await view();
  }

  // Renders the header element for the application.
  renderHeader() {
    const header = document.createElement('header');
    header.innerHTML = `
        <nav class="navbar">
        <div class="logo"><a href="/">Real Time Forum</a></div>

        <div class="hamburger-menu">
          <div class="bar"></div>
          <div class="bar"></div>
          <div class="bar"></div>
        </div>

        <div class="menu-content">
          <div class="theme-toggler">
            <span class="tooltip-text">Toggle Mode</span>
            <box-icon class="sun" name="sun"></box-icon>
            <box-icon class="moon" name="moon"></box-icon>
          </div>

          <div class="user-profile">
            <box-icon name='user-circle'></box-icon>
            <p>Profile</p>
          </div>

          <div class="settings">
            <box-icon name='cog'  ></box-icon>
            <p>Settings</p>
          </div>
        </div>
      </nav>
    `;

    // Insert the header into the DOM before the app element or append to body
    const app = document.getElementById('app');
    if (app && app.parentNode) {
      app.parentNode.insertBefore(header, app);
    } else {
      document.body.appendChild(header);
    }
  }
}

