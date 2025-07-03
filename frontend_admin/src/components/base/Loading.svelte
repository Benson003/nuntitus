<script lang="ts">
    import { onMount, onDestroy } from "svelte";

    // Default props via $props pattern
    // You can override these by passing props to the component
    let { size = 6 }
    let size: number = 6;
    let colors: string[] = [
        "bg-blue-900",
        "bg-blue-600",
        "bg-cyan-400",
        "bg-blue-200",
    ];
    let intervalMs: number = 180;

    let active = 0;
    const sequence = [0, 1, 2, 3, 2, 1];

    let idx = 0;
    let interval: ReturnType<typeof setInterval>;

    onMount(() => {
        interval = setInterval(() => {
            active = sequence[idx];
            idx = (idx + 1) % sequence.length;
        }, intervalMs);
    });

    onDestroy(() => {
        clearInterval(interval);
    });
</script>

<div
    role="status"
    class="flex space-x-2 items-center justify-center p-4"
    aria-live="polite"
    aria-label="Loading animation"
>
    {#each colors as color, i}
        <div
            class={`rounded transition-all duration-200 w-${size} h-${size} ${color} ${
                i === active ? "scale-125 brightness-125" : ""
            }`}
        ></div>
    {/each}
</div>
