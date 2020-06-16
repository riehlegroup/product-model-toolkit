#! /bin/sh

# SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
#
# SPDX-License-Identifier: Apache-2.0

docker-compose rm -f db
docker-compose build db