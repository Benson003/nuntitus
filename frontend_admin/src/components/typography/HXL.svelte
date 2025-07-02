<script lang="ts">
    let { children, class: className = "" } = $props<{
        class?: string;
        children: () => any;
    }>();
</script>

<h1
    class={`
        text-5xl md:text-6xl lg:text-7xl
        font-extrabold
        tracking-tight
        text-center
        
        ${className}
    `}
>
    {@render children()}
</h1>
