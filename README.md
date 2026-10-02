# go-aws-transaction-logger

A simple go backend that logs a transaction for learning AWS.

The planned flow:

1. Validate a transaction (account reference and amount in pence).
2. Store it in **RDS Postgres** using `pgx`.
3. Generate a plain-text receipt.
4. Upload the receipt to **S3** using `aws-sdk-go-v2` and return a presigned download URL.
5. Log each step (CloudWatch).

Money is stored as integer pence (`int64`) to avoid rounding errors.

## Run

```
go test ./...
go run .
```