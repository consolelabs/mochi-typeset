# mochi-typeset

mochi-typeset defined shared model between go repositories

Usecase

1. API response shape (pagination / overall structure)
2. kafka message for event-sourcing systems
3. constant between system (e.g ERROR_CODE, NOTIFICATION_TYPE)


<!-- consolidation-hardening: dev-docs -->
## Development & docs

This repo was reindexed in the Console Labs org-consolidation hardening pass (2026-06).

- `CLAUDE.md` , guidance for AI agents + humans (stack, conventions, commands).
- `docs/ARCHITECTURE.md` , what's here and how it fits together.
- `docs/SECURITY-AUDIT-2026-06-25.md` , secret-scan + dependency baseline.
- CI: `.github/workflows/security.yml` runs gitleaks + a dependency audit on every PR.

Build / test:

```
go build ./...
go test ./...
```

Secrets come from env / the platform, never hardcoded.
