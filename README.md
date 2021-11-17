[![CircleCI](https://circleci.com/gh/giantswarm/{APP-NAME}-app.svg?style=shield)](https://circleci.com/gh/giantswarm/{APP-NAME}-app)

# AWS Load Balancer Controller chart

AWS Load Balancer controller Helm chart for Giant Swarm clusters

## Introduction
AWS Load Balancer controller manages the following AWS resources
- Application Load Balancers to satisfy Kubernetes ingress objects
- Network Load Balancers to satisfy Kubernetes service objects of type LoadBalancer with appropriate annotations

## Security updates
**Note**: Deployed chart does not receive security updates automatically. You need to manually upgrade to a newer chart.

## Prerequisites
- kiam-app installed

The controller runs on the worker nodes, so it needs access to the AWS ALB/NLB resources via IAM permissions. The
IAM permissions can be setup through the kiam-app.

Download IAM policy for the AWS Load Balancer Controller
    ```
    curl -o iam-policy.json https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/main/docs/install/iam_policy.json
    ```

For a thorough explanation on how to create the IAM policy and role please refer to the [upstream charts README.md.](helm/aws-load-balancer-controller/README.md)

## Installing

There are 3 ways to install this app onto a workload cluster.

1. [Using our web interface](https://docs.giantswarm.io/ui-api/web/app-platform/#installing-an-app)
2. [Using our API](https://docs.giantswarm.io/api/#operation/createClusterAppV5)
3. Directly creating the [App custom resource](https://docs.giantswarm.io/ui-api/management-api/crd/apps.application.giantswarm.io/) on the management cluster.

## Configuring

### values.yaml
**This is an example of a values file you could upload using our web interface.**
```
# values.yaml

```

### Sample App CR and ConfigMap for the management cluster
If you have access to the Kubernetes API on the management cluster, you could create
the App CR and ConfigMap directly.

Here is an example that would install the app to
workload cluster `abc12`:

```
# appCR.yaml

```

```
# user-values-configmap.yaml


```

See our [full reference page on how to configure applications](https://docs.giantswarm.io/app-platform/app-configuration/) for more details.

## Compatibility

This app has been tested to work with the following workload cluster release versions:

*

## Limitations

Some apps have restrictions on how they can be deployed.
Not following these limitations will most likely result in a broken deployment.

*

## Credit

* {APP HELM REPOSITORY}
