package main

// "bufio"
// "io"

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// check(e: error): void
// Check if an error has occured,
// and if it has run panic.
func check(e error) {
	
	// If error is not null
	if e != nil {

		// Run panic function
		panic(e)
	}
}

// isdir(string) : bool
// Given a file path, returns if 
// the file is a folder 
// or a standard file.
func isdir(path string) (b bool) {

	// Test the folder path
	fi, err := os.Stat(path);

	// Ensure no error occured
	check(err)

	// Switch on file mode
	switch mode := fi.Mode(); {
		
		// If the file 
		// is a directory
		case mode.IsDir():
			return true

		// Otherwise, file
		// is a standard file
		default:
			return false
	}
}

// children(): void
// Gets all of the files and 
// folders in a given file path.
func children(root string) (f []string) {

	// Empty array, will be
	// populated and returned
	var files []string;

	// Walk through all of the files in the 'root' directory
	filepath.Walk(root, func(path string, info os.FileInfo, err error) (error) {

		// Append the 'path' string to the files array
		files = append(files,path);

		// Return nil if successful
		return nil
	})

	// Return the files
	// to the terminal
	return files;
}

// write(path: string, str: string): void
// Given a file path and a string, writes the
// string to the given file path
func write(path, str string) {

	// Create a folder to write the old string to
	file, err := os.Create(path);

	// If the file failed to create
	if (err != nil) {

		// Do nothing

	} else {

		// Trim the leading and trailing newlines
		content := strings.Trim(str,"\r\n ");

		// Write the string to the file 
		_, err = file.WriteString(content);

		// If string failed to write
		if (err != nil) {

			// Do nothing

		}

	}

	// Defer the file handle closing
	defer file.Close();
}

