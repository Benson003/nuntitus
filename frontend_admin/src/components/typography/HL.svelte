<script lang="ts">
    let { children, class: className = "" } = $props<{
        class?: string;
        children: () => any;
    }>();
</script>

<h1
    class={`
        text-4xl md:text-5xl lg:text-6xl
        m-0.5
        p-2
        font-bold
        tracking-tight
        text-center

        ${className}
    `}
>
    {@render children()}
</h1>
