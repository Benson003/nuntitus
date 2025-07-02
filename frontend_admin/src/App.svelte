<script lang="ts">

    import BackDropContainer from "./components/base/BackDropContainer.svelte";
    import Home from "./pages/Home.svelte";
    import Login from "./pages/Login.svelte";
    import Page404 from "./pages/Page404.svelte";
    import SignUp from "./pages/SignUp.svelte";
    import Router from "svelte-spa-router";

    let routes = {
        "/": Home,
        "/signup": SignUp,
        "/login": Login,
        "*": Page404,
    };
</script>

<main>
    <BackDropContainer class="dark:bg-gray-800 bg-gray-300 dark:text-white text-black">
        <Router {routes} />
    </BackDropContainer>
</main>
