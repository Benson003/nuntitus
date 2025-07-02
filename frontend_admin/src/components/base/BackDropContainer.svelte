<script lang="ts">
    let { children, class: className = "" } = $props<{
        children: () => any;
        class?: string;
    }>();
</script>

<div
    class={`
        flex justify-center items-center
        w-svw h-svh
       font-medium text-xl
        ${className}
        `}
>
    <div
        class="flex flex-col justify-center items-center text-center space-y-4"
    >
        {@render children()}
    </div>
</div>
