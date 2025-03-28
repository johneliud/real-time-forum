export async function createPostBtn() {
  const app = document.getElementById("app");

  const div = document.createElement("div");
  div.className = "floating-create-post-btn-container";
  div.innerHTML = `
        <p>Create Post</p>
        
        <button class="floating-create-post-btn">
          <img src="frontend/static/assets/plus-solid-24.png" alt="Create Post" />
        </button>
    `;
  app.appendChild(div);
}
