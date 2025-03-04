document.addEventListener("DOMContentLoaded", () => {
  const nickName = document.getElementById("nick-name");
  const emailInput = document.getElementById("email");
  const passwordInput = document.getElementById("password");
  const confirmPasswordInput = document.getElementById("confirmed-password");
  const signupForm = document.getElementById("signup-form");

  // Feedback elements
  const nickNameFeedBack = document.createElement("p");
  nickNameFeedBack.className = "feedback-message";
  nickName.parentNode.appendChild(nickNameFeedBack);

  const emailFeedback = document.createElement("p");
  emailFeedback.className = "feedback-message";
  emailInput.parentNode.appendChild(emailFeedback);

  // Check credentials availability
  async function checkAvailability(field, value, feedbackElement) {
    if (!value.trim()) {
      feedbackElement.textContent = "";
      return;
    }

    try {
      const response = await fetch(
        `/validate?${field}=${encodeURIComponent(value)}`
      );
      const data = await response.json();

      if (data.available) {
        feedbackElement.textContent = `${
          field.charAt(0).toUpperCase() + field.slice(1)
        } is available`;
        feedbackElement.style.color = "green";
      } else {
        feedbackElement.textContent = `${
          field.charAt(0).toUpperCase() + field.slice(1)
        } is taken`;
        feedbackElement.style.color = "red";
      }
    } catch (error) {
      console.error("Error validating input:", error);
    }
  }

  // Prevent excessive calls using debounce
  function debounce(func, delay) {
    let timeout;
    return (...args) => {
      clearTimeout(timeout);
      timeout = setTimeout(() => func(...args), delay);
    };
  }

  nickName.addEventListener(
    "input",
    debounce(
      () => checkAvailability("nick-name", nickName.value, nickNameFeedBack),
      500
    )
  );
  emailInput.addEventListener(
    "input",
    debounce(
      () => checkAvailability("email", emailInput.value, emailFeedback),
      500
    )
  );

  function validatePasswordStength(password) {
    if (password.length < 8) return "Must contain at least 8 characters.";
    if (!/[A-Z]/.test(password))
      return "Include at least one uppercase letter.";
    if (!/[a-z]/.test(password))
      return "Include at least one lowercase letter.";
    if (!/[0-9]/.test(password)) return "Include at least one number.";
    if (!/[!,.:;(){}?_@#$%^&*]/.test(password))
      return "Include at least one special character.";
    return "";
  }

  // Show password strength validation
  passwordInput.addEventListener("input", () => {
    const passwordError = validatePasswordStength(passwordInput.value);
    passwordInput.setCustomValidity(passwordError);
    passwordInput.reportValidity();
  });

  // Confirm password validation
  confirmPasswordInput.addEventListener("input", () => {
    if (passwordInput.value !== confirmPasswordInput.value) {
      confirmPasswordInput.setCustomValidity("Passwords do not match.");
    } else {
      confirmPasswordInput.setCustomValidity("");
    }
    confirmPasswordInput.reportValidity();
  });

  signupForm.addEventListener("submit", (e) => {
    if (!signupForm.checkValidity()) {
      e.preventDefault();
    }
  });

  // Toggle password visibility
  document.querySelectorAll(".toggle-password-visibility").forEach((button) => {
    button.addEventListener("click", () => {
      const input = document.getElementById(button.dataset.target);
      if (input.type === "password") {
        input.type = "text";
      } else {
        input.type = "password";
      }
    });
  });
});
