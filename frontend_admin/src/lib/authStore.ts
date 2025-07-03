// frontend/src/lib/authStore.ts
import { push } from "svelte-spa-router";
import { writable } from "svelte/store";

const storedToken = localStorage.getItem("jwt");
const jwt = writable<string | null>(storedToken);
jwt.subscribe(value => {
    if (value) {
        localStorage.setItem("jwt", value);
    } else {
        localStorage.removeItem("jwt");
    }
});

function decodePayload(token: string) {
    try {
        const payload = JSON.parse(atob(token.split('.')[1]));
        return payload;
    } catch (e) {
        console.error("Invalid JWT payload", e);
        return null;
    }
}

function logout() {
    jwt.set(null);
    push("/login");
}

export { jwt, decodePayload, logout };