// main(void): void
// Main function
func main(){

	// Get the arguments
	// (Excluding the file name)
	args := os.Args[1:];

	// File name + File Path
	// path := os.Args[0];

	// If there is (at least)
	// command line argument
	if len(args) > 1 {

		// Dereference the action
		action := args[0];

		// First Argument
		source := args[1];

		// Second argument
		var target string;

		// Two possible actions:
		// import, export

		if action == "import" {

			// If a second argument is provided
			if len(args) > 2 {

				// Second Argument
				target = args[2];

			} else {

				// Use default (input folder plus extension)
				target = source + ".sd";
			}

			// psmanage import [sourcefile] [targetfolder]
			// e.g. psmanage import teams.sd teams

			// String storing all of the teams
			var library string;

			// Get the child files in the source folder
			folder := children(source);

			// Get the folder path depth 
			// in the source folder
			sdepth := strings.Count(source, "\\") + 1;

			// Iterate over all of the 
			// files in the folder
			for _, file := range folder {

				// Get the folder path depth 
				// of the current folder
				depth := strings.Count(file,"\\");

				// Get difference between
				// the depth and source depth
				diff := depth - sdepth;

				// One layer of depth
				// - Format Level Depth
				if diff == 1 {

					// If the file is not a folder AND
					// If the file has the '.sd' extension
					if !isdir(file) && strings.HasSuffix(file,".sd") {

						// Read the content from the file
						data, err := ioutil.ReadFile(file);

						// Check that no error has occured
						check(err);

						// Split filename and foldername 
						// on the backslash (windows separator)
						files := strings.Split(strings.Replace(file,".sd","",1),"\\");

						// Get the string from the raw data
						// Create the team name from the filename + folder name
						content := "=== [" + files[len(files)-2] + "] " + files[len(files)-1] + " ===\n\n" + string(data);

						// Append the content from the page to your teams library
						// Purge the leading and trailing spaces before adding
						library += strings.TrimSpace(content) + "\n\n";

					}

				// Two layers of depth
				// - Folder Level Depth
				} else if diff == 2 {

					// If the file is not a folder, and has the '.sd' extension
					if !isdir(file) && strings.HasSuffix(file,".sd") {

						// Read the content from the file
						data, err := ioutil.ReadFile(file);

						// Check that no error has occured
						check(err);

						// Split filename and foldername 
						// on the backslash (windows separator)
						files := strings.Split(strings.Replace(file,".sd","",1),"\\");

						// Get the string from the raw data
						// Create the team name from the file name + folder name + subfolder name
						content := "=== [" + files[len(files)-3] + "] " + files[len(files)-2] + " / " + files[len(files)-1] + " ===\n\n" + string(data);

						// Append the content from the page to your teams library
						// Purge the leading and trailing spaces before adding
						library += strings.TrimSpace(content) + "\n\n";
					}

				// No layer(s) of depth or
				// more than two levels
				} else {

					// Can ignore it

				}
			}

			// Open the output file 
			outfile, err := os.Create(target);

			// Check for any errors
			check(err);

			// Write the imported library to the target file
			// Strim leading / trailing whitespace, newline
			_, err = outfile.WriteString(strings.Trim(library,"\r\n "));

			// Check for any errors
			check(err);

			// Close the output file
			defer outfile.Close();

		} else if action == "export" {

			// If a second argument is provided
			if len(args) > 2 {

				// Second Argument
				target = args[2];

			} else {

				// Use default (input filename minus extension)
				target = strings.Split(source,".")[0];
			}

			// psmanage export [sourcefolder] [targetfile]
			// e.g. psmanage export teams teams.sd

			data, err := ioutil.ReadFile(source);

			// Check that no error has occured
			check(err);

			// Get the string from the raw data
			content := string(data);

			// Get the current working directory
			cwd, err := os.Getwd();

			// Split the file into lines
			split := strings.Split(content,"\n");

			// Content to write to the next file
			outstr := "";

			// New file which will be created
			outfile := "";

			// Iterate over lines in the file
			for _, line := range split {

				// Current working directory + target folder
				fpath := cwd + "\\" + target;

				// If the path does not exist
				if _, err := os.Stat(fpath); os.IsNotExist(err) {

					// Create a folder for the target
					err := os.Mkdir(fpath, os.ModeDir);
					
					// If an error has occured
					if(err != nil) {
						// Report error to terminal
						fmt.Println("Error occured creating file",fpath,":",err);
					}

					// Non terminating, continue
				}

				// If the file path does exist
				if _, err := os.Stat(fpath); !os.IsNotExist(err) {

					// If the line contains '='
					// This only gets inserted at
					// the start of teams
					if strings.Contains(line,"=") {

						// If there is any content sitting 
						// in the output string
						if len(outfile) > 0 && len(outstr) > 0 {

							// Write to the outfile
							write(outfile, outstr);

							// Null the outfile
							outfile = "";

							// Null the outstr
							outstr = "";
						}

						// Get the category from the title
						category := strings.TrimSpace(strings.Split(strings.Split(line,"[")[1],"]")[0]);

						// Get the path for the category folder
						path := fpath + "\\" + category;

						// If a folder does not exist for this category
						if _, err := os.Stat(path); os.IsNotExist(err) {

							// Create the missing folder
							err := os.Mkdir(path,os.ModeDir);

							// Check to see if there is any errors
							check(err);
						}

						// Get the name of the team
						name := strings.TrimSpace(strings.Split(strings.Split(line,"]")[1],"=")[0]);

						// Regex only accepts letters and numbers
						reg, err := regexp.Compile("[^a-zA-Z0-9 -]+");

						// If the error message is not null
						check(err);

						// If the string contains '/' (folder separator)
						if strings.Contains(name,"/") {

							// Split the name into name / folder on '/'
							subsplit := strings.Split(name,"/");

							// subsplit[0] = folder
							// Remove all special characters which may break the file name
							folder := strings.TrimSpace(reg.ReplaceAllString(subsplit[0],""));

							// Remove all special characters which may break the file name
							name = strings.TrimSpace(reg.ReplaceAllString(subsplit[1],""));

							// Need to create / check for another folder 'base/target/folder'
							fullpath := path + "\\" + folder;

							// If a folder does not exist for this showdown folder
							if _, err := os.Stat(fullpath); os.IsNotExist(err) {

								// Create the missing folder
								err := os.Mkdir(fullpath,os.ModeDir);

								// Check to see if there is any errors
								check(err);
							}

							// Designate the filename of the new team
							outfile = fullpath + "\\" + strings.Title(strings.ToLower(name)) + ".sd";

						} else {

							// No extra folder, just insert into the base folder
							name = strings.TrimSpace(reg.ReplaceAllString(name,""));

							// Designate the filename of the new team
							outfile = path + "\\" + strings.Title(strings.ToLower(name)) + ".sd";

						}

						// Folder should now exist, has
						// been created if it does not

						// Convert the line to uniform title case
						// (e.g. all lower case, except first letter of string)
						// line = strings.Title(strings.ToLower(line));

					} else {

						// Add the line to the output 
						outstr += line;
					}

				} else {
					fmt.Println("Cannot write to file",fpath," as it does not exist. Skipping ...");
				}
			} 

			// Done now, create the last file

			// If there is a file left to write
			if len(outfile) > 0 && len(outstr) > 0 {

				// Write to the outfile
				write(outfile, outstr);
			}

		} else {

			// Unrecognised command
			fmt.Println("Unrecognised command: ",action);

			// Report correct arguments to terminal
			fmt.Println("psmanage import [source file] [target folder]");
			fmt.Println("psmanage export [source folder] [target file]");
		}

	} else {

		// Not enough arguments, 
		fmt.Println("psmanage import [source file] [target folder]");
		fmt.Println("psmanage export [source folder] [target file]");
	}
}