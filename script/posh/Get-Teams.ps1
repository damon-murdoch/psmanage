Param(
  # [String] Source folder (name only, not path) to pull teams from
  [Alias()][Parameter(Mandatory=$False)][String]$SourceName = $TeamsPath,

  # [String] Target file (name only, not path) to save the teams to
  [Alias()][Parameter(Mandatory=$False)][String]$TargetName = $TeamsFile
);

# Get the latest version of the repository
# If the repository does not exist or it 
# has not been configured, this method will
# handle the configuration.
Get-Repository $SourceName;

# Get the current location
$Location = Get-Location;

# Set the location to the path above the target path
Set-Location (Split-Path -Path $TargetName -Parent);

# Import the folder teams into the target file
psmanage import (Split-Path -Path $SourceName -Leaf) (Split-Path -Path $TargetName -Leaf);

# Set back to the previous path
Set-Location $Location;

# Open the target file in vs code
Invoke-Command { Code $TargetName };