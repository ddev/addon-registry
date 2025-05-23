document.addEventListener('DOMContentLoaded', () => {
    const themeSwitch = document.getElementById("themeSwitch");
    const knob = document.getElementById("switchKnob");
    const options = themeSwitch.querySelectorAll(".theme-option");

    function setKnobPosition(theme) {
        const index = ["auto", "light", "dark"].indexOf(theme);
        const position = index * 108;
        knob.style.transform = `translateX(${position}%)`;

        options.forEach((option, i) => {
            option.setAttribute('tabindex', '6');
            option.setAttribute('aria-checked', i === index ? 'true' : 'false');
        });
    }

    function applyTheme(theme) {
        document.documentElement.removeAttribute('data-theme');

        if (theme === "dark") {
            document.documentElement.setAttribute('data-theme', 'dark');
            updateGiscusTheme('dark_dimmed');
        } else if (theme === "light") {
            document.documentElement.setAttribute('data-theme', 'light');
            updateGiscusTheme('light');
        } else if (theme === "auto") {
            if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
                document.documentElement.setAttribute('data-theme', 'dark');
                updateGiscusTheme('dark_dimmed');
            } else {
                document.documentElement.setAttribute('data-theme', 'light');
                updateGiscusTheme('light');
            }
        }

        setKnobPosition(theme);
        localStorage.setItem("theme", theme);
    }

    function updateGiscusTheme(theme) {
        const iframe = document.querySelector('iframe.giscus-frame');
        if (!iframe) return;

        iframe.contentWindow.postMessage({
            giscus: {
                setConfig: {
                    theme: theme
                }
            }
        }, 'https://giscus.app');
    }

    // Listen for Giscus to load
    window.addEventListener('message', (event) => {
        if (event.origin !== 'https://giscus.app') return;
        if (!(typeof event.data === 'object' && event.data.giscus)) return;

        // When Giscus is ready, set the initial theme
        const savedTheme = localStorage.getItem("theme") || "auto";
        if (savedTheme === 'light') {
            updateGiscusTheme('light');
        } else if (savedTheme === 'dark' || (savedTheme === 'auto' && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
            updateGiscusTheme('dark_dimmed');
        }
    });

    // Initialize theme
    const savedTheme = localStorage.getItem("theme") || "auto";
    applyTheme(savedTheme);

    // Theme switch event listeners
    options.forEach(option => {
        option.addEventListener("click", () => {
            const theme = option.dataset.theme;
            applyTheme(theme);
        });

        // Keyboard accessibility
        option.addEventListener("keydown", (e) => {
            if (e.key === 'Enter' || e.key === ' ') {
                e.preventDefault();
                const theme = option.dataset.theme;
                applyTheme(theme);
            }
        });
    });

    // System theme change listener
    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', () => {
        const currentTheme = localStorage.getItem("theme") || "auto";
        if (currentTheme === "auto") {
            applyTheme("auto");
        }
    });
});
