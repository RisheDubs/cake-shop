# Infrastructure

Everything on AWS is defined as code in `template.yaml` (AWS SAM). Nothing is clicked together by hand except the one-time account setup below.

**Region:** Asia Pacific (Mumbai) `ap-south-1`, closest to the client in Bangalore. See [ADR 0006](../docs/architecture/adr/0006-region-and-account.md).

## Accounts

| Account | What it is | Region |
|---|---|---|
| Main (management) account | Your existing account. Owns the AWS Organization, Identity Center and the bill. Your other projects stay here. | any |
| `cakeshop-dev` | Member account only for this project. All dev resources. | `ap-south-1` |
| `cakeshop-prod` | Client-owned account, created or handed over in Sprint 12. | `ap-south-1` |

## One-time setup (done in Sprint 0, Day 1)

1. AWS Organizations enabled in the main account; `cakeshop-dev` created as a member account.
2. `cakeshop-dev` root user: password set, MFA on, IAM access to billing activated.
3. IAM Identity Center: your user has `AdministratorAccess` on `cakeshop-dev`.
4. Budget `cakeshop-monthly` ($25, alerts at 40/80/100% actual and 100% forecast) created **inside** `cakeshop-dev`.
5. CLI profile `cakeshop-dev` via `aws configure sso`, default region `ap-south-1`.
6. Sprint 2: GitHub OIDC role so GitHub Actions can deploy without stored keys.

Never create IAM users or access keys in `cakeshop-dev`.

## Every session

```bash
export AWS_PROFILE=cakeshop-dev
aws sso login
aws sts get-caller-identity     # Account must be the cakeshop-dev ID
```

## Deploy

```bash
cd infra
sam build
sam deploy --config-env dev \
  --parameter-overrides "Stage=dev LambdaAdapterLayerArn=<arn-from-adapter-readme>"
```

`samconfig.toml` already sets the profile, region and the `project=cakeshop` tags.

## Region notes

- Everything runs in `ap-south-1` **except** the SSL certificate for a CloudFront custom domain, which AWS requires in `us-east-1`. So never block other regions with an organization policy.
- Claude on Bedrock from Mumbai: Haiku 4.5 can run inside India; Haiku 5.5 only through the global cross-region profile. Decided in SPIKE-02.

## Stacks

| Stack | Branch | Account | Who uses it |
|---|---|---|---|
| `cakeshop-dev` | `develop` | `cakeshop-dev` | You, testing |
| `cakeshop-prod` | `main` | client's account | The client's customers |
