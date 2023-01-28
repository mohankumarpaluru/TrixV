# Contributing to tube

We welcome contributions from the community and are grateful for any help you can provide!

## Pull Request Conventions

We are using tools to process pull requests. When submitting a pull request,
please follow these conventions. That way can automatically generate the [CHANGELOG.md](CHANGELOG.md).

- Please put a brief description of the changes made in the pull request title.
- If you want the changes to show up in the changelog start the title with
  one of the following words:
    - "Fix" for bug fixes,
    - "Add" for new features,
    - "Document" for changes to the documentation,
    - "Update" for updates of dependencies,
- Include a description of the changes in the pull request body.
  If the changes are not obvious, try to concentrate on the "why" instead of the "what".
  Later developers can always look at the code and see what is going on.
  They can not look into your mind and see "why" you decided to implement it a certain way.
- Include links to the issues that are closed by this PR using the
  format `- Closes #issue_number` or `- Resolves #issue_number`.

### Example Pull Request
#### Title
```
Fix missing validation of input in login screen
```

#### Body

```
The issue was caused by a missing validation check in the login function.
It allowed users to put in values that will be declined by the login handler anyway.
I have added the necessary validation check and added the appropriate error messages,
so the users have instant feedback instead of an error message after submitting the
login form.

- Closes #123
```

## Squashing Commits

Please note that all pull requests will be squashed before merging.
This means that multiple commits on a branch will be combined into
a single commit on the master branch. This helps keep the master
branch history clean and easy to understand.

The pull request's title and body will be used for that squashed commit.
It will become part of the project's history and a help the next developer
to understand the changes that happened.

## Code of Conduct

This project does not have a written code of conduct. Be nice to each other!
If you need further instructions ask your parents or the Dalai Lama. ;-)

## Thank you

Thank you for considering contributing to tube. We appreciate your help and support!
