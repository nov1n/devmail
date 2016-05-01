#What
A script to get a user's name and email address from a GitHub repository he or she contributed to.

#How
Use `go build` to create an executable. Then run `./devmail https://github.com/foo/bar`, that's it!

#Optional flags
- --keep=true, does not remove the cloned repository after the script exists
- --dir=/foo/bar, overwrite the default directory location
- --name=Bob, searches for users named Bob
