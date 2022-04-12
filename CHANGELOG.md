# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Enable PodDisruptionBudget to prevent unavailability of the admission webhook during cluster maintenances.

## [1.1.0] - 2022-03-17

### Changed

- Synchronize with upstream helm chart version 1.4.1 containing AWS Load Balancer Controller [v2.4.1](https://github.com/kubernetes-sigs/aws-load-balancer-controller/releases/tag/v2.4.1)

## [1.0.3] - 2021-12-17

### Changed

- Change README link to absolute path to avoid problems when it is shown in Happa.

## [1.0.2] - 2021-12-14

### Added

- Adds annotation `giantswarm.io/monitoring-app-label` to metrics (status) service as a persistent identifier for monitoring

## [1.0.1] - 2021-12-02

- Metrics service
- Giant Swarm monitoring labels to deployment template

## [1.0.0] - 2021-11-19

### Added

- Implement network policy to allow egress traffic

### Updated
- values.yaml with resource limits/requests
- values.yaml validation schema

## [0.1.0] - 2021-11-19

### Added
- Upstream helm chart
- Automatic clusterName
- App description in README.md

[Unreleased]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v1.1.0...HEAD
[1.1.0]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v1.0.3...v1.1.0
[1.0.3]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v1.0.2...v1.0.3
[1.0.2]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v1.0.1...v1.0.2
[1.0.1]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v0.1.0...v1.0.0
[0.1.0]: https://github.com/giantswarm/aws-load-balancer-controller-app/releases/tag/v0.1.0
