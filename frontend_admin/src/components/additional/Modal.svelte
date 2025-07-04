<script lang="ts">
    let {
        open = false,
        onClose,
        children,
        class: className = "",
    } = $props<{
        open: boolean;
        onClose?: () => void;
        children?: () => any;
        class?: string;
    }>();

    function handleOverlayClick() {
        onClose?.();
    }
</script>

{#if open}
    <div
        class="fixed inset-0 bg-gray-600 flex items-center justify-center z-50"
        onclick={handleOverlayClick}
    >
        <div
            class={`bg-white dark:bg-gray-800 rounded-lg p-6 shadow-xl max-w-md w-full ${className}`}
            onclick={(e) => e.stopPropagation()}
        >
            {@render children()}
        </div>
    </div>
{/if}
