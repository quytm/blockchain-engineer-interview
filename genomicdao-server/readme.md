# genomicdao-server
This project simulates the backend of genomicDAO system.

The system runs step-by-step as follows:
0. Init system
1. Register user
2. Sample Recording
3. laboratory processes the saliva sample and generates the user’s raw genetic data 
4. Report then return riskScore
5. Upload Data to Blockchain 
6. Confirm in Blockchain 
7. Get PCSP balance of user

## Setup
Require: `golang >= 1.21`

1. Clone `config.tmp.yml` then save as `config.yml`. Fill your config into `config.yml`. 

2. Generate Golang code from smart contracts
```shell
make sm-generate # clean and generate go code
```
3. Build project
```shell
go mod tidy
```
4. Run code
```shell
make run
```

