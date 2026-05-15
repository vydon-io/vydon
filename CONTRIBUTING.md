# Contributing

Thanks for opening a PR. The bar is high but the rules are short.

## Setup

```bash
npm install
npx lefthook install
git config commit.template .gitmessage
```

`lefthook install` wires the local `commit-msg` hook. The same check runs in CI.

## Commits

[Conventional Commits 1.0](https://www.conventionalcommits.org/en/v1.0.0/),
enforced by `commitlint` locally and on PRs.

```
<type>(<scope>): <subject>

<body>

<footer>
```

- `type` — one of `feat`, `fix`, `perf`, `refactor`, `docs`, `test`,
  `build`, `ci`, `chore`, `style`, `revert`.
- `subject` — imperative, lowercase, no trailing period, **≤ 72 chars**.
- `body` — blank line above. Explain **why**. Wrap at 72 chars.
- `footer` — `Closes: #123`, `BREAKING CHANGE: <description>`.

One commit, one intent. If the subject contains "and", split.

## Pull requests

- Branch off `main`. PR title follows the same rules as the subject above.
- Tests required for new behavior.
- CI must be green. One approving review. Squash-merge only.

## Bug reports

Open a GitHub issue with reproduction steps, expected vs. actual, version, and
environment.

## Security

Do not open public issues for security reports. See `SECURITY.md`.
