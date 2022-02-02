/**
 * @param { import ('probot').Probot } app
 */

module.exports = (app) => {
  app.log("ci-integration-bot is loaded!");

  // Listen for a pull request being opened or synchronized
  app.on("pull_request", async (context) => {
    const pr = context.payload.pull_request;
    const repo = context.payload.repository;

    const payload = {
      owner: repo.owner.login,
      repo: repo.name,
      sha: pr.head.sha,
      state: "success",
      context: "continuous-integration/tekton/pr",
      description: "Build Started",
    };

    // Create the status object
    return await context.octokit.rest.repos.createCommitStatus(payload);
  });
};
