Addon means an extension of ks-devops. For intance: [Jenkins](http://jenkins.io/), [Argo CD](https://github.com/argoproj/argo-cd/),
[SonarQube](https://www.sonarqube.org/), [ks-releaser](https://github.com/kubesphere-sigs/ks-releaser/), etc.

There are two CRDs:

* `Addon`, it represents a component
* `AddonStrategy`, it describes how to install an addon.

## Supported addons

| Name                                                           | Description                                                               |
|----------------------------------------------------------------|---------------------------------------------------------------------------|
| [ks-releaser](https://github.com/kubesphere-sigs/ks-releaser/) | Help to release a project which especially has multiple git repositories. |
| [Argo CD](https://github.com/argoproj/argo-cd/)                | Declarative continuous deployment for Kubernetes.                         |

## How to use?

The controller will create the corresponding CRD (.e.g `ArgoCD`) once you created the `Addon` CRD.

Find more sample files from [here](../config/samples/addon).

## Support more?

Want to support more addons? It would be easy if you can find it from the [operator hub](https://operatorhub.io/).

> Restriction:
> * Require install desired operator manually.
> * [Hard code](../controllers/addon/operator_controller.go) about the supported addons
