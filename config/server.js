// SPDX-FileCopyrightText: 2015-2022 Strapi Solutions SAS
// SPDX-License-Identifier: MIT

module.exports = ({ env }) => ({
  host: env('HOST', '0.0.0.0'),
  port: env.int('PORT', 1337),
  app: {
    keys: env.array('APP_KEYS'),
  },
});
