# Changelog

Format: [Keep a Changelog](https://keepachangelog.com). Versions follow semver.

## [Unreleased]

### Changed
- Localised for a Bangalore client: prices in INR (paise), Diwali occasion, Indian payment gateways, English-first assistant
- Region Mumbai `ap-south-1` in a dedicated `cakeshop-dev` AWS account (ADR 0006); SAM profile, region and `project=cakeshop` tags in `samconfig.toml`

### Added
- Architecture blueprint v1, data model, API contract, ADRs 0001–0005
- Go API skeleton: health, cake catalogue, server-side quotes, tests
- SAM template (DynamoDB, S3, Cognito, HTTP API, Lambda), CI workflow
- Solo Scrumban process, sprint plan, backlog import script
