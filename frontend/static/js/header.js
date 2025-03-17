import { checkAuthStatus, logout } from './auth.js';

// Renders the header element for the application.
export async function renderHeader(router) {
  // Check authentication status to show appropriate header options
  const { authenticated } = await checkAuthStatus();

  // Create header element
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

        ${
          authenticated
            ? `
        <div class="user-profile">
          <box-icon name='user-circle'></box-icon>
          <p>Profile</p>
        </div>

        <div class="settings">
          <box-icon name='cog'></box-icon>
          <p>Settings</p>
        </div>
        
        <div class="log-out" id="logout-btn">
          <box-icon name='log-out'></box-icon>
          <p>Logout</p>
        </div>
        `
            : ``
        }
      </div>
    </nav>
  `;

  // Insert the header into the DOM
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
          // Redirect to sign in after logout
          router.navigateTo('/sign-in');
          router.handleLocation();
        }
      });
    }
  }
}
