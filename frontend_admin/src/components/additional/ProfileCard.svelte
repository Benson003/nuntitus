<script lang="ts">
    import { push } from "svelte-spa-router";
    import { logout } from "../../lib/authStore";
    import Button from "../base/Button.svelte";
    import PM from "../typography/PM.svelte";
    import PS from "../typography/PS.svelte";

    let {
        username = "username",
        email = "user@example.com",
        first_name = "First",
        last_name = "Last",
    } = $props<{
        username: string;
        email: string;
        first_name: string;
        last_name: string;
    }>();

    function handleEdit() {
        push("/settings");
    }

    function handleLogout() {
        logout();
    }
</script>

<div
    class="
        w-full max-w-xs p-4 m-4 rounded-2xl border border-gray-200 dark:border-gray-800
        bg-transparent
        relative left-3
        transition duration-300 ease-in-out
        hover:shadow-lg hover:shadow-green-200 dark:hover:shadow-green-900
    "
>
    <div class="flex flex-col space-y-1 text-left">
        <PM class="text-gray-900 dark:text-gray-100">
            {first_name}
            {last_name}
        </PM>
        <PS class="text-gray-600 dark:text-gray-400">
            @{username}
        </PS>
        <PS class="text-gray-700 dark:text-gray-300 truncate">
            {email}
        </PS>
    </div>

    <div class="flex space-x-2 mt-4">
        <Button
            onClick={handleEdit}
            class="flex-1 bg-blue-500 hover:bg-blue-600 text-white rounded px-3 py-1"
        >
            Edit
        </Button>
        <Button
            onClick={handleLogout}
            class="flex-1 bg-red-500 hover:bg-red-600 text-white rounded px-3 py-1"
        >
            Logout
        </Button>
    </div>
</div>
