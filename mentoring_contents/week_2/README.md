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

## Utilizing GitHub in LoxiLB

### GitHub Action 

GitHub provides a **CI/CD framework** through GitHub Actions, which allows you to automate workflows for building, testing, and deploying code. 

- **Triggering Events**: GitHub Actions can be triggered by various events such as **Commit**, **Push**, **Manual**, **Periodic**.
    - **Commit**: Automatically triggered when code is committed to the repository.
    - **Push**: Executes when code is pushed to a specific branch.
    - **Manual**: Developers can manually trigger workflows.
    - **Periodic**: Scheduled workflows that run at set intervals. (e.g., daily, weekly).

- **Platform Support**: GitHub Actions supports running workflows on different platforms, allowing for flexibility in testing environments
    - **Linux**, **Windows**, and **macOS** virtual machines are available for running jobs. This makes it possible to ensure that code works correctly across different operating systems and environments. 

#### Key Concepts

- **Workflows**: Automated processes defined in .github/workflows, which specify a series of actions to perform. These workflows define how code is built, tested, and deployed.

- **Event, On**: The activities that trigger the execution of a workflow. These can include events like **Pull Request**, **Commit**, **Issue**.
    - **Pull Request**: When a pull request is opened or updated.
    - **Commit**: When code is pushed to the repository.
    - **Issue**: When an issue is created or commented on.

- **Runs-On**: The server environment where the workflow is executed. GitHub provides several options. 
    - **Ubuntu Linux**: Default and commonly used.
    - **Microsoft Windows**: For Windows-specific builds.
    - **macOS**: For macOS-specific builds and testing.

- **Jobs**: A collection of shell scripts or tasks that run on a specific runner. Jobs are the core execution units within workflows, and each job consists of multiple steps.

- **Actions**: Reusable functions or tasks that can be executed as part of a workflow. Actions are used to simplify repetitive tasks, such as setting up environments, checking out code, or running specific commands.

> **CI/CD** refers to the combination of Continuous Integration (CI) and Continuous Delivery (CD). CI/CD automates the development, testing, and deployment processing, making it easier to detect software bugs early and enabling faster release cycles.

- **CI**: Code changes are continuously integrated into the main codebase, and tests are automatically run after each change to ensure that the new code works with the existing system. 

- **CD**: After the code is integrated and tested, it is automatically deployed to production (or another environment) at any desired time, ensuring that the latest stable version of the software is always ready for deployment.

#### GitHub Actions - Workflow Structure 

The workflow is defined in YAML format, making it easy to configure and read.

```
name: TEST

on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v4
    - name: Run a one-line script
      run: echo Hello, world!
```

- **Event Trigger - on**: The **on** field defines what events trigger the workflow.
    - Example. A **push** is made to the main branch and a **pull request** is opened or updated in the main branch.

- **Jobs**: The jobs section defines what tasks will be executed. In this case, there is a build job. 
    - **runs-on** specifies the environment or runner where the job will be executed. 
    - Here, it is set to **ubuntu-latest**.

- **Steps**: Steps define the individual tasks that the job will execute. 
    - **uses**: The first step checks out the repository using the actions/checkout@v4.
    - **run**: The second step runs a simple shell script (echo Hello, world!).

#### GitHub Action - How to Schedule Events 

In GitHub Actions, you can use the **schedule event with cron expressions** to automate workflows at specific times or intervals.

```
name: CI

on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]
  schedule:
    - cron: '0 23 * * *'  # Triggers every day at 23:00 (11:00 PM)
```

- Cron syntax is used to specify the schedule. It follows the format ```* * * * *```, where each asterisk represents a specific time unit.
    - **Minute** (0-59)
    - **Hour** (0-23)
    - **Day of the Month** (1-31)
    - **Month** (1-12 or JAN-DEC)
    - **Day of the Week** (0-6 where 0 is Sunday or SUN-SAT)
- Specify more detailed scheduling options using special symbols like ```* , - /```.
    - **\***: (Asterisk) Represents every value.
        - Example. ```* * * * *```: Every minute of every hour.
    - **,**: (Comma) Used to specify multiple values.
        - Example. ```5,10 * * * *```: Triggers at both the 5th and 10th minute of every hour.
    - **-**: (Hyphen) Defines a range of values.
        - Example. ```* 4-6 * * *```: Runs every minute during hours 4, 5, and 6 (i.e., from 04:00 to 06:59).
    - **/**: (Slash) Specifies a step value. It means "every nth" occurrence within a range.
        - Example. ```*/15 * * * *```: Runs every 15 minutes (i.e., 00, 15, 30, 45 minutes past the hour).

**Example**

```
cron: '0 12 * * MON'
```
- Every Monday at 12:00 (noon)

```
cron: '*/15 * * * *'
```
- Every 15 minutes

```
cron: '15 * * * *'
```
- At 15 minutes past every hour

#### GitHub Action - Jobs 

In GitHub Actions, jobs are define the tasks that will run in the workflow. A job consists of multiple steps, and you can specify the environment (runner) on which the job will execute.

```
jobs:
  build:
    runs-on: ubuntu-latest   # Specifies the environment as Ubuntu

    steps:
    - uses: actions/checkout@v4  # Checks out the repository
    - name: Run a one-line script # Single-line Scripts
      run: echo Hello, world!     # Runs a one-line shell command
    - name: Run a multi-line script # Multi-line Scripts
      run: |
        echo Add other actions to build,
        echo test, and deploy your project.  # Runs multiple shell commands
