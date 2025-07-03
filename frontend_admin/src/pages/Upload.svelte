<script lang="ts">
    import Container from "../components/base/Container.svelte";
    import Button from "../components/base/Button.svelte";
    import Loading from "../components/base/Loading.svelte";
    import { get } from "svelte/store";
    import { jwt } from "../lib/authStore";
    import { push } from "svelte-spa-router";
    import MarkDownViewer from "../components/additional/MarkDownViewer.svelte";

    let file: File | null = null;
    let markdownContent = "";
    let error = "";
    let loading = false;

    // Blog metadata
    let title = "";
    let summary = "";
    let visibility = "true"; // string "true" or "false"
    let publishDate = new Date().toISOString().slice(0, 10); // "YYYY-MM-DD"

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
        error = "";
        file = selected;
        const reader = new FileReader();
        reader.onload = () => {
            markdownContent = reader.result as string;
        };
        reader.readAsText(selected);
    }

    function formatPublishDate(dateStr: string) {
        return new Date(dateStr + "T00:00:00Z").toISOString();
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
            const token = get(jwt);
            if (!token) throw new Error("Not authenticated");

            // Step 1: Create blog metadata (without content)
            const createRes = await fetch("/api/v1/blogs", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify({
                    title: title.trim(),
                    summary: summary.trim(),
                    visibility: visibility === "true",
                    publish_time: formatPublishDate(publishDate),
                    content: "", // empty on creation
                }),
            });

            if (!createRes.ok) {
                throw new Error("Failed to create blog metadata");
            }

            const createData = await createRes.json();
            const blogId = createData.data?.blog_id;
            if (!blogId) {
                throw new Error("Blog ID not returned from server");
            }

            // Step 2: Upload markdown content to separate endpoint
            const formData = new FormData();
            formData.append("file", file);

            const uploadRes = await fetch(`/api/v1/blogs/${blogId}/upload`, {
                method: "POST",
                headers: {
                    Authorization: `Bearer ${token}`,
                },
                body: formData,
            });

            if (!uploadRes.ok) {
                throw new Error("Failed to upload markdown content");
            }

            // Success! Redirect or clear form
            push("/dashboard");
        } catch (e) {
            error = e instanceof Error ? e.message : "Upload error";
        } finally {
            loading = false;
        }
    }
</script>

<Container class="w-screen h-screen bg-white dark:bg-gray-950 p-6">
    <div class="flex flex-row w-full h-full space-x-8">
        <!-- Left: Upload + Form -->
        <section
            class="w-2/5 flex flex-col space-y-6 overflow-auto p-6 bg-gray-100 dark:bg-gray-900 rounded-lg shadow-lg"
            tabindex="0"
        >
            <h2 class="text-3xl font-bold text-gray-900 dark:text-gray-100">
                Create New Blog
            </h2>

            <div class="flex flex-col space-y-4">
                <label
                    class="block text-gray-700 dark:text-gray-300 font-semibold"
                    >Title *</label
                >
                <input
                    type="text"
                    class="input input-bordered w-full"
                    bind:value={title}
                    placeholder="Enter blog title"
                />

                <label
                    class="block text-gray-700 dark:text-gray-300 font-semibold"
                    >Summary</label
                >
                <textarea
                    class="textarea textarea-bordered w-full"
                    rows="3"
                    bind:value={summary}
                    placeholder="Brief summary of the blog"
                ></textarea>

                <label
                    class="block text-gray-700 dark:text-gray-300 font-semibold"
                    >Visibility</label
                >
                <select
                    class="select select-bordered w-full dark:bg-gray-800"
                    bind:value={visibility}
                    aria-label="Blog visibility"
                >
                    <option value="true" class="text-black dark:text-white"
                        >Public</option
                    >
                    <option value="false" class="text-black dark:text-white"
                        >Private</option
                    >
                </select>

                <label
                    class="block text-gray-700 dark:text-gray-300 font-semibold"
                    >Publish Date</label
                >
                <input
                    type="date"
                    class="input input-bordered w-full"
                    bind:value={publishDate}
                    aria-label="Publish date"
                />

                <label
                    class="block text-gray-700 dark:text-gray-300 font-semibold"
                    >Upload Markdown File</label
                >
                <input
                    type="file"
                    accept=".md,text/markdown"
                    on:change={onFileChange}
                    class="file-input file-input-bordered w-full"
                    aria-label="Upload markdown file"
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
            class="flex-1 overflow-auto p-6 bg-white dark:bg-gray-800 rounded-lg shadow-inner"
            tabindex="0"
            aria-label="Markdown preview"
        >
            <MarkDownViewer content={markdownContent} />
        </section>
    </div>
</Container>
