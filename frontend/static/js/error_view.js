// errorView dynamically updates the HTML structure with the encountered error.
export function errorView(code, message) {
  const app = document.getElementById('app');

  app.innerHTML = `
    <div class="error-details">
      <pre class="status-code">${code}</pre>
      <pre class="status-msg">${message}</pre>
      <a class="back-btn" href="/" onclick="event.preventDefault(); router.navigateTo('/');">Back Home</a>
  </div>
  `;

  document.title = `${code} - ${message}`;
}
