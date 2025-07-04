<script lang="ts">
    import { marked } from "marked";
    import hljs from "highlight.js";
    import "highlight.js/styles/github-dark.css";
    import { gfmHeadingId } from "marked-gfm-heading-id";
    import { markedEmoji } from "marked-emoji";
    import DOMPurify from "dompurify";

    let { content = "" } = $props<{ content?: string }>();

    // Huge usable emoji set
    const emojiMap = {
        smile: "https://github.githubassets.com/images/icons/emoji/unicode/1f604.png?v8",
        grin: "https://github.githubassets.com/images/icons/emoji/unicode/1f601.png?v8",
        joy: "https://github.githubassets.com/images/icons/emoji/unicode/1f602.png?v8",
        rofl: "https://github.githubassets.com/images/icons/emoji/unicode/1f923.png?v8",
        wink: "https://github.githubassets.com/images/icons/emoji/unicode/1f609.png?v8",
        blush: "https://github.githubassets.com/images/icons/emoji/unicode/1f60a.png?v8",
        heart_eyes:
            "https://github.githubassets.com/images/icons/emoji/unicode/1f60d.png?v8",
        smirk: "https://github.githubassets.com/images/icons/emoji/unicode/1f60f.png?v8",
        thinking:
            "https://github.githubassets.com/images/icons/emoji/unicode/1f914.png?v8",
        neutral_face:
            "https://github.githubassets.com/images/icons/emoji/unicode/1f610.png?v8",
        expressionless:
            "https://github.githubassets.com/images/icons/emoji/unicode/1f611.png?v8",
        unamused:
            "https://github.githubassets.com/images/icons/emoji/unicode/1f612.png?v8",
        sweat_smile:
            "https://github.githubassets.com/images/icons/emoji/unicode/1f605.png?v8",
        sob: "https://github.githubassets.com/images/icons/emoji/unicode/1f62d.png?v8",
        cry: "https://github.githubassets.com/images/icons/emoji/unicode/1f622.png?v8",
        angry: "https://github.githubassets.com/images/icons/emoji/unicode/1f620.png?v8",
        rage: "https://github.githubassets.com/images/icons/emoji/unicode/1f621.png?v8",
        thumbs_up:
            "https://github.githubassets.com/images/icons/emoji/unicode/1f44d.png?v8",
        thumbs_down:
            "https://github.githubassets.com/images/icons/emoji/unicode/1f44e.png?v8",
        clap: "https://github.githubassets.com/images/icons/emoji/unicode/1f44f.png?v8",
        ok_hand:
            "https://github.githubassets.com/images/icons/emoji/unicode/1f44c.png?v8",
        wave: "https://github.githubassets.com/images/icons/emoji/unicode/1f44b.png?v8",
        pray: "https://github.githubassets.com/images/icons/emoji/unicode/1f64f.png?v8",
        muscle: "https://github.githubassets.com/images/icons/emoji/unicode/1f4aa.png?v8",
        eyes: "https://github.githubassets.com/images/icons/emoji/unicode/1f440.png?v8",
        fire: "https://github.githubassets.com/images/icons/emoji/unicode/1f525.png?v8",
        sparkles:
            "https://github.githubassets.com/images/icons/emoji/unicode/2728.png?v8",
        star: "https://github.githubassets.com/images/icons/emoji/unicode/2b50.png?v8",
        heart: "https://github.githubassets.com/images/icons/emoji/unicode/2764.png?v8",
        broken_heart:
            "https://github.githubassets.com/images/icons/emoji/unicode/1f494.png?v8",
        "100": "https://github.githubassets.com/images/icons/emoji/unicode/1f4af.png?v8",
        tada: "https://github.githubassets.com/images/icons/emoji/unicode/1f389.png?v8",
        balloon:
            "https://github.githubassets.com/images/icons/emoji/unicode/1f388.png?v8",
        gift: "https://github.githubassets.com/images/icons/emoji/unicode/1f381.png?v8",
        rocket: "https://github.githubassets.com/images/icons/emoji/unicode/1f680.png?v8",
        boom: "https://github.githubassets.com/images/icons/emoji/unicode/1f4a5.png?v8",
        poop: "https://github.githubassets.com/images/icons/emoji/unicode/1f4a9.png?v8",
        see_no_evil:
            "https://github.githubassets.com/images/icons/emoji/unicode/1f648.png?v8",
        hear_no_evil:
            "https://github.githubassets.com/images/icons/emoji/unicode/1f649.png?v8",
        speak_no_evil:
            "https://github.githubassets.com/images/icons/emoji/unicode/1f64a.png?v8",
        skull: "https://github.githubassets.com/images/icons/emoji/unicode/1f480.png?v8",
        coffee: "https://github.githubassets.com/images/icons/emoji/unicode/2615.png?v8",
        tea: "https://github.githubassets.com/images/icons/emoji/unicode/1f375.png?v8",
        beer: "https://github.githubassets.com/images/icons/emoji/unicode/1f37a.png?v8",
        pizza: "https://github.githubassets.com/images/icons/emoji/unicode/1f355.png?v8",
        cake: "https://github.githubassets.com/images/icons/emoji/unicode/1f370.png?v8",
        burger: "https://github.githubassets.com/images/icons/emoji/unicode/1f354.png?v8",
        fries: "https://github.githubassets.com/images/icons/emoji/unicode/1f35f.png?v8",
        candy: "https://github.githubassets.com/images/icons/emoji/unicode/1f36c.png?v8",
        cookie: "https://github.githubassets.com/images/icons/emoji/unicode/1f36a.png?v8",
        apple: "https://github.githubassets.com/images/icons/emoji/unicode/1f34e.png?v8",
        lemon: "https://github.githubassets.com/images/icons/emoji/unicode/1f34b.png?v8",
        watermelon:
            "https://github.githubassets.com/images/icons/emoji/unicode/1f349.png?v8",
        cherries:
            "https://github.githubassets.com/images/icons/emoji/unicode/1f352.png?v8",
        strawberry:
            "https://github.githubassets.com/images/icons/emoji/unicode/1f353.png?v8",
    };

    marked.use([markedEmoji({ emojis: emojiMap }), gfmHeadingId()]);

    marked.setOptions({
        gfm: true,
        breaks: true,
        headerIds: true,
        mangle: false,
        highlight: function (code, lang) {
            if (lang && hljs.getLanguage(lang)) {
                return hljs.highlight(code, { language: lang }).value;
            } else {
                return hljs.highlightAuto(code).value;
            }
        },
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

    renderer.hr = function () {
        return `<hr class="my-6 border-gray-300 dark:border-gray-600" />`;
    };

    marked.use({ renderer });
</script>

<section
    tabindex="0"
    aria-label="Markdown content preview"
    class="prose dark:prose-invert max-w-full p-4 sm:p-6 overflow-y-auto font-sans leading-relaxed
      text-gray-900 dark:text-gray-100 bg-transparent"
>
    {@html DOMPurify.sanitize(marked(content))}
</section>

<style>
    section {
        text-align: left;
    }

    pre code {
        display: block;
        overflow-x: auto;
        padding: 1em;
        border-radius: 0.375rem;
    }

    pre {
        background: #0d1117; /* GitHub dark */
        color: #c9d1d9;
        border-radius: 0.375rem;
        padding: 1em;
        margin-bottom: 1em;
    }

    code {
        background: rgba(110, 118, 129, 0.4);
        padding: 0.2em 0.4em;
        border-radius: 0.375rem;
        font-family: "Fira Mono", monospace;
        font-size: 0.9em;
    }

    a {
        text-decoration: underline;
    }

    blockquote {
        border-left: 4px solid #d1d5db;
        padding-left: 1rem;
        color: #6b7280;
        font-style: italic;
        margin-bottom: 1rem;
    }
    :global(.dark) blockquote {
        border-left-color: #4b5563;
        color: #9ca3af;
    }

    hr {
        border: none;
        border-top: 1px solid #d1d5db;
        margin: 2rem 0;
    }
    :global(.dark) hr {
        border-top-color: #4b5563;
    }
</style>
