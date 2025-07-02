<script lang="ts">
    let { children, class: className = "" } = $props<{
        class?: string;
        children: () => any;
    }>();
</script>

<p
    class={`
        text-sm md:text-base
        text-center

        ${className}
    `}
>
    {@render children()}
</p>
