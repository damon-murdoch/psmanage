<#
    .SYNOPSIS
    Opens a visual studio code window of an empty
    file, which should have the Pokemon Showdown teams
    export pasted into it, saved and closed.

    .DESCRIPTION
    Opens a visual studio code window of an empty
    file, which should have the Pokemon Showdown teams
    export pasted into it, saved and closed. Once the file
    has been closed, it is exported to a file structure using
    psmanage.exe and uploaded to the remote git repository.

    # Requirements:
    # - Have psmanage.exe added to path
    # - Have Visual Studio Code Installed
    # - Have Git Command Line Installed

    .PARAMETER SourceFile
    Text file which will be opened, 
    and converted to a directory

    .PARAMETER TargetPath
    Folder directory which will be the 
    location of the converted directory

    .INPUTS
    [String][String]

    .OUTPUTS
    [None]

    .NOTES
    Author: Damon Murdoch
    Last Updated: 18/02/2021
#>

param(

    # Text file which will be opened, 
    # and converted to a directory
    [Alias()][Parameter(Mandatory=$False)][String]
    $SourceFile = $ENV:PSManageFile,

    # Folder directory which will be the 
    # location of the converted directory
    [Alias()][Parameter(Mandatory=$False)][String]
    $TargetPath = $ENV:PSManagePath
);

Try 
{
    # If the $ENV:PSManageFile path
    # has not been assigned
    If (-not $SourceFile)
    {
        # Prompt the user for a path
        Write-Host "No source file path has been assigned.";
        Write-Host "If the file provided does not exist, it will be created.";
        $SourceFile = Read-Host "Source File";

        # Save argument as environment variables for future use
        [System.Environment]::SetEnvironmentVariable("PSManageFile",$SourceFile);
    }

    # If the $ENV:PSManagePath path
    # has not been assigned
    If (-not $TargetPath)
    {
        # Prompt the user for a path
        Write-Host "No target path has been assigned.";
        Write-Host "If the path provided does not exist, it will be created.";
        $TargetPath = Read-Host "Target Path";

        # Save argument as environment variable for future use
        [System.Environment]::SetEnvironmentVariable("PSManagePath",$TargetPath);
    }

    # Open the source file 
    # in visual studio code

    Code $SourceFile --Wait;

    # Test to see if the target path already exists

    # If the path does already exist
    If (Test-Path -Path $TargetPath)
    {
        # Do nothing, get repository should already exist
    }
    Else # Path does not exist, no git repo
    {
        # Create the new folder
        New-Item -ItemType Directory -Path $TargetPath;

        # Move to the target path
        Set-Location $TargetPath;

        # Create a git repository in the folder
        Git init .

        Write-Host "No Repository Set. Please provide one now.";
        Write-Host "Repo Link (e.g. https://github.com/scrubbs/teams)";
        
        # Prompt the user to provide repo link
        $Repository = Read-Host "Link";

        # Add the repository as the remote link
        Git remote add origin $Repository;
    }

    # Attempt to convert the 
    # source file to a folder
    # at the target path

    psmanage export $SourceFile $TargetPath;

    Try
    {
        # Attempt to move to the new directory
        Set-Location $TargetPath -ErrorAction Stop;

        # Add all teams in the folder
        Git add .

        # No -m switch, as we want to provide a message manually
        Git commit

        # Push the repository to the master branch
        Git push origin master
    }
    Catch
    {
        Write-Output "Failed to push '$TargetPath'! Reason: $($_.Exception.Message)";
    }
}
Catch 
{
    Write-Output "Failed to save teams! Reason: $($_.Exception.Message)";
}