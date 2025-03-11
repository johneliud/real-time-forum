export function initThemeToggler() {
  const savedTheme = localStorage.getItem('theme') || 'light';
  applyTheme(savedTheme);

  document.body.addEventListener('click', (e) => {
    if (e.target.closest('.theme-toggler')) {
      toggleTheme();
    }
  });

  function applyTheme(theme) {
    if (theme === 'dark') {
      document.body.classList.add('dark-theme');
    } else {
      document.body.classList.remove('dark-theme');
    }
  }

  function toggleTheme() {
    const currentTheme = document.body.classList.contains('dark-theme')
      ? 'dark'
      : 'light';
    const newTheme = currentTheme === 'dark' ? 'light' : 'dark';

    applyTheme(newTheme);
    localStorage.setItem('theme', newTheme);
  }
}

export function showMessage(message, isSuccess) {
  const messagePopup = document.getElementById('message-popup');

  messagePopup.textContent = message;
  messagePopup.classList.remove('success', 'error');

  messagePopup.classList.add('show', isSuccess ? 'success' : 'error');

  setTimeout(() => {
    messagePopup.classList.remove('show', 'success', 'error');
  }, 3000);
}

export function toggleHamburgerMenu() {
  const hamburgerMenu = document.querySelector('.hamburger-menu');
  const menuContent = document.querySelector('.menu-content');

  if (hamburgerMenu) {
    hamburgerMenu.addEventListener('click', () => {
      console.log("Clicked");
      
      hamburgerMenu.classList.toggle('active');
      menuContent.style.display =
        menuContent.style.display === 'block' ? 'none' : 'block';
    });
  }
}
