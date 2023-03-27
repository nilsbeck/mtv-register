# mtv-register
This tiny tool lets you automatically register for courses at MTV Braunschweig.

## Setup

* Create env variables for `MTV_MAIL` and `MTV_PASS` to login. Write them into your bash/zsh profile.

  ```
  export MTV_MAIL="mail@mail.com"
  export MTV_PASS="myPass"
  ```
* build with:

  ```
  go build main.go
  ```
 
* run binary with the course number, you extract it from the course url and a time selector. Counting starts with 1 from the drop down. If no time needs to be selected because the course is only offered once on a day, pass `-1`.

 ```
 ./main 297 1
 ```
 The above will inscribe you to Turnzwerge II on a Thursday (`297`) for 15:30 (`1`).
 
 ## Crontab
 
 To run as a cron job, add the PATH variable to your crontab:
 
```
 PATH=<paths from `.bashrc`>
```

and run the binary as the following example shows:

```
30 16 * * 4 /bin/bash -l -c '<ABS_PATH_TO_BINARY>/main 297 2 >> <ABS_PATH_TO_LOGS>/logging 2>&1'
```

# Troubleshooting

If you encounter problems, there is sample code in the main.go file that you can use to create screenshots of the browser's viewport.
