/* SPDX-License-Identifier: (GPL-2.0-only OR BSD-2-Clause) */
/* Copyright Authors of Cilium */

#ifndef _VRRP_H
#define _VRRP_H

#include <linux/types.h>

#define VRRP_VERSION_2		2		/* VRRP version 2 -- rfc2338.5.3.1 */
#define VRRP_VERSION_3		3		/* VRRP version 3 -- rfc5798.5.2.1 */
#define VRRP_PKT_ADVERT		1		/* packet type -- rfc2338.5.3.2 */

struct vrrphdr {
  __u8		vers_type;
  __u8		vrid;
  __u8		priority;
  __u8      naddr;
  union {
	struct {
		__u8	auth_type;
		__u8	adver_int;
	} v2;
	struct {
		__u16	adver_int;
	} v3;
  };
  __u16 checksum;
};

#endif /* _VRRP_H */
