document.addEventListener('DOMContentLoaded', function () {
    let addonList = new List('addon-wrapper', {
        valueNames: [
            'name',
            'stars',
            'description',
            'type',
            'updated_at',
        ],
        pagination: {
            innerWindow: 5,
            outerWindow: 1,
            item: "<li><a class='page' href='javascript:;' tabindex='4'></a></li>",
        },
        page: 30,
    });

    // Scroll to bottom on page click
    document.addEventListener('click', function (e) {
        if (e.target.classList.contains('page')) {
            e.preventDefault();
            const el = document.scrollingElement || document.documentElement;
            el.scrollTo({ top: el.scrollHeight });
        }
    });

    // Hide pagination if not needed
    if ((addonList.size() / addonList.page) <= 1) {
        const pagination = document.querySelector('.pagination');
        if (pagination) pagination.style.display = 'none';
    }

    // URL parameter handling
    const searchInput = document.getElementById('addon-search');

    function updateURL(searchTerm) {
        const url = new URL(window.location);
        if (searchTerm) {
            url.searchParams.set('search', searchTerm);
        } else {
            url.searchParams.delete('search');
        }
        window.history.replaceState({}, '', url);
    }

    // Fuzzy match algorithm (Bitap-like)
    function fuzzyMatch(str, pattern, options) {
        options = options || {};
        const location = options.location || 0;
        const distance = options.distance || 100;
        const threshold = options.threshold || 0.4;

        if (pattern === str) return true;
        if (pattern.length > 32) return false;

        // Build alphabet pattern
        const patternAlphabet = {};
        for (let i = 0; i < pattern.length; i++) {
            patternAlphabet[pattern.charAt(i)] = 0;
        }
        for (let i = 0; i < pattern.length; i++) {
            patternAlphabet[pattern.charAt(i)] |= 1 << (pattern.length - i - 1);
        }

        function bitapScore(errors, x) {
            const accuracy = errors / pattern.length;
            const proximity = Math.abs(location - x);
            if (!distance) return proximity ? 1.0 : accuracy;
            return accuracy + (proximity / distance);
        }

        let scoreThreshold = threshold;
        let bestLoc = str.indexOf(pattern, location);
        if (bestLoc !== -1) {
            scoreThreshold = Math.min(bitapScore(0, bestLoc), scoreThreshold);
            bestLoc = str.lastIndexOf(pattern, location + pattern.length);
            if (bestLoc !== -1) {
                scoreThreshold = Math.min(bitapScore(0, bestLoc), scoreThreshold);
            }
        }

        const matchmask = 1 << (pattern.length - 1);
        bestLoc = -1;

        let binMin, binMid;
        let binMax = pattern.length + str.length;
        let lastRd;

        for (let d = 0; d < pattern.length; d++) {
            binMin = 0;
            binMid = binMax;
            while (binMin < binMid) {
                if (bitapScore(d, location + binMid) <= scoreThreshold) {
                    binMin = binMid;
                } else {
                    binMax = binMid;
                }
                binMid = Math.floor((binMax - binMin) / 2 + binMin);
            }

            binMax = binMid;
            let start = Math.max(1, location - binMid + 1);
            const finish = Math.min(location + binMid, str.length) + pattern.length;

            const rd = Array(finish + 2);
            rd[finish + 1] = (1 << d) - 1;

            for (let j = finish; j >= start; j--) {
                const charMatch = patternAlphabet[str.charAt(j - 1)];
                if (d === 0) {
                    rd[j] = ((rd[j + 1] << 1) | 1) & charMatch;
                } else {
                    rd[j] = (((rd[j + 1] << 1) | 1) & charMatch) | (((lastRd[j + 1] | lastRd[j]) << 1) | 1) | lastRd[j + 1];
                }
                if (rd[j] & matchmask) {
                    const score = bitapScore(d, j - 1);
                    if (score <= scoreThreshold) {
                        scoreThreshold = score;
                        bestLoc = j - 1;
                        if (bestLoc <= location) {
                            break;
                        }
                        start = Math.max(1, 2 * location - bestLoc);
                    }
                }
            }

            if (bitapScore(d + 1, location) > scoreThreshold) {
                break;
            }
            lastRd = rd;
        }

        return bestLoc >= 0;
    }

    // Custom search with fuzzy matching
    function performFuzzySearch(searchTerm) {
        if (!searchTerm) {
            addonList.filter();
            addonList.search();
            return;
        }

        // Replace "/" with space so "owner/repo" becomes separate terms
        searchTerm = searchTerm.replace(/\//g, ' ');

        const terms = searchTerm.trim().split(/\s+/);

        addonList.filter(function(item) {
            const name = item.values().name.toLowerCase();
            const description = item.values().description.toLowerCase();

            // Check if ALL terms match (either in name or description)
            for (let term of terms) {
                let termLower = term.toLowerCase();

                // Strip all 'ddev-' or 'ddev' prefixes (e.g., "ddev/ddev-redis" -> "redis")
                termLower = termLower.replace(/ddev-?/g, '');

                // Skip empty terms after stripping
                if (!termLower) continue;

                // Try exact match first (faster)
                if (name.includes(termLower) || description.includes(termLower)) {
                    continue; // This term matches, check next term
                }

                // Try fuzzy match
                const nameMatch = fuzzyMatch(name, termLower, {
                    location: 0,
                    distance: 100,
                    threshold: 0.4
                });
                const descMatch = fuzzyMatch(description, termLower, {
                    location: 0,
                    distance: 100,
                    threshold: 0.4
                });

                if (!nameMatch && !descMatch) {
                    return false; // This term doesn't match, filter out this item
                }
            }

            // All terms matched
            return true;
        });
    }

    // Manual fuzzy search with URL updates
    if (searchInput) {
        searchInput.addEventListener('input', function (e) {
            const searchTerm = e.target.value;
            performFuzzySearch(searchTerm);
            updateURL(searchTerm);
        });
    }

    // Check for search parameter in URL on page load
    const urlParams = new URLSearchParams(window.location.search);
    const searchParam = urlParams.get('search');
    if (searchParam && searchInput) {
        searchInput.value = searchParam;
        performFuzzySearch(searchParam);
    }
});
