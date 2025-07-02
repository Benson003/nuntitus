<script lang="ts">
    let { children, class: className = "" } = $props<{
        class?: string;
        children: () => any;
    }>();
</script>

<p
    class={`
        text-base md:text-lg lg:text-xl
        m-0.5
        p-1
        text-center

        ${className}
    `}
>
    {@render children()}
</p>
