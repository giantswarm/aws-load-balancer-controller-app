# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.3.4] - 2023-08-28

### Fixed

- PSP name

### Changed

- Migration from monitoring labels to ServiceMonitor.

### Added

- Added a new psp for CAPA support.

## [1.3.3] - 2023-07-13

### Added

- `Values.image.registry`

### Changed

- Template for image reference in deployment.yaml
- `application.giantswarm.io/team` annotation to Phoenix

## [1.3.2] - 2023-06-30

### Changed

- Add VPA

### Fixed

- Detect China region to fix IRSA role

## [1.3.1] - 2023-04-24

### Changed

- Metadata: use SVG icon.

## [1.3.0] - 2023-04-12

### Added

- Add IRSA annotation automatically based on Account ID and Cluster ID.

## [1.2.1] - 2023-01-31

### Changed

- Add helpers for setting values `region`, `vpcId` and IAM role annotations automatically from Giant Swarm provided default cluster values.

## [1.2.0] - 2022-04-12

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

[Unreleased]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v1.3.4...HEAD
[1.3.4]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v1.3.3...v1.3.4
[1.3.3]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v1.3.2...v1.3.3
[1.3.2]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v1.3.1...v1.3.2
[1.3.1]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v1.3.0...v1.3.1
[1.3.0]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v1.2.1...v1.3.0
[1.2.1]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v1.2.0...v1.2.1
[1.2.0]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v1.0.3...v1.1.0
[1.0.3]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v1.0.2...v1.0.3
[1.0.2]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v1.0.1...v1.0.2
[1.0.1]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/giantswarm/aws-load-balancer-controller-app/compare/v0.1.0...v1.0.0
[0.1.0]: https://github.com/giantswarm/aws-load-balancer-controller-app/releases/tag/v0.1.0
