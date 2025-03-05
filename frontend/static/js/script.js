document.addEventListener('DOMContentLoaded', () => {
  function applyTheme(theme) {
    if (theme === 'dark') {
      document.body.classList.add('dark-theme');
    } else {
      document.body.classList.remove('dark-theme');
    }
  }

  function toggleTheme() {
    let currentTheme = '';
    let newTheme = '';

    if (document.body.classList.contains('dark-theme')) {
      currentTheme = 'dark';
    } else {
      currentTheme = 'light';
    }

    if (currentTheme === 'dark') {
      newTheme = 'light';
    } else {
      newTheme = 'dark';
    }

    applyTheme(newTheme);
    localStorage.setItem('theme', newTheme);
  }
  const savedTheme = localStorage.getItem('theme') || 'light';
  applyTheme(savedTheme);

  const themeToggler = document.querySelector('.theme-toggler');

  if (themeToggler) {
    themeToggler.addEventListener('click', toggleTheme);
  }

  const hamburgerMenu = document.querySelector('.hamburger-menu');
  const menuContent = document.querySelector('.menu-content');

  if (hamburgerMenu) {
    hamburgerMenu.addEventListener('click', () => {
      hamburgerMenu.classList.toggle('active');
      menuContent.style.display =
        menuContent.style.display === 'block' ? 'none' : 'block';
    });

    document.addEventListener('click', function (event) {
      if (
        !hamburgerMenu.contains(event.target) &&
        !menuContent.contains(event.target)
      ) {
        menuContent.style.display = 'none';
        hamburgerMenu.classList.remove('active');
      }
    });
  }
});
