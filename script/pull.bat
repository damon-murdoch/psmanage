rem Change directory to the teams folder
rem UPDATE THIS TO USE YOUR OWN FOLDER PATH!
cd "C:/users/sirsc/teams/teams"

rem Get the latest version from the repository
git pull origin master

rem Change directory to the parent folder
rem UPDATE THIS TO USE YOUR OWN FOLDER PATH!
cd "C:/users/sirsc/teams"

rem Import the latest version into raw text
psmanage import "teams" "teams.sd"

rem Open the latest version in VS Code
start cmd /k code "teams.sd"