# PennyWise
Wisdom with few pennies brings wisdom with many dollars. 

# Go Setup Basics

## Installation
- Follow the official docs to install for your machine. [Go Installation](https://golang.org/doc/install)

## Environment Setup
- When creating a new project, run `Go mod init <module_name>` to initialize a new module. Usually the name of your project works here.
- Then you can install Packages using `go get <package_name>`.

# Features:
- csv import for transactions
- login page
- transction (CRUD)
    - table view
    - form view
    - detail view
- categories
    - categorize transactions
    - view budget by category (% consumed)
- basic graphical display of cash flow (maybe hold off for time being)
- OAuth?!?
- Create a turso db
- Setup basic connection from Go to db

## Responsibilities:
- CJ
    - 
- Landan
    - csv import for transactions


# Git Workflow

### 1. Switch to main & update:
In the integrated terminal, ensure you're on main and pull the latest changes:

`git checkout main`

`git pull origin main`

### 2. Create a new branch:
Create and switch to your UI feature branch:

`git checkout -b feature/ui`

### 3. Work on the UI feature:
Open your files (like index.html) in VS Code, make changes, and save.
Stage and commit your changes:

`git add .`

`git commit -m "Implement initial UI for landing screen"`

### 4. Check changes from other developers:

Use the Source Control pane in VS Code to review incoming changes.
Or in the terminal, you can fetch updates:

`git fetch origin`

Compare your branch with main:

`git diff main`

To see commit history, run:

`git log --oneline --graph --all`

### 5. Merge changes from main into your feature branch:
Before merging back, pull the latest changes from main into your branch:

`git checkout main`

`git pull origin main`

`git checkout feature/ui`

`git merge main`

Resolve any conflicts, then commit the merge.

### 6. Merge your feature branch back into main safely:
Once your feature is complete and tested, switch back to main and merge your branch:

`git checkout main`

`git merge feature/ui`

`git push origin main`

Alternatively, you can push your branch and open a Pull Request so that other devs can review your changes.

## Git Merge vs. Git Rebase

### Git Merge:

- Combines changes from one branch into another by creating a merge commit.
- Keeps the complete history of both branches.
- Useful when you want to preserve the context of the feature branch and maintain a record of how branches have diverged and come back together.
  
Example:

`git checkout main`

`git merge feature/ui`

### Git Rebase:

- Moves or "replays" commits from one branch onto another, creating a linear history.
- Results in a cleaner, more linear project history, but rewrites commit history.
- Best used for local changes before sharing them, not for branches that others are using.
  
Example:

`git checkout feature/ui`

`git rebase main`

### When to Use Each:

**Use Merge when:**
- You want to preserve the exact history and context of the feature branch.
- Working in a collaborative environment where rewriting history is risky.
- Merging finished features in to the main branch.

**Use Rebase when:**
- You want a cleaner, linear history.
- You're working on local changes and want to update your feature branch with the latest main branch commits.
- You need to resolve conflicts on your local branch before sharing your changes.