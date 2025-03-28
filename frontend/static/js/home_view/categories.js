// Renders the categories view of the application.
export async function categoriesView() {
  const sidebar = document.getElementById("sidebar");

  if (sidebar) {
    sidebar.innerHTML = `
    <h2>Filter By:</h2>
      <form class="filter-form" action="/filter" method="get">
        <fieldset>
          <label
            ><input type="checkbox" name="category" value="Technology" />
            Technology</label
          >
          <label
            ><input type="checkbox" name="category" value="Health" />
            Health</label
          >
          <label
            ><input type="checkbox" name="category" value="Education" />
            Education</label
          >
          <label
            ><input type="checkbox" name="category" value="Sports" />
            Sports</label
          >
          <label
            ><input type="checkbox" name="category" value="Entertainment" />
            Entertainment</label
          >
          <label
            ><input type="checkbox" name="category" value="Finance" />
            Finance</label
          >
          <label
            ><input type="checkbox" name="category" value="Travel" />
            Travel</label
          >
          <label
            ><input type="checkbox" name="category" value="Food" /> Food</label
          >
          <label
            ><input type="checkbox" name="category" value="Lifestyle" />
            Lifestyle</label
          >
          <label
            ><input type="checkbox" name="category" value="Science" />
            Science</label
          >
        </fieldset>

        <button class="apply">Apply Filter</button>
      </form>

      <form class="filter-form" action="/filter" method="get">
        <ul class="sidebar-links">
          <li>
            <button type="submit" name="filter" value="created">Created Posts</button>
          </li>
          <li>
            <button type="submit" name="filter" value="liked">Liked Posts</button>
          </li>
        </ul>
      </form>
    `;
  }
}
