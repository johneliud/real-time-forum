// Renders the post creation menu
export async function postCreationView() {
  const app = document.getElementById("app");

  const section = document.createElement("section");
  section.className = "create-post hidden";
  section.innerHTML = `
        <h2>Create a New Post</h2>
        <form
          name="upload"
          enctype="multipart/form-data"
          action="/upload"
          method="POST"
        >
          <label for="post-title">Title</label>
          <input
            type="text"
            id="post-title"
            name="post-title"
            placeholder="Enter your post title"
            required
          />

          <label for="post-content">Content</label>
          <textarea
            id="post-content"
            name="post-content"
            placeholder="Write your post here..."
            required
          ></textarea>

          <fieldset class="categories" name="categories">
            <legend>Select Category</legend>
            <label>
              <input type="checkbox" name="category[]" value="Technology" />
              Technology
            </label>
            <label>
              <input type="checkbox" name="category[]" value="Health" />
              Health
            </label>
            <label>
              <input type="checkbox" name="category[]" value="Education" />
              Education
            </label>
            <label>
              <input type="checkbox" name="category[]" value="Sports" />
              Sports
            </label>
            <label>
              <input type="checkbox" name="category[]" value="Entertainment" />
              Entertainment
            </label>
            <label>
              <input type="checkbox" name="category[]" value="Finance" />
              Finance
            </label>
            <label>
              <input type="checkbox" name="category[]" value="Travel" />
              Travel
            </label>
            <label>
              <input type="checkbox" name="category[]" value="Food" />
              Food
            </label>
            <label>
              <input type="checkbox" name="category[]" value="Lifestyle" />
              Lifestyle
            </label>
            <label>
              <input type="checkbox" name="category[]" value="Science" />
              Science
            </label>
          </fieldset>

          <div class="post-operations">
            <input type="file" name="uploaded-file" />

            <button style="color: white;" type="submit">Post</button>
          </div>
        </form>
    `;
  app.appendChild(section);
}
