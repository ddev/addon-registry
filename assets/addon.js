function toggleContent(expand = false) {
    const content = document.getElementById("hidden-content");
    const btn = document.getElementById("read-more-btn");

    if (content.style.display === "none" || expand) {
        content.style.display = "block";
        btn.textContent = "Read less";
        btn.title = "Read less about add-on";
        btn.ariaLabel = "Read less about add-on";
    } else {
        content.style.display = "none";
        btn.textContent = "Read more";
        btn.title = "Read more about add-on";
        btn.ariaLabel = "Read more about add-on";
    }
    btn.blur();
}

function expandHash(hash) {
    if (!hash) return;

    const target = document.querySelector(hash);
    const content = document.getElementById("hidden-content");

    if (target && content && content.contains(target)) {
        toggleContent(true);

        // Wait for layout update, then scroll
        setTimeout(() => {
            target.scrollIntoView({ behavior: "smooth", block: "start" });
        }, 100);
    }
}

document.addEventListener("DOMContentLoaded", function () {
    // Add permalinks to headings
    document.querySelectorAll('h1, h2, h3, h4, h5').forEach((heading) => {
        if (heading.id) {
            const link = document.createElement('a');
            link.style.textDecoration = 'none';
            link.style.color = 'inherit';
            link.href = `#${heading.id}`;

            // Move all heading content inside the link
            while (heading.firstChild) {
                link.appendChild(heading.firstChild);
            }

            heading.appendChild(link);
        }
    });

    // Open external links in new tab
    document.querySelectorAll('a[href^="http"]').forEach((link) => {
        if (link.hostname !== window.location.hostname) {
            link.setAttribute("target", "_blank");
            link.setAttribute("rel", "noopener noreferrer");
        }
    });

    // Expand hash on load
    expandHash(window.location.hash);

    // Intercept internal anchor clicks
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
        anchor.addEventListener("click", (e) => {
            const hash = anchor.getAttribute("href");
            if (hash.length > 1 && document.querySelector(hash)) {
                e.preventDefault();
                history.pushState(null, '', hash); // update URL without reloading
                expandHash(hash);
            }
        });
    });
});

// Handle hash changes (e.g. if user edits URL manually)
window.addEventListener("hashchange", () => {
    expandHash(window.location.hash);
});
