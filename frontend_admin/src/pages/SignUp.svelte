<script lang="ts">
    import Button from "../components/base/Button.svelte";
    import Container from "../components/base/Container.svelte";
    import TextInput from "../components/base/inputs/TextInput.svelte";
    import PasswordInput from "../components/base/inputs/PasswordInput.svelte";
    import EmailInput from "../components/base/inputs/EmailInput.svelte";
    import Links from "../components/base/Links.svelte";
    import Loading from "../components/base/Loading.svelte";

    import { signupUser } from "../lib/api";
    import { jwt } from "../lib/authStore";

    import HXL from "../components/typography/HXL.svelte";
    import PL from "../components/typography/PL.svelte";
    import { push } from "svelte-spa-router";

    let username = "";
    let email = "";
    let password = "";
    let error = "";
    let loading = false;

    async function handleSubmit() {
        error = "";

        if (!username.trim()) {
            error = "Username is required.";
            return;
        }
        if (!email.trim()) {
            error = "Email is required.";
            return;
        }
        if (password.length < 6) {
            error = "Password must be at least 6 characters.";
            return;
        }

        loading = true;
        try {
            const token = await signupUser(username, email, password);
            jwt.set(token);
            push("/dashboard");
        } catch (e) {
            error =
                e instanceof Error ? e.message : "Signup failed. Try again.";
        } finally {
            loading = false;
        }
    }
</script>

<Container
    class="flex flex-col items-center justify-center p-6 w-svw h-svh bg-white dark:bg-gray-950"
>
    <div
        class="w-full max-w-md p-8 rounded-xl shadow-xl bg-gray-100 dark:bg-gray-900"
    >
        <HXL class="text-center mb-6 text-gray-900 dark:text-gray-100">
            Create Account
        </HXL>

        <div class="space-y-4">
            <TextInput
                label="Username"
                value={username}
                onInput={(v) => (username = v)}
                placeholder="RootLogger"
            />
            <EmailInput
                label="Email"
                value={email}
                onInput={(v) => (email = v)}
                placeholder="test@mail.com"
            />
            <PasswordInput
                label="Password"
                value={password}
                onInput={(v) => (password = v)}
                placeholder="*********"
            />
        </div>

        {#if error}
            <PL class="mt-2 text-red-600 text-center dark:text-red-400">
                {error}
            </PL>
        {/if}

        <Button
            class="w-full mt-6 bg-green-600 hover:bg-green-700 dark:bg-green-700 dark:hover:bg-green-800 text-white transition-transform duration-150 hover:scale-105 active:scale-95"
            onClick={handleSubmit}
            disabled={loading}
        >
            {#if loading}
                <Loading />
            {:else}
                Sign Up
            {/if}
        </Button>

        <div class="mt-4 text-center">
            <Links href="/#/login" class="text-blue-600 hover:underline">
                Log In
            </Links>
        </div>
    </div>
</Container>
