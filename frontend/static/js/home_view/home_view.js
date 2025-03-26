import { categoriesView } from './categories.js';

// Renders the home view of the application.
export async function homeView() {
  const app = document.getElementById('app');

  const aside = document.createElement('aside');
  aside.id = 'sidebar';
  aside.className = 'sidebar';

  app.parentNode.insertBefore(aside, app);

  categoriesView();

  // Set the inner HTML of the app element to display the home view content
  app.innerHTML = `
          <h1>Real Time Forum</h1>
  `;
}
