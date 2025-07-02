<script lang="ts">
    let { children, class: className = "" } = $props<{
        children?: () => any;
        class?: string;
    }>();
</script>

<div
    class={`
      max-w-4xl mx-auto p-4
      rounded-lg
      shadow-lg
      border border-gray-400 dark:border-gray-700
      ${className}
      `}
>
    {@render children()}
</div>
