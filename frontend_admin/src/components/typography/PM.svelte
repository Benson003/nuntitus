<script lang="ts">
    let { children, class: className = "" } = $props<{
        class?: string;
        children: () => any;
    }>();
</script>

<p
    class={`
        text-base md:text-lg lg:text-xl
        text-center

        ${className}
    `}
>
    {@render children()}
</p>
