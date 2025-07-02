<script lang="ts">
    let { children, class: className = "" } = $props<{
        class?: string;
        children: () => any;
    }>();
</script>

<h2
    class={`
        text-3xl md:text-4xl lg:text-5xl
        font-bold
        tracking-tight
        text-center
        ${className}
    `}
>
    {@render children()}
</h2>
