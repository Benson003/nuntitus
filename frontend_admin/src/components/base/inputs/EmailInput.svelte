<script lang="ts">
    import Ps from "../../typography/PS.svelte";

    let {
        label = "",
        value = "",
        placeholder = "",
        type = "email",
        onInput,
        class: className = "",
    } = $props<{
        label?: string;
        value?: string;
        placeholder?: string;
        type?: string;
        onInput?: (v: string) => void;
        class?: string;
    }>();
    let is_valid_email: boolean = $state(false);
    let display_text: string = $state("");
    function handleInput(e: Event) {
        const input = e.target as HTMLInputElement;
        const value = input.value;
        if (isValidEmail(value) && value != "") {
            onInput?.(value);
            is_valid_email = true
            display_text="Valid email address"
        } else if(value === "") {
            is_valid_email = false
            display_text = "No email entered"
        }else {
            is_valid_email = false
            display_text = "Invalid email format must follow 'test@mail.com'"
        }
    }

    function isValidEmail(email: string): boolean {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailRegex.test(email);
    }
</script>

<label class="block m-4 text-sm font-bold text-black dark:text-white">
    {label}
    <input
        {type}
        {value}
        {placeholder}
        oninput={handleInput}
        class={`
        w-full p-2 rounded-md border border-gray-400 dark:border-gray-700
        focus:outline-none focus:ring-2 focus:ring-blue-500 shadow-md
        ${className}
      `}
    />
</label>
<Ps class={is_valid_email? "text-green-600": "text-red-600"}>{display_text}</Ps>
