<script lang="ts">
    let { children, class: className = "" } = $props<{
        class?: string;
        children: () => any;
    }>();
</script>

<h3
    class={`
        text-2xl md:text-3xl lg:text-4xl
        m-0.5
        p-2
        font-semibold
        tracking-tight
        text-center
        ${className}
    `}
>
    {@render children()}
</h3>
