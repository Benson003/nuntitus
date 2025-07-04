<script lang="ts">
    import Container from "../components/base/Container.svelte";
    import Loading from "../components/base/Loading.svelte";
    import Button from "../components/base/Button.svelte";
    import { onMount } from "svelte";
    import { fetchWithAuth } from "../lib/fetchWithAuth";
    import EmailInput from "../components/base/inputs/EmailInput.svelte";
    import TextInput from "../components/base/inputs/TextInput.svelte";
    import PasswordInput from "../components/base/inputs/PasswordInput.svelte";

    let loading = $state(true);
    let error = $state("");
    let success = $state("");

    let email = $state("");
    let username = $state("");
    let first_name = $state("");
    let last_name = $state("");
    let password = $state("");

    onMount(async () => {
        try {
            loading = true;
            const res = await fetchWithAuth("/api/v1/auth/user");
            if (!res.ok) throw new Error("Failed to fetch profile data");
            const data = await res.json();
            const user = data.data;

            email = user.email || "";
            username = user.username || "";
            first_name = user.first_name || "";
            last_name = user.last_name || "";
        } catch (e) {
            error = e instanceof Error ? e.message : "Failed to load profile";
        } finally {
            loading = false;
        }
    });

    async function handleUpdate() {
        try {
            loading = true;
            error = "";
            success = "";

            const updateData: any = {
                email,
                username,
                first_name,
                last_name,
            };

            if (password.trim() !== "") {
                updateData.password = password;
            }

            const res = await fetchWithAuth("/api/v1/auth/user", {
                method: "PUT",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(updateData),
            });

            if (!res.ok) throw new Error("Failed to update profile");
            success = "Profile updated successfully.";
            password = ""; // clear after submission
        } catch (e) {
            error = e instanceof Error ? e.message : "Failed to update profile";
        } finally {
            loading = false;
        }
    }
</script>

<Container
    class="min-h-screen w-full flex flex-col items-center justify-start p-6 bg-white dark:bg-gray-950 overflow-auto"
>
    <div
        class="w-full max-w-2xl p-6 bg-gray-50 dark:bg-gray-900 rounded-lg shadow space-y-6"
    >
        <h1 class="text-3xl font-bold text-gray-900 dark:text-gray-100">
            Settings
        </h1>

        {#if loading}
            <div class="flex justify-center">
                <Loading size={8} intervalMs={150} />
            </div>
        {:else}
            <div class="flex flex-col space-y-4">
                {#if error}
                    <p class="text-red-600 dark:text-red-400">{error}</p>
                {/if}

                {#if success}
                    <p class="text-green-600 dark:text-green-400">{success}</p>
                {/if}

                <EmailInput
                    label="Ëmail"
                    value={email}
                    onInput={(v) => (email = v)}
                ></EmailInput>
                <TextInput
                    label="Username"
                    value={username}
                    onInput={(v) => (username = v)}
                ></TextInput>
                <TextInput
                    label="First Name"
                    value={first_name}
                    onInput={(v) => (first_name = v)}
                ></TextInput>

                <TextInput
                    label="Last Name"
                    value={last_name}
                    onInput={(v) => (last_name = v)}
                ></TextInput>

                <PasswordInput
                    label="Password (optional)"
                    value={password}
                    onInput={(v) => (password = v)}
                ></PasswordInput>

                <Button
                    class="mt-4 bg-blue-600 hover:bg-blue-700 text-white"
                    onClick={handleUpdate}
                    disabled={loading}
                >
                    {#if loading}
                        <Loading size={4} />
                    {:else}
                        Update Profile
                    {/if}
                </Button>
            </div>
        {/if}
    </div>
</Container>
