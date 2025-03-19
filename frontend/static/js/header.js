import { checkAuthStatus, logout } from './auth.js';

// Renders the header element for the application.
export async function renderHeader(router) {
  // Check authentication status to show appropriate header options
  const { authenticated, username } = await checkAuthStatus();

  // Create header element
  const header = document.createElement('header');
  header.innerHTML = `
    <nav class="navbar">
      <div class="logo"><a href="/">Real Time Forum</a></div>

      <div class="hamburger-menu">
        <div class="bar"></div>
        <div class="bar"></div>
        <div class="bar"></div>
      </div>

      <div class="menu-content">
        <div class="theme-toggler">
          <span class="tooltip-text">Toggle Mode</span>
          <box-icon class="sun" name="sun"></box-icon>
          <box-icon class="moon" name="moon"></box-icon>
        </div>

        ${
          authenticated
            ? `
        <div class="user-profile">
          <box-icon name='user-circle'></box-icon>
          <p>Profile</p>
        </div>

        <div class="inbox">
          <box-icon name='envelope'></box-icon>
          <p>Inbox</p>
        </div>

        <div class="settings">
          <box-icon name='cog'></box-icon>
          <p>Settings</p>
        </div>
        
        <div class="log-out" id="logout-btn">
          <box-icon name='log-out'></box-icon>
          <p>Logout</p>
        </div>
        `
            : ``
        }
      </div>
    </nav>
  `;

  // Insert the header into the DOM
  const app = document.getElementById('app');
  if (app && app.parentNode) {
    app.parentNode.insertBefore(header, app);
  } else {
    document.body.appendChild(header);
  }

  // Add logout functionality if user is authenticated
  if (authenticated) {
    const logoutBtn = document.getElementById('logout-btn');
    if (logoutBtn) {
      logoutBtn.addEventListener('click', async () => {
        const result = await logout();
        if (result.success) {
          // Redirect to sign in after logout
          router.navigateTo('/sign-in');
          router.handleLocation();
        }
      });
    }
  }

  // Add click event listener for the inbox
  const inbox = header.querySelector('.inbox');
  if (inbox) {
    inbox.addEventListener('click', async () => {
      // Create a modal for displaying messages
      const modal = document.createElement('div');
      modal.classList.add('chat-modal');
      modal.innerHTML = `<div class='modal-content'>
            <span class='close'>&times;</span>
            <h2>Chat Messages</h2>
            <div class='message-list'></div>
        </div>`;

      document.body.appendChild(modal);

      // Fetch messages from the server
      const response = await fetch('/api/messages');
      if (!response.ok) {
        console.error('Failed to fetch messages:', response.statusText);
        return;
      }
      const messages = await response.json();

      console.log('Fetched messages:', messages);

      const messageList = modal.querySelector('.message-list');
      if (!messageList) {
        console.error('Message list not found');
        return;
      }
      if (messages && messages.length > 0) {
        messages.forEach((message) => {
          const messageItem = document.createElement('div');
          messageItem.textContent = `${message.sender}: ${message.content}`;
          messageList.appendChild(messageItem);
        });
      } else {
        const noMessagesItem = document.createElement('div');
        noMessagesItem.textContent = 'No messages available.';
        messageList.appendChild(noMessagesItem);
      }

      // Close modal functionality
      modal.querySelector('.close').onclick = function () {
        modal.remove();
      };
    });
  }

  const profileDiv = header.querySelector('.user-profile');
  if (profileDiv) {
    profileDiv.addEventListener('click', async () => {
      const modal = document.createElement('div');
      modal.classList.add('profile-modal');
      modal.innerHTML = `<div class='modal-content'>
          <span class='close'>&times;</span>
          <h2>User Profile</h2>
          <div class='profile-details'></div>
          <img id='profileImagePreview' src='' alt='Profile Image' style='display:none;' />
          <input type='file' id='profileImage' accept='image/*' />
          <button id='uploadImage'>Upload Image</button>
      </div>`;

      document.getElementById('app').appendChild(modal);

      // Close modal on click
      modal.querySelector('.close').onclick = function () {
        modal.remove();
      };

      // Fetch user profile data
      const response = await fetch('/api/profile');
      if (response.ok) {
        const profileData = await response.json();
        const profileDetailsDiv = modal.querySelector('.profile-details');
        profileDetailsDiv.innerHTML = `
          <p>Name: ${profileData.name}</p>
          <p>Email: ${profileData.email}</p>
        `;
        if (profileData.profileImage) {
          const profileImagePreview = modal.querySelector(
            '#profileImagePreview'
          );
          profileImagePreview.src = profileData.profileImage;
          profileImagePreview.style.display = 'block';
        }
      } else {
        console.error('Failed to fetch profile data');
      }

      // Handle image upload
      const uploadButton = modal.querySelector('#uploadImage');
      uploadButton.addEventListener('click', async () => {
        const fileInput = modal.querySelector('#profileImage');
        const file = fileInput.files[0];
        if (file) {
          const formData = new FormData();
          formData.append('profileImage', file);

          const uploadResponse = await fetch('/api/profile/image', {
            method: 'POST',
            body: formData,
          });

          if (uploadResponse.ok) {
            const result = await uploadResponse.json();
            const profileImagePreview = modal.querySelector(
              '#profileImagePreview'
            );
            profileImagePreview.src = result.profileImage;
            profileImagePreview.style.display = 'block';
          } else {
            console.error('Failed to upload image');
          }
        }
      });
    });
  }
}
