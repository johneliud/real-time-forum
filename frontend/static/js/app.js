import { initThemeToggler, toggleHamburgerMenu } from './script.js';
import { homeView } from './home_view.js';
import { signUpView } from './signup_view.js';
import { signInView } from './signin_view.js';
import { checkAuthStatus, logout } from './auth.js';

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
    // Define routes and their authentication requirements
    this.routes = {
      '/': {
        view: homeView,
        requiresAuth: false
      },
      '/sign-up': {
        view: signUpView,
        requiresAuth: false
      },
      '/sign-in': {
        view: signInView,
        requiresAuth: false
      },
      // Add additional protected routes here
      '/profile': {
        view: null, // Replace with your profile view when implemented
        requiresAuth: true
      }
    };
  }

  navigateTo(url) {
    history.pushState(null, null, url);
    this.handleLocation();
  }

  // Handles the current browser location and renders the appropriate view.
  async handleLocation() {
    // Render header first (this also checks auth status)
    await this.renderHeader();

    const path = window.location.pathname;
    const route = this.routes[path] || this.routes['/'];
    
    // Check if route requires authentication
    if (route.requiresAuth) {
      const { authenticated } = await checkAuthStatus();
      if (!authenticated) {
        // Redirect to sign-in page if not authenticated
        history.pushState(null, null, '/sign-in');
        // Call the view function for the sign-in page
        await this.routes['/sign-in'].view();
        return;
      }
    }

    // Await the result of the view function to ensure it renders correctly
    if (route.view) {
      await route.view();
    }
  }

  // Renders the header element for the application.
  async renderHeader() {
    // Check authentication status to show appropriate header options
    const { authenticated } = await checkAuthStatus();
    
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

          ${authenticated ? `
          <div class="user-profile">
            <box-icon name='user-circle'></box-icon>
            <p>Profile</p>
          </div>

          <div class="settings">
            <box-icon name='cog'></box-icon>
            <p>Settings</p>
          </div>
          
          <div class="logout" id="logout-btn">
            <box-icon name='log-out'></box-icon>
            <p>Logout</p>
          </div>
          ` : `
          <div class="auth-links">
            <a href="/sign-in" data-link>Sign In</a>
            <a href="/sign-up" data-link>Sign Up</a>
          </div>
          `}
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
    
    // Add logout functionality if user is authenticated
    if (authenticated) {
      const logoutBtn = document.getElementById('logout-btn');
      if (logoutBtn) {
        logoutBtn.addEventListener('click', async () => {
          const result = await logout();
          if (result.success) {
            // Redirect to home page after logout
            this.navigateTo('/sign-in');
            // Force a refresh of the application to update authentication state
            this.handleLocation();
          }
        });
      }
    }
  }
}

