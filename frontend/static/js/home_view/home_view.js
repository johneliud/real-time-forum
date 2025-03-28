import { categoriesView } from "./categories.js";
import { postCreationView } from "./post_creation.js";
import { createPostBtn } from "./create_post_btn.js";

// Renders the home view of the application.
export async function homeView() {
  const app = document.getElementById("app");

  const aside = document.createElement("aside");
  aside.id = "sidebar";
  aside.className = "sidebar";

  app.parentNode.insertBefore(aside, app);

  categoriesView();

  postCreationView();

  createPostBtn();

}
