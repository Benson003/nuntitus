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
        class="fixed inset-0 bg-opacity-50 flex items-center justify-center z-50"
        onclick={handleOverlayClick}
    >
        <div
            class={`
          p-6 rounded-lg shadow-lg max-w-lg w-full ${className}
        `}
            onclick={(e) => e.stopPropagation()}
        >
            {@render children()}
        </div>
    </div>
{/if}
