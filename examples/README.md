# ChainUp Custody Go SDK - Examples

This directory contains example code demonstrating how to use the ChainUp Custody Go SDK.

## Examples

### WaaS API Examples

See [waas_example.go](waas_example.go) for examples of:

- User registration (mobile/email)
- Account balance queries
- Deposit address management
- Withdrawals
- Transfers
- Coin list retrieval
- Async notification handling

### MPC API Examples

See [mpc_example.go](mpc_example.go) for examples of:

- Wallet creation and management
- Address creation
- Asset queries
- Withdrawals
- Deposit record retrieval
- Web3 transactions
- Auto-sweep operations
- Workspace operations
- TRON resource delegation
- Async notification handling

## Running the Examples

1. Update the credentials in the example files:

   - Replace `your-app-id` with your actual app ID
   - Replace the private and public key placeholders with your actual keys
   - Replace `your-api-key` with your actual API key (for MPC)

2. Run the WaaS example:

```bash
go run examples/waas_example.go
```

3. Run the MPC example:

```bash
go run examples/mpc_example.go
```

## Important Notes

- Never commit your actual credentials to version control
- Always use environment variables or configuration files for credentials in production
- Enable debug mode (`SetDebug(true)`) only in development environments
- Test with small amounts first before using in production

## API Documentation

For complete API documentation, see the main [README](../README.md).
