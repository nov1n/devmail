#What
A script to get names and email from all contributors of a git repository.

#Why
- To contact a developer whose email you cannot find online.
- To build a mailing list targeting developers.

#How
Git stores the user's name and email for each commit.

#Running
Use `go build` to create an executable. Then run `./devmail git@github.com:foo/bar.git`.

#Optional flags
- --keep=true, does not remove the cloned repository after the script exits
- --dir=/foo/bar, overwrite the default directory location
- --name=Bob, searches for users named Bob
- --help, print usage

#Write to a file
To save the output in a file use `./devmail https://github.com/foo/bar > devmails.txt`

#Example
The following command `./devmail git@github.com:facebook/react.git` produces the following output:
```
Adam Solove <asolove@gmail.com>
Adam Stankiewicz <sheerun@sher.pl>
Adam Zapletal <adamzap@gmail.com>
adraeth <jerzy.mirecki@gmail.com>
Adrian Sieber <mail@adriansieber.com>
Ahmad Wali Sidiqi <wali-s@users.noreply.github.com>
Alan deLevie <adelevie@gmail.com>
Alan Plum <me@pluma.io>
Alan Souza <alansouzati@gmail.com>
Alastair Hole <afhole@gmail.com>
...
```
