<script lang="ts">
    let {
        label = "",
        value = "",
        placeholder = "",
        type = "text",
        disabled = false,
        onInput,
        class: className = "",
        requried = false,
    } = $props<{
        label?: string;
        value?: string;
        placeholder?: string;
        type?: string;
        disabled?: boolean;
        onInput?: (v: string) => void;
        class?: string;
        requried?: boolean;
    }>();

    function handleInput(e: Event) {
        onInput?.((e.target as HTMLInputElement).value);
    }
</script>

<label class="block m-4 text-sm font-bold text-black dark:text-white">
    {label}
    <input
        required={requried}
        {type}
        {value}
        {disabled}
        {placeholder}
        oninput={handleInput}
        class={`
        w-full p-2 rounded-md border border-gray-400 dark:border-gray-700
        focus:outline-none focus:ring-2 focus:ring-blue-500 shadow-md
        ${className}
      `}
    />
</label>
