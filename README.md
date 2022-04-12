[![CircleCI](https://circleci.com/gh/giantswarm/aws-load-balancer-controller-app/tree/main.svg?style=svg)](https://circleci.com/gh/giantswarm/aws-load-balancer-controller-app/tree/main)

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

Download the recommended IAM policy for the AWS Load Balancer Controller
    ```
    curl -o iam-policy.json https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/main/docs/install/iam_policy.json
    ```

For a thorough explanation on how to create the IAM policy and role please refer to the [upstream charts README.md.](https://github.com/giantswarm/aws-load-balancer-controller-app/blob/main/helm/aws-load-balancer-controller/README.md)

## Installing

There are 3 ways to install this app onto a workload cluster.

1. [Using our web interface](https://docs.giantswarm.io/ui-api/web/app-platform/#installing-an-app)
2. [Using our API](https://docs.giantswarm.io/api/#operation/createClusterAppV5)
3. Directly creating the [App custom resource](https://docs.giantswarm.io/ui-api/management-api/crd/apps.application.giantswarm.io/) on the management cluster.

## Configuring
Additionally to the IAM role, the region (e.g. eu-west-1) and the VPC ID are required.

By default, a PodDisruptionBudget is configured so the admission webhook does not become unreachable, possibly blocking scheduling other pods or cluster maintenances.

### values.yaml
**This is an example of a values file you could upload using our web interface.**
```
# RBAC
serviceAccount:
    create: true

# Deployment
podAnnotations:
    iam.amazonaws.com/role: AWSLoadBalancerControllerIAMRole # Will be picked up by KIAM to associate the pod with the given role
vpcId: vpc-0c7dc1da1ca5b1819
region: eu-west-1

# PodDisruptionBudget
podDisruptionBudget:
  minAvailable: 1
  # maxUnavailable: 1
```

See our [full reference page on how to configure applications](https://docs.giantswarm.io/app-platform/app-configuration/) for more details.

## Credit

* https://github.com/giantswarm/aws-load-balancer-controller-app/tree/main/helm/aws-load-balancer-controller
