<script lang="ts">
    import Modal from "../additional/Modal.svelte";
    import Button from "../base/Button.svelte";
    import Loading from "../base/Loading.svelte";
    import MarkDownViewer from "../additional/MarkDownViewer.svelte";
    import TextInput from "../base/inputs/TextInput.svelte";

    import { uploadBlogContent, updateBlog } from "../../lib/api";
    import { marked } from "marked";
    import { createEventDispatcher, onMount } from "svelte";

    let {
        open = false,
        onClose,
        blog = null,
    } = $props<{
        open: boolean;
        onClose: () => void;
        blog: any;
    }>();

    const dispatch = createEventDispatcher();

    let file: File | null = null;
    let markdownContent = "";
    let loading = false;
    let error = "";

    let title = "";
    let summary = "";
    let visibility = "true";
    let publishDate = new Date().toISOString().slice(0, 10);

    marked.setOptions({
        breaks: true,
        headerIds: true,
        mangle: false,
    });

    onMount(() => {
        if (blog) {
            title = blog.title ?? "";
            summary = blog.summary ?? "";
            visibility = blog.visibility ? "true" : "false";
            publishDate = new Date(blog.publish_time)
                .toISOString()
                .slice(0, 10);
            markdownContent = blog.content ?? "";
        }
    });

    function onFileChange(event: Event) {
        const input = event.target as HTMLInputElement;
        if (!input.files || input.files.length === 0) {
            file = null;
            markdownContent = "";
            return;
        }
        const selected = input.files[0];
        if (
            !selected.name.endsWith(".md") &&
            !selected.type.includes("markdown")
        ) {
            error = "Please upload a Markdown (.md) file.";
            file = null;
            markdownContent = "";
            return;
        }
        file = selected;
        error = "";

        const reader = new FileReader();
        reader.onload = () => {
            markdownContent = reader.result as string;
        };
        reader.readAsText(selected);
    }

    async function handleEdit() {
        if (!title.trim()) {
            error = "Title is required.";
            return;
        }
        loading = true;
        error = "";

        try {
            const publish_time = new Date(
                `${publishDate}T00:00:00Z`
            ).toISOString();
            await updateBlog(blog.blog_id, {
                title,
                summary,
                visibility: visibility === "true",
                publish_time,
            });

            if (file) {
                await uploadBlogContent(blog.blog_id, file);
            }

            dispatch("edited");
            onClose();
        } catch (e) {
            error = e instanceof Error ? e.message : "Error updating blog";
        } finally {
            loading = false;
        }
    }
</script>

<Modal {open} {onClose} class="!p-0">
    <div
        class="fixed inset-0 flex flex-col md:flex-row bg-white dark:bg-gray-950 overflow-hidden z-50"
    >
        <!-- Left: Edit Form -->
        <section
            class="w-full md:w-1/3 h-full overflow-auto p-6 space-y-4 border-r border-gray-300 dark:border-gray-700"
        >
            <h2 class="text-2xl font-bold text-gray-900 dark:text-gray-100">
                Edit Blog
            </h2>

            <TextInput
                label="Title *"
                required
                value={title}
                onInput={(v) => (title = v)}
            />

            <label
                class="text-sm font-semibold text-gray-700 dark:text-gray-300"
                >Summary</label
            >
            <textarea
                class="p-2 rounded border border-gray-400 dark:border-gray-700 bg-white dark:bg-gray-800 w-full"
                rows="3"
                bind:value={summary}
                placeholder="Brief summary of the blog"
            ></textarea>

            <label
                class="text-sm font-semibold text-gray-700 dark:text-gray-300"
                >Visibility</label
            >
            <select
                class="p-2 rounded border border-gray-400 dark:border-gray-700 bg-white dark:bg-gray-800 w-full"
                bind:value={visibility}
            >
                <option value="true">Public</option>
                <option value="false">Private</option>
            </select>

            <TextInput
                label="Publish Date"
                date="date"
                value={publishDate}
                onInput={(v) => (publishDate = v)}
            />

            <label
                class="text-sm font-semibold text-gray-700 dark:text-gray-300"
                >Upload New Markdown File (optional)</label
            >
            <input
                type="file"
                accept=".md,text/markdown"
                on:change={onFileChange}
                class="file-input file-input-bordered w-full"
            />

            {#if file}
                <p class="text-sm text-gray-700 dark:text-gray-300">
                    Selected file: {file.name}
                </p>
            {/if}

            {#if error}
                <p class="text-sm text-red-600 dark:text-red-400">{error}</p>
            {/if}

            <Button
                class="bg-green-600 hover:bg-green-700 text-white w-full mt-2"
                onClick={handleEdit}
                disabled={loading}
            >
                {#if loading}
                    <Loading />
                {:else}
                    Save Changes
                {/if}
            </Button>

            <Button
                class="bg-gray-300 hover:bg-gray-400 dark:bg-gray-700 dark:hover:bg-gray-600 text-gray-900 dark:text-gray-100 w-full"
                onClick={onClose}
            >
                Cancel
            </Button>
        </section>

        <!-- Right: Markdown Preview -->
        <section
            class="flex-1 h-full overflow-auto p-6 bg-gray-50 dark:bg-gray-900"
        >
            <h3
                class="text-lg font-semibold text-gray-900 dark:text-gray-100 mb-2"
            >
                Markdown Preview
            </h3>
            <MarkDownViewer content={markdownContent} />
        </section>
    </div>
</Modal>
