<script lang="ts">
    let {
        href,
        children,
        newTab,
        class: className = "",
    } = $props<{
        href: string;
        children: () => any;
        newTab?: boolean;
        class?: string;
    }>();
</script>

<a
    {href}
    target={newTab ? "_blank" : undefined}
    rel={newTab ? "noopener noreferrer" : undefined}
    class={
    `font-mono
        text-black dark:text-white
        p-1
        hover:underline
        hover:text-blue-700 dark:hover:text-blue-400
        transition-colors duration-200
        cursor-pointer
        rounded
        focus:outline-none focus:ring-2 focus:ring-blue-500
        ${className}
        `}>{@render children()}</a
>
