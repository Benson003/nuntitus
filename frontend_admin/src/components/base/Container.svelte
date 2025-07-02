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
      shadow-sm
      ${className}
      `}
>
    {@render children()}
</div>
