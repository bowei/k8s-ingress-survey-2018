# Kubernetes Ingress Survey

This survey is to provide Kubernetes SIG-NETWORK with data about how the wider community is using the Ingress API.

## Users

### I currently use Ingress (check all that apply)...
* in production workloads.
* for development/prototyping.
* for testing the feature. (Still evaluating the Ingress API).

### Who authors and manages Ingress definitions in your organization? (Check all that apply.)
* Developers (who implement the service/application)
* Devops (who deploy the application)
* Network ops (who manage cluster networking and security policies)
* Other…

### What Ingress providers do you use? (Check all that apply.)
* Ingress provided by my cloud environment
* Off-the-shelf proxy-based provider (e.g. Nginx, HAProxy, Envoy, Apache TrafficServer, ...)
* Other…

## Configuration
For the following questions, describe your typical Ingress deployment. If your deployments are extremely varied, describe your more important use case.

### How many Ingress objects per cluster?
* 1 - 10
* 11 - 50
* 51 or more

### How many paths in the Ingress definition? (Across all `hosts` sections.)
* 1 - 5
* 6 - 20
* 21 or more

### How many different Services referenced by a single Ingress?
* 1 - 3
* 4 - 10
* 11 or more

### Do you have different Ingresses pointing to the same service?
* Yes
* No

### If yes, what is your use case?

### Do you have different paths within an Ingress pointing to the same service?
* Yes
* No

### If yes, what is your use case?

### How many annotations in your Ingress object? (Not including status annotations populated by the controller).
* No annotations.
* 1 - 3
* 3 - 10
* 11 or more

### What are your annotations used for? (Check all that apply.)
* To configure timeouts/healthcheck parameters.
* To configure client termination settings (e.g. TLS, extra protocols, keepalive)
* To enhance path routing (e.g. match requests on headers, cookies)
* To use features specific to the Ingress provider.
* Other…

## Portability
### Do you expect the Ingress spec (without annotations) to be transparently portable across environments?
* Yes
* No

### Would you prefer a less portable spec that is more expressive or a less expressive but more portable spec?
* More expressive, less portable.
* More portable, less expressive.

### Have you migrated an Ingress configuration between providers?
* Yes
* No

## Features

### Is the vanilla Ingress (no annotations) expressive enough for your application?
* Yes
* No

### If you could add one feature/change one thing about the current Ingress specification, what would that be?
