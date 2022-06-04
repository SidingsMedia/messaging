// SPDX-FileCopyrightText: 2015-2022 Strapi Solutions SAS
// SPDX-License-Identifier: MIT

module.exports = ({ env }) => ({
  auth: {
    secret: env('ADMIN_JWT_SECRET'),
  },
  apiToken: {
    salt: env('API_TOKEN_SALT'),
  },
});
