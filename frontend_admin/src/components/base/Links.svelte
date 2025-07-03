<script lang="ts">
    let {
        href,
        children,
        newTab,
        class: className = "",
        disabled = false,
    } = $props<{
        href: string;
        children: () => any;
        newTab?: boolean;
        class?: string;
        disabled?: boolean;
    }>();
</script>

<a
    {href}
    target={newTab ? "_blank" : undefined}
    rel={newTab ? "noopener noreferrer" : undefined}
    class={`font-mono
        text-black dark:text-white
        p-1
        hover:underline
        hover:text-blue-700 dark:hover:text-blue-400
        transition-colors duration-200
        cursor-pointer
        rounded
        inline-flex
        focus:outline-none focus:ring-2 focus:ring-blue-500
        ${className}
        `}
    >{@render children()}
    {#if newTab}
        <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
            class="lucide lucide-external-link"
        >
            <path
                d="M18 13v6a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6"
            />
            <polyline points="15 3 21 3 21 9" />
            <line x1="10" y1="14" x2="21" y2="3" />
        </svg>
    {/if}
</a>
