// frontend/src/lib/api.ts
import { fetchWithAuth } from "./fetchWithAuth";

const API_BASE = "/api/v1";

// === AUTH ===

export async function loginUser(username_or_email: string, password: string) {
    const response = await fetch(`${API_BASE}/auth/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username_or_email, password }),
    });
    if (!response.ok) {
        throw new Error("Failed to login");
    }
    const data = await response.json();

    // Extract token from nested data object
    if (!data.data || !data.data.token) {
        throw new Error("Token not found in response");
    }

    return data.data.token; // return the JWT token string only
}

export async function signupUser(username: string, email: string, password: string) {
    const response = await fetch(`${API_BASE}/auth/signup`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, email, password }),
    });
    if (!response.ok) {
        throw new Error("Failed to signup");
    }
    const data = await response.json();

    if (!data.data || !data.data.token) {
        throw new Error("Token not found in signup response");
    }

    return data.data.token;  // Return just the token string
}



// === BLOGS ===

export async function fetchUserBlogs(page = 1, limit = 10) {
    const response = await fetchWithAuth(`${API_BASE}/blogs/user?page=${page}&limit=${limit}`);
    if (!response.ok) {
        throw new Error("Failed to fetch user blogs");
    }
    return await response.json(); // [{...}, {...}]
}

export async function fetchBlogById(id: string) {
    const response = await fetchWithAuth(`${API_BASE}/blogs/${id}`);
    if (!response.ok) {
        throw new Error("Failed to fetch blog");
    }
    return await response.json();
}
export async function uploadBlogContent(blogId: string, content: string) {
    const response = await fetchWithAuth(`${API_BASE}/blogs/${blogId}/upload`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ content }),
    });
    if (!response.ok) {
        throw new Error("Failed to upload blog content");
    }
    return await response.json();
}


export async function createBlog(blogData: {
    title: string;
    summary: string;
    visibility: boolean;
    publish_time: any; // ISO date string
}) {
    const response = await fetchWithAuth(`${API_BASE}/blogs`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(blogData),
    });
    if (!response.ok) {
        throw new Error("Failed to create blog");
    }
    return await response.json(); // Expect this to include blog_id in data
}

export async function updateBlog(id: string, blogData: {
    title: string;
    summary: string;
    content: string;
    visibility: boolean;
    publish_time: string; // ISO date string
}) {
    const response = await fetchWithAuth(`/api/v1/blogs/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(blogData),
    });
    if (!response.ok) {
        throw new Error("Failed to update blog");
    }
    return await response.json();
}

export async function deleteBlog(id: string) {
    const response = await fetchWithAuth(`${API_BASE}/blogs/${id}`, {
        method: "DELETE",
    });
    if (!response.ok) {
        throw new Error("Failed to delete blog");
    }
    return await response.json();
}
