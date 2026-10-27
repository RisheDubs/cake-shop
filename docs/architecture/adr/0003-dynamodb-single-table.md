# 0003. DynamoDB single-table design

- Status: accepted
- Date: 2026-10-27

## Context
Need a database with no server, near-zero cost at low volume, and backups.

## Decision
One DynamoDB table, on-demand billing, PK/SK + one GSI, point-in-time recovery on. See ../data-model.md.

## Consequences
- ~$1–3/month, no maintenance.
- Access patterns must be designed up front; ad-hoc reporting is harder. If the owner later wants rich analytics, export to S3 + Athena.
- Relational alternative (Aurora Serverless / RDS) rejected on cost for this scale.
