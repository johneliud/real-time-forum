import { checkAuthStatus, logout } from './auth/auth.js';

// Renders the header element for the application.
export async function renderHeader(router) {
  // Check authentication status to show appropriate header options
  const { authenticated, username } = await checkAuthStatus();

  // Create header element
  const header = document.createElement('header');
  header.innerHTML = `
    <nav class="navbar">
      <div class="logo"><a href="/">Real Time Forum</a></div>

      <div class="menu-content">
        ${
          authenticated
            ? `
        <div class="user-profile">
          <span class="tooltip-text">Profile</span>
          <box-icon name='user-circle'></box-icon>
        </div>

        <div class="inbox">
          <span class="tooltip-text">Inbox</span>
          <box-icon name='envelope'></box-icon>
        </div>

        <div class="settings">
          <span class="tooltip-text">Settings</span>
          <box-icon name='cog'></box-icon>
        </div>
        
        <div class="log-out" id="logout-btn">
          <span class="tooltip-text">Log Out</span>
          <box-icon name='log-out'></box-icon>
        </div>

        <div class="theme-toggler">
          <span class="tooltip-text">Toggle Theme</span>
          <box-icon class="sun" name="sun"></box-icon>
          <box-icon class="moon" name="moon"></box-icon>
        </div>
        `
            : `
        <div style="visibility: none; opacity: 0;" class="user-profile">
          <span class="tooltip-text">Profile</span>
          <box-icon name='user-circle'></box-icon>
        </div>

        <div style="visibility: none; opacity: 0;" class="inbox">
          <span class="tooltip-text">Inbox</span>
          <box-icon name='envelope'></box-icon>
        </div>

        <div style="visibility: none; opacity: 0;" class="settings">
          <span class="tooltip-text">Settings</span>
          <box-icon name='cog'></box-icon>
        </div>
        
        <div style="visibility: none; opacity: 0;" class="log-out" id="logout-btn">
          <span class="tooltip-text">Log Out</span>
          <box-icon name='log-out'></box-icon>
        </div>

        <div class="theme-toggler">
          <span class="tooltip-text">Toggle Theme</span>
          <box-icon class="sun" name="sun"></box-icon>
          <box-icon class="moon" name="moon"></box-icon>
        </div>
            `
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

      modal.innerHTML = `
      <div class='modal-content'>
        <span class='close'>&times;</span>
        <h2>User Profile</h2>
        
        <div id="profileImageContainer" class="text-center">
          <img id='profileImagePreview' src='/images/default-avatar.png' alt='Profile Image' />
        </div>
        
        <div class='profile-details'>
          <p><strong>Loading profile data...</strong></p>
        </div>
        
        <div class="image-upload-container">
          <label for="profileImage" class="custom-file-upload">
            Choose Image
          </label>
          <input type='file' id='profileImage' accept='image/*' />
          <button id='uploadImage' disabled>Upload Image</button>
        </div>
      </div>
    `;

      document.body.appendChild(modal);

      const closeBtn = modal.querySelector('.close');
      closeBtn.onclick = function () {
        modal.remove();
      };

      modal.addEventListener('click', (e) => {
        if (e.target === modal) {
          modal.remove();
        }
      });

      // Fetch user profile data
      try {
        const response = await fetch('/api/profile');
        if (response.ok) {
          const profileData = await response.json();
          const profileDetailsDiv = modal.querySelector('.profile-details');

          // Display profile information
          profileDetailsDiv.innerHTML = `
          <p><strong>Username:</strong> ${username}</p>
          <p><strong>Name:</strong> ${profileData.name || 'Not set'}</p>
          <p><strong>Email:</strong> ${profileData.email || 'Not set'}</p>
          <p><strong>Member since:</strong> ${new Date(
            profileData.joinDate || Date.now()
          ).toLocaleDateString()}</p>
        `;

          if (profileData.profileImage) {
            const profileImagePreview = modal.querySelector(
              '#profileImagePreview'
            );
            profileImagePreview.src = profileData.profileImage;
          }
        } else {
          console.error('Failed to fetch profile data');
          const profileDetailsDiv = modal.querySelector('.profile-details');
          profileDetailsDiv.innerHTML = `<p>Error loading profile data. Please try again later.</p>`;
        }
      } catch (error) {
        console.error('Error fetching profile data:', error);
      }

      const fileInput = modal.querySelector('#profileImage');
      const uploadButton = modal.querySelector('#uploadImage');

      fileInput.addEventListener('change', () => {
        const file = fileInput.files[0];
        if (file) {
          uploadButton.disabled = false;

          // Show image preview
          const reader = new FileReader();
          reader.onload = (e) => {
            const profileImagePreview = modal.querySelector(
              '#profileImagePreview'
            );
            profileImagePreview.src = e.target.result;
          };
          reader.readAsDataURL(file);
        } else {
          uploadButton.disabled = true;
        }
      });

      // Handle image upload
      uploadButton.addEventListener('click', async () => {
        const file = fileInput.files[0];
        if (file) {
          // Show loading state
          uploadButton.textContent = 'Uploading...';
          uploadButton.disabled = true;

          const formData = new FormData();
          formData.append('profileImage', file);

          try {
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

              // Show success message
              alert('Profile image updated successfully!');
            } else {
              console.error('Failed to upload image');
              alert('Failed to upload image. Please try again.');
            }
          } catch (error) {
            console.error('Error uploading image:', error);
            alert('Error uploading image. Please try again.');
          } finally {
            // Reset button state
            uploadButton.textContent = 'Upload Image';
            uploadButton.disabled = false;
          }
        }
      });
    });
  }
}
