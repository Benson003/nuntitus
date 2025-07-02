<script lang="ts">
    let { label = '', value = '', placeholder = '', type = 'text', onInput,class:className = "" } = $props<{
      label?: string;
      value?: string;
      placeholder?: string;
      type?: string;
      onInput?: (v: string) => void;
      class?: string;
    }>();
    
    function handleInput(e: Event) {
      onInput?.((e.target as HTMLInputElement).value);
    }
  </script>

  <label class="block m-4 text-sm font-bold text-black dark:text-white">
    {label}
    <input
      type={type}
      value={value}
      placeholder={placeholder}
      oninput={handleInput}
      class={`
        w-full p-2 rounded-md border border-gray-300 dark:border-gray-700
        focus:outline-none focus:ring-2 focus:ring-blue-500
        ${className}
      `}
    />
  </label>
