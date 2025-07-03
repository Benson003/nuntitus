<script lang="ts">
    import { derived } from "svelte/store";

    let {
        label = "",
        value = "",
        placeholder = "",
        onInput,
        class: className = "",
    } = $props<{
        label?: string;
        value?: string;
        placeholder?: string;
        onInput?: (v: string) => void;
        class?: string;
    }>();

    let showPassword = $state(false);

    function handleInput(e: Event) {
        onInput?.((e.target as HTMLInputElement).value);
    }

    let inputType = $derived(showPassword ? "text" : "password");
</script>

<label class="block m-4 text-sm font-bold text-black dark:text-white">
    {label}
    <input
        type={inputType}
        value={value}
        placeholder={placeholder}
        oninput={handleInput}
        class={`w-full p-2 rounded-md border border-gray-400 dark:border-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-500 shadow-md ${className}`}
    />

    <label
        class="flex items-center space-x-2 text-sm font-medium text-black dark:text-white mt-2 cursor-pointer select-none"
    >
        <input
            type="checkbox"
            checked={showPassword}
            oninput={(e) =>
                (showPassword = (e.target as HTMLInputElement).checked)}
        />
        <span>Show password</span>
    </label>
</label>
