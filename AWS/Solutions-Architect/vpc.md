# Amazon Virtual Private Cloud (VPC)

A virtual private cloud is an on-demand configurable pool of shared resources allocated within a pulic cloud environment, providing a certain level of isolation between the different organizations using the resources. The isolation between one VPC user and all other users of the same cloud is achieved normally through allocation of a private IP subnet and a virtual communication construct per user. In a VPC, the previouly described mechanism, providing isolation within the cloud, is accompained with a VPN function that secures, by means of authentication and encryption, the remote access of the organization to its VPC resources.

A virtual private cloud (VPC) is a private cloud hosted within a public cloud.

A virtual private cloud (VPC) is a secure, isolated private cloud hosted within a public cloud. VPC customers can run code, store data, host websites, and do anything else they could do in an ordinary private cloud, but the private cloud is hosted remotely by a public cloud provider. VPCs combine the scalability and convenience of public cloud computing with the data isolation of private cloud computing. A private cloud is single-tenant. A private cloud is a cloud service that is exclusively offered to one organization. A virtual private cloud (VPC) is a private cloud within a public cloud; no one else shares the VPC with the VPC customer.

Amazon Virtual Private Cloud (VPC) is a service that lets you launch AWS resources in a logically isolated virtual network that you define. You have complete control over your virtual networking environment, including selection of your own IP address range, creation of subnets, and configuration of route tables and networking gateways. You can use both IPv4 and IPv6 for most resources in your VPC, helping to ensure secure and easy access to resources and applications.

## Features

Flow Logs: You can monitor your VPC flow logs delivered to Amazon Simple Storage Service (Amazon S3) or Amazon CloudWatch to gain operational visibility into your network dependecies and traffic patterns, detect anomalies and prevent data leakage, and troubleshoot network connectivity and configuration issues. The enriched metadata in flow logs helps you learn more about who initiated your TCP connections and the packet-level source and destination for traffic flowing through intermidiate layers (such as NAT gateway). You can also archive your flow logs to assist in meeting certain compliance requirements. TLDR: Logs into S3 of CloudWatch all network logs.

IP Address Manager (IPAM): IPAM makes it easieir for you to plan, track and monitor IP addresses for your AWS workloads. IPAM automates IP addresses assigments to your Amazon VPC, removing the need to use homegrown or spreadsheet-based planning applications. It also enhances your network observability by showing IP usage across multiple accounts and VPCs in a unified operational view.

IP Addressing: IP addressing enable resources in your VPC to communicate with each other and with resources over the internet. Amazon VPC supports both IPv4 and IPv6 addressing protocols. In a VPC, you can create IPv4-only, dual stack, and IPv6-only subnets and launch Amazon EC2 instances in these subnets. Amazon gives you multiple options to assign public IP addresses to your instances. You can use the Amazon provided IPv4 addresses, Elastic IPv4 addresses, or an IP address from the Amazon provided CIDRs. Apart from this, you have the option to bring your own IPv4 or IPv6 addresses within the Amazon VPC that can be assigned to these instances.

Ingress Routing: With this feature, you can route all incoming and outgoing traffic flowing to/from an internet gateway or virtual private gateway to a specific Amazon EC2 instance's elastic network interface. Configure your virtual private cloud to send all traffic to a gateway or an Amazon EC2 instance before it reaches your business workloads.

Network Access Analyzer: Network Access Analyzer helps you verify that your network on AWS conforms to your network security and compliance requirements. You can use Network Access Analyzer to understand network access to your resources, helping you identify improvements to your cloud security posture and easily demonstrate compliance.

Network Access Control List: A network access control list (network ACL) is an optional layer of security fro your VPC that acts as a firewall for controlling traffic in and out of one or more subnets. You might set up network ACLs with rules similar to those of your security groups.

Reachability Analyzer: This static configuration analysis tool enables you to analyze and debug network reachability between two resources in your VPC. After you specify the source and destination resources, Reachability Analyzer produces hop-by-hop details of the virtual path between them when they are reachable, and identifies the blocking component when they are unreachable.

Security Groups: Create security groups to act as firewall for associated Amazon EC2 instances, controlling inbound and outbound traffic at the instance level. When you launch an instance, you can associate it with one or more security groups. If you don't specify a group, the instance is automatically associated with the VPC's default group. Each instance in your VPC can belong to a different set of groups.

