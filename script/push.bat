rem Change directory to the parent folder
rem UPDATE THIS TO USE YOUR OWN FOLDER PATH!
cd "C:/users/sirsc/teams"

rem Open the raw text teams file in vs code
start cmd /k code "teams.sd"

rem Wait for the user to finish editing
set /p wait=Hit ENTER to continue...

rem Export the new teams to a repository
psmanage export "teams.sd" "teams"

rem Change directory to the new teams folder
rem UPDATE THIS TO USE YOUR OWN FOLDER PATH!
cd "C:/users/sirsc/teams/teams"

rem [if it does not already exist] create a new git repository

if exist ".git" (
  rem No need to do anything
) else (
  rem Create a new git repository
  git init .

  rem Prompt for a remote url path
  set /p origin="Set repository url path:"

  rem Add the remote url path to the repository
  git remote add origin %origin%
)

rem Add all of the new files to the repository
git add .

rem Prompt the user for a commit message
set /p message="Commit message:"

rem "Add the new files to the repository"
git commit -m "%message%"

rem "Pushing repository to master ..."
git push origin master

rem Wait for the user confirmation
set /p wait=Execution finished. Press enter to quit.