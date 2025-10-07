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

    // Intercept internal anchor clicks
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
        anchor.addEventListener("click", (e) => {
            const hash = anchor.getAttribute("href");
            if (hash.length > 1 && document.querySelector(hash)) {
                e.preventDefault();
                history.pushState(null, '', hash); // update URL without reloading
            }
        });
    });
});
