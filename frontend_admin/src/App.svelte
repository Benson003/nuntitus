<script lang="ts">
    import BackDropContainer from "./components/base/BackDropContainer.svelte";
    import Login from "./pages/Login.svelte";
    import Page404 from "./pages/Page404.svelte";
    import SignUp from "./pages/SignUp.svelte";
    import Router, { push } from "svelte-spa-router";
    import { onMount } from "svelte";
    import Dashboard from "./pages/Dashboard.svelte";
    import EditBlog from "./pages/EditBlog.svelte";
    import Upload from "./pages/Upload.svelte";
    let logged_in: boolean = false;
    onMount(() => {
        if (!logged_in) {
            push(`/login`);
        }
    });

    let routes = {
        "/edit": EditBlog,
        "/upload": Upload,
        "/dashboard": Dashboard,
        "/signup": SignUp,
        "/login": Login,
        "*": Page404,
    };
</script>

<main>
    <BackDropContainer
        class="dark:bg-gray-800 bg-gray-300 dark:text-white text-black"
    >
        <Router {routes} />
    </BackDropContainer>
</main>
