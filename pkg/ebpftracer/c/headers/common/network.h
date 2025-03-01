#ifndef __COMMON_NETWORK_H__
#define __COMMON_NETWORK_H__

#include <vmlinux.h>
#include <vmlinux_flavors.h>

#include <bpf/bpf_endian.h>

#include <common/common.h>

// clang-format off

// TYPES

typedef union iphdrs_t {
    struct iphdr iphdr;
    struct ipv6hdr ipv6hdr;
} iphdrs;

// NOTE: proto header structs need full type in vmlinux.h (for correct skb copy)

typedef union protohdrs_t {
    struct tcphdr tcphdr;
    struct udphdr udphdr;
    struct icmphdr icmphdr;
    struct icmp6hdr icmp6hdr;
    union {
        u8 tcp_extra[40]; // data offset might set it up to 60 bytes
    };
} protohdrs;

typedef struct nethdrs_t {
    iphdrs iphdrs;
    protohdrs protohdrs;
} nethdrs;

// cgroupctxmap

typedef enum net_packet {
    CAP_NET_PACKET = 1 << 0,
    // Layer 3
    SUB_NET_PACKET_IP = 1 << 1,
    // Layer 4
    SUB_NET_PACKET_TCP = 1 << 2,
    SUB_NET_PACKET_UDP = 1<<3,
    SUB_NET_PACKET_ICMP = 1 <<4,
    SUB_NET_PACKET_ICMPV6 = 1<<5,
    // Layer 7
    SUB_NET_PACKET_DNS = 1<< 6,
    SUB_NET_PACKET_HTTP = 1<<7,
} net_packet_t;

typedef struct net_event_contextmd {
    u8 submit; // keeping this for a TODO (check should_submit_net_event)
    u32 header_size;
    u8 captured;
} __attribute__((__packed__)) net_event_contextmd_t;

typedef struct net_event_context {
    event_context_t eventctx;
    u8 argnum;
    struct { // event arguments (needs packing), use anonymous struct to ...
        u8 index0;
        u32 bytes;
        // ... (payload sent by bpf_perf_event_output)
    } __attribute__((__packed__)); // ... avoid address-of-packed-member warns
    // members bellow this point are metadata (not part of event to be sent)
    net_event_contextmd_t md;
} __attribute__((__packed__)) net_event_context_t;

// network related maps

typedef struct {
    u64 ts;
    u16 ip_csum;
    struct in6_addr src;
    struct in6_addr dst;
} indexer_t;

typedef struct {
    __uint(type, BPF_MAP_TYPE_LRU_HASH);
    __uint(max_entries, 4096); // 800 KB    // simultaneous cgroup/skb ingress/eggress progs
    __type(key, indexer_t);                 // layer 3 header fields used as indexer
    __type(value, net_event_context_t);     // event context built so cgroup/skb can use
} cgrpctxmap_t;

cgrpctxmap_t cgrpctxmap_in SEC(".maps");    // saved info SKB caller <=> SKB ingress
cgrpctxmap_t cgrpctxmap_eg SEC(".maps");    // saved info SKB caller <=> SKB egress

// inodemap

typedef struct net_task_context {
    struct task_struct *task;
    task_context_t taskctx;
    u64 matched_policies;
    int syscall;
} net_task_context_t;

struct {
    __uint(type, BPF_MAP_TYPE_LRU_HASH);
    __uint(max_entries, 65535); // 9 MB     // simultaneous sockets being traced
    __type(key, u64);                       // socket inode number ...
    __type(value, struct net_task_context); // ... linked to a task context
} inodemap SEC(".maps");                    // relate sockets and tasks

// sockmap (map two cloned "socket" representation structs ("sock"))

struct {
    __uint(type, BPF_MAP_TYPE_LRU_HASH);
    __uint(max_entries, 65535); // 9 MB     // simultaneous sockets being cloned
    __type(key, u64);                       // *(struct sock *newsock) ...
    __type(value, u64);                     // ... old sock->socket inode number
} sockmap SEC(".maps");                     // relate a cloned sock struct with

// entrymap

typedef struct entry {
    long unsigned int args[6];
} entry_t;

struct {
    __uint(type, BPF_MAP_TYPE_LRU_HASH);
    __uint(max_entries, 2048);              // simultaneous tasks being traced for entry/exit
    __type(key, u32);                       // host thread group id (tgid or tid) ...
    __type(value, struct entry);            // ... linked to entry ctx->args
} entrymap SEC(".maps");                    // can't use args_map (indexed by existing events only)

// network capture events

struct {
    __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
    __uint(max_entries, 10240);
    __type(key, u32);
    __type(value, u32);
} net_cap_events SEC(".maps");

// scratch area

