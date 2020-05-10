/* SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
 *
 * SPDX-License-Identifier: Apache-2.0
 */
 
\connect product_model;

CREATE TABLE public.product (
    id SERIAL PRIMARY KEY,
    name TEXT,
    version TEXT,
    vcs TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE public.component(
    id SERIAL PRIMARY KEY,
    name TEXT,
    version TEXT,
    package TEXT,
    license TEXT,
    product_id INTEGER NOT NULL REFERENCES public.product(id)

);