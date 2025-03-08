export function initThemeToggler() {
    const savedTheme = localStorage.getItem("theme") || "light";
    applyTheme(savedTheme);
  
    document.body.addEventListener("click", (e) => {
      if (e.target.closest(".theme-toggler")) {
        toggleTheme();
      }
    });
  
    function applyTheme(theme) {
      if (theme === "dark") {
        document.body.classList.add("dark-theme");
      } else {
        document.body.classList.remove("dark-theme");
      }
    }
  
    function toggleTheme() {
      const currentTheme = document.body.classList.contains("dark-theme")
        ? "dark"
        : "light";
      const newTheme = currentTheme === "dark" ? "light" : "dark";
  
      applyTheme(newTheme);
      localStorage.setItem("theme", newTheme);
    }
  }