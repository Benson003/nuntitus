<script lang="ts">
    import Button from "../components/base/Button.svelte";
    import Container from "../components/base/Container.svelte";
    import Hl from "../components/typography/HL.svelte";
    import Pm from "../components/typography/PM.svelte";
    import { push } from "svelte-spa-router";

    function goHome() {
        push(`/home`);
    }
</script>

<Container
    class="flex flex-col items-center justify-center text-center p-8 space-y-6 animate-fade-in"
>
    <Hl
        class="text-6xl font-extrabold text-red-500 dark:text-red-400 animate-pulse"
    >
        404
    </Hl>
    <Pm class="text-xl text-gray-700 dark:text-gray-300 max-w-md">
        Oops, the page you are looking for does not exist or has been moved.
    </Pm>
    <Button class="bg-blue-500 hover:bg-blue-600 text-white" onClick={goHome}>
        Go Home
    </Button>
</Container>
