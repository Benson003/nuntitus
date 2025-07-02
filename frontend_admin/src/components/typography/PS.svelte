<script lang="ts">
    let { children, class: className = "" } = $props<{
        class?: string;
        children: () => any;
    }>();
</script>

<p
    class={`
        text-sm md:text-base
        m-0.5
        p-1
        text-center

        ${className}
    `}
>
    {@render children()}
</p>
