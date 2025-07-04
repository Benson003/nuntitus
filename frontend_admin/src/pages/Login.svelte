<script lang="ts">
    import Button from "../components/base/Button.svelte";
    import Container from "../components/base/Container.svelte";
    import TextInput from "../components/base/inputs/TextInput.svelte";
    import PasswordInput from "../components/base/inputs/PasswordInput.svelte";
    import Links from "../components/base/Links.svelte";
    import Loading from "../components/base/Loading.svelte";

    import { push } from "svelte-spa-router";
    import { jwt } from "../lib/authStore";
    import { loginUser } from "../lib/api";

    import HXL from "../components/typography/HXL.svelte";
    import PL from "../components/typography/PL.svelte";
    import { onMount } from "svelte";

    let username = "";
    let password = "";
    let error: string | null = null;
    let loading = false;

    onMount(() => {
        if ($jwt) {
            push("/dashboard");
        }
    });

    async function handleSubmit() {
        error = null;
        if (username.trim() && password.trim()) {
            loading = true;
            try {
                const token = await loginUser(username, password);
                jwt.set(token);
                push("/dashboard");
            } catch (e) {
                console.error(e);
                error = "Invalid credentials or server error.";
            } finally {
                loading = false;
            }
        } else {
            error = "Please enter both username and password.";
        }
    }
</script>

<Container
    class="flex flex-col items-center justify-center p-6 w-full h-svh bg-white dark:bg-gray-950"
>
    <div
        class="w-full max-w-md p-8 rounded-xl shadow-xl bg-gray-100 dark:bg-gray-900"
    >
        <HXL class="text-center mb-6 text-gray-900 dark:text-gray-100">
            Welcome Back
        </HXL>

        {#if error}
            <PL class="text-center text-red-500 mb-2">{error}</PL>
        {/if}

        {#if loading}
            <div class="flex justify-center my-4">
                <Loading />
            </div>
        {/if}

        <div class="space-y-4">
            <TextInput
                label="Username"
                value={username}
                onInput={(v) => (username = v)}
                disabled={loading}
            />
            <PasswordInput
                label="Password"
                value={password}
                onInput={(v) => (password = v)}
                disabled={loading}
            />
        </div>

        <Button
            class="w-full mt-6 bg-green-600 hover:bg-green-700 dark:bg-green-700 dark:hover:bg-green-800 text-white transition-transform duration-150 hover:scale-105 active:scale-95"
            onClick={handleSubmit}
            disabled={loading}
        >
            {#if loading}
                <Loading />
            {:else}
                Log In
            {/if}
        </Button>
    </div>

    <div class="mt-4 text-center">
        <Links href="/#/signup" class="text-blue-600 hover:underline">
            Sign Up
        </Links>
    </div>
</Container>
