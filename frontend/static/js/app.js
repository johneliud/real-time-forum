import { initSignup } from './signup.js';

async function loadStylesheet(href) {
  return new Promise((resolve, reject) => {
    const link = document.createElement('link');
    link.rel = 'stylesheet';
    link.href = href;
    link.onload = () => resolve(link);
    link.onerror = reject;

    const stylesContainer = document.getElementById('dynamic-styles');
    stylesContainer.parentNode.insertBefore(link, stylesContainer);
  });
}

function navigateTo(path) {
  window.history.pushState({}, path, window.location.origin + path);
  router();
}

async function router() {
  const path = window.location.pathname;

  if (path === '/sign-up') {
    await loadStylesheet('/frontend/static/css/sign-up.css');
    initSignup();
  }
}

window.addEventListener('popstate', router);

document.addEventListener('DOMContentLoaded', () => {
  document.body.addEventListener('click', (e) => {
    if (e.target.matches('[data-link]')) {
      e.preventDefault();
      navigateTo(e.target.href);
    }
  });

  router();
});

initSignup().catch(console.error);
