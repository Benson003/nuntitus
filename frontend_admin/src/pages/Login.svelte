<script lang="ts">
    import Button from "../components/base/Button.svelte";
    import Container from "../components/base/Container.svelte";
    import TextInput from "../components/base/inputs/TextInput.svelte";
    import PasswordInput from "../components/base/inputs/PasswordInput.svelte";
    import { push } from "svelte-spa-router";
    import Links from "../components/base/Links.svelte";

    let username = $state("");
    let password = $state("");

    function handleSubmit() {
        if(username != "" && password != ""){
        console.log("Username:", username);
        console.log("Password:", password);
        // TODO: Hook into your auth system
        push(`/home`)
    }
    }
</script>

<Container
    class="flex flex-col items-center justify-center p-6 w-svw h-svh bg-white dark:bg-gray-950"
>
    <div
        class="w-full max-w-md p-8 rounded-xl shadow-xl bg-gray-100 dark:bg-gray-900"
    >
        <h1
            class="text-3xl font-extrabold text-center mb-6 text-gray-900 dark:text-gray-100"
        >
            Welcome Back
        </h1>

        <div class="space-y-4">
            <TextInput label="Username" value={username} onInput={(v) => username = v}></TextInput>
            <PasswordInput label="Password" value={password} onInput={(v) => password = v}></PasswordInput>
        </div>

        <Button
            class="w-full mt-6 bg-blue-500 hover:bg-blue-600 text-white transition-transform duration-150 hover:scale-105 active:scale-95"
            onClick={handleSubmit}
        >
            Submit
        </Button>
    </div>
    <Links href="/#/signup">Sign Up</Links>
</Container>
