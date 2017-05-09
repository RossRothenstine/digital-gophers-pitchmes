---

### Golang Microservice Configuration

Ross Rothenstine   
Sr. Software Engineer    
Digital Engine - GE Current

---

### Configuration Overview

 - `configuration` => Any external property of a system that when modified changes behavior of the system.
 - URIs and credentials are two of the most common externalized items.
 - Many ways to inject configuration into an application. No one-size-fits-all solution.
 - Ease of local development and ease of cloud deployment are musts.
 - CloudFoundry has its own influences.
 
---

### Best Practices 

 - Configuration names should be descriptive and document themselves.
   - Be consistent with naming (URI vs URL).
   - Verbose is generally better than abbreviation e.g. `MY_MICROSERVICE_URL` vs `MMS_URL`.
   - Prefer to group via name e.g. `REDIS_HOST` and `REDIS_PORT`
 - Prefer simple values as opposed to complex structures (embedded JSON, yaml)
   - `FOO: { "bar": 42 }` vs `FOO_BAR: 42`
   - Lists are okay, just use a sane delimiter - `FOO: bar baz bang`
 - Prefer to use structs to hold configurations rather than variables for each property.
   
---

### Configuration Systems

 - Environment Variables
 - File-based
 - Centralized Configuration Server
 
---

### Environment Variables

 - Pros
   - Simplest from an implementation point-of-view.
   - No library or file parsing overhead.
   - Easy to change.
   - Integrates well with CloudFoundry.\*
 - Cons
   - Changing multiple values is difficult.
   - Initialization without a library can be tricky. 
   
<sub>\* - YMMV</sub>

---

### Environment Variables Demo

---

### CloudFoundry Notes

 - Credentials are exposed through `VCAP_SERVICES`.
 - VCAP_SERVICES is a JSON encoded environment variable.
 - While libraries exist to parse and use VCAP_SERVICES, we can do better. :)
 
---

### JQ for VCAP_SERVICES

 - `jq` => command line json queries and transforming.
 - Allows to parse VCAP_SERVICES and set environment variables before the application runs.
 - Configuration code gets to stay succinct and beautiful.
 

