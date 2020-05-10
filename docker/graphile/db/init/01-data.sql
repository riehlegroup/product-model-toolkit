/* SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
 *
 * SPDX-License-Identifier: Apache-2.0
 */

\connect product_model;

/*Create some dummy products*/
INSERT INTO public.product (name, version, vcs) VALUES
('First product', '0.1.0', 'github.com/first-product'),
('Second product', '0.2.0', 'github.com/second-product');

/*Create some dummy components*/
INSERT INTO public.component (name, version, package, license, product_id) VALUES 
('Comp A', '1.0.0', 'org.a', 'MIT', 1),
('Comp B', '1.1.0', 'org.b', 'MIT', 1),
('Comp C', '2.0.0', 'org.c', 'GPL', 1),
('Comp D', '1.2.0', 'org.d', 'MIT', 1),
('Comp X', '1.0.0-beta', 'org.x', 'Apache', 2),
('Comp Y', '3.2.0', 'org.y', 'unkown', 2);