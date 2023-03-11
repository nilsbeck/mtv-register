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
 
* run binary with the course number, you extract it from the course url and a time selector. Counting starts with 1 from the drop down

 ```
 ./main 297 1
 ```
 The above will inscribe you to Turnzwerge II on a Thursday (`297`) for 15:30 (`1`).
