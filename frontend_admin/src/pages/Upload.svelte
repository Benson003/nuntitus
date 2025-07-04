<script lang="ts">
    import Container from "../components/base/Container.svelte";
    import Button from "../components/base/Button.svelte";
    import Loading from "../components/base/Loading.svelte";
    import MarkDownViewer from "../components/additional/MarkDownViewer.svelte";
    import { push } from "svelte-spa-router";
    import { createBlog, uploadBlogContent } from "../lib/api";
    import { marked } from "marked";
    import TextInput from "../components/base/inputs/TextInput.svelte";

    let file: File | null = null;
    let markdownContent = "";
    let error = "";
    let loading = false;

    let title = "";
    let summary = "";
    let visibility = "true";
    let publishDate = new Date().toISOString().slice(0, 10); // YYYY-MM-DD

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

    async function handleUpload() {
        if (!file) {
            error = "Please select a Markdown file first.";
            return;
        }
        if (!title.trim()) {
            error = "Title is required.";
            return;
        }

        loading = true;
        error = "";

        try {
            // 1️⃣ Create blog metadata first
            const publish_time = new Date(
                publishDate + "T00:00:00Z"
            ).toISOString();

            const response = await createBlog({
                title,
                summary,
                visibility: visibility === "true",
                publish_time,
            });

            const blogId = response.data.blog_id;
            if (!blogId) {
                throw new Error("Blog ID not returned from server");
            }

            // 2️⃣ Upload the markdown file
            await uploadBlogContent(blogId, file);

            // Redirect to dashboard or blog view
            push("/dashboard");
        } catch (e) {
            error = e instanceof Error ? e.message : "Upload error";
        } finally {
            loading = false;
        }
    }

    marked.setOptions({
        breaks: true,
        headerIds: true,
        mangle: false,
    });
</script>

<Container class="w-full h-screen bg-white dark:bg-gray-950 p-6">
    <div class="flex flex-row w-full h-full space-x-8">
        <!-- Left: Upload Form -->
        <section
            class="w-1/5 flex flex-col space-y-6 overflow-auto p-6 bg-gray-100 dark:bg-gray-900 rounded-lg shadow-lg"
            tabindex="0"
        >
            <h2 class="text-3xl font-bold text-gray-900 dark:text-gray-100">
                Create New Blog
            </h2>

            <div class="flex flex-col space-y-4">
                <TextInput
                    label="Title *"
                    requried
                    value={title}
                    onInput={(v) => (title = v)}
                ></TextInput>

                <label
                    class="block text-gray-700 dark:text-gray-300 font-semibold"
                    >Summary</label
                >
                <textarea
                    class="textarea textarea-bordered w-full p-2 rounded-md border border-gray-400 dark:border-gray-700
        focus:outline-none focus:ring-2 focus:ring-blue-500 shadow-md"
                    rows="3"
                    bind:value={summary}
                    placeholder="Brief summary of the blog"
                ></textarea>

                <label
                    class="block text-gray-700 dark:text-gray-300 font-semibold"
                    >Visibility</label
                >
                <select
                    class=" p-4 rounded select select-bordered w-full dark:bg-gray-800"
                    bind:value={visibility}
                >
                    <option value="true">Public</option>
                    <option value="false">Private</option>
                </select>

                <TextInput
                    label="Publish Date"
                    type="date"
                    value={publishDate}
                    onInput={(v) => (publishDate = v)}
                ></TextInput>

                <label
                    class="block text-gray-700 dark:text-gray-300 font-semibold"
                    >Upload Markdown File</label
                >
                <input
                    type="file"
                    accept=".md,text/markdown"
                    on:change={onFileChange}
                    class="file-input file-input-bordered w-full"
                />

                {#if file}
                    <p class="text-gray-700 dark:text-gray-300">
                        Selected file: {file.name}
                    </p>
                {/if}

                {#if error}
                    <p class="text-red-600 dark:text-red-400">{error}</p>
                {/if}

                <Button
                    class="mt-4 bg-green-600 hover:bg-green-700 text-white"
                    onClick={handleUpload}
                    disabled={loading}
                >
                    {#if loading}
                        <Loading />
                    {:else}
                        Upload Blog
                    {/if}
                </Button>
            </div>
        </section>

        <!-- Right: Markdown Preview -->
        <section
            class="flex-1 w-3/5 overflow-auto p-6 bg-transparent"
            tabindex="0"
        >
            <MarkDownViewer content={markdownContent} />
        </section>
    </div>
</Container>