struct {
    __uint(type, BPF_MAP_TYPE_PERCPU_ARRAY);
    __uint(max_entries, 1);                 // simultaneous softirqs running per CPU (?)
    __type(key, u32);                       // per cpu index ... (always zero)
    __type(value, scratch_t);               // ... linked to a scratch area
} net_heap_scratch SEC(".maps");

struct {
    __uint(type, BPF_MAP_TYPE_PERCPU_ARRAY);
    __uint(max_entries, 1);                 // simultaneous softirqs running per CPU (?)
    __type(key, u32);                       // per cpu index ... (always zero)
    __type(value, event_data_t);            // ... linked to a scratch area
} net_heap_event SEC(".maps");

// CONSTANTS

// network retval values
#define family_ipv4     (1 << 0)
#define family_ipv6     (1 << 1)
#define proto_http_req  (1 << 2)
#define proto_http_resp (1 << 3)

// payload size: full packets, only headers
#define FULL    65536       // 1 << 16
#define HEADERS 0           // no payload

// when guessing by src/dst ports, declare at network.h
#define UDP_PORT_DNS 53
#define TCP_PORT_DNS 53

// layer 7 parsing related constants
#define http_min_len 7 // longest http command is "DELETE "

typedef union  {
	u32 v4addr;
	unsigned __int128 v6addr;
}  __attribute__((packed)) addr_t;

typedef struct {
    addr_t saddr;
    addr_t daddr;
    u16 sport;
	u16 dport;
	u16 family;
} __attribute__((packed)) tuple_t;

// PROTOTYPES

statfunc u32 get_inet_rcv_saddr(struct inet_sock *);
statfunc u32 get_inet_saddr(struct inet_sock *);
statfunc u32 get_inet_daddr(struct inet_sock *);
statfunc u16 get_inet_sport(struct inet_sock *);
statfunc u16 get_inet_num(struct inet_sock *);
statfunc u16 get_inet_dport(struct inet_sock *);
statfunc struct sock *get_socket_sock(struct socket *);
statfunc u16 get_sock_family(struct sock *);
statfunc u16 get_sock_protocol(struct sock *);
statfunc u16 get_sockaddr_family(struct sockaddr *);
statfunc struct in6_addr get_sock_v6_rcv_saddr(struct sock *);
statfunc struct in6_addr get_ipv6_pinfo_saddr(struct ipv6_pinfo *);
statfunc struct in6_addr get_sock_v6_daddr(struct sock *);
statfunc volatile unsigned char get_sock_state(struct sock *);
statfunc struct ipv6_pinfo *get_inet_pinet6(struct inet_sock *);
statfunc struct sockaddr_un get_unix_sock_addr(struct unix_sock *);
statfunc int get_network_details_from_sock_v4(struct sock *, net_conn_v4_t *, int);
statfunc struct ipv6_pinfo *inet6_sk_own_impl(struct sock *, struct inet_sock *);
statfunc int get_network_details_from_sock_v6(struct sock *, net_conn_v6_t *, int);
statfunc int get_local_sockaddr_in_from_network_details(struct sockaddr_in *, net_conn_v4_t *, u16);
statfunc int get_remote_sockaddr_in_from_network_details(struct sockaddr_in *, net_conn_v4_t *, u16);
statfunc int get_local_sockaddr_in6_from_network_details(struct sockaddr_in6 *, net_conn_v6_t *, u16);
statfunc int get_remote_sockaddr_in6_from_network_details(struct sockaddr_in6 *, net_conn_v6_t *, u16);
statfunc int get_local_net_id_from_network_details_v4(struct sock *, net_id_t *, net_conn_v4_t *, u16);
statfunc int get_local_net_id_from_network_details_v6(struct sock *, net_id_t *, net_conn_v6_t *, u16);
statfunc bool fill_tuple(struct sock *, tuple_t *);

// clang-format on

// FUNCTIONS

//
// Regular events related to network
//

statfunc u32 get_inet_rcv_saddr(struct inet_sock *inet)
{
    return BPF_CORE_READ(inet, inet_rcv_saddr);
}

statfunc u32 get_inet_saddr(struct inet_sock *inet)
{
    return BPF_CORE_READ(inet, inet_saddr);
}

statfunc u32 get_inet_daddr(struct inet_sock *inet)
{
    return BPF_CORE_READ(inet, inet_daddr);
}

statfunc u16 get_inet_sport(struct inet_sock *inet)
{
    return BPF_CORE_READ(inet, inet_sport);
}

statfunc u16 get_inet_num(struct inet_sock *inet)
{
    return BPF_CORE_READ(inet, inet_num);
}

statfunc u16 get_inet_dport(struct inet_sock *inet)
{
    return BPF_CORE_READ(inet, inet_dport);
}

statfunc struct sock *get_socket_sock(struct socket *socket)
{
    return BPF_CORE_READ(socket, sk);
}

