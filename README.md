# AltLife
Second chance for impact driven projects

# Problem We Solve
AltLife provides a decentralized safety net for critical projects by:

- Protecting Innovative Code: Prevent unilateral removal of open-source projects
- Preserving Digital Sovereignty: Keep projects alive even if centralized platforms block or delete them
- Mitigating Platform Risk: Ensure continuous access to code and sites regardless of geopolitical or corporate pressures
- Enabling Free Expression: Create a resilient hosting environment for projects that might challenge status quo
- Continuous Availability: Store project artifacts permanently on walrus, making takedowns virtually impossible

### Key Use Cases:

1. Alternate media and website hosting
2. Open-source privacy tools
3. DeSci - Controversial research platforms
4. Academic and scientific repositories at risk of suppression
5. Independent journalism platforms

Inspired by real-world scenarios to ensure technological freedom and innovation cannot be arbitrarily silenced.

# Solution
Decentralized Project Preservation Workflow
Technical Architecture:

- GitHub Action Trigger: Custom workflow script initiates backup process
- Backend API:
    - Clones repository
    - Compresses files using tar.gz
- Decentralized Storage: Stores compressed repo on Walrus
- Blockchain Anchoring: 
    - Stores blob ID, object ID, and file hash on Starknet
    - Mapped against user's wallet address

Key Mechanism: Automated, immutable backup that ensures project code, media and website survival across decentralized infrastructure, making censorship or unilateral takedown practically impossible.

# How to Use

- Install Rust, Sui, Walrus and other dependencies on your system
- Clone/Fork this repo
- Start the go server
- Setup a SSH reverse proxy to open the port over a https URL
- In another repo, create a workflow as shown in the `test.yml` file:
    - If the repo is generic codebase, pass type as `generic`
    - If the repo is of a static website with all the html, js and css files, pass `type` as `static`
    - If the repo is of a react project, pass the `type` as `react`