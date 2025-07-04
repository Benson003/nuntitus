// frontend/src/lib/fetchWithAuth.ts
import { get } from "svelte/store";
import { jwt, logout } from "./authStore";

export async function fetchWithAuth(
    url: string,
    options: RequestInit = {}
): Promise<Response> {
    const token = get(jwt);
    const headers = new Headers(options.headers || {});
    if (token) {
        headers.set("Authorization", `Bearer ${token}`);
    }
    try {
        const response = await fetch(url, { ...options, headers });
        if (response.status === 401) {
            console.warn("Unauthorized, logging out");
            logout();
        }
        return response;
    } catch (e) {
        console.error("Network error:", e);
        throw e;
    }
}
