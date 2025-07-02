<script lang="ts">
    let { children, class: className = "" } = $props<{
        children: () => any;
        class?: string;
    }>();
</script>

<h1
    class={`
        text-4xl md:text-5xl lg:text-6xl
        font-extrabold
        tracking-tight
        text-center
        mb-4
        ${className}
    `}
>
    {@render children()}
</h1>
