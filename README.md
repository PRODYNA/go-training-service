# Go Training Service

---

## Purpose

The purpose of this project is to allow participants of the Go training sessions to have a hands-on experience, developing and 
deploying a simple Go service.

## Usage

Using this project to develop and deploy a small Go service is quite simple and straightforward. To create and deploy a
new service simply follow the steps outlined below:

1. Pick a name for the folder where your source code will live (under the `src` directory).
2. Start by branching off of the project's `main` branch. Your new branch should follow this convention `team/<name_from_step_1>`.
3. Copy all the contents from the `template` folder. 
4. Make sure to rename all occurrences of `template` to match the name you selected at step 1.
5. Perform all the desired code changes and push your local branch to the remote
6. Open up a `PR` and verify that your code can be built without any issues.
7. Merge your `PR` to `main` and wait for the deployment process to complete
8. Verify you service is reachable by using the following link `<name_from_step_1>.20.71.73.61.nip.io`.

:warning: **When merging your `PR` to the `main` branch please make sure to avoid**