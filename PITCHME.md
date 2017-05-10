---

### Golang Microservice Configuration

Ross Rothenstine   
Sr. Software Engineer    
Digital Engine - GE Current

---

GE Current + Go

 - WebSocket Push Server for Events in real time.
 - HTTP Load Testing

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
 - Don't overconfigure.
 - Prefer to use structs to hold configurations rather than variables for each property.
   
---

### Configuration Systems

 - Environment Variables
 - File-based
 - Custom User Provided Service (CUPS)
 - Centralized Configuration Server
 
---

### Environment Variables

 - Pros
   - Simplest from an implementation point-of-view.
   - No library or file parsing overhead.
   - Easy to change single values.
   - Integrates well with CloudFoundry.\*
 - Cons
   - Changing multiple values is difficult.
   - Initialization without a library can be tricky. 
   
<sub>\* - YMMV</sub>

---

### Environment Variables Demo

---

 - Raw environment configuration is boring and tedious.
 - Packages like [envconfig](https://github.com/kelseyhightower/envconfig) ease the process tremendously.

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
 
--- 

### File-based

 - Pros
   - Very readable.
   - Easy to change multiple values, point to different files.
 - Cons
   - Doesn't play well with CloudFoundry.

--- 

### File-based Demo

---

### CUPS

 - Pros
   - Integrates well with CloudFoundry.
   - Define in one space, reuse across multiple applications.
   - No run script necessary.
 - Cons
   - Tight integration with CloudFoundry
   - Probably will have headaches doing local development.
   - Updating CUPS isn't necessarily easy.
   
---

```json
"user-provided": [
   {
    "credentials": {
     "clientId": "my_client",
     "clientSecret": "super_secret",
     "issuerUri": "http://foobar.com/token",
     "uri": "http://foobar.com"
    },
    "label": "user-provided",
    "name": "my-microservice",
    "syslog_drain_url": "",
    "tags": [],
    "volume_mounts": []
   },
   {
    "credentials": {
     "foobar": {
      "password": "my_password",
      "username": "my_user"
     }
    },
    "label": "user-provided",
    "name": "another-service",
    "syslog_drain_url": "",
    "tags": [],
    "volume_mounts": []
   }
  ]

```

---

### CUPS demo

---

### Centralized Configuration Service

A service that is usually backed by some storage like Redis or Postgres that will be the central authority on configuration values.

Configurations then become keys e.g., `dev:redis` or `dev:microserviceA`, which are then queried for.

[Consul](https://www.consul.io)

---

  - Pros
    - Centralized.
    - Ability for zero-down time when a configuration value is changed.
  - Cons
    - Centralized. If configuration server is down, then apps cannot start.
    - Maximum overhead.
    - Not integrated well with CloudFoundry.
    
---

Thanks!

Q/A
