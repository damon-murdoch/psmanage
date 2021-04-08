Param(
  # [String] Source file (name only, not path) to pull teams from
  [Alias()][Parameter(Mandatory=$False)][String]$SourceName = $TeamsFile,

  # [String] Target folder (name only, not path) to save the teams to
  [Alias()][Parameter(Mandatory=$False)][String]$TargetName = $TeamsPath
);

# Open a VS Code window to pull teams from
Invoke-Command { Code $SourceName };

# Wait for the user to confirm
Read-Host "Press ENTER to continue ...";

# Get the current location
$Location = Get-Location;

# Set the location to the path above the target path
Set-Location (Split-Path -Path $TargetName -Parent);

# Export the source file to the target file
psmanage export (Split-Path -Path $SourceName -Leaf) (Split-Path -Path $TargetName -Leaf);

# Set back to the previous path
Set-Location $Location;

# Push the repository to the remote path
# If the repository does not exist / has not 
# been configured, this command will handle
# the configuration.
Push-Repository -Path $TargetName;