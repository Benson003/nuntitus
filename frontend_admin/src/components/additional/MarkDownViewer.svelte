<script lang="ts">
    import { marked } from "marked";

    let { content = "" } = $props<{ content?: string }>();

    marked.setOptions({
        breaks: true,
        headerIds: true,
        mangle: false,
    });

    const renderer = new marked.Renderer();
    renderer.link = function (href, title, text) {
        const link = marked.Renderer.prototype.link.call(
            this,
            href,
            title,
            text
        );
        return link.replace(
            /^<a /,
            '<a target="_blank" rel="noopener noreferrer" '
        );
    };
    marked.use({ renderer });
</script>

<section
    tabindex="0"
    aria-label="Markdown content preview"
    class="mx-auto max-w-[83.333333%] p-6 overflow-y-auto font-sans leading-relaxed
      text-black bg-transparent
      dark:text-white"
    style="max-width: 83.333333%;"
    >
    {@html marked(content)}
</section>

<style>
    section {
        margin-top: 0;
        margin-bottom: 0;
        /* No text centering */
        text-align: left;
    }

    h1,
    h2,
    h3,
    h4,
    h5,
    h6 {
        font-weight: 600;
        margin-top: 1.5rem;
        margin-bottom: 0.75rem;
        color: inherit;
    }

    p {
        margin-bottom: 1rem;
    }

    a {
        color: #3b82f6;
        text-decoration: underline;
    }
    a:hover {
        color: #2563eb;
    }
    :global(.dark) a {
        color: #60a5fa;
    }
    :global(.dark) a:hover {
        color: #3b82f6;
    }

    code {
        background: #f5f5f5;
        padding: 0.2rem 0.4rem;
        border-radius: 4px;
        font-family: "Fira Mono", monospace;
    }
    :global(.dark) code {
        background: #2d2d2d;
        color: #f8f8f2;
    }

    pre {
        background: #1e1e1e;
        color: #d4d4d4;
        padding: 1rem;
        border-radius: 6px;
        overflow-x: auto;
        font-family: "Fira Mono", monospace;
        margin-bottom: 1rem;
    }
    :global(:not(.dark)) pre {
        background: #f5f5f5;
        color: #333;
    }

    ul,
    ol {
        margin-left: 1.25rem;
        margin-bottom: 1rem;
    }

    blockquote {
        border-left: 4px solid #ddd;
        padding-left: 1rem;
        color: #666;
        font-style: italic;
        margin-bottom: 1rem;
    }
    :global(.dark) blockquote {
        border-left-color: #555;
        color: #aaa;
    }
</style>
