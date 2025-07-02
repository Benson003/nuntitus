<script lang="ts">
    let { children, class: className = "" } = $props<{
        class?: string;
        children: () => any;
    }>();
</script>

<h1
    class={`
        text-4xl md:text-5xl lg:text-6xl
        font-bold
        tracking-tight
        text-center
       
        ${className}
    `}
>
    {@render children()}
</h1>
