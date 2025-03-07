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

  // Initial route handling
  router.handleLocation();
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

  async handleLocation() {
    const path = window.location.pathname;
    const view = this.routes[path] || this.routes['/'];

    const app = document.getElementById('app');
    app.innerHTML = '';

    // Render navigation bar
    this.renderHeader();

    // Render the view content
    await view();
  }

  renderHeader() {
    const header = document.createElement('header');
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

    document.body.insertBefore(header, document.getElementById('app'));
  }
}

function initThemeToggler() {
  const savedTheme = localStorage.getItem('theme') || 'light';
  applyTheme(savedTheme);

  // Set up toggle event
  document.body.addEventListener('click', (e) => {
    if (e.target.closest('.theme-toggler')) {
      toggleTheme();
    }
  });

  function applyTheme(theme) {
    if (theme === 'dark') {
      document.body.classList.add('dark-theme');
    } else {
      document.body.classList.remove('dark-theme');
    }
  }

  function toggleTheme() {
    const currentTheme = document.body.classList.contains('dark-theme')
      ? 'dark'
      : 'light';
    const newTheme = currentTheme === 'dark' ? 'light' : 'dark';

    applyTheme(newTheme);
    localStorage.setItem('theme', newTheme);
  }
}

// View functions
async function homeView() {
  const app = document.getElementById('app');
  app.innerHTML = `
    <div class="home-container">
      <h1>Welcome to Real Time Forum</h1>
    </div>
  `;
}

async function signUpView() {
  const app = document.getElementById('app');

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

  initSignupValidation();
}

async function signInView() {
  const app = document.getElementById('app');

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

  initSigninForm();
}
