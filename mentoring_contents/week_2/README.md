# Week 2

## Basic Features of a Load Balancer 

### NAT 

Load Balancer uses NAT(Network Address Translation) technology to modify the source or destination IP addresses of packets as they pass through the Load Balancer. This allows the Load Balancer to efficiently distribute traffic to backend servers while hiding the internal network structure from external users.

NAT at the Layer 4 of OSI model.
- NAT operates at Layer 4 of the OSI model, which means it is closely tied to port numbers at the transport layer, as well as IP addresses at Layer 3.  

> The type of NAT is determined by how the **Source IP** is modified during the process. 

#### DNAT (Destination Network Address Translation)

DNAT modifies the destination IP address of incoming packets as they pass through a load balancer or router.

- Typically used in load balancer to forward client requests to different backend servers based on the destination IP.
- Example: A public IP is mapped to multiple private backend servers, and the load balancer translates the destination IP to route traffic accordingly.

#### One-ARM Load Balancer 

In a one-arm configuration, both incoming and outgoing traffic flows through the load balancer using a single network interface. The load balancer acts like a proxy.

- Often used when the load balancer and the servers are on the same subnet, and the goal is to avoid routing complexities.
- **Advantage**: Simplifies network setup, as all traffic flows through one interface.
- **Disadvantage**: Can become a bottleneck if traffic increases significantly.

#### FullNAT 

FullNAT translates both source and destination IP addresses as well as port numbers of the packets.

- Provides complete flexibility in controlling both ends of the connection, often used when both the source and destination addresses need to be hidden or modified for routing.
- **Advantage**: More control over the network flow, suitable for complex multi-network setups.

#### DSR (Direct Server Return)

**DSR is similar to DNAT in that both are used in load balancing, but the way they handle traffic is different.**

In DSR, incoming traffic is load balanced by the load balancer, but the response traffic from the backend servers bypasses the load balancer and goes directly back to the client. 

- This reduces the load on the load balancer, allowing it to handle a higher number of requests.
- **Advantage**: Highly efficient for high-throughput applications since the load balancer only manages incoming traffic.
- **Disadvantage**: Requires special network setup as the backend servers need to send traffic directly back to the client without passing through the load balancer again.

### Load Balancing Algorithms

#### Round-Robin 

Requests are distributed sequentially across all available servers in a cyclic order.

- Simple and effective when the servers have similar capabilities and there is a uniform traffic load.

#### Weighted Round-Robin

Similar to Round-Robin, but each server is assigned a weight based on its capacity. Servers with higher weights receive more requests.

- Useful when some servers have more resources (CPU, memory, etc.) than others, allowing for a more balanced distributed of traffic based on capacity.

#### Flow Hash 

Requests are distributed based on a hash of the 5-tuple (Source IP, Destination IP, Source Port, Destination Port, Protocol). This ensures that a given flow consistently goes to the same server. 

- Ensures session persistence, which is important for scenarios like web sessions or transactions where it's crucial to maintain the same connection between client and server. 

#### Least-Connection 

Traffic is directed to the server with the fewest active connections at the time of the request. 

- Ideal for scenarios where request loads vary significantly, ensuring that no single server gets overwhelmed by more connections than it can handle.