statfunc u16 get_sock_family(struct sock *sock)
{
    return BPF_CORE_READ(sock, sk_family);
}

statfunc u16 get_sock_protocol(struct sock *sock)
{
    u16 protocol = 0;

    // commit bf9765145b85 ("sock: Make sk_protocol a 16-bit value")
    struct sock___old *check = NULL;
    if (bpf_core_field_exists(check->__sk_flags_offset)) {
        check = (struct sock___old *) sock;
        bpf_core_read(&protocol, 1, (void *) (&check->sk_gso_max_segs) - 3);
    } else {
        protocol = BPF_CORE_READ(sock, sk_protocol);
    }

    return protocol;
}

statfunc u16 get_sockaddr_family(struct sockaddr *address)
{
    return BPF_CORE_READ(address, sa_family);
}

statfunc struct in6_addr get_sock_v6_rcv_saddr(struct sock *sock)
{
    return BPF_CORE_READ(sock, sk_v6_rcv_saddr);
}

statfunc struct in6_addr get_ipv6_pinfo_saddr(struct ipv6_pinfo *np)
{
    return BPF_CORE_READ(np, saddr);
}

statfunc struct in6_addr get_sock_v6_daddr(struct sock *sock)
{
    return BPF_CORE_READ(sock, sk_v6_daddr);
}

statfunc volatile unsigned char get_sock_state(struct sock *sock)
{
    volatile unsigned char sk_state_own_impl;
    bpf_probe_read(
        (void *) &sk_state_own_impl, sizeof(sk_state_own_impl), (const void *) &sock->sk_state);
    return sk_state_own_impl;
}

statfunc struct ipv6_pinfo *get_inet_pinet6(struct inet_sock *inet)
{
    struct ipv6_pinfo *pinet6_own_impl;
    bpf_probe_read(&pinet6_own_impl, sizeof(pinet6_own_impl), &inet->pinet6);
    return pinet6_own_impl;
}

statfunc struct sockaddr_un get_unix_sock_addr(struct unix_sock *sock)
{
    struct unix_address *addr = BPF_CORE_READ(sock, addr);
    int len = BPF_CORE_READ(addr, len);
    struct sockaddr_un sockaddr = {};
    if (len <= sizeof(struct sockaddr_un)) {
        bpf_probe_read(&sockaddr, len, addr->name);
    }
    return sockaddr;
}

statfunc int get_network_details_from_sock_v4(struct sock *sk, net_conn_v4_t *net_details, int peer)
{
    struct inet_sock *inet = inet_sk(sk);

    if (!peer) {
        net_details->local_address = get_inet_rcv_saddr(inet);
        net_details->local_port = bpf_ntohs(get_inet_num(inet));
        net_details->remote_address = get_inet_daddr(inet);
        net_details->remote_port = get_inet_dport(inet);
    } else {
        net_details->remote_address = get_inet_rcv_saddr(inet);
        net_details->remote_port = bpf_ntohs(get_inet_num(inet));
        net_details->local_address = get_inet_daddr(inet);
        net_details->local_port = get_inet_dport(inet);
    }

    return 0;
}

statfunc struct ipv6_pinfo *inet6_sk_own_impl(struct sock *__sk, struct inet_sock *inet)
{
    volatile unsigned char sk_state_own_impl;
    sk_state_own_impl = get_sock_state(__sk);

    struct ipv6_pinfo *pinet6_own_impl;
    pinet6_own_impl = get_inet_pinet6(inet);

    bool sk_fullsock = (1 << sk_state_own_impl) & ~(TCPF_TIME_WAIT | TCPF_NEW_SYN_RECV);
    return sk_fullsock ? pinet6_own_impl : NULL;
}

statfunc int get_network_details_from_sock_v6(struct sock *sk, net_conn_v6_t *net_details, int peer)
{
    // inspired by 'inet6_getname(struct socket *sock, struct sockaddr *uaddr, int peer)'
    // reference: https://elixir.bootlin.com/linux/latest/source/net/ipv6/af_inet6.c#L509

    struct inet_sock *inet = inet_sk(sk);
    struct ipv6_pinfo *np = inet6_sk_own_impl(sk, inet);

    struct in6_addr addr = {};
    addr = get_sock_v6_rcv_saddr(sk);
    if (ipv6_addr_any(&addr)) {
        addr = get_ipv6_pinfo_saddr(np);
    }

    // the flowinfo field can be specified by the user to indicate a network flow. how it is used by
    // the kernel, or whether it is enforced to be unique is not so obvious.  getting this value is
    // only supported by the kernel for outgoing packets using the 'struct ipv6_pinfo'.  in any
    // case, leaving it with value of 0 won't affect our representation of network flows.
    net_details->flowinfo = 0;

    // the scope_id field can be specified by the user to indicate the network interface from which
    // to send a packet. this only applies for link-local addresses, and is used only by the local
    // kernel.  getting this value is done by using the 'ipv6_iface_scope_id(const struct in6_addr
    // *addr, int iface)' function.  in any case, leaving it with value of 0 won't affect our
    // representation of network flows.
    net_details->scope_id = 0;

    if (peer) {
        net_details->local_address = get_sock_v6_daddr(sk);
        net_details->local_port = get_inet_dport(inet);
        net_details->remote_address = addr;
        net_details->remote_port = get_inet_sport(inet);
    } else {
        net_details->local_address = addr;
        net_details->local_port = get_inet_sport(inet);
        net_details->remote_address = get_sock_v6_daddr(sk);
        net_details->remote_port = get_inet_dport(inet);
    }

    return 0;
}

