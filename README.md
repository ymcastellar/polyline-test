# polyline-test

## The Assignment 🔧
### Problem
#### Intro
There is one of the apps we've developed which uses geo-location information to build some business value for one of our clients. This app generates log files (plain) on the mobile device of the user which are useful to debug information from the users.

These logs are written in plain-text as the one you can find [here](https://gist.github.com/paganotoni/139a3a9af08c29393087405e1c23b0ac), and apart from geo-location information it container other application-related info.

### The task
We would like you to write a Go CLI that receives the path to one of these log files and based on it generates and print an encoded polyline with the locations it found in the log files.

Encoded Polylines look like this:

```
_p~iF~ps|U_ulLnnqC_mqNvxq`@
```

### How to run

1. Extract .rar file
2. Enter to poly-test folder
3. Open the Terminal.
4. Type ``` go run .\main.go example.log ```

## Author ✒️

_Developed by_

**Yoyman Castellar** :computer: :man: 

- Github - [ymcastellar](https://github.com/ymcastellar)
- Twitter - [@CastellarYoyman](https://twitter.com/CastellarYoyman)
- LinkedIn - [yoyman-castellar](https://www.linkedin.com/in/yoyman-castellar/)