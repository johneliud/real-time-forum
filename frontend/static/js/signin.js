document
  .getElementById('signin-form')
  .addEventListener('submit', async function (e) {
    e.preventDefault();
    const formData = new FormData(this);
    const response = await fetch('/sign-in', {
      method: 'POST',
      body: formData,
    });

    if (response.ok) {
      window.location.href = '/home'; // Redirect on success
    } else {
      alert('Sign-in failed! Check your credentials.');
    }
  });
