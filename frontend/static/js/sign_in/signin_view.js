import { initSigninValidation } from './signin_validation.js';

// Function to render the sign in view.
export async function signInView() {
  const app = document.getElementById('app');

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
  setTimeout(() => {
    initSigninValidation();
  }, 0);
}
