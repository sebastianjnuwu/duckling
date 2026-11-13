---
name: github-workflow
description: >-
  Use this skill whenever committing, pushing, opening pull requests, or merging changes to GitHub.
  Enforces creating a dedicated feature/fix branch, running pre-commit verification and changelog updates,
  committing with Conventional Commits, opening a Pull Request via GitHub CLI (`gh`), and merging it back into `idle`.
---

# GitHub Branch, PR & Merge Workflow

Never commit or push directly to the default branch (`idle`). Every change must go through a dedicated topic branch and a Pull Request.

## 1. Pre-Flight Verification

Before staging or committing any files, verify all project constraints:

1. **Line count check**: Ensure all created or modified `.go` files have at most 100 lines (ideally ≤75 lines).
2. **Build & Vet**: Run the Go compiler and static analyzer:
   ```powershell
   go build ./...
   go vet ./...
   ```
3. **Changelog**: Update `CHANGELOG.md` under `## [Unreleased]` describing the change. Never commit without updating `CHANGELOG.md`.
4. **Clean up**: Ensure no temporary test binaries or scratch files remain outside `tmp/`.

## 2. Create a Topic Branch

Sync with the base branch (`idle`) and create a descriptive branch using `<type>/<short-slug>`:
- `feat/<slug>` — New feature or gameplay mechanic
- `fix/<slug>` — Bug fix
- `docs/<slug>` — Documentation or agent config changes
- `refactor/<slug>` — Code restructuring without behavior changes
- `chore/<slug>` — Build, tooling, or maintenance

```powershell
git checkout -b <type>/<short-slug>
```

*(If you are already on `idle` with uncommitted changes, `git checkout -b <type>/<short-slug>` carries your working tree changes onto the new branch.)*

## 3. Stage, Commit & Push

Always ensure the repository Git author is configured as `sebastianjnuwu` (`sebastianjnuwu@gmail.com`), stage only the relevant files (including `CHANGELOG.md`), and commit using Conventional Commits (`<type>(<scope>): <summary>`):

```powershell
git config user.name "sebastianjnuwu"
git config user.email "sebastianjnuwu@gmail.com"
git status
git add <files...> CHANGELOG.md
git commit -m "<type>(<scope>): <concise summary>"
git push -u origin <type>/<short-slug>
```

## 4. Open Pull Request & Merge

Use the GitHub CLI (`gh`) to create the Pull Request targeting `idle`, and merge it once verified:

1. **Create the PR**:
   ```powershell
   gh pr create --base idle --head <type>/<short-slug> --title "<type>(<scope>): <concise summary>" --body "## Summary`n- <bullet points of changes>`n`n## Verification`n- `go build ./... && go vet ./...` passed"
   ```
2. **Merge the PR & Delete Remote/Local Branch**:
   ```powershell
   gh pr merge --squash --delete-branch
   ```
   *(Use `--merge` instead of `--squash` if preserving individual commits is preferred by the user.)*
3. **Verify Local State**:
   ```powershell
   git checkout idle
   git pull origin idle
   ```
