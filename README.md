# valorant-go
A Go API and CLI for the game Valorant

![image](https://user-images.githubusercontent.com/3121306/104110386-763baa00-52d7-11eb-9b70-7645c6e5d05b.png)

This module includes two packages: 
- A CLI in `cmd` 
- A Valorant API wrapper in `pkg` 

**Current status:** 
- The API wrapper is very fundamental and lacks things like proper error handling. :construction:
- Only feature is displaying the MMR / Elo changes of your latest matches
- Todo: Create binaries via CI

### Usage
Download binaries of the program here: https://github.com/mammuth/valorant-go/releases/

After this, run `valorant.exe matches --region eu --username <your-name> --password <your-password>` to get your match history

You can create a config file at `$HOME/.valorant-go.yaml` like this:
```yaml
region: eu
username: dummy
password: verysecret
```