```

- **runs-on**: This specifies the runner where the job will run. In this example, it's set to ```ubuntu-lastest```, meaning the job will run on an Ubuntu virtual machine.

- **steps**: The steps define the sequence of tasks within the job. Each step can perform an action, such as running a shell script or using a predefined GitHub Action.
    - Example. The first step uses ```actions/checkout@v4``` to check out the repository.
    - Single-line Scripts
        - ```run: command```
    - Multi-line Scripts
        - ```run: |``` 

### Package - Container Registry

- **GitHub Docs**: GitHub provides documentation on how to work with GitHub's built-in package registry, including container registries for managing Docker Images.

- **GitHub Package Registry**: GitHub offers a container registry service that functions similarly to Docker Hub, allowing users to store, manage, and distribute container images directly from their GitHub repositories.
    - With GitHub's container registry, developers can manage Docker images alongside their code and workflows, enabling better integration with GitHub Actions for CI/CD.

### GitHub Pages 

- **Using GitHub to Deploy Documentation**: GitHub Pages allow you to host and deploy static websites, which is commonly used for publishing documentation. 
- **Markdown(MD) Format**: You can write your documentation in Markdown format(.md) to create simple and easy-to-read single-page manuals or guides.
- **Multi-page Setup with MkDocs**: MkDocs is a static site generator that converts Markdown files into well-structured websites.
    - You define the site's structure using a mkdocs.yml file, which connects different pages, allowing navigation between them.

## LoxiLB Code Analysis 

- **.github/workflows**: Contains YAML files that define workflow for GitHub Actions. These workflows automate tasks like testing, building, and deploying the LoxiLB project.

- **A collection of scripts used for Continuous Integration/Continuous Deployment**: This includes setting up, validating, and tearing down the testing environments. 

- **api**: 
    - **Swagger-based REST API server**: Provides RESTful APIs for managing LoxiLB. 
    - Integrations with modules such as Kubernetes and Prometheus to enable management and monitoring within cloud-native environments.

- **common**: Contains shared variables and functions that facilitate communication between the API and Loxinet, the core networking module of LoxiLB.

- **pkg**: Contains modules related to Linux networking, such as Loxinet (the core networking engine) and bfd (Bidirectional Forwarding Detection), providing funtionality for low-level networking tasks.

- **options**: Stores a collection of execution options needed for configuring and running LoxiLB. These options define how LoxiLB operates during startup and execution. 

- **loxilb-ebpf**: This submodule handles the Data Plane, where actual packet processing and traffic handling occur. It uses eBPF (Extended Berkeley Packet Filter) to perform high-performance, kernel-level packet processing in the LoxiLB system. 

## About eBPF

eBPF is a technology that allows programs to run in a privileged context, such as within the operating system kernel. 

- It is the successor to the BPF, with the "e" originally standing for "extended". While it was initially designed for packet filtering, eBPF has evolved and is now used in many non-networking parts of the Linux kernel as well, providing a powerful mechanism for in-kernel processing and performance monitoring.


### key eBPF Control Actions in Networking

In the Linux networking stack, eBPF operates at various stages to control and process network traffic **before it enters the network stack**. It can use actions like PASS, TX, and REDIRECT to manage how packets are handled.

- **PASS**: This action allows the packet to pass through to the next stage of the Linux network stack without modification. It means the packet is accepted and continues its normal flow through the system.

- **TX**: This action sends the packet out through the same interface it was received on or a specified interface, essentially transmitting the packet immediately.

- **REDIRECT**: The REDIRECT action forwards the packet to a different network interface, bypassing the rest of the networking stack. This is often used for load balancing or traffic routing scenarios where packets are forwarded to another destination based on custom logic.

### eBPF in LoxiLB 

- **eBPF uses C**
- **Calling C Functions in Golang**
    - ```import "C"``` is used to call C functions. This allows LoxiLB to integrate and execute C-based logic within its Go-based architecture.
- **Hooking Before Entering the Linux Network Stack**
    - C functions are utilized in LoxiLB to hook into the network stack before packets fully enter it. This enables LoxiLB to control and process network traffic at the kernel level, applying custom logic for packet filtering, redirection, or other actions.
- **Submodule Due to Licensing**
    - Due to licensing restrictions, the C components used in LoxiLB are managed as submodules. This allows LoxiLB to include and use C-based functionality while keeping it seperate from the main Go codebase to ensure compliance with open-source licenses.
