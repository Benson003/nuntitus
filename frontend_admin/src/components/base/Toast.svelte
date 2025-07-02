<script lang="ts">
    let {
        message = "",
        show = false,
        class: className = "",
    } = $props<{
        message?: string;
        show: boolean;
        class?: string;
    }>();
</script>

{#if show}
    <div
        class={`
        fixed bottom-4 right-4
        px-4 py-2 rounded shadow-lg
        transition-opacity duration-300
        ${className}
        `}
    >
        {message}
    </div>
{/if}
