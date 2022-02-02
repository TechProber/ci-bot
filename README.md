# ci-bot

## Local Setup

Install dependencies

```bash
npm install
```

Start the server

```bash
npm run dev
```

## How it works

The `api/github/webhooks/index.js` file is handling requests to `POST` `/api/github/webhooks`, make sure to configure your GitHub App registration's webhook URL accordingly.

## References

- [Probot Documentation](https://probot.github.io/docs/)
- [Probot interacting-with-github](https://probot.github.io/docs/github-api/)
- [Octokit Rest Documentation](https://octokit.github.io/rest.js/v18)
- [Octokit repos-create-commit-status](https://octokit.github.io/rest.js/v18#repos-create-commit-status)
- [Smee.io - Webhook payload delivery service](https://smee.io)
