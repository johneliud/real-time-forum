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

async function initApp() {
    await Promise.all([
        loadStylesheet('/frontend/static/css/style.css'),
        loadStylesheet('/frontend/static/css/sign-up.css')
    ]);

    initSignup();
}

initApp().catch(console.error);