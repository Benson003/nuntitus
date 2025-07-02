<script lang="ts">
    let { children, class: className = "" } = $props<{
        class?: string;
        children: () => any;
    }>();
</script>

<p
    class={`
        text-lg md:text-xl lg:text-2xl
        text-center

        ${className}
    `}
>
    {@render children()}
</p>
