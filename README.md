[![CircleCI](https://circleci.com/gh/giantswarm/aws-load-balancer-controller-app/tree/main.svg?style=svg)](https://circleci.com/gh/giantswarm/aws-load-balancer-controller-app/tree/main)

# AWS Load Balancer Controller chart

AWS Load Balancer controller Helm chart for Giant Swarm clusters

## Introduction
AWS Load Balancer controller manages the following AWS resources
- Application Load Balancers to satisfy Kubernetes ingress objects
- Network Load Balancers to satisfy Kubernetes service objects of type LoadBalancer with appropriate annotations

## Index
- [Prerequisites](#prerequisites)
- [Security Updates](#security-updates)
- [Installing](#installing)
- [Configuring](#configuring)
  - [values.yaml](#valuesyaml)
- [Release Process](#release-process)
- [Contributing & Reporting Bugs](#contributing--reporting-bugs)
- [Credit](#credit)

## Prerequisites
- kiam-app installed

The controller runs on the worker nodes, so it needs access to the AWS ALB/NLB resources via IAM permissions. The
IAM permissions can be setup through the kiam-app.

Download the recommended IAM policy for the AWS Load Balancer Controller
```bash
curl -o iam-policy.json https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/main/docs/install/iam_policy.json
```

## Security Updates
**Note**: Deployed chart does not receive security updates automatically. You need to manually upgrade to a newer chart.

For a thorough explanation on how to create the IAM policy and role please refer to the [upstream charts README.md.](https://github.com/giantswarm/aws-load-balancer-controller-app/blob/main/helm/aws-load-balancer-controller/README.md)

## Installing

There are 3 ways to install this app onto a workload cluster.

1. [Using our web interface](https://docs.giantswarm.io/ui-api/web/app-platform/#installing-an-app)
2. [Using our API](https://docs.giantswarm.io/api/#operation/createClusterAppV5)
3. Directly creating the [App custom resource](https://docs.giantswarm.io/ui-api/management-api/crd/apps.application.giantswarm.io/) on the management cluster.

To automatically configure the correct KIAM annotation on the namespace, you can specify additional annotations directly in your App CR:

Starting with [Giant Swarm Release 18.2.0](https://docs.giantswarm.io/changes/workload-cluster-releases-aws/releases/aws-v18.2.0/), aws-load-balancer-controller can be installed without specifying any additional configuration:

```yaml
apiVersion: application.giantswarm.io/v1alpha1
kind: App
metadata:
  name: aws-load-balancer-controller
  namespace: <your-cluster-id>
spec:
  catalog: giantswarm
  kubeConfig:
    inCluster: false
  name: aws-load-balancer-controller
  namespace: aws-load-balancer-controller
  namespaceConfig:
    annotations:
      iam.amazonaws.com/permitted: .*
  version: 1.2.1
```

For all other releases, specify at least these values (Don't forget to reference your ConfigMap in the App CRs `spec.userConfig`):

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: aws-load-balancer-controller-userconfig
  namespace: <your-cluster-id>
data:
  values: |
    podAnnotations:
        # don't forget to create the role and policy before trying to use them
        iam.amazonaws.com/role: gs-<your-cluster-id>-ALBController-Role
    vpcId: vpc-0c7dc1da1ca5b1819 # the VPC Id of your cluster
    region: eu-west-1 # The AWS region your cluster is running in
```

## Configuring
Additionally to the IAM role, the region (e.g. eu-west-1) and the VPC ID are required.

By default, a [PodDisruptionBudget](https://kubernetes.io/docs/tasks/run-application/configure-pdb) is configured so the admission webhook does not become unreachable, possibly blocking scheduling other pods or cluster maintenances.

### values.yaml
**This is an example of a values file you could upload using our web interface.**
```
# Deployment
podAnnotations:
    iam.amazonaws.com/role: AWSLoadBalancerControllerIAMRole # Will be picked up by KIAM to associate the pod with the given role
vpcId: vpc-0c7dc1da1ca5b1819
region: eu-west-1
```

See our [full reference page on how to configure applications](https://docs.giantswarm.io/app-platform/app-configuration/) for more details.

## Contributing & Reporting Bugs
If you have suggestions for how `aws-load-balancer-controller` could be improved, or want to report a bug, open an issue! We'd love all and any contributions.

Check out the [Contributing Guide](https://github.com/giantswarm/aws-load-balancer-controller-app/blob/main/CONTRIBUTING.md) for details on the contribution workflow, submitting patches, and reporting bugs.


## Credit

* https://github.com/giantswarm/aws-load-balancer-controller-app/tree/main/helm/aws-load-balancer-controller
