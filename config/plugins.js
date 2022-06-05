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
});