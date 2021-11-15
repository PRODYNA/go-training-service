# Go Training Service

## Purpose

The purpose of this project is to allow participants of the Go training sessions to have a hands-on experience, developing and 
deploying a simple Go service.

## Usage

Using this project to develop and deploy a small Go service is quite simple and straightforward. To create and deploy a
new service simply follow the steps outlined below:

### Code preparation

1. Start by creating a branch from `main`. **Make sure that your branch adheres to this naming convention `team/<username>`**.
For example, a valid branch name is `team/dkrizic`.
2. Copy all the contents of the `src/template/` folder in a new folder named `/src/<username>`. **Make sure that the
username defined in the branch name is the same as the name of your folder.**
3. Rename all occurrences of `template` within the copied code to match your **username**.
4. Do all the necessary code changes to your Go code.

### Pull Requests

1. Once all the code changes are done, you can go ahead and open up a pull-request targeting the `main` branch.
2. Opening up a pull request will trigger a CI run for your code. This will build your code and verify its integrity.
3. You can perform any number of commits whilst the PR is open, each triggering a new CI run.

### Merging

1. Upon finalizing your code changes and given that the newly added code gets built properly, you can go ahead and merge to `main`.
2. :warning: **Make sure that you do not tamper with GitHub's default commit message as it contains the branch name**. The latter is
required for the retrieving deployment related information. If a custom message must be provided, please make sure to include
the branch name in it in lowercase.
3. After merging the changes and once the CI is finished, you can verify the liveliness of your service by visiting the following
link `<username>.20.71.73.61.nip.io`