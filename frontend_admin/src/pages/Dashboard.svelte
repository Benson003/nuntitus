<script lang="ts">
    import Container from "../components/base/Container.svelte";
    import TableView from "../components/additional/TableView.svelte";
    import Loading from "../components/base/Loading.svelte";
    import ProfileCard from "../components/additional/ProfileCard.svelte";
    import Modal from "../components/additional/Modal.svelte";
    import EditBlogModal from "../components/additional/EditBlog.svelte";
    import Button from "../components/base/Button.svelte";

    import { onMount } from "svelte";
    import { fetchUserBlogs, API_BASE, deleteBlog } from "../lib/api";
    import { fetchWithAuth } from "../lib/fetchWithAuth";

    let blogs: any = $state([]);
    let page = $state(1);
    let totalPages = $state(1);
    let loading = $state(true);
    let error = $state("");

    let userData: any = $state(null);
    let userLoading = $state(true);
    let userError = $state("");

    let showDeleteModal = $state(false);
    let showEditModal = $state(false);
    let selectedBlogId = $state("");
    let selectedBlog = $state(null);

    async function loadBlogs() {
        loading = true;
        error = "";
        try {
            const res = await fetchUserBlogs(page, 10);
            blogs = res.data.blogs;
            totalPages = res.data.total_pages ?? 1;
        } catch (e) {
            error = e instanceof Error ? e.message : "Failed to load blogs";
        } finally {
            loading = false;
        }
    }

    async function loadUser() {
        userLoading = true;
        userError = "";
        try {
            const res = await fetchWithAuth(`${API_BASE}/auth/user`);
            if (!res.ok) throw new Error("Failed to get user details");
            userData = await res.json();
        } catch (e) {
            userError = e instanceof Error ? e.message : "Failed to load user";
        } finally {
            userLoading = false;
        }
    }

    function nextPage() {
        if (page < totalPages) {
            page += 1;
            loadBlogs();
        }
    }

    function previousPage() {
        if (page > 1) {
            page -= 1;
            loadBlogs();
        }
    }

    async function handleDelete(id: string) {
        try {
            await deleteBlog(id);
            await loadBlogs();
            showDeleteModal = false;
        } catch (e) {
            console.error(e);
            alert("Failed to delete blog.");
        }
    }

    function handleEdit(id: string) {
        selectedBlog = blogs.find((b) => b.blog_id === id);
        if (selectedBlog) {
            showEditModal = true;
        } else {
            console.error("Blog not found for editing:", id);
        }
    }

    onMount(() => {
        loadUser();
        loadBlogs();
    });
</script>

<!-- Edit Modal -->
<EditBlogModal
    open={showEditModal}
    onClose={() => (showEditModal = false)}
    blog={selectedBlog}
    on:edited={loadBlogs}
/>

<!-- Delete Confirmation Modal -->
<Modal open={showDeleteModal} onClose={() => (showDeleteModal = false)}>
    <div class="flex flex-col space-y-4 text-center">
        <h2 class="text-lg font-semibold text-gray-900 dark:text-gray-100">
            Confirm Delete
        </h2>
        <p class="text-gray-700 dark:text-gray-300">
            Are you sure you want to delete this blog?
        </p>
        <div class="flex justify-center space-x-4">
            <Button
                class="px-4 py-2 rounded bg-red-500 hover:bg-red-600 text-white"
                onClick={() => handleDelete(selectedBlogId)}
            >
                Confirm
            </Button>
            <Button
                class="px-4 py-2 rounded bg-gray-300 hover:bg-gray-400 dark:bg-gray-700 dark:hover:bg-gray-600 text-gray-900 dark:text-gray-100"
                onClick={() => (showDeleteModal = false)}
            >
                Cancel
            </Button>
        </div>
    </div>
</Modal>

<!-- ProfileCard -->
{#if userLoading}
    <div class="flex justify-center mt-4">
        <Loading size={6} intervalMs={120} />
    </div>
{:else if userError}
    <p class="text-red-600 dark:text-red-400 text-center mt-4">{userError}</p>
{:else if userData}
    <ProfileCard
        username={userData.data.username}
        email={userData.data.email}
        first_name={userData.data.first_name}
        last_name={userData.data.last_name}
    />
{/if}

<!-- Blogs Table -->
<Container class="p-6 w-full h-full space-y-4">
    <h1 class="text-3xl font-bold mb-4 text-gray-900 dark:text-gray-100">
        Your Blogs
    </h1>

    {#if loading}
        <div class="flex justify-center items-center h-48">
            <Loading size={8} intervalMs={150} />
        </div>
    {:else if error}
        <p class="text-red-600 dark:text-red-400 text-center">{error}</p>
    {:else}
        <TableView
            data={blogs}
            {page}
            {totalPages}
            onNext={nextPage}
            onPrevious={previousPage}
            onEdit={(id) => handleEdit(id)}
            onDelete={(id) => {
                selectedBlogId = id;
                showDeleteModal = true;
            }}
        />
    {/if}
</Container>
