# Infrastructure

Everything on AWS is defined as code in `template.yaml` (AWS SAM). Nothing is clicked together by hand except the one-time account setup below.

## One-time account setup (Sprint 0)

1. Create the AWS account, turn on MFA for the root user, then stop using root.
2. Create an IAM Identity Center user for yourself (admin) and log in with `aws configure sso`.
3. **AWS Budgets:** create a monthly cost budget with email alerts at $10 and $25.
4. Pick the region (spike SPIKE-01) and put it in `samconfig.toml`.
5. Create a GitHub OIDC role so GitHub Actions can deploy without stored keys (Sprint 2).

## Deploy

```bash
cd infra
sam build
sam deploy --config-env dev \
  --parameter-overrides "Stage=dev LambdaAdapterLayerArn=<arn-from-adapter-readme>"
```

Prod is the same with `--config-env prod`, only from the `main` branch.

## Stacks

| Stack | Branch | Who uses it |
|---|---|---|
| `cakeshop-dev` | `develop` | You, testing |
| `cakeshop-prod` | `main` | The client's customers |
