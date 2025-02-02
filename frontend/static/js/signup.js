document
  .getElementById('signup-form')
  .addEventListener('submit', async function (e) {
    e.preventDefault();
    const formData = new FormData(this);

    try {
      const response = await fetch('/sign-up', {
        method: 'POST',
        body: formData,
      });

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(errorText);
      }

      window.location.href = '/sign-in'; // Redirect on success
    } catch (error) {
      alert('Signup failed! ' + error.message);
    }
  });

document.getElementById('signup-form').addEventListener('submit', function (e) {
  e.preventDefault();
  document.getElementById('success-message').style.display = 'block';
  this.reset(); // Clear form fields

  setTimeout(() => {
    document.getElementById('success-message').style.display = 'none';
  }, 3000);
});
