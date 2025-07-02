<script lang="ts">
    import Button from "../components/base/Button.svelte";
    import Container from "../components/base/Container.svelte";
    import Inputs from "../components/base/Inputs.svelte";

    let username = $state("");
    let password = $state("");
</script>
<h1 class="text-3xl font-extrabold text-center text-black dark:text-white">
    Welcome Back
</h1>
<Container>
    <div class="w-full max-w-md mx-auto p-8 space-y-6 bg-white dark:bg-gray-900 rounded-xl shadow-lg">


        <div class="space-y-4">
            <Inputs label="Username" type="text" value={username} />
            <Inputs label="Password" type="password" value={password} />
        </div>

        <Button class="w-full transition-transform duration-150 hover:scale-105 active:scale-95">
            Submit
        </Button>
    </div>
</Container>
