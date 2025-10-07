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
        page: 27,
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

    // Search input filtering
    let previousValue = "";
    const searchInput = document.getElementById('addon-search');

    if (searchInput) {
        searchInput.addEventListener('input', function (e) {
            const currentValue = e.target.value;
            if (previousValue && currentValue === '') {
                addonList.filter();
            } else {
                const searchTerm = currentValue.toLowerCase();
                addonList.filter(function (item) {
                    return item.values().name.toLowerCase().includes(searchTerm) ||
                        item.values().description.toLowerCase().includes(searchTerm);
                });
            }
            previousValue = currentValue;
        });
    }
});
