// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

// SPDXLicense represents license known by SPDX
type SPDXLicense struct {
	Name        string
	URL         string
	OSIApproved bool
}

// AvailableSPDXLicenses is a map of SPDX identifiers and their license information
var AvailableSPDXLicenses = map[string]SPDXLicense{
	"0BSD": {
		Name:        "BSD Zero Clause License",
		URL:         "http://landley.net/toybox/license.html",
		OSIApproved: true,
	},
	"AAL": {
		Name:        "Attribution Assurance License",
		URL:         "https://opensource.org/licenses/attribution",
		OSIApproved: true,
	},
	"ADSL": {
		Name:        "Amazon Digital Services License",
		URL:         "https://fedoraproject.org/wiki/Licensing/AmazonDigitalServicesLicense",
		OSIApproved: false,
	},
	"AFL-1.1": {
		Name:        "Academic Free License v1.1",
		URL:         "http://opensource.linux-mirror.org/licenses/afl-1.1.txt",
		OSIApproved: true,
	},
	"AFL-1.2": {
		Name:        "Academic Free License v1.2",
		URL:         "http://opensource.linux-mirror.org/licenses/afl-1.2.txt",
		OSIApproved: true,
	},
	"AFL-2.0": {
		Name:        "Academic Free License v2.0",
		URL:         "http://wayback.archive.org/web/20060924134533/http://www.opensource.org/licenses/afl-2.0.txt",
		OSIApproved: true,
	},
	"AFL-2.1": {
		Name:        "Academic Free License v2.1",
		URL:         "http://opensource.linux-mirror.org/licenses/afl-2.1.txt",
		OSIApproved: true,
	},
	"AFL-3.0": {
		Name:        "Academic Free License v3.0",
		URL:         "http://www.rosenlaw.com/AFL3.0.htm",
		OSIApproved: true,
	},
	"AGPL-1.0": {
		Name:        "Affero General Public License v1.0",
		URL:         "http://www.affero.org/oagpl.html",
		OSIApproved: false,
	},
	"AGPL-1.0-only": {
		Name:        "Affero General Public License v1.0 only",
		URL:         "http://www.affero.org/oagpl.html",
		OSIApproved: false,
	},
	"AGPL-1.0-or-later": {
		Name:        "Affero General Public License v1.0 or later",
		URL:         "http://www.affero.org/oagpl.html",
		OSIApproved: false,
	},
	"AGPL-3.0": {
		Name:        "GNU Affero General Public License v3.0",
		URL:         "https://www.gnu.org/licenses/agpl.txt",
		OSIApproved: true,
	},
	"AGPL-3.0-only": {
		Name:        "GNU Affero General Public License v3.0 only",
		URL:         "https://www.gnu.org/licenses/agpl.txt",
		OSIApproved: true,
	},
	"AGPL-3.0-or-later": {
		Name:        "GNU Affero General Public License v3.0 or later",
		URL:         "https://www.gnu.org/licenses/agpl.txt",
		OSIApproved: true,
	},
	"AMDPLPA": {
		Name:        "AMD's plpa_map.c License",
		URL:         "https://fedoraproject.org/wiki/Licensing/AMD_plpa_map_License",
		OSIApproved: false,
	},
	"AML": {
		Name:        "Apple MIT License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Apple_MIT_License",
		OSIApproved: false,
	},
	"AMPAS": {
		Name:        "Academy of Motion Picture Arts and Sciences BSD",
		URL:         "https://fedoraproject.org/wiki/Licensing/BSD#AMPASBSD",
		OSIApproved: false,
	},
	"ANTLR-PD": {
		Name:        "ANTLR Software Rights Notice",
		URL:         "http://www.antlr2.org/license.html",
		OSIApproved: false,
	},
	"APAFML": {
		Name:        "Adobe Postscript AFM License",
		URL:         "https://fedoraproject.org/wiki/Licensing/AdobePostscriptAFM",
		OSIApproved: false,
	},
	"APL-1.0": {
		Name:        "Adaptive Public License 1.0",
		URL:         "https://opensource.org/licenses/APL-1.0",
		OSIApproved: true,
	},
	"APSL-1.0": {
		Name:        "Apple Public Source License 1.0",
		URL:         "https://fedoraproject.org/wiki/Licensing/Apple_Public_Source_License_1.0",
		OSIApproved: true,
	},
	"APSL-1.1": {
		Name:        "Apple Public Source License 1.1",
		URL:         "http://www.opensource.apple.com/source/IOSerialFamily/IOSerialFamily-7/APPLE_LICENSE",
		OSIApproved: true,
	},
	"APSL-1.2": {
		Name:        "Apple Public Source License 1.2",
		URL:         "http://www.samurajdata.se/opensource/mirror/licenses/apsl.php",
		OSIApproved: true,
	},
	"APSL-2.0": {
		Name:        "Apple Public Source License 2.0",
		URL:         "http://www.opensource.apple.com/license/apsl/",
		OSIApproved: true,
	},
	"Abstyles": {
		Name:        "Abstyles License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Abstyles",
		OSIApproved: false,
	},
	"Adobe-2006": {
		Name:        "Adobe Systems Incorporated Source Code License Agreement",
		URL:         "https://fedoraproject.org/wiki/Licensing/AdobeLicense",
		OSIApproved: false,
	},
	"Adobe-Glyph": {
		Name:        "Adobe Glyph List License",
		URL:         "https://fedoraproject.org/wiki/Licensing/MIT#AdobeGlyph",
		OSIApproved: false,
	},
	"Afmparse": {
		Name:        "Afmparse License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Afmparse",
		OSIApproved: false,
	},
	"Aladdin": {
		Name:        "Aladdin Free Public License",
		URL:         "http://pages.cs.wisc.edu/~ghost/doc/AFPL/6.01/Public.htm",
		OSIApproved: false,
	},
	"Apache-1.0": {
		Name:        "Apache License 1.0",
		URL:         "http://www.apache.org/licenses/LICENSE-1.0",
		OSIApproved: false,
	},
	"Apache-1.1": {
		Name:        "Apache License 1.1",
		URL:         "http://apache.org/licenses/LICENSE-1.1",
		OSIApproved: true,
	},
	"Apache-2.0": {
		Name:        "Apache License 2.0",
		URL:         "http://www.apache.org/licenses/LICENSE-2.0",
		OSIApproved: true,
	},
	"Artistic-1.0": {
		Name:        "Artistic License 1.0",
		URL:         "https://opensource.org/licenses/Artistic-1.0",
		OSIApproved: true,
	},
	"Artistic-1.0-Perl": {
		Name:        "Artistic License 1.0 (Perl)",
		URL:         "http://dev.perl.org/licenses/artistic.html",
		OSIApproved: true,
	},
	"Artistic-1.0-cl8": {
		Name:        "Artistic License 1.0 w/clause 8",
		URL:         "https://opensource.org/licenses/Artistic-1.0",
		OSIApproved: true,
	},
	"Artistic-2.0": {
		Name:        "Artistic License 2.0",
		URL:         "http://www.perlfoundation.org/artistic_license_2_0",
		OSIApproved: true,
	},
	"BSD-1-Clause": {
		Name:        "BSD 1-Clause License",
		URL:         "https://svnweb.freebsd.org/base/head/include/ifaddrs.h?revision=326823",
		OSIApproved: false,
	},
	"BSD-2-Clause": {
		Name:        "BSD 2-Clause \"Simplified\" License",
		URL:         "https://opensource.org/licenses/BSD-2-Clause",
		OSIApproved: true,
	},
	"BSD-2-Clause-FreeBSD": {
		Name:        "BSD 2-Clause FreeBSD License",
		URL:         "http://www.freebsd.org/copyright/freebsd-license.html",
		OSIApproved: false,
	},
	"BSD-2-Clause-NetBSD": {
		Name:        "BSD 2-Clause NetBSD License",
		URL:         "http://www.netbsd.org/about/redistribution.html#default",
		OSIApproved: false,
	},
	"BSD-2-Clause-Patent": {
		Name:        "BSD-2-Clause Plus Patent License",
		URL:         "https://opensource.org/licenses/BSDplusPatent",
		OSIApproved: true,
	},
	"BSD-3-Clause": {
		Name:        "BSD 3-Clause \"New\" or \"Revised\" License",
		URL:         "https://opensource.org/licenses/BSD-3-Clause",
		OSIApproved: true,
	},
	"BSD-3-Clause-Attribution": {
		Name:        "BSD with attribution",
		URL:         "https://fedoraproject.org/wiki/Licensing/BSD_with_Attribution",
		OSIApproved: false,
	},
	"BSD-3-Clause-Clear": {
		Name:        "BSD 3-Clause Clear License",
		URL:         "http://labs.metacarta.com/license-explanation.html#license",
		OSIApproved: false,
	},
	"BSD-3-Clause-LBNL": {
		Name:        "Lawrence Berkeley National Labs BSD variant license",
		URL:         "https://fedoraproject.org/wiki/Licensing/LBNLBSD",
		OSIApproved: true,
	},
	"BSD-3-Clause-No-Nuclear-License": {
		Name:        "BSD 3-Clause No Nuclear License",
		URL:         "http://download.oracle.com/otn-pub/java/licenses/bsd.txt?AuthParam=1467140197_43d516ce1776bd08a58235a7785be1cc",
		OSIApproved: false,
	},
	"BSD-3-Clause-No-Nuclear-License-2014": {
		Name:        "BSD 3-Clause No Nuclear License 2014",
		URL:         "https://java.net/projects/javaeetutorial/pages/BerkeleyLicense",
		OSIApproved: false,
	},
	"BSD-3-Clause-No-Nuclear-Warranty": {
		Name:        "BSD 3-Clause No Nuclear Warranty",
		URL:         "https://jogamp.org/git/?p=gluegen.git;a=blob_plain;f=LICENSE.txt",
		OSIApproved: false,
	},
	"BSD-3-Clause-Open-MPI": {
		Name:        "BSD 3-Clause Open MPI variant",
		URL:         "https://www.open-mpi.org/community/license.php",
		OSIApproved: false,
	},
	"BSD-4-Clause": {
		Name:        "BSD 4-Clause \"Original\" or \"Old\" License",
		URL:         "http://directory.fsf.org/wiki/License:BSD_4Clause",
		OSIApproved: false,
	},
	"BSD-4-Clause-UC": {
		Name:        "BSD-4-Clause (University of California-Specific)",
		URL:         "http://www.freebsd.org/copyright/license.html",
		OSIApproved: false,
	},
	"BSD-Protection": {
		Name:        "BSD Protection License",
		URL:         "https://fedoraproject.org/wiki/Licensing/BSD_Protection_License",
		OSIApproved: false,
	},
	"BSD-Source-Code": {
		Name:        "BSD Source Code Attribution",
		URL:         "https://github.com/robbiehanson/CocoaHTTPServer/blob/master/LICENSE.txt",
		OSIApproved: false,
	},
	"BSL-1.0": {
		Name:        "Boost Software License 1.0",
		URL:         "http://www.boost.org/LICENSE_1_0.txt",
		OSIApproved: true,
	},
	"Bahyph": {
		Name:        "Bahyph License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Bahyph",
		OSIApproved: false,
	},
	"Barr": {
		Name:        "Barr License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Barr",
		OSIApproved: false,
	},
	"Beerware": {
		Name:        "Beerware License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Beerware",
		OSIApproved: false,
	},
	"BitTorrent-1.0": {
		Name:        "BitTorrent Open Source License v1.0",
		URL:         "http://sources.gentoo.org/cgi-bin/viewvc.cgi/gentoo-x86/licenses/BitTorrent?r1=1.1&r2=1.1.1.1&diff_format=s",
		OSIApproved: false,
	},
	"BitTorrent-1.1": {
		Name:        "BitTorrent Open Source License v1.1",
		URL:         "http://directory.fsf.org/wiki/License:BitTorrentOSL1.1",
		OSIApproved: false,
	},
	"BlueOak-1.0.0": {
		Name:        "Blue Oak Model License 1.0.0",
		URL:         "https://blueoakcouncil.org/license/1.0.0",
		OSIApproved: false,
	},
	"Borceux": {
		Name:        "Borceux license",
		URL:         "https://fedoraproject.org/wiki/Licensing/Borceux",
		OSIApproved: false,
	},
	"CATOSL-1.1": {
		Name:        "Computer Associates Trusted Open Source License 1.1",
		URL:         "https://opensource.org/licenses/CATOSL-1.1",
		OSIApproved: true,
	},
	"CC-BY-1.0": {
		Name:        "Creative Commons Attribution 1.0 Generic",
		URL:         "https://creativecommons.org/licenses/by/1.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-2.0": {
		Name:        "Creative Commons Attribution 2.0 Generic",
		URL:         "https://creativecommons.org/licenses/by/2.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-2.5": {
		Name:        "Creative Commons Attribution 2.5 Generic",
		URL:         "https://creativecommons.org/licenses/by/2.5/legalcode",
		OSIApproved: false,
	},
	"CC-BY-3.0": {
		Name:        "Creative Commons Attribution 3.0 Unported",
		URL:         "https://creativecommons.org/licenses/by/3.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-4.0": {
		Name:        "Creative Commons Attribution 4.0 International",
		URL:         "https://creativecommons.org/licenses/by/4.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-NC-1.0": {
		Name:        "Creative Commons Attribution Non Commercial 1.0 Generic",
		URL:         "https://creativecommons.org/licenses/by-nc/1.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-NC-2.0": {
		Name:        "Creative Commons Attribution Non Commercial 2.0 Generic",
		URL:         "https://creativecommons.org/licenses/by-nc/2.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-NC-2.5": {
		Name:        "Creative Commons Attribution Non Commercial 2.5 Generic",
		URL:         "https://creativecommons.org/licenses/by-nc/2.5/legalcode",
		OSIApproved: false,
	},
	"CC-BY-NC-3.0": {
		Name:        "Creative Commons Attribution Non Commercial 3.0 Unported",
		URL:         "https://creativecommons.org/licenses/by-nc/3.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-NC-4.0": {
		Name:        "Creative Commons Attribution Non Commercial 4.0 International",
		URL:         "https://creativecommons.org/licenses/by-nc/4.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-NC-ND-1.0": {
		Name:        "Creative Commons Attribution Non Commercial No Derivatives 1.0 Generic",
		URL:         "https://creativecommons.org/licenses/by-nd-nc/1.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-NC-ND-2.0": {
		Name:        "Creative Commons Attribution Non Commercial No Derivatives 2.0 Generic",
		URL:         "https://creativecommons.org/licenses/by-nc-nd/2.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-NC-ND-2.5": {
		Name:        "Creative Commons Attribution Non Commercial No Derivatives 2.5 Generic",
		URL:         "https://creativecommons.org/licenses/by-nc-nd/2.5/legalcode",
		OSIApproved: false,
	},
	"CC-BY-NC-ND-3.0": {
		Name:        "Creative Commons Attribution Non Commercial No Derivatives 3.0 Unported",
		URL:         "https://creativecommons.org/licenses/by-nc-nd/3.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-NC-ND-4.0": {
		Name:        "Creative Commons Attribution Non Commercial No Derivatives 4.0 International",
		URL:         "https://creativecommons.org/licenses/by-nc-nd/4.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-NC-SA-1.0": {
		Name:        "Creative Commons Attribution Non Commercial Share Alike 1.0 Generic",
		URL:         "https://creativecommons.org/licenses/by-nc-sa/1.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-NC-SA-2.0": {
		Name:        "Creative Commons Attribution Non Commercial Share Alike 2.0 Generic",
		URL:         "https://creativecommons.org/licenses/by-nc-sa/2.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-NC-SA-2.5": {
		Name:        "Creative Commons Attribution Non Commercial Share Alike 2.5 Generic",
		URL:         "https://creativecommons.org/licenses/by-nc-sa/2.5/legalcode",
		OSIApproved: false,
	},
	"CC-BY-NC-SA-3.0": {
		Name:        "Creative Commons Attribution Non Commercial Share Alike 3.0 Unported",
		URL:         "https://creativecommons.org/licenses/by-nc-sa/3.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-NC-SA-4.0": {
		Name:        "Creative Commons Attribution Non Commercial Share Alike 4.0 International",
		URL:         "https://creativecommons.org/licenses/by-nc-sa/4.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-ND-1.0": {
		Name:        "Creative Commons Attribution No Derivatives 1.0 Generic",
		URL:         "https://creativecommons.org/licenses/by-nd/1.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-ND-2.0": {
		Name:        "Creative Commons Attribution No Derivatives 2.0 Generic",
		URL:         "https://creativecommons.org/licenses/by-nd/2.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-ND-2.5": {
		Name:        "Creative Commons Attribution No Derivatives 2.5 Generic",
		URL:         "https://creativecommons.org/licenses/by-nd/2.5/legalcode",
		OSIApproved: false,
	},
	"CC-BY-ND-3.0": {
		Name:        "Creative Commons Attribution No Derivatives 3.0 Unported",
		URL:         "https://creativecommons.org/licenses/by-nd/3.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-ND-4.0": {
		Name:        "Creative Commons Attribution No Derivatives 4.0 International",
		URL:         "https://creativecommons.org/licenses/by-nd/4.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-SA-1.0": {
		Name:        "Creative Commons Attribution Share Alike 1.0 Generic",
		URL:         "https://creativecommons.org/licenses/by-sa/1.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-SA-2.0": {
		Name:        "Creative Commons Attribution Share Alike 2.0 Generic",
		URL:         "https://creativecommons.org/licenses/by-sa/2.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-SA-2.5": {
		Name:        "Creative Commons Attribution Share Alike 2.5 Generic",
		URL:         "https://creativecommons.org/licenses/by-sa/2.5/legalcode",
		OSIApproved: false,
	},
	"CC-BY-SA-3.0": {
		Name:        "Creative Commons Attribution Share Alike 3.0 Unported",
		URL:         "https://creativecommons.org/licenses/by-sa/3.0/legalcode",
		OSIApproved: false,
	},
	"CC-BY-SA-4.0": {
		Name:        "Creative Commons Attribution Share Alike 4.0 International",
		URL:         "https://creativecommons.org/licenses/by-sa/4.0/legalcode",
		OSIApproved: false,
	},
	"CC-PDDC": {
		Name:        "Creative Commons Public Domain Dedication and Certification",
		URL:         "https://creativecommons.org/licenses/publicdomain/",
		OSIApproved: false,
	},
	"CC0-1.0": {
		Name:        "Creative Commons Zero v1.0 Universal",
		URL:         "https://creativecommons.org/publicdomain/zero/1.0/legalcode",
		OSIApproved: false,
	},
	"CDDL-1.0": {
		Name:        "Common Development and Distribution License 1.0",
		URL:         "https://opensource.org/licenses/cddl1",
		OSIApproved: true,
	},
	"CDDL-1.1": {
		Name:        "Common Development and Distribution License 1.1",
		URL:         "http://glassfish.java.net/public/CDDL+GPL_1_1.html",
		OSIApproved: false,
	},
	"CDLA-Permissive-1.0": {
		Name:        "Community Data License Agreement Permissive 1.0",
		URL:         "https://cdla.io/permissive-1-0",
		OSIApproved: false,
	},
	"CDLA-Sharing-1.0": {
		Name:        "Community Data License Agreement Sharing 1.0",
		URL:         "https://cdla.io/sharing-1-0",
		OSIApproved: false,
	},
	"CECILL-1.0": {
		Name:        "CeCILL Free Software License Agreement v1.0",
		URL:         "http://www.cecill.info/licences/Licence_CeCILL_V1-fr.html",
		OSIApproved: false,
	},
	"CECILL-1.1": {
		Name:        "CeCILL Free Software License Agreement v1.1",
		URL:         "http://www.cecill.info/licences/Licence_CeCILL_V1.1-US.html",
		OSIApproved: false,
	},
	"CECILL-2.0": {
		Name:        "CeCILL Free Software License Agreement v2.0",
		URL:         "http://www.cecill.info/licences/Licence_CeCILL_V2-en.html",
		OSIApproved: false,
	},
	"CECILL-2.1": {
		Name:        "CeCILL Free Software License Agreement v2.1",
		URL:         "http://www.cecill.info/licences/Licence_CeCILL_V2.1-en.html",
		OSIApproved: true,
	},
	"CECILL-B": {
		Name:        "CeCILL-B Free Software License Agreement",
		URL:         "http://www.cecill.info/licences/Licence_CeCILL-B_V1-en.html",
		OSIApproved: false,
	},
	"CECILL-C": {
		Name:        "CeCILL-C Free Software License Agreement",
		URL:         "http://www.cecill.info/licences/Licence_CeCILL-C_V1-en.html",
		OSIApproved: false,
	},
	"CERN-OHL-1.1": {
		Name:        "CERN Open Hardware Licence v1.1",
		URL:         "https://www.ohwr.org/project/licenses/wikis/cern-ohl-v1.1",
		OSIApproved: false,
	},
	"CERN-OHL-1.2": {
		Name:        "CERN Open Hardware Licence v1.2",
		URL:         "https://www.ohwr.org/project/licenses/wikis/cern-ohl-v1.2",
		OSIApproved: false,
	},
	"CNRI-Jython": {
		Name:        "CNRI Jython License",
		URL:         "http://www.jython.org/license.html",
		OSIApproved: false,
	},
	"CNRI-Python": {
		Name:        "CNRI Python License",
		URL:         "https://opensource.org/licenses/CNRI-Python",
		OSIApproved: true,
	},
	"CNRI-Python-GPL-Compatible": {
		Name:        "CNRI Python Open Source GPL Compatible License Agreement",
		URL:         "http://www.python.org/download/releases/1.6.1/download_win/",
		OSIApproved: false,
	},
	"CPAL-1.0": {
		Name:        "Common Public Attribution License 1.0",
		URL:         "https://opensource.org/licenses/CPAL-1.0",
		OSIApproved: true,
	},
	"CPL-1.0": {
		Name:        "Common Public License 1.0",
		URL:         "https://opensource.org/licenses/CPL-1.0",
		OSIApproved: true,
	},
	"CPOL-1.02": {
		Name:        "Code Project Open License 1.02",
		URL:         "http://www.codeproject.com/info/cpol10.aspx",
		OSIApproved: false,
	},
	"CUA-OPL-1.0": {
		Name:        "CUA Office Public License v1.0",
		URL:         "https://opensource.org/licenses/CUA-OPL-1.0",
		OSIApproved: true,
	},
	"Caldera": {
		Name:        "Caldera License",
		URL:         "http://www.lemis.com/grog/UNIX/ancient-source-all.pdf",
		OSIApproved: false,
	},
	"ClArtistic": {
		Name:        "Clarified Artistic License",
		URL:         "http://gianluca.dellavedova.org/2011/01/03/clarified-artistic-license/",
		OSIApproved: false,
	},
	"Condor-1.1": {
		Name:        "Condor Public License v1.1",
		URL:         "http://research.cs.wisc.edu/condor/license.html#condor",
		OSIApproved: false,
	},
	"Crossword": {
		Name:        "Crossword License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Crossword",
		OSIApproved: false,
	},
	"CrystalStacker": {
		Name:        "CrystalStacker License",
		URL:         "https://fedoraproject.org/wiki/Licensing:CrystalStacker?rd=Licensing/CrystalStacker",
		OSIApproved: false,
	},
	"Cube": {
		Name:        "Cube License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Cube",
		OSIApproved: false,
	},
	"D-FSL-1.0": {
		Name:        "Deutsche Freie Software Lizenz",
		URL:         "http://www.dipp.nrw.de/d-fsl/lizenzen/",
		OSIApproved: false,
	},
	"DOC": {
		Name:        "DOC License",
		URL:         "http://www.cs.wustl.edu/~schmidt/ACE-copying.html",
		OSIApproved: false,
	},
	"DSDP": {
		Name:        "DSDP License",
		URL:         "https://fedoraproject.org/wiki/Licensing/DSDP",
		OSIApproved: false,
	},
	"Dotseqn": {
		Name:        "Dotseqn License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Dotseqn",
		OSIApproved: false,
	},
	"ECL-1.0": {
		Name:        "Educational Community License v1.0",
		URL:         "https://opensource.org/licenses/ECL-1.0",
		OSIApproved: true,
	},
	"ECL-2.0": {
		Name:        "Educational Community License v2.0",
		URL:         "https://opensource.org/licenses/ECL-2.0",
		OSIApproved: true,
	},
	"EFL-1.0": {
		Name:        "Eiffel Forum License v1.0",
		URL:         "http://www.eiffel-nice.org/license/forum.txt",
		OSIApproved: true,
	},
	"EFL-2.0": {
		Name:        "Eiffel Forum License v2.0",
		URL:         "http://www.eiffel-nice.org/license/eiffel-forum-license-2.html",
		OSIApproved: true,
	},
	"EPL-1.0": {
		Name:        "Eclipse Public License 1.0",
		URL:         "http://www.eclipse.org/legal/epl-v10.html",
		OSIApproved: true,
	},
	"EPL-2.0": {
		Name:        "Eclipse Public License 2.0",
		URL:         "https://www.eclipse.org/legal/epl-2.0",
		OSIApproved: true,
	},
	"EUDatagrid": {
		Name:        "EU DataGrid Software License",
		URL:         "http://eu-datagrid.web.cern.ch/eu-datagrid/license.html",
		OSIApproved: true,
	},
	"EUPL-1.0": {
		Name:        "European Union Public License 1.0",
		URL:         "http://ec.europa.eu/idabc/en/document/7330.html",
		OSIApproved: false,
	},
	"EUPL-1.1": {
		Name:        "European Union Public License 1.1",
		URL:         "https://joinup.ec.europa.eu/software/page/eupl/licence-eupl",
		OSIApproved: true,
	},
	"EUPL-1.2": {
		Name:        "European Union Public License 1.2",
		URL:         "https://joinup.ec.europa.eu/page/eupl-text-11-12",
		OSIApproved: true,
	},
	"Entessa": {
		Name:        "Entessa Public License v1.0",
		URL:         "https://opensource.org/licenses/Entessa",
		OSIApproved: true,
	},
	"ErlPL-1.1": {
		Name:        "Erlang Public License v1.1",
		URL:         "http://www.erlang.org/EPLICENSE",
		OSIApproved: false,
	},
	"Eurosym": {
		Name:        "Eurosym License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Eurosym",
		OSIApproved: false,
	},
	"FSFAP": {
		Name:        "FSF All Permissive License",
		URL:         "https://www.gnu.org/prep/maintain/html_node/License-Notices-for-Other-Files.html",
		OSIApproved: false,
	},
	"FSFUL": {
		Name:        "FSF Unlimited License",
		URL:         "https://fedoraproject.org/wiki/Licensing/FSF_Unlimited_License",
		OSIApproved: false,
	},
	"FSFULLR": {
		Name:        "FSF Unlimited License (with License Retention)",
		URL:         "https://fedoraproject.org/wiki/Licensing/FSF_Unlimited_License#License_Retention_Variant",
		OSIApproved: false,
	},
	"FTL": {
		Name:        "Freetype Project License",
		URL:         "http://freetype.fis.uniroma2.it/FTL.TXT",
		OSIApproved: false,
	},
	"Fair": {
		Name:        "Fair License",
		URL:         "http://fairlicense.org/",
		OSIApproved: true,
	},
	"Frameworx-1.0": {
		Name:        "Frameworx Open License 1.0",
		URL:         "https://opensource.org/licenses/Frameworx-1.0",
		OSIApproved: true,
	},
	"FreeImage": {
		Name:        "FreeImage Public License v1.0",
		URL:         "http://freeimage.sourceforge.net/freeimage-license.txt",
		OSIApproved: false,
	},
	"GFDL-1.1": {
		Name:        "GNU Free Documentation License v1.1",
		URL:         "https://www.gnu.org/licenses/old-licenses/fdl-1.1.txt",
		OSIApproved: false,
	},
	"GFDL-1.1-only": {
		Name:        "GNU Free Documentation License v1.1 only",
		URL:         "https://www.gnu.org/licenses/old-licenses/fdl-1.1.txt",
		OSIApproved: false,
	},
	"GFDL-1.1-or-later": {
		Name:        "GNU Free Documentation License v1.1 or later",
		URL:         "https://www.gnu.org/licenses/old-licenses/fdl-1.1.txt",
		OSIApproved: false,
	},
	"GFDL-1.2": {
		Name:        "GNU Free Documentation License v1.2",
		URL:         "https://www.gnu.org/licenses/old-licenses/fdl-1.2.txt",
		OSIApproved: false,
	},
	"GFDL-1.2-only": {
		Name:        "GNU Free Documentation License v1.2 only",
		URL:         "https://www.gnu.org/licenses/old-licenses/fdl-1.2.txt",
		OSIApproved: false,
	},
	"GFDL-1.2-or-later": {
		Name:        "GNU Free Documentation License v1.2 or later",
		URL:         "https://www.gnu.org/licenses/old-licenses/fdl-1.2.txt",
		OSIApproved: false,
	},
	"GFDL-1.3": {
		Name:        "GNU Free Documentation License v1.3",
		URL:         "https://www.gnu.org/licenses/fdl-1.3.txt",
		OSIApproved: false,
	},
	"GFDL-1.3-only": {
		Name:        "GNU Free Documentation License v1.3 only",
		URL:         "https://www.gnu.org/licenses/fdl-1.3.txt",
		OSIApproved: false,
	},
	"GFDL-1.3-or-later": {
		Name:        "GNU Free Documentation License v1.3 or later",
		URL:         "https://www.gnu.org/licenses/fdl-1.3.txt",
		OSIApproved: false,
	},
	"GL2PS": {
		Name:        "GL2PS License",
		URL:         "http://www.geuz.org/gl2ps/COPYING.GL2PS",
		OSIApproved: false,
	},
	"GPL-1.0": {
		Name:        "GNU General Public License v1.0 only",
		URL:         "https://www.gnu.org/licenses/old-licenses/gpl-1.0-standalone.html",
		OSIApproved: false,
	},
	"GPL-1.0+": {
		Name:        "GNU General Public License v1.0 or later",
		URL:         "https://www.gnu.org/licenses/old-licenses/gpl-1.0-standalone.html",
		OSIApproved: false,
	},
	"GPL-1.0-only": {
		Name:        "GNU General Public License v1.0 only",
		URL:         "https://www.gnu.org/licenses/old-licenses/gpl-1.0-standalone.html",
		OSIApproved: false,
	},
	"GPL-1.0-or-later": {
		Name:        "GNU General Public License v1.0 or later",
		URL:         "https://www.gnu.org/licenses/old-licenses/gpl-1.0-standalone.html",
		OSIApproved: false,
	},
	"GPL-2.0": {
		Name:        "GNU General Public License v2.0 only",
		URL:         "https://www.gnu.org/licenses/old-licenses/gpl-2.0-standalone.html",
		OSIApproved: true,
	},
	"GPL-2.0+": {
		Name:        "GNU General Public License v2.0 or later",
		URL:         "https://www.gnu.org/licenses/old-licenses/gpl-2.0-standalone.html",
		OSIApproved: true,
	},
	"GPL-2.0-only": {
		Name:        "GNU General Public License v2.0 only",
		URL:         "https://www.gnu.org/licenses/old-licenses/gpl-2.0-standalone.html",
		OSIApproved: true,
	},
	"GPL-2.0-or-later": {
		Name:        "GNU General Public License v2.0 or later",
		URL:         "https://www.gnu.org/licenses/old-licenses/gpl-2.0-standalone.html",
		OSIApproved: true,
	},
	"GPL-2.0-with-GCC-exception": {
		Name:        "GNU General Public License v2.0 w/GCC Runtime Library exception",
		URL:         "https://gcc.gnu.org/git/?p=gcc.git;a=blob;f=gcc/libgcc1.c;h=762f5143fc6eed57b6797c82710f3538aa52b40b;hb=cb143a3ce4fb417c68f5fa2691a1b1b1053dfba9#l10",
		OSIApproved: false,
	},
	"GPL-2.0-with-autoconf-exception": {
		Name:        "GNU General Public License v2.0 w/Autoconf exception",
		URL:         "http://ac-archive.sourceforge.net/doc/copyright.html",
		OSIApproved: false,
	},
	"GPL-2.0-with-bison-exception": {
		Name:        "GNU General Public License v2.0 w/Bison exception",
		URL:         "http://git.savannah.gnu.org/cgit/bison.git/tree/data/yacc.c?id=193d7c7054ba7197b0789e14965b739162319b5e#n141",
		OSIApproved: false,
	},
	"GPL-2.0-with-classpath-exception": {
		Name:        "GNU General Public License v2.0 w/Classpath exception",
		URL:         "https://www.gnu.org/software/classpath/license.html",
		OSIApproved: false,
	},
	"GPL-2.0-with-font-exception": {
		Name:        "GNU General Public License v2.0 w/Font exception",
		URL:         "https://www.gnu.org/licenses/gpl-faq.html#FontException",
		OSIApproved: false,
	},
	"GPL-3.0": {
		Name:        "GNU General Public License v3.0 only",
		URL:         "https://www.gnu.org/licenses/gpl-3.0-standalone.html",
		OSIApproved: true,
	},
	"GPL-3.0+": {
		Name:        "GNU General Public License v3.0 or later",
		URL:         "https://www.gnu.org/licenses/gpl-3.0-standalone.html",
		OSIApproved: true,
	},
	"GPL-3.0-only": {
		Name:        "GNU General Public License v3.0 only",
		URL:         "https://www.gnu.org/licenses/gpl-3.0-standalone.html",
		OSIApproved: true,
	},
	"GPL-3.0-or-later": {
		Name:        "GNU General Public License v3.0 or later",
		URL:         "https://www.gnu.org/licenses/gpl-3.0-standalone.html",
		OSIApproved: true,
	},
	"GPL-3.0-with-GCC-exception": {
		Name:        "GNU General Public License v3.0 w/GCC Runtime Library exception",
		URL:         "https://www.gnu.org/licenses/gcc-exception-3.1.html",
		OSIApproved: true,
	},
	"GPL-3.0-with-autoconf-exception": {
		Name:        "GNU General Public License v3.0 w/Autoconf exception",
		URL:         "https://www.gnu.org/licenses/autoconf-exception-3.0.html",
		OSIApproved: false,
	},
	"Giftware": {
		Name:        "Giftware License",
		URL:         "http://liballeg.org/license.html#allegro-4-the-giftware-license",
		OSIApproved: false,
	},
	"Glide": {
		Name:        "3dfx Glide License",
		URL:         "http://www.users.on.net/~triforce/glidexp/COPYING.txt",
		OSIApproved: false,
	},
	"Glulxe": {
		Name:        "Glulxe License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Glulxe",
		OSIApproved: false,
	},
	"HPND": {
		Name:        "Historical Permission Notice and Disclaimer",
		URL:         "https://opensource.org/licenses/HPND",
		OSIApproved: true,
	},
	"HPND-sell-variant": {
		Name:        "Historical Permission Notice and Disclaimer - sell variant",
		URL:         "https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/net/sunrpc/auth_gss/gss_generic_token.c?h=v4.19",
		OSIApproved: false,
	},
	"HaskellReport": {
		Name:        "Haskell Language Report License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Haskell_Language_Report_License",
		OSIApproved: false,
	},
	"IBM-pibs": {
		Name:        "IBM PowerPC Initialization and Boot Software",
		URL:         "http://git.denx.de/?p=u-boot.git;a=blob;f=arch/powerpc/cpu/ppc4xx/miiphy.c;h=297155fdafa064b955e53e9832de93bfb0cfb85b;hb=9fab4bf4cc077c21e43941866f3f2c196f28670d",
		OSIApproved: false,
	},
	"ICU": {
		Name:        "ICU License",
		URL:         "http://source.icu-project.org/repos/icu/icu/trunk/license.html",
		OSIApproved: false,
	},
	"IJG": {
		Name:        "Independent JPEG Group License",
		URL:         "http://dev.w3.org/cvsweb/Amaya/libjpeg/Attic/README?rev=1.2",
		OSIApproved: false,
	},
	"IPA": {
		Name:        "IPA Font License",
		URL:         "https://opensource.org/licenses/IPA",
		OSIApproved: true,
	},
	"IPL-1.0": {
		Name:        "IBM Public License v1.0",
		URL:         "https://opensource.org/licenses/IPL-1.0",
		OSIApproved: true,
	},
	"ISC": {
		Name:        "ISC License",
		URL:         "https://www.isc.org/downloads/software-support-policy/isc-license/",
		OSIApproved: true,
	},
	"ImageMagick": {
		Name:        "ImageMagick License",
		URL:         "http://www.imagemagick.org/script/license.php",
		OSIApproved: false,
	},
	"Imlib2": {
		Name:        "Imlib2 License",
		URL:         "http://trac.enlightenment.org/e/browser/trunk/imlib2/COPYING",
		OSIApproved: false,
	},
	"Info-ZIP": {
		Name:        "Info-ZIP License",
		URL:         "http://www.info-zip.org/license.html",
		OSIApproved: false,
	},
	"Intel": {
		Name:        "Intel Open Source License",
		URL:         "https://opensource.org/licenses/Intel",
		OSIApproved: true,
	},
	"Intel-ACPI": {
		Name:        "Intel ACPI Software License Agreement",
		URL:         "https://fedoraproject.org/wiki/Licensing/Intel_ACPI_Software_License_Agreement",
		OSIApproved: false,
	},
	"Interbase-1.0": {
		Name:        "Interbase Public License v1.0",
		URL:         "https://web.archive.org/web/20060319014854/http://info.borland.com/devsupport/interbase/opensource/IPL.html",
		OSIApproved: false,
	},
	"JPNIC": {
		Name:        "Japan Network Information Center License",
		URL:         "https://gitlab.isc.org/isc-projects/bind9/blob/master/COPYRIGHT#L366",
		OSIApproved: false,
	},
	"JSON": {
		Name:        "JSON License",
		URL:         "http://www.json.org/license.html",
		OSIApproved: false,
	},
	"JasPer-2.0": {
		Name:        "JasPer License",
		URL:         "http://www.ece.uvic.ca/~mdadams/jasper/LICENSE",
		OSIApproved: false,
	},
	"LAL-1.2": {
		Name:        "Licence Art Libre 1.2",
		URL:         "http://artlibre.org/licence/lal/licence-art-libre-12/",
		OSIApproved: false,
	},
	"LAL-1.3": {
		Name:        "Licence Art Libre 1.3",
		URL:         "https://artlibre.org/",
		OSIApproved: false,
	},
	"LGPL-2.0": {
		Name:        "GNU Library General Public License v2 only",
		URL:         "https://www.gnu.org/licenses/old-licenses/lgpl-2.0-standalone.html",
		OSIApproved: true,
	},
	"LGPL-2.0+": {
		Name:        "GNU Library General Public License v2 or later",
		URL:         "https://www.gnu.org/licenses/old-licenses/lgpl-2.0-standalone.html",
		OSIApproved: true,
	},
	"LGPL-2.0-only": {
		Name:        "GNU Library General Public License v2 only",
		URL:         "https://www.gnu.org/licenses/old-licenses/lgpl-2.0-standalone.html",
		OSIApproved: true,
	},
	"LGPL-2.0-or-later": {
		Name:        "GNU Library General Public License v2 or later",
		URL:         "https://www.gnu.org/licenses/old-licenses/lgpl-2.0-standalone.html",
		OSIApproved: true,
	},
	"LGPL-2.1": {
		Name:        "GNU Lesser General Public License v2.1 only",
		URL:         "https://www.gnu.org/licenses/old-licenses/lgpl-2.1-standalone.html",
		OSIApproved: true,
	},
	"LGPL-2.1+": {
		Name:        "GNU Library General Public License v2.1 or later",
		URL:         "https://www.gnu.org/licenses/old-licenses/lgpl-2.1-standalone.html",
		OSIApproved: true,
	},
	"LGPL-2.1-only": {
		Name:        "GNU Lesser General Public License v2.1 only",
		URL:         "https://www.gnu.org/licenses/old-licenses/lgpl-2.1-standalone.html",
		OSIApproved: true,
	},
	"LGPL-2.1-or-later": {
		Name:        "GNU Lesser General Public License v2.1 or later",
		URL:         "https://www.gnu.org/licenses/old-licenses/lgpl-2.1-standalone.html",
		OSIApproved: true,
	},
	"LGPL-3.0": {
		Name:        "GNU Lesser General Public License v3.0 only",
		URL:         "https://www.gnu.org/licenses/lgpl-3.0-standalone.html",
		OSIApproved: true,
	},
	"LGPL-3.0+": {
		Name:        "GNU Lesser General Public License v3.0 or later",
		URL:         "https://www.gnu.org/licenses/lgpl-3.0-standalone.html",
		OSIApproved: true,
	},
	"LGPL-3.0-only": {
		Name:        "GNU Lesser General Public License v3.0 only",
		URL:         "https://www.gnu.org/licenses/lgpl-3.0-standalone.html",
		OSIApproved: true,
	},
	"LGPL-3.0-or-later": {
		Name:        "GNU Lesser General Public License v3.0 or later",
		URL:         "https://www.gnu.org/licenses/lgpl-3.0-standalone.html",
		OSIApproved: true,
	},
	"LGPLLR": {
		Name:        "Lesser General Public License For Linguistic Resources",
		URL:         "http://www-igm.univ-mlv.fr/~unitex/lgpllr.html",
		OSIApproved: false,
	},
	"LPL-1.0": {
		Name:        "Lucent Public License Version 1.0",
		URL:         "https://opensource.org/licenses/LPL-1.0",
		OSIApproved: true,
	},
	"LPL-1.02": {
		Name:        "Lucent Public License v1.02",
		URL:         "http://plan9.bell-labs.com/plan9/license.html",
		OSIApproved: true,
	},
	"LPPL-1.0": {
		Name:        "LaTeX Project Public License v1.0",
		URL:         "http://www.latex-project.org/lppl/lppl-1-0.txt",
		OSIApproved: false,
	},
	"LPPL-1.1": {
		Name:        "LaTeX Project Public License v1.1",
		URL:         "http://www.latex-project.org/lppl/lppl-1-1.txt",
		OSIApproved: false,
	},
	"LPPL-1.2": {
		Name:        "LaTeX Project Public License v1.2",
		URL:         "http://www.latex-project.org/lppl/lppl-1-2.txt",
		OSIApproved: false,
	},
	"LPPL-1.3a": {
		Name:        "LaTeX Project Public License v1.3a",
		URL:         "http://www.latex-project.org/lppl/lppl-1-3a.txt",
		OSIApproved: false,
	},
	"LPPL-1.3c": {
		Name:        "LaTeX Project Public License v1.3c",
		URL:         "http://www.latex-project.org/lppl/lppl-1-3c.txt",
		OSIApproved: true,
	},
	"Latex2e": {
		Name:        "Latex2e License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Latex2e",
		OSIApproved: false,
	},
	"Leptonica": {
		Name:        "Leptonica License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Leptonica",
		OSIApproved: false,
	},
	"LiLiQ-P-1.1": {
		Name:        "Licence Libre du Québec – Permissive version 1.1",
		URL:         "https://forge.gouv.qc.ca/licence/fr/liliq-v1-1/",
		OSIApproved: true,
	},
	"LiLiQ-R-1.1": {
		Name:        "Licence Libre du Québec – Réciprocité version 1.1",
		URL:         "https://www.forge.gouv.qc.ca/participez/licence-logicielle/licence-libre-du-quebec-liliq-en-francais/licence-libre-du-quebec-reciprocite-liliq-r-v1-1/",
		OSIApproved: true,
	},
	"LiLiQ-Rplus-1.1": {
		Name:        "Licence Libre du Québec – Réciprocité forte version 1.1",
		URL:         "https://www.forge.gouv.qc.ca/participez/licence-logicielle/licence-libre-du-quebec-liliq-en-francais/licence-libre-du-quebec-reciprocite-forte-liliq-r-v1-1/",
		OSIApproved: true,
	},
	"Libpng": {
		Name:        "libpng License",
		URL:         "http://www.libpng.org/pub/png/src/libpng-LICENSE.txt",
		OSIApproved: false,
	},
	"Linux-OpenIB": {
		Name:        "Linux Kernel Variant of OpenIB.org license",
		URL:         "https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/drivers/infiniband/core/sa.h",
		OSIApproved: false,
	},
	"MIT": {
		Name:        "MIT License",
		URL:         "https://opensource.org/licenses/MIT",
		OSIApproved: true,
	},
	"MIT-0": {
		Name:        "MIT No Attribution",
		URL:         "https://github.com/aws/mit-0",
		OSIApproved: false,
	},
	"MIT-CMU": {
		Name:        "CMU License",
		URL:         "https://fedoraproject.org/wiki/Licensing:MIT?rd=Licensing/MIT#CMU_Style",
		OSIApproved: false,
	},
	"MIT-advertising": {
		Name:        "Enlightenment License (e16)",
		URL:         "https://fedoraproject.org/wiki/Licensing/MIT_With_Advertising",
		OSIApproved: false,
	},
	"MIT-enna": {
		Name:        "enna License",
		URL:         "https://fedoraproject.org/wiki/Licensing/MIT#enna",
		OSIApproved: false,
	},
	"MIT-feh": {
		Name:        "feh License",
		URL:         "https://fedoraproject.org/wiki/Licensing/MIT#feh",
		OSIApproved: false,
	},
	"MITNFA": {
		Name:        "MIT +no-false,-attribs license",
		URL:         "https://fedoraproject.org/wiki/Licensing/MITNFA",
		OSIApproved: false,
	},
	"MPL-1.0": {
		Name:        "Mozilla Public License 1.0",
		URL:         "http://www.mozilla.org/MPL/MPL-1.0.html",
		OSIApproved: true,
	},
	"MPL-1.1": {
		Name:        "Mozilla Public License 1.1",
		URL:         "http://www.mozilla.org/MPL/MPL-1.1.html",
		OSIApproved: true,
	},
	"MPL-2.0": {
		Name:        "Mozilla Public License 2.0",
		URL:         "http://www.mozilla.org/MPL/2.0/",
		OSIApproved: true,
	},
	"MPL-2.0-no-copyleft-exception": {
		Name:        "Mozilla Public License 2.0 (no copyleft exception)",
		URL:         "http://www.mozilla.org/MPL/2.0/",
		OSIApproved: true,
	},
	"MS-PL": {
		Name:        "Microsoft Public License",
		URL:         "http://www.microsoft.com/opensource/licenses.mspx",
		OSIApproved: true,
	},
	"MS-RL": {
		Name:        "Microsoft Reciprocal License",
		URL:         "http://www.microsoft.com/opensource/licenses.mspx",
		OSIApproved: true,
	},
	"MTLL": {
		Name:        "Matrix Template Library License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Matrix_Template_Library_License",
		OSIApproved: false,
	},
	"MakeIndex": {
		Name:        "MakeIndex License",
		URL:         "https://fedoraproject.org/wiki/Licensing/MakeIndex",
		OSIApproved: false,
	},
	"MirOS": {
		Name:        "The MirOS Licence",
		URL:         "https://opensource.org/licenses/MirOS",
		OSIApproved: true,
	},
	"Motosoto": {
		Name:        "Motosoto License",
		URL:         "https://opensource.org/licenses/Motosoto",
		OSIApproved: true,
	},
	"MulanPSL-1.0": {
		Name:        "Mulan Permissive Software License, Version 1",
		URL:         "https://license.coscl.org.cn/MulanPSL/",
		OSIApproved: false,
	},
	"Multics": {
		Name:        "Multics License",
		URL:         "https://opensource.org/licenses/Multics",
		OSIApproved: true,
	},
	"Mup": {
		Name:        "Mup License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Mup",
		OSIApproved: false,
	},
	"NASA-1.3": {
		Name:        "NASA Open Source Agreement 1.3",
		URL:         "http://ti.arc.nasa.gov/opensource/nosa/",
		OSIApproved: true,
	},
	"NBPL-1.0": {
		Name:        "Net Boolean Public License v1",
		URL:         "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=37b4b3f6cc4bf34e1d3dec61e69914b9819d8894",
		OSIApproved: false,
	},
	"NCSA": {
		Name:        "University of Illinois/NCSA Open Source License",
		URL:         "http://otm.illinois.edu/uiuc_openSource",
		OSIApproved: true,
	},
	"NGPL": {
		Name:        "Nethack General Public License",
		URL:         "https://opensource.org/licenses/NGPL",
		OSIApproved: true,
	},
	"NLOD-1.0": {
		Name:        "Norwegian Licence for Open Government Data",
		URL:         "http://data.norge.no/nlod/en/1.0",
		OSIApproved: false,
	},
	"NLPL": {
		Name:        "No Limit Public License",
		URL:         "https://fedoraproject.org/wiki/Licensing/NLPL",
		OSIApproved: false,
	},
	"NOSL": {
		Name:        "Netizen Open Source License",
		URL:         "http://bits.netizen.com.au/licenses/NOSL/nosl.txt",
		OSIApproved: false,
	},
	"NPL-1.0": {
		Name:        "Netscape Public License v1.0",
		URL:         "http://www.mozilla.org/MPL/NPL/1.0/",
		OSIApproved: false,
	},
	"NPL-1.1": {
		Name:        "Netscape Public License v1.1",
		URL:         "http://www.mozilla.org/MPL/NPL/1.1/",
		OSIApproved: false,
	},
	"NPOSL-3.0": {
		Name:        "Non-Profit Open Software License 3.0",
		URL:         "https://opensource.org/licenses/NOSL3.0",
		OSIApproved: true,
	},
	"NRL": {
		Name:        "NRL License",
		URL:         "http://web.mit.edu/network/isakmp/nrllicense.html",
		OSIApproved: false,
	},
	"NTP": {
		Name:        "NTP License",
		URL:         "https://opensource.org/licenses/NTP",
		OSIApproved: true,
	},
	"NTP-0": {
		Name:        "NTP No Attribution",
		URL:         "https://github.com/tytso/e2fsprogs/blob/master/lib/et/et_name.c",
		OSIApproved: false,
	},
	"Naumen": {
		Name:        "Naumen Public License",
		URL:         "https://opensource.org/licenses/Naumen",
		OSIApproved: true,
	},
	"Net-SNMP": {
		Name:        "Net-SNMP License",
		URL:         "http://net-snmp.sourceforge.net/about/license.html",
		OSIApproved: false,
	},
	"NetCDF": {
		Name:        "NetCDF license",
		URL:         "http://www.unidata.ucar.edu/software/netcdf/copyright.html",
		OSIApproved: false,
	},
	"Newsletr": {
		Name:        "Newsletr License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Newsletr",
		OSIApproved: false,
	},
	"Nokia": {
		Name:        "Nokia Open Source License",
		URL:         "https://opensource.org/licenses/nokia",
		OSIApproved: true,
	},
	"Noweb": {
		Name:        "Noweb License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Noweb",
		OSIApproved: false,
	},
	"Nunit": {
		Name:        "Nunit License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Nunit",
		OSIApproved: false,
	},
	"OCCT-PL": {
		Name:        "Open CASCADE Technology Public License",
		URL:         "http://www.opencascade.com/content/occt-public-license",
		OSIApproved: false,
	},
	"OCLC-2.0": {
		Name:        "OCLC Research Public License 2.0",
		URL:         "http://www.oclc.org/research/activities/software/license/v2final.htm",
		OSIApproved: true,
	},
	"ODC-By-1.0": {
		Name:        "Open Data Commons Attribution License v1.0",
		URL:         "https://opendatacommons.org/licenses/by/1.0/",
		OSIApproved: false,
	},
	"ODbL-1.0": {
		Name:        "ODC Open Database License v1.0",
		URL:         "http://www.opendatacommons.org/licenses/odbl/1.0/",
		OSIApproved: false,
	},
	"OFL-1.0": {
		Name:        "SIL Open Font License 1.0",
		URL:         "http://scripts.sil.org/cms/scripts/page.php?item_id=OFL10_web",
		OSIApproved: false,
	},
	"OFL-1.0-RFN": {
		Name:        "SIL Open Font License 1.0 with Reserved Font Name",
		URL:         "http://scripts.sil.org/cms/scripts/page.php?item_id=OFL10_web",
		OSIApproved: false,
	},
	"OFL-1.0-no-RFN": {
		Name:        "SIL Open Font License 1.0 with no Reserved Font Name",
		URL:         "http://scripts.sil.org/cms/scripts/page.php?item_id=OFL10_web",
		OSIApproved: false,
	},
	"OFL-1.1": {
		Name:        "SIL Open Font License 1.1",
		URL:         "http://scripts.sil.org/cms/scripts/page.php?item_id=OFL_web",
		OSIApproved: true,
	},
	"OFL-1.1-RFN": {
		Name:        "SIL Open Font License 1.1 with Reserved Font Name",
		URL:         "http://scripts.sil.org/cms/scripts/page.php?item_id=OFL_web",
		OSIApproved: true,
	},
	"OFL-1.1-no-RFN": {
		Name:        "SIL Open Font License 1.1 with no Reserved Font Name",
		URL:         "http://scripts.sil.org/cms/scripts/page.php?item_id=OFL_web",
		OSIApproved: true,
	},
	"OGL-Canada-2.0": {
		Name:        "Open Government Licence - Canada",
		URL:         "https://open.canada.ca/en/open-government-licence-canada",
		OSIApproved: false,
	},
	"OGL-UK-1.0": {
		Name:        "Open Government Licence v1.0",
		URL:         "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/1/",
		OSIApproved: false,
	},
	"OGL-UK-2.0": {
		Name:        "Open Government Licence v2.0",
		URL:         "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/2/",
		OSIApproved: false,
	},
	"OGL-UK-3.0": {
		Name:        "Open Government Licence v3.0",
		URL:         "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/3/",
		OSIApproved: false,
	},
	"OGTSL": {
		Name:        "Open Group Test Suite License",
		URL:         "http://www.opengroup.org/testing/downloads/The_Open_Group_TSL.txt",
		OSIApproved: true,
	},
	"OLDAP-1.1": {
		Name:        "Open LDAP Public License v1.1",
		URL:         "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=806557a5ad59804ef3a44d5abfbe91d706b0791f",
		OSIApproved: false,
	},
	"OLDAP-1.2": {
		Name:        "Open LDAP Public License v1.2",
		URL:         "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=42b0383c50c299977b5893ee695cf4e486fb0dc7",
		OSIApproved: false,
	},
	"OLDAP-1.3": {
		Name:        "Open LDAP Public License v1.3",
		URL:         "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=e5f8117f0ce088d0bd7a8e18ddf37eaa40eb09b1",
		OSIApproved: false,
	},
	"OLDAP-1.4": {
		Name:        "Open LDAP Public License v1.4",
		URL:         "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=c9f95c2f3f2ffb5e0ae55fe7388af75547660941",
		OSIApproved: false,
	},
	"OLDAP-2.0": {
		Name:        "Open LDAP Public License v2.0 (or possibly 2.0A and 2.0B)",
		URL:         "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=cbf50f4e1185a21abd4c0a54d3f4341fe28f36ea",
		OSIApproved: false,
	},
	"OLDAP-2.0.1": {
		Name:        "Open LDAP Public License v2.0.1",
		URL:         "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=b6d68acd14e51ca3aab4428bf26522aa74873f0e",
		OSIApproved: false,
	},
	"OLDAP-2.1": {
		Name:        "Open LDAP Public License v2.1",
		URL:         "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=b0d176738e96a0d3b9f85cb51e140a86f21be715",
		OSIApproved: false,
	},
	"OLDAP-2.2": {
		Name:        "Open LDAP Public License v2.2",
		URL:         "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=470b0c18ec67621c85881b2733057fecf4a1acc3",
		OSIApproved: false,
	},
	"OLDAP-2.2.1": {
		Name:        "Open LDAP Public License v2.2.1",
		URL:         "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=4bc786f34b50aa301be6f5600f58a980070f481e",
		OSIApproved: false,
	},
	"OLDAP-2.2.2": {
		Name:        "Open LDAP Public License 2.2.2",
		URL:         "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=df2cc1e21eb7c160695f5b7cffd6296c151ba188",
		OSIApproved: false,
	},
	"OLDAP-2.3": {
		Name:        "Open LDAP Public License v2.3",
		URL:         "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=d32cf54a32d581ab475d23c810b0a7fbaf8d63c3",
		OSIApproved: false,
	},
	"OLDAP-2.4": {
		Name:        "Open LDAP Public License v2.4",
		URL:         "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=cd1284c4a91a8a380d904eee68d1583f989ed386",
		OSIApproved: false,
	},
	"OLDAP-2.5": {
		Name:        "Open LDAP Public License v2.5",
		URL:         "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=6852b9d90022e8593c98205413380536b1b5a7cf",
		OSIApproved: false,
	},
	"OLDAP-2.6": {
		Name:        "Open LDAP Public License v2.6",
		URL:         "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=1cae062821881f41b73012ba816434897abf4205",
		OSIApproved: false,
	},
	"OLDAP-2.7": {
		Name:        "Open LDAP Public License v2.7",
		URL:         "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=47c2415c1df81556eeb39be6cad458ef87c534a2",
		OSIApproved: false,
	},
	"OLDAP-2.8": {
		Name:        "Open LDAP Public License v2.8",
		URL:         "http://www.openldap.org/software/release/license.html",
		OSIApproved: false,
	},
	"OML": {
		Name:        "Open Market License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Open_Market_License",
		OSIApproved: false,
	},
	"OPL-1.0": {
		Name:        "Open Public License v1.0",
		URL:         "http://old.koalateam.com/jackaroo/OPL_1_0.TXT",
		OSIApproved: false,
	},
	"OSET-PL-2.1": {
		Name:        "OSET Public License version 2.1",
		URL:         "http://www.osetfoundation.org/public-license",
		OSIApproved: true,
	},
	"OSL-1.0": {
		Name:        "Open Software License 1.0",
		URL:         "https://opensource.org/licenses/OSL-1.0",
		OSIApproved: true,
	},
	"OSL-1.1": {
		Name:        "Open Software License 1.1",
		URL:         "https://fedoraproject.org/wiki/Licensing/OSL1.1",
		OSIApproved: false,
	},
	"OSL-2.0": {
		Name:        "Open Software License 2.0",
		URL:         "http://web.archive.org/web/20041020171434/http://www.rosenlaw.com/osl2.0.html",
		OSIApproved: true,
	},
	"OSL-2.1": {
		Name:        "Open Software License 2.1",
		URL:         "http://web.archive.org/web/20050212003940/http://www.rosenlaw.com/osl21.htm",
		OSIApproved: true,
	},
	"OSL-3.0": {
		Name:        "Open Software License 3.0",
		URL:         "https://web.archive.org/web/20120101081418/http://rosenlaw.com:80/OSL3.0.htm",
		OSIApproved: true,
	},
	"OpenSSL": {
		Name:        "OpenSSL License",
		URL:         "http://www.openssl.org/source/license.html",
		OSIApproved: false,
	},
	"PDDL-1.0": {
		Name:        "ODC Public Domain Dedication & License 1.0",
		URL:         "http://opendatacommons.org/licenses/pddl/1.0/",
		OSIApproved: false,
	},
	"PHP-3.0": {
		Name:        "PHP License v3.0",
		URL:         "http://www.php.net/license/3_0.txt",
		OSIApproved: true,
	},
	"PHP-3.01": {
		Name:        "PHP License v3.01",
		URL:         "http://www.php.net/license/3_01.txt",
		OSIApproved: false,
	},
	"PSF-2.0": {
		Name:        "Python Software Foundation License 2.0",
		URL:         "https://opensource.org/licenses/Python-2.0",
		OSIApproved: false,
	},
	"Parity-6.0.0": {
		Name:        "The Parity Public License 6.0.0",
		URL:         "https://paritylicense.com/versions/6.0.0.html",
		OSIApproved: false,
	},
	"Plexus": {
		Name:        "Plexus Classworlds License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Plexus_Classworlds_License",
		OSIApproved: false,
	},
	"PostgreSQL": {
		Name:        "PostgreSQL License",
		URL:         "http://www.postgresql.org/about/licence",
		OSIApproved: true,
	},
	"Python-2.0": {
		Name:        "Python License 2.0",
		URL:         "https://opensource.org/licenses/Python-2.0",
		OSIApproved: true,
	},
	"QPL-1.0": {
		Name:        "Q Public License 1.0",
		URL:         "http://doc.qt.nokia.com/3.3/license.html",
		OSIApproved: true,
	},
	"Qhull": {
		Name:        "Qhull License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Qhull",
		OSIApproved: false,
	},
	"RHeCos-1.1": {
		Name:        "Red Hat eCos Public License v1.1",
		URL:         "http://ecos.sourceware.org/old-license.html",
		OSIApproved: false,
	},
	"RPL-1.1": {
		Name:        "Reciprocal Public License 1.1",
		URL:         "https://opensource.org/licenses/RPL-1.1",
		OSIApproved: true,
	},
	"RPL-1.5": {
		Name:        "Reciprocal Public License 1.5",
		URL:         "https://opensource.org/licenses/RPL-1.5",
		OSIApproved: true,
	},
	"RPSL-1.0": {
		Name:        "RealNetworks Public Source License v1.0",
		URL:         "https://helixcommunity.org/content/rpsl",
		OSIApproved: true,
	},
	"RSA-MD": {
		Name:        "RSA Message-Digest License ",
		URL:         "http://www.faqs.org/rfcs/rfc1321.html",
		OSIApproved: false,
	},
	"RSCPL": {
		Name:        "Ricoh Source Code Public License",
		URL:         "http://wayback.archive.org/web/20060715140826/http://www.risource.org/RPL/RPL-1.0A.shtml",
		OSIApproved: true,
	},
	"Rdisc": {
		Name:        "Rdisc License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Rdisc_License",
		OSIApproved: false,
	},
	"Ruby": {
		Name:        "Ruby License",
		URL:         "http://www.ruby-lang.org/en/LICENSE.txt",
		OSIApproved: false,
	},
	"SAX-PD": {
		Name:        "Sax Public Domain Notice",
		URL:         "http://www.saxproject.org/copying.html",
		OSIApproved: false,
	},
	"SCEA": {
		Name:        "SCEA Shared Source License",
		URL:         "http://research.scea.com/scea_shared_source_license.html",
		OSIApproved: false,
	},
	"SGI-B-1.0": {
		Name:        "SGI Free Software License B v1.0",
		URL:         "http://oss.sgi.com/projects/FreeB/SGIFreeSWLicB.1.0.html",
		OSIApproved: false,
	},
	"SGI-B-1.1": {
		Name:        "SGI Free Software License B v1.1",
		URL:         "http://oss.sgi.com/projects/FreeB/",
		OSIApproved: false,
	},
	"SGI-B-2.0": {
		Name:        "SGI Free Software License B v2.0",
		URL:         "http://oss.sgi.com/projects/FreeB/SGIFreeSWLicB.2.0.pdf",
		OSIApproved: false,
	},
	"SHL-0.5": {
		Name:        "Solderpad Hardware License v0.5",
		URL:         "https://solderpad.org/licenses/SHL-0.5/",
		OSIApproved: false,
	},
	"SHL-0.51": {
		Name:        "Solderpad Hardware License, Version 0.51",
		URL:         "https://solderpad.org/licenses/SHL-0.51/",
		OSIApproved: false,
	},
	"SISSL": {
		Name:        "Sun Industry Standards Source License v1.1",
		URL:         "http://www.openoffice.org/licenses/sissl_license.html",
		OSIApproved: true,
	},
	"SISSL-1.2": {
		Name:        "Sun Industry Standards Source License v1.2",
		URL:         "http://gridscheduler.sourceforge.net/Gridengine_SISSL_license.html",
		OSIApproved: false,
	},
	"SMLNJ": {
		Name:        "Standard ML of New Jersey License",
		URL:         "https://www.smlnj.org/license.html",
		OSIApproved: false,
	},
	"SMPPL": {
		Name:        "Secure Messaging Protocol Public License",
		URL:         "https://github.com/dcblake/SMP/blob/master/Documentation/License.txt",
		OSIApproved: false,
	},
	"SNIA": {
		Name:        "SNIA Public License 1.1",
		URL:         "https://fedoraproject.org/wiki/Licensing/SNIA_Public_License",
		OSIApproved: false,
	},
	"SPL-1.0": {
		Name:        "Sun Public License v1.0",
		URL:         "https://opensource.org/licenses/SPL-1.0",
		OSIApproved: true,
	},
	"SSH-OpenSSH": {
		Name:        "SSH OpenSSH license",
		URL:         "https://github.com/openssh/openssh-portable/blob/1b11ea7c58cd5c59838b5fa574cd456d6047b2d4/LICENCE#L10",
		OSIApproved: false,
	},
	"SSH-short": {
		Name:        "SSH short notice",
		URL:         "https://github.com/openssh/openssh-portable/blob/1b11ea7c58cd5c59838b5fa574cd456d6047b2d4/pathnames.h",
		OSIApproved: false,
	},
	"SSPL-1.0": {
		Name:        "Server Side Public License, v 1",
		URL:         "https://www.mongodb.com/licensing/server-side-public-license",
		OSIApproved: false,
	},
	"SWL": {
		Name:        "Scheme Widget Library (SWL) Software License Agreement",
		URL:         "https://fedoraproject.org/wiki/Licensing/SWL",
		OSIApproved: false,
	},
	"Saxpath": {
		Name:        "Saxpath License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Saxpath_License",
		OSIApproved: false,
	},
	"Sendmail": {
		Name:        "Sendmail License",
		URL:         "http://www.sendmail.com/pdfs/open_source/sendmail_license.pdf",
		OSIApproved: false,
	},
	"Sendmail-8.23": {
		Name:        "Sendmail License 8.23",
		URL:         "https://www.proofpoint.com/sites/default/files/sendmail-license.pdf",
		OSIApproved: false,
	},
	"SimPL-2.0": {
		Name:        "Simple Public License 2.0",
		URL:         "https://opensource.org/licenses/SimPL-2.0",
		OSIApproved: true,
	},
	"Sleepycat": {
		Name:        "Sleepycat License",
		URL:         "https://opensource.org/licenses/Sleepycat",
		OSIApproved: true,
	},
	"Spencer-86": {
		Name:        "Spencer License 86",
		URL:         "https://fedoraproject.org/wiki/Licensing/Henry_Spencer_Reg-Ex_Library_License",
		OSIApproved: false,
	},
	"Spencer-94": {
		Name:        "Spencer License 94",
		URL:         "https://fedoraproject.org/wiki/Licensing/Henry_Spencer_Reg-Ex_Library_License",
		OSIApproved: false,
	},
	"Spencer-99": {
		Name:        "Spencer License 99",
		URL:         "http://www.opensource.apple.com/source/tcl/tcl-5/tcl/generic/regfronts.c",
		OSIApproved: false,
	},
	"StandardML-NJ": {
		Name:        "Standard ML of New Jersey License",
		URL:         "http://www.smlnj.org//license.html",
		OSIApproved: false,
	},
	"SugarCRM-1.1.3": {
		Name:        "SugarCRM Public License v1.1.3",
		URL:         "http://www.sugarcrm.com/crm/SPL",
		OSIApproved: false,
	},
	"TAPR-OHL-1.0": {
		Name:        "TAPR Open Hardware License v1.0",
		URL:         "https://www.tapr.org/OHL",
		OSIApproved: false,
	},
	"TCL": {
		Name:        "TCL/TK License",
		URL:         "http://www.tcl.tk/software/tcltk/license.html",
		OSIApproved: false,
	},
	"TCP-wrappers": {
		Name:        "TCP Wrappers License",
		URL:         "http://rc.quest.com/topics/openssh/license.php#tcpwrappers",
		OSIApproved: false,
	},
	"TMate": {
		Name:        "TMate Open Source License",
		URL:         "http://svnkit.com/license.html",
		OSIApproved: false,
	},
	"TORQUE-1.1": {
		Name:        "TORQUE v2.5+ Software License v1.1",
		URL:         "https://fedoraproject.org/wiki/Licensing/TORQUEv1.1",
		OSIApproved: false,
	},
	"TOSL": {
		Name:        "Trusster Open Source License",
		URL:         "https://fedoraproject.org/wiki/Licensing/TOSL",
		OSIApproved: false,
	},
	"TU-Berlin-1.0": {
		Name:        "Technische Universitaet Berlin License 1.0",
		URL:         "https://github.com/swh/ladspa/blob/7bf6f3799fdba70fda297c2d8fd9f526803d9680/gsm/COPYRIGHT",
		OSIApproved: false,
	},
	"TU-Berlin-2.0": {
		Name:        "Technische Universitaet Berlin License 2.0",
		URL:         "https://github.com/CorsixTH/deps/blob/fd339a9f526d1d9c9f01ccf39e438a015da50035/licences/libgsm.txt",
		OSIApproved: false,
	},
	"UCL-1.0": {
		Name:        "Upstream Compatibility License v1.0",
		URL:         "https://opensource.org/licenses/UCL-1.0",
		OSIApproved: true,
	},
	"UPL-1.0": {
		Name:        "Universal Permissive License v1.0",
		URL:         "https://opensource.org/licenses/UPL",
		OSIApproved: true,
	},
	"Unicode-DFS-2015": {
		Name:        "Unicode License Agreement - Data Files and Software (2015)",
		URL:         "https://web.archive.org/web/20151224134844/http://unicode.org/copyright.html",
		OSIApproved: false,
	},
	"Unicode-DFS-2016": {
		Name:        "Unicode License Agreement - Data Files and Software (2016)",
		URL:         "http://www.unicode.org/copyright.html",
		OSIApproved: false,
	},
	"Unicode-TOU": {
		Name:        "Unicode Terms of Use",
		URL:         "http://www.unicode.org/copyright.html",
		OSIApproved: false,
	},
	"Unlicense": {
		Name:        "The Unlicense",
		URL:         "https://unlicense.org/",
		OSIApproved: false,
	},
	"VOSTROM": {
		Name:        "VOSTROM Public License for Open Source",
		URL:         "https://fedoraproject.org/wiki/Licensing/VOSTROM",
		OSIApproved: false,
	},
	"VSL-1.0": {
		Name:        "Vovida Software License v1.0",
		URL:         "https://opensource.org/licenses/VSL-1.0",
		OSIApproved: true,
	},
	"Vim": {
		Name:        "Vim License",
		URL:         "http://vimdoc.sourceforge.net/htmldoc/uganda.html",
		OSIApproved: false,
	},
	"W3C": {
		Name:        "W3C Software Notice and License (2002-12-31)",
		URL:         "http://www.w3.org/Consortium/Legal/2002/copyright-software-20021231.html",
		OSIApproved: true,
	},
	"W3C-19980720": {
		Name:        "W3C Software Notice and License (1998-07-20)",
		URL:         "http://www.w3.org/Consortium/Legal/copyright-software-19980720.html",
		OSIApproved: false,
	},
	"W3C-20150513": {
		Name:        "W3C Software Notice and Document License (2015-05-13)",
		URL:         "https://www.w3.org/Consortium/Legal/2015/copyright-software-and-document",
		OSIApproved: false,
	},
	"WTFPL": {
		Name:        "Do What The F*ck You Want To Public License",
		URL:         "http://sam.zoy.org/wtfpl/COPYING",
		OSIApproved: false,
	},
	"Watcom-1.0": {
		Name:        "Sybase Open Watcom Public License 1.0",
		URL:         "https://opensource.org/licenses/Watcom-1.0",
		OSIApproved: true,
	},
	"Wsuipa": {
		Name:        "Wsuipa License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Wsuipa",
		OSIApproved: false,
	},
	"X11": {
		Name:        "X11 License",
		URL:         "http://www.xfree86.org/3.3.6/COPYRIGHT2.html#3",
		OSIApproved: false,
	},
	"XFree86-1.1": {
		Name:        "XFree86 License 1.1",
		URL:         "http://www.xfree86.org/current/LICENSE4.html",
		OSIApproved: false,
	},
	"XSkat": {
		Name:        "XSkat License",
		URL:         "https://fedoraproject.org/wiki/Licensing/XSkat_License",
		OSIApproved: false,
	},
	"Xerox": {
		Name:        "Xerox License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Xerox",
		OSIApproved: false,
	},
	"Xnet": {
		Name:        "X.Net License",
		URL:         "https://opensource.org/licenses/Xnet",
		OSIApproved: true,
	},
	"YPL-1.0": {
		Name:        "Yahoo! Public License v1.0",
		URL:         "http://www.zimbra.com/license/yahoo_public_license_1.0.html",
		OSIApproved: false,
	},
	"YPL-1.1": {
		Name:        "Yahoo! Public License v1.1",
		URL:         "http://www.zimbra.com/license/yahoo_public_license_1.1.html",
		OSIApproved: false,
	},
	"ZPL-1.1": {
		Name:        "Zope Public License 1.1",
		URL:         "http://old.zope.org/Resources/License/ZPL-1.1",
		OSIApproved: false,
	},
	"ZPL-2.0": {
		Name:        "Zope Public License 2.0",
		URL:         "http://old.zope.org/Resources/License/ZPL-2.0",
		OSIApproved: true,
	},
	"ZPL-2.1": {
		Name:        "Zope Public License 2.1",
		URL:         "http://old.zope.org/Resources/ZPL/",
		OSIApproved: false,
	},
	"Zed": {
		Name:        "Zed License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Zed",
		OSIApproved: false,
	},
	"Zend-2.0": {
		Name:        "Zend License v2.0",
		URL:         "https://web.archive.org/web/20130517195954/http://www.zend.com/license/2_00.txt",
		OSIApproved: false,
	},
	"Zimbra-1.3": {
		Name:        "Zimbra Public License v1.3",
		URL:         "http://web.archive.org/web/20100302225219/http://www.zimbra.com/license/zimbra-public-license-1-3.html",
		OSIApproved: false,
	},
	"Zimbra-1.4": {
		Name:        "Zimbra Public License v1.4",
		URL:         "http://www.zimbra.com/legal/zimbra-public-license-1-4",
		OSIApproved: false,
	},
	"Zlib": {
		Name:        "zlib License",
		URL:         "http://www.zlib.net/zlib_license.html",
		OSIApproved: true,
	},
	"blessing": {
		Name:        "SQLite Blessing",
		URL:         "https://www.sqlite.org/src/artifact/e33a4df7e32d742a?ln=4-9",
		OSIApproved: false,
	},
	"bzip2-1.0.5": {
		Name:        "bzip2 and libbzip2 License v1.0.5",
		URL:         "http://bzip.org/1.0.5/bzip2-manual-1.0.5.html",
		OSIApproved: false,
	},
	"bzip2-1.0.6": {
		Name:        "bzip2 and libbzip2 License v1.0.6",
		URL:         "https://github.com/asimonov-im/bzip2/blob/master/LICENSE",
		OSIApproved: false,
	},
	"copyleft-next-0.3.0": {
		Name:        "copyleft-next 0.3.0",
		URL:         "https://github.com/copyleft-next/copyleft-next/blob/master/Releases/copyleft-next-0.3.0",
		OSIApproved: false,
	},
	"copyleft-next-0.3.1": {
		Name:        "copyleft-next 0.3.1",
		URL:         "https://github.com/copyleft-next/copyleft-next/blob/master/Releases/copyleft-next-0.3.1",
		OSIApproved: false,
	},
	"curl": {
		Name:        "curl License",
		URL:         "https://github.com/bagder/curl/blob/master/COPYING",
		OSIApproved: false,
	},
	"diffmark": {
		Name:        "diffmark license",
		URL:         "https://fedoraproject.org/wiki/Licensing/diffmark",
		OSIApproved: false,
	},
	"dvipdfm": {
		Name:        "dvipdfm License",
		URL:         "https://fedoraproject.org/wiki/Licensing/dvipdfm",
		OSIApproved: false,
	},
	"eCos-2.0": {
		Name:        "eCos license version 2.0",
		URL:         "https://www.gnu.org/licenses/ecos-license.html",
		OSIApproved: false,
	},
	"eGenix": {
		Name:        "eGenix.com Public License 1.1.0",
		URL:         "http://www.egenix.com/products/eGenix.com-Public-License-1.1.0.pdf",
		OSIApproved: false,
	},
	"etalab-2.0": {
		Name:        "Etalab Open License 2.0",
		URL:         "https://github.com/DISIC/politique-de-contribution-open-source/blob/master/LICENSE.pdf",
		OSIApproved: false,
	},
	"gSOAP-1.3b": {
		Name:        "gSOAP Public License v1.3b",
		URL:         "http://www.cs.fsu.edu/~engelen/license.html",
		OSIApproved: false,
	},
	"gnuplot": {
		Name:        "gnuplot License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Gnuplot",
		OSIApproved: false,
	},
	"iMatix": {
		Name:        "iMatix Standard Function Library Agreement",
		URL:         "http://legacy.imatix.com/html/sfl/sfl4.htm#license",
		OSIApproved: false,
	},
	"libpng-2.0": {
		Name:        "PNG Reference Library version 2",
		URL:         "http://www.libpng.org/pub/png/src/libpng-LICENSE.txt",
		OSIApproved: false,
	},
	"libselinux-1.0": {
		Name:        "libselinux public domain notice",
		URL:         "https://github.com/SELinuxProject/selinux/blob/master/libselinux/LICENSE",
		OSIApproved: false,
	},
	"libtiff": {
		Name:        "libtiff License",
		URL:         "https://fedoraproject.org/wiki/Licensing/libtiff",
		OSIApproved: false,
	},
	"mpich2": {
		Name:        "mpich2 License",
		URL:         "https://fedoraproject.org/wiki/Licensing/MIT",
		OSIApproved: false,
	},
	"psfrag": {
		Name:        "psfrag License",
		URL:         "https://fedoraproject.org/wiki/Licensing/psfrag",
		OSIApproved: false,
	},
	"psutils": {
		Name:        "psutils License",
		URL:         "https://fedoraproject.org/wiki/Licensing/psutils",
		OSIApproved: false,
	},
	"wxWindows": {
		Name:        "wxWindows Library License",
		URL:         "https://opensource.org/licenses/WXwindows",
		OSIApproved: false,
	},
	"xinetd": {
		Name:        "xinetd License",
		URL:         "https://fedoraproject.org/wiki/Licensing/Xinetd_License",
		OSIApproved: false,
	},
	"xpp": {
		Name:        "XPP License",
		URL:         "https://fedoraproject.org/wiki/Licensing/xpp",
		OSIApproved: false,
	},
	"zlib-acknowledgement": {
		Name:        "zlib/libpng License with Acknowledgement",
		URL:         "https://fedoraproject.org/wiki/Licensing/ZlibWithAcknowledgement",
		OSIApproved: false,
	},
}
