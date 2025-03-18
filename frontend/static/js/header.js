import { checkAuthStatus, logout } from "./auth.js";

// Renders the header element for the application.
export async function renderHeader(router) {
  // Check authentication status to show appropriate header options
  const { authenticated, username } = await checkAuthStatus();

  // Create header element
  const header = document.createElement("header");
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
  const app = document.getElementById("app");
  if (app && app.parentNode) {
    app.parentNode.insertBefore(header, app);
  } else {
    document.body.appendChild(header);
  }

  // Add logout functionality if user is authenticated
  if (authenticated) {
    const logoutBtn = document.getElementById("logout-btn");
    if (logoutBtn) {
      logoutBtn.addEventListener("click", async () => {
        const result = await logout();
        if (result.success) {
          // Redirect to sign in after logout
          router.navigateTo("/sign-in");
          router.handleLocation();
        }
      });
    }
  }

  // Add click event listener for the inbox
  const inbox = header.querySelector(".inbox");
  if (inbox) {
    inbox.addEventListener("click", async () => {
      // Create a modal for displaying messages
      const modal = document.createElement("div");
      modal.classList.add("chat-modal");
      modal.innerHTML = `<div class='modal-content'>
            <span class='close'>&times;</span>
            <h2>Chat Messages</h2>
            <div class='message-list'></div>
        </div>`;

      document.body.appendChild(modal);

      // Fetch messages from the server
      const response = await fetch("/api/messages");
      if (!response.ok) {
        console.error("Failed to fetch messages:", response.statusText);
        return;
      }
      const messages = await response.json();

      console.log("Fetched messages:", messages);

      const messageList = modal.querySelector(".message-list");
      if (!messageList) {
        console.error("Message list not found");
        return;
      }
      if (messages && messages.length > 0) {
        messages.forEach((message) => {
          const messageItem = document.createElement("div");
          messageItem.textContent = `${message.sender}: ${message.content}`;
          messageList.appendChild(messageItem);
        });
      } else {
        const noMessagesItem = document.createElement("div");
        noMessagesItem.textContent = "No messages available.";
        messageList.appendChild(noMessagesItem);
      }

      // Close modal functionality
      modal.querySelector(".close").onclick = function () {
        modal.remove();
      };
    });
  }

  const profileDiv = header.querySelector('.user-profile');
  if (profileDiv) {
    profileDiv.addEventListener('click', async () => {
      // Create a modal for displaying profile details
      const modal = document.createElement('div');
      modal.classList.add('profile-modal');
      modal.innerHTML = `<div class='modal-content'>
          <span class='close'>&times;</span>
          <h2>User Profile</h2>
          <div class='profile-details'></div>
          <input type='file' id='profileImage' accept='image/*' />
          <button id='uploadImage'>Upload Image</button>
      </div>`;

      document.body.appendChild(modal);

      // Fetch user profile data
      const response = await fetch('/api/user/profile');
      console.log('Profile response:', response);
      if (response.ok) {
        const responseText = await response.text();
        console.log('Profile response body:', responseText);
        const userProfile = JSON.parse(responseText);
        const profileDetails = modal.querySelector('.profile-details');
        profileDetails.innerHTML = `
            <p>Username: ${userProfile.username}</p>
            <p>Email: ${userProfile.email}</p>
            <p>First Name: ${userProfile.firstName}</p>
            <p>Last Name: ${userProfile.lastName}</p>
            <p>Age: ${userProfile.age}</p>
            <p>Gender: ${userProfile.gender}</p>
        `;
      } else {
        console.error('Failed to fetch profile:', response.statusText);
      }

      // Handle image upload
      const uploadButton = modal.querySelector('#uploadImage');
      uploadButton.addEventListener('click', () => {
        const imageInput = modal.querySelector('#profileImage');
        const file = imageInput.files[0];
        if (file) {
          const formData = new FormData();
          formData.append('profileImage', file);
          fetch('/api/user/upload-image', {
            method: 'POST',
            body: formData,
          }).then(res => {
            if (res.ok) {
              console.log('Image uploaded successfully');
            } else {
              console.error('Image upload failed');
            }
          });
        }
      });

      // Close modal functionality
      modal.querySelector('.close').onclick = function() {
        modal.remove();
      };
    });
  }
}
