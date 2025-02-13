document
  .getElementById('signin-form')
  .addEventListener('submit', async function (e) {
    e.preventDefault();

    const formData = new URLSearchParams(new FormData(this));

    const response = await fetch('/sign-in', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      body: formData,
    });

    if (response.ok) {
      const successMessage = document.getElementById('success-message');
      successMessage.classList.add('show-success');
      setTimeout(() => {
        successMessage.classList.remove('show');
      }, 2000);
      window.location.href = '/';
    } else {
      alert('Sign-in failed! Check your credentials.');
    }
  });

// Toggle password visibility
document.querySelectorAll('.toggle-password-visibility').forEach((button) => {
  button.addEventListener('click', () => {
    const input = document.getElementById(button.dataset.target);
    if (input.type === 'password') {
      input.type = 'text';
    } else {
      input.type = 'password';
    }
  });
});
