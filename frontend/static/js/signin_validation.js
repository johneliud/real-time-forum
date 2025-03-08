export function initSigninValidation() {
  const signinForm = document.getElementById("signin-form");
  const messagePopup = document.getElementById("message-popup");

  if (!signinForm || !messagePopup) {
    console.error("Required elements not found");
    return;
  }

  // Password visibility toggle
  document.querySelectorAll(".toggle-password-visibility").forEach((button) => {
    button.addEventListener("click", () => {
      const input = document.getElementById(button.dataset.target);
      input.type = input.type === "password" ? "text" : "password";
    });
  });

  function showMessage(message, isSuccess) {
    messagePopup.textContent = message;
    messagePopup.classList.remove("success", "error");

    messagePopup.classList.add("show", isSuccess ? "success" : "error");

    setTimeout(() => {
      messagePopup.classList.remove("show", "success", "error");
    }, 3000);
  }

  // Form submission
  signinForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    // Clear previous messages
    if (messagePopup) {
      messagePopup.textContent = "";
      messagePopup.style.display = "none";
    }

    const identifier = document.getElementById("email-or-nickname").value;
    const password = document.getElementById("password").value;

    // Validate form before submission
    if (!signinForm.checkValidity()) {
      showMessage("Please check your form values again!", "error");
      return;
    }

    // Prepare signin data
    const signinData = {
      identifier: identifier,
      password: password,
    };

    try {
      const response = await fetch("/api/sign-in", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(signinData),
      });

      const result = await response.json();

      if (result.success) {
        showMessage(result.message || "Sign in successful!", "success");

        // Store user data in localStorage
        localStorage.setItem("user", JSON.stringify(result.userData));

        // Store token if not using HttpOnly cookies
        if (result.token) {
          localStorage.setItem("token", result.token);
        }

        // Redirect after successful signin
        setTimeout(() => {
          window.location.href = "/";
        }, 2000);
      } else {
        showMessage(result.message || "Sign in failed.", "error");
      }
    } catch (error) {
      console.error("Signin error:", error);
      showMessage("An unexpected error occurred. Please try again.", "error");
    }
  });
}
