export async function createPostBtn() {
  const app = document.getElementById("app");

  const div = document.createElement("div");
  div.className = "floating-create-post-btn-container";
  div.innerHTML = `
        <p>Create Post</p>
        

        <button class="floating-create-post-btn">
          <box-icon name='plus' color='#ffffff' ></box-icon>
        </button>
    `;
  app.appendChild(div);
}
