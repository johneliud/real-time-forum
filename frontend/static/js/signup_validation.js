import { showMessage } from "./script.js";

export function initSignupValidation() {
  const signupForm = document.getElementById("signup-form");
  const messagePopup = document.getElementById("message-popup");

  if (!signupForm || !messagePopup) {
    console.error("Required elements not found");
    return;
  }

  const firstNameInput = document.getElementById("first-name");
  const lastNameInput = document.getElementById("last-name");
  const nickNameInput = document.getElementById("nick-name");
  const genderInput = document.getElementById("gender");
  const ageInput = document.getElementById("age");
  const emailInput = document.getElementById("email");
  const passwordInput = document.getElementById("password");
  const confirmPasswordInput = document.getElementById("confirmed-password");

  // Create and attach feedback elements
  function createFeedbackElement(parentNode) {
    const feedbackElement = document.createElement("p");
    feedbackElement.className = "feedback-message";
    parentNode.appendChild(feedbackElement);
    return feedbackElement;
  }

  const nickNameFeedback = createFeedbackElement(nickNameInput.parentNode);
  const emailFeedback = createFeedbackElement(emailInput.parentNode);

  // Debounce function to prevent excessive calls
  function debounce(func, delay) {
    let timeout;
    return (...args) => {
      clearTimeout(timeout);
      timeout = setTimeout(() => func(...args), delay);
    };
  }

  // Check availability of nickname and email
  async function checkAvailability(field, value) {
    if (!value.trim()) return null;

    try {
      const response = await fetch(
        `/api/validate?${field}=${encodeURIComponent(value)}`
      );
      const data = await response.json();
      return data.available;
    } catch (error) {
      console.error("Error validating input:", error);
      return null;
    }
  }

  // Event listeners for availability checks
  nickNameInput.addEventListener(
    "input",
    debounce(async () => {
      const isAvailable = await checkAvailability(
        "nick-name",
        nickNameInput.value
      );
      if (isAvailable !== null) {
        nickNameFeedback.textContent = isAvailable
          ? "Nickname is available"
          : "Nickname is taken";
        nickNameFeedback.style.color = isAvailable ? "green" : "red";
      }
    }, 1000)
  );

  emailInput.addEventListener(
    "input",
    debounce(async () => {
      const isAvailable = await checkAvailability("email", emailInput.value);
      if (isAvailable !== null) {
        emailFeedback.textContent = isAvailable
          ? "Email is available"
          : "Email is taken";
        emailFeedback.style.color = isAvailable ? "green" : "red";
      }
    }, 1000)
  );

  // Password strength validation
  function validatePasswordStrength(password) {
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

  // Password validation
  passwordInput.addEventListener("input", () => {
    const passwordError = validatePasswordStrength(passwordInput.value);
    passwordInput.setCustomValidity(passwordError);
    passwordInput.reportValidity();
  });

  confirmPasswordInput.addEventListener("input", () => {
    if (passwordInput.value !== confirmPasswordInput.value) {
      confirmPasswordInput.setCustomValidity("Passwords do not match.");
    } else {
      confirmPasswordInput.setCustomValidity("");
    }
    confirmPasswordInput.reportValidity();
  });

  // Password visibility toggle
  document.querySelectorAll(".toggle-password-visibility").forEach((button) => {
    button.addEventListener("click", () => {
      const input = document.getElementById(button.dataset.target);
      input.type = input.type === "password" ? "text" : "password";
    });
  });

  // Form submission
  signupForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    // Clear previous messages
    if (messagePopup) {
      messagePopup.textContent = "";
      messagePopup.classList.remove("show", "success", "error");
    }

    // Validate form before submission
    if (!signupForm.checkValidity()) {
      showMessage("Please check your form values again!", false);
      return;
    }

    // Prepare signup data
    const signupData = {
      firstName: firstNameInput.value,
      lastName: lastNameInput.value,
      nickName: nickNameInput.value,
      gender: genderInput.value,
      age: parseInt(ageInput.value, 10),
      email: emailInput.value,
      password: passwordInput.value,
      confirmedPassword: confirmPasswordInput.value,
    };

    try {
      const response = await fetch("/api/sign-up", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(signupData),
      });

      const result = await response.json();

      if (result.success) {
        showMessage(result.message || "Sign up successful!", true);

        // Redirect after successful signup
        setTimeout(() => {
          window.location.href = "/sign-in";
        }, 2000);
      } else {
        showMessage(result.message || "Sign up failed.", false);
      }
    } catch (error) {
      console.error("Signup error:", error);
      showMessage("An unexpected error occurred. Please try again.", false);
    }
  });
}
