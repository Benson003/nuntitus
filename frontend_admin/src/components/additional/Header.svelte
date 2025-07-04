<script lang="ts">
    import { jwt, logout } from "../../lib/authStore";
    import Links from "../base/Links.svelte";
    import Button from "../base/Button.svelte";

    // derived store to track if logged in
    // no need to manually subscribe if using $ shorthand in template
</script>

<nav
    class="flex items-center justify-between p-2 bg-gray-100 dark:bg-gray-900 text-gray-900 dark:text-gray-100 w-full sticky top-0 z-50 shadow"
>
    <div class="text-xl font-bold">Nuntius</div>
    <div class="space-x-4 flex items-center">
        {#if $jwt}
            <Links href="/#/dashboard" class="hover:underline">Dashboard</Links>
            <Links href="/#/upload" class="hover:underline">Upload</Links>
            <Links href="/#/settings" class="hover:underline">Settings</Links>
            <Button
                onClick={logout}
                class="bg-red-500 hover:bg-red-600 text-white px-3 py-1 rounded"
            >
                Logout
            </Button>
        {:else}
            <Links href="/#/login" class="hover:underline">Login</Links>
            <Links href="/#/signup" class="hover:underline">Sign Up</Links>
        {/if}
    </div>
</nav>