statfunc int get_local_sockaddr_in_from_network_details(struct sockaddr_in *addr,
                                                        net_conn_v4_t *net_details,
                                                        u16 family)
{
    addr->sin_family = family;
    addr->sin_port = net_details->local_port;
    addr->sin_addr.s_addr = net_details->local_address;

    return 0;
}

statfunc int get_remote_sockaddr_in_from_network_details(struct sockaddr_in *addr,
                                                         net_conn_v4_t *net_details,
                                                         u16 family)
{
    addr->sin_family = family;
    addr->sin_port = net_details->remote_port;
    addr->sin_addr.s_addr = net_details->remote_address;

    return 0;
}

statfunc int get_local_sockaddr_in6_from_network_details(struct sockaddr_in6 *addr,
                                                         net_conn_v6_t *net_details,
                                                         u16 family)
{
    addr->sin6_family = family;
    addr->sin6_port = net_details->local_port;
    addr->sin6_flowinfo = net_details->flowinfo;
    addr->sin6_addr = net_details->local_address;
    addr->sin6_scope_id = net_details->scope_id;

    return 0;
}

statfunc int get_remote_sockaddr_in6_from_network_details(struct sockaddr_in6 *addr,
                                                          net_conn_v6_t *net_details,
                                                          u16 family)
{
    addr->sin6_family = family;
    addr->sin6_port = net_details->remote_port;
    addr->sin6_flowinfo = net_details->flowinfo;
    addr->sin6_addr = net_details->remote_address;
    addr->sin6_scope_id = net_details->scope_id;

    return 0;
}

statfunc int get_local_net_id_from_network_details_v4(struct sock *sk,
                                                      net_id_t *connect_id,
                                                      net_conn_v4_t *net_details,
                                                      u16 family)
{
    connect_id->address.s6_addr32[3] = net_details->local_address;
    connect_id->address.s6_addr16[5] = 0xffff;
    connect_id->port = net_details->local_port;
    connect_id->protocol = get_sock_protocol(sk);

    return 0;
}

statfunc int get_local_net_id_from_network_details_v6(struct sock *sk,
                                                      net_id_t *connect_id,
                                                      net_conn_v6_t *net_details,
                                                      u16 family)
{
    connect_id->address = net_details->local_address;
    connect_id->port = net_details->local_port;
    connect_id->protocol = get_sock_protocol(sk);

    return 0;
}

statfunc bool fill_tuple(struct sock *sk, tuple_t *tuple)
{
    u16 family = BPF_CORE_READ(sk, __sk_common.skc_family);
    tuple->family = family;

	switch (family) {
	case AF_INET:
		BPF_CORE_READ_INTO(&tuple->saddr.v4addr, sk, __sk_common.skc_rcv_saddr);
		if (tuple->saddr.v4addr == 0)
			return false;

		BPF_CORE_READ_INTO(&tuple->daddr.v4addr, sk, __sk_common.skc_daddr);
		if (tuple->daddr.v4addr == 0)
			return false;

		break;
	case AF_INET6:
		BPF_CORE_READ_INTO(&tuple->saddr.v6addr, sk, __sk_common.skc_v6_rcv_saddr.in6_u.u6_addr32);
		if (tuple->saddr.v6addr == 0)
			return false;
		BPF_CORE_READ_INTO(&tuple->daddr.v6addr, sk, __sk_common.skc_v6_daddr.in6_u.u6_addr32);
		if (tuple->daddr.v6addr == 0)
			return false;

		break;

	default:
		return false;
	}

	//BPF_CORE_READ_INTO(&tuple->sport, sockp, inet_sport);
	BPF_CORE_READ_INTO(&tuple->sport, sk, __sk_common.skc_num);
	if (tuple->sport == 0)
	    return false;

    BPF_CORE_READ_INTO(&tuple->dport, sk, __sk_common.skc_dport);
    if (tuple->dport == 0)
        return false;
    tuple->dport = bpf_ntohs(tuple->dport);

	return true;
}

#endif
