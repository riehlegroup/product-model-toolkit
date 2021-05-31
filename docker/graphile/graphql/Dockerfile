# SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
#
# SPDX-License-Identifier: Apache-2.0

FROM node:alpine

RUN npm install -g postgraphile
RUN npm install -g postgraphile-plugin-connection-filter
RUN npm install -g postgraphile-audit-plugin tslib pg

EXPOSE 5000
ENTRYPOINT ["postgraphile", "-n", "0.0.0.0"]