Traffic Mirroring: This feature allows you to copy network from an elastic network interface of Amazon EC2 instances and send it to out-of-band security and monitoring appliances for deep packet inspection. You can detect network and security anomalies, gain operational insights, implement compliance and security controls, and troubleshoot issues. Traffic Mirroring gives you direct access to the network packets flowing through your VPC.

## Classless Inter-Domain Routing (CIDR)

Classless Inter-Domain Routing (CIDR) is a method for allocating IP addresses. They are used in Security Groups rules and AWS networking in general.

CIDRs help to define an IP address range. It consists of two components:
- Based IP
- Subnet Mask

The Base IP represents an IP contained in the range (XX.XX.XX.XX). Example: 10.0.0.0, 192.168.0.0.

The Subnet Mask defines how many bits can change in the IP. Example: /0, /24, /32. Subnet Masks can be represented in two forms: 
- /8 == 255.0.0.0
- /16 == 255.255.0.0
- /24 == 255.255.255.0
- /32 == 255.255.255.255

The Subnet Mask basically allows part of the underlying IP to get additional next values from the base IP.
- 192.168.0.0/32 => allows for 1 IP (2^0) - 192.168.0.0
- 192.168.0.0/31 => allows for 2 IPs (2^1) - 192.168.0.0 -> 192.168.0.1
- 192.168.0.0/30 => allows for 4 IPs (2^2) - 192.168.0.0 -> 192.168.0.3
- 192.168.0.0/29 => allows for 8 IPs (2^3) - 192.168.0.0 -> 192.168.0.7
- 192.168.0.0/28 => allows for 16 IPs (2^4) - 192.168.0.0 -> 192.168.0.15
- 192.168.0.0/27 => allows for 32 IPs (2^5) - 192.168.0.0 -> 192.168.0.31
- 192.168.0.0/26 => allows for 64 IPs (2^6) - 192.168.0.0 -> 192.168.0.63
- 192.168.0.0/25 => allows for 128 IPs (2^7) - 192.168.0.0 -> 192.168.0.127
- 192.168.0.0/24 => allows for 255 IPs (2^8) - 192.168.0.0 -> 192.168.0.255

An IPv4 i composed of four octets.
- /32 - no octect can change
- /24 - last octet can change
- /16 - last 2 octets can change
- /8 - last 3 octets can change
- /0 - all octets can change

### Public vs. Private IP (IPv4)

The Internet Assigned Numbers Authority (IANA) established certain blocks of IPv4 addresses for the use of private (LAN) and public (Internet) addresses.

Private IP can only allow certain values:
- 10.0.0.0 - 10.255.255.255 (10.0.0.0/8)
- 172.16.0.0 - 172.31.255.255 (172.16.0.0/12) <- AWS default VPC is in this range
- 192.168.0.0 - 192.168.255.255 (192.168.0.0/16) <- e.g., home networks

All the rest of the IP addresses on the Internet are Public

## Default VPC

- All new AWS accounts have a default VPC
- New EC2 instances are launched into the default VPC if no subnet is specified
- Default VPC has Internet connectivity and all EC2 instances inside it have public IPv4 addresses
- We also get a public and a private IPv4 DNS names

## VPC Overview

- VPC - Virtual Private Cloud
- You can have multiple VPCs in an AWS region (max 5 per region - soft limit)
- Max CIDR per VPC is 5, for each CIDR:
  - Min size is /28 (16 IP addresses)
  - Max size is /16 (65536 IP addresses)
- Because VPC is private, only the Private IPv4 ranges are allowed:
  - 10.0.0.0 - 10.255.255.255 (10.0.0.0/8)
  - 172.16.0.0 - 172.31.255.255 (172.16.0.0/12)
  - 192.168.0.0 - 192.168.255.255 (192.168.0.0/16)
- Your VPC CIDR should not overlap with your other networks

## Subnet Overview

- AWS reserves 5 IP addresses (first 4 & last 1) in each subnet
- These 5 IP addresses are not available for use and can't be assigned to an EC2 instance
- Example: if CIDR block 10.0.0.0/24, then reserved IP addresses are:
   - 10.0.0.0 - Network Address
   - 10.0.0.1 - reserved by AWS for the VPC router
   - 10.0.0.2 - reserved by AWS for mapping to Amazon-provided DNS
   - 10.0.0.3 - reserved by AWS for future use
   - 10.0.0.255 - Network Broadcast Address. AWS does not support broadcast in a VPC, therefore the address is reserved