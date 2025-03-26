import { showMessage } from '../script.js';

export function initSigninValidation() {
  const signinForm = document.getElementById('signin-form');
  const messagePopup = document.getElementById('message-popup');

  if (!signinForm || !messagePopup) {
    console.error('Required elements not found');
    return;
  }

  // Password visibility toggle
  document.querySelectorAll('.toggle-password-visibility').forEach((button) => {
    button.addEventListener('click', () => {
      const input = document.getElementById(button.dataset.target);
      input.type = input.type === 'password' ? 'text' : 'password';
    });
  });

  let isSubmitting = false;

  // Form submission
  signinForm.addEventListener('submit', async (e) => {
    e.preventDefault();

    if (isSubmitting) return;
    isSubmitting = true;

    if (messagePopup) {
      messagePopup.textContent = '';
      messagePopup.classList.remove('show', 'success', 'error');
    }

    const identifier = document.getElementById('email-or-nickname').value;
    const password = document.getElementById('password').value;

    // Validate form before submission
    if (!signinForm.checkValidity()) {
      showMessage('Please check your form values and try again!', false);
      isSubmitting = false;
      return;
    }

    // Prepare signin data
    const signinData = {
      identifier: identifier,
      password: password,
    };

    try {
      const response = await fetch('/api/sign-in', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'X-Requested-With': 'XMLHttpRequest',
        },
        body: JSON.stringify(signinData),
      });

      const result = await response.json();

      if (!result.success) {
        showMessage(result.message || 'Sign in failed.', false);
        isSubmitting = false;
        return;
      } else if (result.success) {
        sessionStorage.setItem('session_token', result.sessionToken);
        showMessage(result.message || 'Sign in successful!', true);
        signinForm.reset();

        setTimeout(() => {
          window.location.href = '/';
        }, 1000);
      }
    } catch (error) {
      console.error('Signin error:', error);
      showMessage('An unexpected error occurred. Please try again.', false);
      isSubmitting = false;
    }
  });
}
