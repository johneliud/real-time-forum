// Renders the home view of the application.
export async function homeView() {
  const app = document.getElementById('app');

  // Set the inner HTML of the app element to display the home view content
  app.innerHTML = `
      <div class="home-container">
          <h1>Welcome to Real Time Forum</h1>
      </div>
  `;
}