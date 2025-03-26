import { initSignupValidation } from './signup_validation.js';

// Function to render the sign up view.
export async function signUpView() {
  const app = document.getElementById('app');

  // Render sign up form
  app.innerHTML = `
      <p class="message-popup" id="message-popup"></p>
      <div class="form-container">
          <h2>Sign Up</h2>
          <form action="/sign-up" id="signup-form" method="POST" >
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
  setTimeout(() => {
    initSignupValidation();
  }, 0);
}