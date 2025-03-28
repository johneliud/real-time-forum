export async function postContainerView() {
  const app = document.getElementById("app");

  // Fetch posts
  try {
    const response = await fetch("/api/posts");
    if (!response.ok) {
      throw new Error("Failed to fetch posts");
    }

    const posts = await response.json();

    // Render each post
    posts.forEach((post) => {
      const article = document.createElement("article");
      article.className = "post";
      article.innerHTML = `
          <div class="post-header">
            <p class="post-author">@${post.UserName}</p>
            <p class="post-time">
              Posted: <time datetime="${post.CreatedOn}">${new Date(
        post.CreatedOn
      ).toLocaleString()}</time>
            </p>
          </div>
          <h3>${post.PostTitle}</h3>
          <p>${post.Body}</p>
  
          ${
            post.MediaURL
              ? `<img class="uploaded-file" src="${post.MediaURL}" alt="${post.PostTitle}" />`
              : ""
          }
  
          <div class="category-div">
            ${post.Categories.map(
              (category) =>
                `<p class="post-category"><span>${category.CategoryName}</span></p>`
            ).join("")}
          </div>
  
          <div class="post-actions" data-post-id="${post.ID}">
            <button
              data-posted-id="${post.ID}"
              class="like-button"
              data-reaction="Like"
              aria-label="Like this post"
            >
              <img
                class="icon"
                style="
                  height: 25px;
                  width: 1.2rem;
                  filter: invert(17%) sepia(27%) saturate(7051%)
                    hue-rotate(205deg) brightness(90%) contrast(99%);
                  margin-right: 5px;
                "
                src="/frontend/static/assets/thumbs-up-regular.svg"
                alt="thumbs-up-regular"
              />
              <span class="like-count">${post.Likes}</span>
            </button>
  
            <button
              data-posted-id="${post.ID}"
              class="dislike-button"
              data-reaction="Dislike"
              aria-label="Dislike this post"
            >
              <img
                class="icon"
                style="
                  height: 25px;
                  width: 1.2rem;
                  filter: invert(17%) sepia(27%) saturate(7051%)
                    hue-rotate(205deg) brightness(90%) contrast(99%);
                  margin-right: 5px;
                "
                src="/frontend/static/assets/thumbs-down-regular.svg"
                class="web-icon"
                alt="thumbs-down-regular"
              />
              <span class="dislike-count">${post.Dislikes}</span>
            </button>
  
            <button class="comment-button" aria-label="View or add comments">
              <img
                class="icon"
                style="
                  height: 25px;
                  width: 1.2rem;
                  filter: invert(17%) sepia(27%) saturate(7051%)
                    hue-rotate(205deg) brightness(90%) contrast(99%);
                  margin-right: 5px;
                "
                src="/frontend/static/assets/comment-regular.svg"
                class="web-icon"
                alt="comment-regular"
              />
              <span class="comment-count">${post.CommentCount}</span>
            </button>
          </div>
  
          <div class="comments-section">
            <h4>Comments</h4>
  
            <div class="comment-input">
              <form action="/comments" method="post">
                <input type="hidden" name="id" value="${post.ID}" />
                <input
                  type="text"
                  name="comment"
                  class="comment-box"
                  placeholder="Write a comment..." required
                />
                <button class="submit-comment">
                  <img
                    style="height: 20px; margin: 0"
                    src="/frontend/static/assets/paper-plane-regular.svg"
                    class="web-icon"
                    alt="paper-plane-regular"
                  />
                </button>
              </form>
            </div>
  
            ${post.Comments.map(
              (comment) => `
              <div class="comment" data-post-id="${comment.ID}">
                <p><strong>${comment.UserName}</strong>: ${comment.Body}</p>
                <div class="comment-actions">
                  <button
                    data-posted-id="${comment.ID}"
                    data-reaction="Like"
                    aria-label="Like this comment"
                    class="like-comment-button"
                  >
                    <img
                      class="icon"
                      style="
                        height: 25px;
                        width: 1.2rem;
                        filter: invert(17%) sepia(27%) saturate(7051%)
                          hue-rotate(205deg) brightness(90%) contrast(99%);
                      "
                      src="/frontend/static/assets/thumbs-up-regular.svg"
                      class="web-icon"
                      alt="thumbs-up-regular"
                    />
                    <span>${comment.Likes}</span>
                  </button>
                  <button
                    data-posted-id="${comment.ID}"
                    data-reaction="Dislike"
                    aria-label="Dislike this comment"
                    class="dislike-comment-button"
                  >
                    <img
                      class="icon"
                      style="
                        height: 25px;
                        width: 1.2rem;
                        filter: invert(17%) sepia(27%) saturate(7051%)
                          hue-rotate(205deg) brightness(90%) contrast(99%);
                      "
                      src="/frontend/static/assets/thumbs-down-regular.svg"
                      class="web-icon"
                      alt="thumbs-down-regular"
                    />
                    <span>${comment.Dislikes}</span>
                  </button>
                </div>
              </div>
            `
            ).join("")}
          </div>
        `;
      app.appendChild(article);
    });
  } catch (error) {
    console.error("Error fetching posts:", error);
    app.innerHTML = "<p>Failed to load posts. Please try again later.</p>";
  }
}
