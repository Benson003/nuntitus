<script lang="ts">
    import Links from "../base/Links.svelte";

    // You can optionally pass year or default to current year
    const year: number = new Date().getFullYear();
    const owner: string = "Nuntius";
    const email: string = "nwankwobenson29@gmail.com";
</script>

<footer
    class="bg-gray-100 dark:bg-gray-900 text-gray-700 dark:text-gray-300 py-4 flex flex-col items-center space-y-2 select-none text-sm w-full"
>
    <div>
        <Links href={`mailto:${email}`} class="hover:underline">
            Contact: {email}
        </Links>
    </div>
    <div>
        &copy; {year} {owner}. All rights reserved.
    </div>
</footer>

