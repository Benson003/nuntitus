<script lang="ts">
    import BackDropContainer from "./components/base/BackDropContainer.svelte";
    import Login from "./pages/Login.svelte";
    import Page404 from "./pages/Page404.svelte";
    import SignUp from "./pages/SignUp.svelte";
    import Router, { push } from "svelte-spa-router";
    import { onMount } from "svelte";
    import Dashboard from "./pages/Dashboard.svelte";
    import EditBlog from "./components/additional/EditBlog.svelte";
    import Upload from "./pages/Upload.svelte";
    import Settings from "./pages/Settings.svelte";
    import Header from "./components/additional/Header.svelte";
    import Footer from "./components/additional/Footer.svelte";
    let logged_in: boolean = false;
    onMount(() => {
        if (!logged_in) {
            push(`/login`);
        }
    });

    let routes = {
        "/edit": EditBlog,
        "/settings": Settings,
        "/upload": Upload,
        "/dashboard": Dashboard,
        "/signup": SignUp,
        "/login": Login,
        "*": Page404,
    };
</script>

<main class="min-h-screen flex flex-col">
    <Header />
    <BackDropContainer
        class="flex-1 dark:bg-gray-800 bg-gray-300 dark:text-white text-black"
    >
        <Router {routes} />
    </BackDropContainer>
    <Footer />
</main>
