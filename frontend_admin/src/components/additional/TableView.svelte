<script lang="ts">
    import Button from "../base/Button.svelte";
    import Pm from "../typography/PM.svelte";
    import Ps from "../typography/PS.svelte";

    let {
        data = [],
        page = 1,
        totalPages = 1,
        onNext = () => {},
        onPrevious = () => {},
        onEdit = (id: string) => {},
        onDelete = (id: string) => {},
    } = $props<{
        data: any[];
        page: number;
        totalPages: number;
        onNext: () => void;
        onPrevious: () => void;
        onEdit: (id: string) => void;
        onDelete: (id: string) => void;
    }>();
</script>

<div class="w-full overflow-x-auto">
    <table class="min-w-full divide-y divide-gray-300 dark:divide-gray-700">
        <thead class="bg-gray-100 dark:bg-gray-800">
            <tr>
                {#each ["Title", "Summary", "Date", "Visibility", "Created At", "Updated At", "Actions"] as header}
                    <th
                        class="px-4 py-2 text-left text-sm font-semibold text-gray-900 dark:text-gray-100"
                    >
                        {header}
                    </th>
                {/each}
            </tr>
        </thead>
        <tbody class="divide-y divide-gray-200 dark:divide-gray-600">
            {#if data.length === 0}
                <tr>
                    <td
                        colspan="7"
                        class="px-4 py-4 text-center text-gray-500 dark:text-gray-400"
                        >No data available</td
                    >
                </tr>
            {:else}
                {#each data as item (item.blog_id)}
                    <tr
                        class="hover:bg-gray-50 dark:hover:bg-gray-700 transition"
                    >
                        <td class="px-4 py-2"><Ps>{item.title}</Ps></td>
                        <td class="px-4 py-2"><Ps>{item.summary}</Ps></td>
                        <td class="px-4 py-2">
                            <Ps
                                >{new Date(
                                    item.publish_time
                                ).toLocaleDateString()}</Ps
                            >
                        </td>
                        <td class="px-4 py-2">
                            <span
                                class={`px-2 py-1 rounded-2xl text-xs font-semibold ${
                                    item.visibility
                                        ? "bg-green-200 text-green-800 dark:bg-green-800 dark:text-green-200"
                                        : "bg-yellow-200 text-yellow-800 dark:bg-yellow-800 dark:text-yellow-200"
                                }`}
                            >
                                {item.visibility ? "Public" : "Private"}
                            </span>
                        </td>
                        <td class="px-4 py-2"
                            ><Ps
                                >{new Date(
                                    item.created_at
                                ).toLocaleDateString()}</Ps
                            ></td
                        >
                        <td class="px-4 py-2"
                            ><Ps
                                >{new Date(
                                    item.updated_at
                                ).toLocaleDateString()}</Ps
                            ></td
                        >
                        <td class="px-4 py-2">
                            <div class="flex space-x-2">
                                <Button
                                    class="bg-blue-500 hover:bg-blue-600 text-white px-2 py-1 rounded text-xs"
                                    onClick={() => onEdit(item.blog_id)}
                                    >Edit</Button
                                >
                                <Button
                                    class="bg-red-500 hover:bg-red-600 text-white px-2 py-1 rounded text-xs"
                                    onClick={() => onDelete(item.blog_id)}
                                    >Delete</Button
                                >
                            </div>
                        </td>
                    </tr>
                {/each}
            {/if}
        </tbody>
    </table>

    <div class="flex justify-between items-center mt-4">
        <Button
            onClick={onPrevious}
            disabled={page <= 1}
            class="px-3 py-1 rounded bg-gray-200 dark:bg-gray-700 text-gray-800 dark:text-gray-100 disabled:opacity-50"
        >
            Previous
        </Button>

        <Pm class="text-sm text-gray-700 dark:text-gray-300">
            Page {page} of {totalPages}
        </Pm>

        <Button
            onClick={onNext}
            disabled={page >= totalPages}
            class="px-3 py-1 rounded bg-gray-200 dark:bg-gray-700 text-gray-800 dark:text-gray-100 disabled:opacity-50"
        >
            Next
        </Button>
    </div>
</div>
