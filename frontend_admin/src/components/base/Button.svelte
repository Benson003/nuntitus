<script lang="ts">
    let {
        onClick,
        children,
        class:className = "",
    } = $props<{
        onClick?: () => void;
        children?: () => any;
        class?: string;
    }>();
</script>

<button
    onclick={(e) => {
        e.stopPropagation();
        e.preventDefault();
        onClick?.();
    }}
    class={`p-4 m-4 rounded-2xl font-bold transition duration-200
        cursor-pointer
        border border-gray-300 dark:border-gray-700
        shadow-sm
        hover:shadow-md
        active:scale-[0.98]
        ${className}

        `}>{@render children()}</button
>
