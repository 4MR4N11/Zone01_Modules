document.addEventListener("DOMContentLoaded", () => {
  // Handle new post submission
  document
    .getElementById("new_post_form")
    ?.addEventListener("submit", async (e) => {
      e.preventDefault();
      const form = e.target;
      const formData = {
        title: form.title.value,
        content: form.content.value,
        tags: form.tags.value.split(",").map((tag) => tag.trim()),
      };

      try {
        const response = await fetch("/api/post", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(formData),
          credentials: "include",
        });

        const data = await response.json();

        if (!response.ok)
          throw new Error(data.error || "Failed to create post");

        form.reset();
        prependNewPost(data.post);
      } catch (error) {
        showError(error.message);
      }
    });

  // Handle reactions (likes/dislikes)
  document.getElementById("posts").addEventListener("click", async (e) => {
    const likeBtn = e.target.closest(".like-up, .like-down");
    if (!likeBtn) return;

    e.preventDefault();
    const postElement = e.target.closest(".post");
    const postId = postElement.dataset.id;
    const isLike = likeBtn.classList.contains("like-up");

    try {
      const response = await fetch("/api/react", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          postId,
          reaction: isLike ? 1 : -1,
        }),
        credentials: "include",
      });

      const data = await response.json();

      if (!response.ok)
        throw new Error(data.error || "Failed to update reaction");

      updateReactionCounts(postElement, data.likes, data.dislikes);
    } catch (error) {
      showError(error.message);
    }
  });

  // Handle filtering
  document
    .getElementById("filter-form")
    .addEventListener("submit", async (e) => {
      e.preventDefault();
      const form = e.target;
      const params = new URLSearchParams({
        query: form.query.value,
        options: form.options.value,
      });

      try {
        const response = await fetch(`/api/filter?${params}`);
        const data = await response.json();

        if (!response.ok)
          throw new Error(data.error || "Failed to filter posts");

        updatePostsSection(data);
      } catch (error) {
        showError(error.message);
      }
    });

  // Handle pagination
  document.getElementById("posts").addEventListener("click", async (e) => {
    const paginationLink = e.target.closest(".pagination a");
    if (!paginationLink) return;

    e.preventDefault();
    const page = new URL(paginationLink.href).searchParams.get("page");

    try {
      const response = await fetch(`/api/posts?page=${page}`);
      const data = await response.json();

      if (!response.ok) throw new Error(data.error || "Failed to load posts");

      updatePostsSection(data);
    } catch (error) {
      showError(error.message);
    }
  });
});

function prependNewPost(post) {
  const postsContainer = document.getElementById("posts");
  const newPostHTML = createPostHTML(post);

  if (postsContainer.querySelector("#new_post_form")) {
    postsContainer.insertAdjacentHTML("beforeend", newPostHTML);
  } else {
    postsContainer.insertAdjacentHTML("afterbegin", newPostHTML);
  }
}

function updateReactionCounts(postElement, likes, dislikes) {
  postElement.querySelector(".like-up .like-count").textContent = likes;
  postElement.querySelector(".like-down .like-count").textContent = dislikes;
}

function updatePostsSection(data) {
  const postsContainer = document.getElementById("posts");

  // Clear existing posts and pagination
  postsContainer
    .querySelectorAll(".post, .pagination")
    .forEach((el) => el.remove());

  // Add new posts
  data.posts.forEach((post) => {
    postsContainer.insertAdjacentHTML("beforeend", createPostHTML(post));
  });

  // Add new pagination
  if (data.totalPages > 1) {
    postsContainer.insertAdjacentHTML("beforeend", createPaginationHTML(data));
  }
}

function createPostHTML(post) {
  return `
        <div data-id="${post.ID}" class="post flex gap-5">
            <div class="like flex flex-colum align-center justify-center">
                <a class="like-up flex flex-colum" href="#">
                    <i class="fa-solid fa-chevron-up"></i>
                    <span class="like-count">${post.Likes}</span>
                </a>
                <a class="like-down flex flex-colum text-center" href="#">
                    <span class="like-count">${post.Dislikes}</span>
                    <i class="fa-solid fa-chevron-down"></i>
                </a>
            </div>
            <div class="content">
                <a href="/post/${post.ID}" class="title">${post.Title}</a>
                <div class="info flex gap-2 align-center">
                    <i class="fa-regular fa-calendar-days"></i>
                    <span>${new Date(post.CreatedAt).toLocaleDateString(
                      "en-US",
                      {
                        year: "numeric",
                        month: "short",
                        day: "2-digit",
                      }
                    )}</span>
                    <p>By <span><a href="#">${post.Username}</a></span></p>
                </div>
                <div class="tags flex gap-2 mb-2">
                    ${post.Tags.map(
                      (tag) =>
                        `<a href="/filter?query=${tag}" class="tag">${tag}</a>`
                    ).join("")}
                </div>
                <p class="post-content">${post.Content}</p>
            </div>
        </div>
    `;
}

function createPaginationHTML(data) {
  return `
        <ul class="pagination flex gap-4 m-auto">
            ${
              data.currentPage > 1
                ? `<a href="?page=${data.currentPage - 1}">Previous</a>`
                : ""
            }
            ${
              data.currentPage < data.totalPages
                ? `<a href="?page=${data.currentPage + 1}">Next</a>`
                : ""
            }
        </ul>
    `;
}

function showError(message) {
  const errorDiv = document.createElement("div");
  errorDiv.className = "error notification";
  errorDiv.textContent = message;

  document.body.prepend(errorDiv);
  setTimeout(() => errorDiv.remove(), 5000);
}

function register() {
  const form = document.getElementById("register-form");
  form.addEventListener("submit", async (e) => {
    e.preventDefault();
    const formData = new FormData(form);
    let data = {
      email: form.get("email").trim(),
      username: form.get("username").trim(),
      age: form.get("age").trim(),
      firstname: form.get("firstname").trim(),
      lastname: form.get("lastname").trim(),
      password: form.get("password").trim(),
    };

    try {
      const response = await fetch("/api/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
        credentials: "include",
      });

      const responseData = await response.json();

      if (!response.ok)
        throw new Error(responseData.error || "Failed to register");

      window.location.href = "/";
    } catch (error) {
      showError(error.message);
    }
  });
}

function login() {
  const form = document.getElementById("login-form");
  form.addEventListener("submit", async (e) => {
    e.preventDefault();
    const formData = new FormData(form);
    let data = {
      username: form.get("username").trim(),
      password: form.get("password").trim(),
    };

    try {
      const response = await fetch("/api/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(data),
        credentials: "include",
      });

      const responseData = await response.json();

      if (!response.ok)
        throw new Error(responseData.error || "Failed to login");

      window.location.href = "/";
    } catch (error) {
      showError(error.message);
    }
  });
}
