// SPDX-FileCopyrightText: 2015-2022 Strapi Solutions SAS
// SPDX-License-Identifier: MIT

module.exports = ({ env }) => ({
    "fuzzy-search": {
        enabled: true,
        config: {
            contentTypes: [
                {
                    uid: "api::article.article",
                    modelName: "Title",
                    queryConstraints: {
                        where: {
                            $and: [
                                {
                                    publishedAt: { $notNull: true },
                                },
                            ],
                        },
                    },
                    fuzzysortOptions: {
                        characterLimit: 300,
                        threshold: -600,
                        limit: 10,
                        keys: [
                            {
                                name: "Title",
                                weight: 100,
                            },
                            {
                                name: "Content",
                                weight: -100,
                            },
                        ],
                    },
                },
            ],
        },
    },
    'strapi-prometheus': {
        enabled: true,
        config: {
            // add prefix to all the prometheus metrics names.
            prefix: '',

            // use full url instead of matched url
            // if true sets path label to `/api/models/1`
            // if false sets path label as `/api/models/:id`
            fullURL: false,

            // include url query in the url label
            // if true sets path label to `/api/models?limit=1`
            // if false sets path label to `/api/models`
            includeQuery: false,

            // collect default metrics of `prom-client`
            defaultMetrics: true,
        }
    }
});