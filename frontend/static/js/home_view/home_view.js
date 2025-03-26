import { categoriesView } from './categories.js';

// Renders the home view of the application.
export async function homeView() {
  const app = document.getElementById('app');

  // Set the inner HTML of the app element to display the home view content
  app.innerHTML = `
          <aside id="sidebar" class="sidebar"></aside>
  `;

  categoriesView();
}
