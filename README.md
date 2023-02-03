# Go Projects

## Some Point to Keep In Mind

- Always check the branch you are working on
- Also always make sure that your branch is up to date with your parent branch i.e to take latest pull of both child and parent branch and then run command `git merge parent-branch`

## Steps to Commit/Checkin the Code

- First add the files to staging that you want to commit
- Make sure you don't add files that are not intended to commit
- Now once files are in staging you can now commit the same using an appropriate commit message
- After commit make sure you push those changes

## Commit using command line

- Add the untarcked file to tracking (if any) usinf command `git add filename`
- Now your file is in staging
- Now you can commit your file using command `git commit -m "Message"`
- Now you could have also done the above 2 steps in single command using `git commit -am "Message"` but avoid this to avoid unnecessary checkins
- Now you have to push those commit as well using command `git push`

## Steps

1. Remember The Branch Name
2. Switch To Current Working Branch
3. Take pull using "git pull" command also make sure you take latest pull of the parent branch before merging it to child branch
4. update the branch from main/parent using "git merge parent-branch-name" command
5. Continue The Modifications Or Updates You Do After It's Done Then Commit All The Changes With Proper Commit Message Take Pull And Then Push The Changes And At Last Create PR

## Information About Go

 1. Go Is Not OOP Nor POP Language
