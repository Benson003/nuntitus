<script lang="ts">
    import Footer from "../additional/Footer.svelte";
    import Header from "../additional/Header.svelte";

    let { children, class: className = "" } = $props<{
        children: () => any;
        class?: string;
    }>();
</script>

<Header />
<div
    class={`
        flex justify-center items-center
        w-svw h-svh
       font-medium text-xl
       p-0 m-0
        ${className}
        `}
>
    <div
        class="flex p-0 m-0 flex-col justify-center items-center text-center space-y-4"
    >
        {@render children()}
    </div>
</div>
<Footer />
