<script lang="ts">
    let { children, class: className = "" } = $props<{
        children: () => any;
        class?: string;
    }>();
</script>

<div
    class={`
        flex flex-col
        min-h-screen w-full
        p-0 m-0
        ${className}
    `}
>
    <div class="flex-grow w-full flex flex-col p-4">
        {@render children()}
    </div>
</div>
