const applyTheme = (theme) => {
  if (theme === 'dark') {
    document.body.classList.add('dark-theme');
  } else {
    document.body.classList.remove('dark-theme');
  }
};

// toggleTheme uses the value stored in the browsers local storage to determine the current theme.
const toggleTheme = () => {
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
};
const savedTheme = localStorage.getItem('theme') || 'light';
applyTheme(savedTheme);

document.querySelector('.theme-toggler').addEventListener('click', toggleTheme);
