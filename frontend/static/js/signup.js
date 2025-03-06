export function initSignup() {
  const app = document.getElementById('app');

  function createElement(tag, attributes = {}, children = []) {
    const element = document.createElement(tag);

    Object.entries(attributes).forEach(([key, value]) => {
      if (key === 'class') {
        element.classList.add(...value.split(' '));
      } else if (key.startsWith('data-')) {
        element.dataset[
          key
            .replace('data-', '')
            .replace(/-([a-z])/g, (g) => g[1].toUpperCase())
        ] = value;
      } else {
        element[key] = value;
      }
    });

    children.forEach((child) => {
      if (typeof child === 'string') {
        element.appendChild(document.createTextNode(child));
      } else {
        element.appendChild(child);
      }
    });

    return element;
  }

  function createSignupForm() {
    const signupForm = createElement(
      'form',
      {
        id: 'signup-form',
        novalidate: true,
      },
      [
        // First Name
        createElement('div', { class: 'input-group' }, [
          createElement('label', { for: 'first-name' }, ['First Name']),
          createElement('input', {
            type: 'text',
            id: 'first-name',
            name: 'first-name',
            required: true,
          }),
        ]),
        // Last Name
        createElement('div', { class: 'input-group' }, [
          createElement('label', { for: 'last-name' }, ['Last Name']),
          createElement('input', {
            type: 'text',
            id: 'last-name',
            name: 'last-name',
            required: true,
          }),
        ]),
        // Nickname
        createElement('div', { class: 'input-group' }, [
          createElement('label', { for: 'nick-name' }, ['Nickname']),
          createElement('input', {
            type: 'text',
            id: 'nick-name',
            name: 'nick-name',
            required: true,
          }),
        ]),
        // Gender
        createElement('div', { class: 'input-group' }, [
          createElement('label', { for: 'gender' }, ['Gender']),
          createElement(
            'select',
            {
              id: 'gender',
              name: 'gender',
              required: true,
            },
            [
              createElement('option', { value: '' }, ['Select Gender']),
              createElement('option', { value: 'male' }, ['Male']),
              createElement('option', { value: 'female' }, ['Female']),
            ]
          ),
        ]),
        // Age
        createElement('div', { class: 'input-group' }, [
          createElement('label', { for: 'age' }, ['Age']),
          createElement('input', {
            type: 'number',
            id: 'age',
            name: 'age',
            required: true,
            min: '13',
            max: '120',
          }),
        ]),
        // Email
        createElement('div', { class: 'input-group' }, [
          createElement('label', { for: 'email' }, ['Email']),
          createElement('input', {
            type: 'email',
            id: 'email',
            name: 'email',
            required: true,
          }),
        ]),
        // Password
        createElement('div', { class: 'input-group' }, [
          createElement('label', { for: 'password' }, ['Password']),
          createElement('input', {
            type: 'password',
            id: 'password',
            name: 'password',
            required: true,
          }),
        ]),
        // Confirm Password
        createElement('div', { class: 'input-group' }, [
          createElement('label', { for: 'confirmed-password' }, [
            'Confirm Password',
          ]),
          createElement('input', {
            type: 'password',
            id: 'confirmed-password',
            name: 'confirmed-password',
            required: true,
          }),
        ]),
        // Submit Button
        createElement(
          'button',
          {
            type: 'submit',
            class: 'sign-up-btn',
          },
          ['Create Account']
        ),
      ]
    );

    return signupForm;
  }

  function attachFormValidation() {
    const form = document.getElementById('signup-form');
    const messagePopup = document.createElement('div');
    messagePopup.id = 'message-popup';
    form.prepend(messagePopup);

    form.addEventListener('submit', async (e) => {
      e.preventDefault();

      // Basic form validation
      if (!form.checkValidity()) {
        form.reportValidity();
        return;
      }

      // Prepare signup data
      const formData = new FormData(form);
      const signupData = {
        firstName: formData.get('first-name'),
        lastName: formData.get('last-name'),
        nickName: formData.get('nick-name'),
        gender: formData.get('gender'),
        age: parseInt(formData.get('age'), 10),
        email: formData.get('email'),
        password: formData.get('password'),
        confirmedPassword: formData.get('confirmed-password'),
      };

      try {
        const response = await fetch('/sign-up', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(signupData),
        });

        const result = await response.json();

        if (result.success) {
          messagePopup.textContent = 'Signup successful!';
          messagePopup.style.color = 'green';
          // Optional: Redirect or clear form
          setTimeout(() => {
            window.location.href = '/sign-in';
          }, 2000);
        } else {
          messagePopup.textContent = result.message || 'Signup failed';
          messagePopup.style.color = 'red';
        }
      } catch (error) {
        console.error('Signup error:', error);
        messagePopup.textContent = 'An error occurred. Please try again.';
        messagePopup.style.color = 'red';
      }
    });
  }

  // Render signup form
  function renderSignupPage() {
    app.innerHTML = ''; // Clear existing content

    const signupContainer = createElement(
      'div',
      { class: 'signup-container' },
      [createElement('h1', {}, ['Sign Up']), createSignupForm()]
    );

    app.appendChild(signupContainer);
    attachFormValidation();
  }

  // Initial render
  renderSignupPage();
}
