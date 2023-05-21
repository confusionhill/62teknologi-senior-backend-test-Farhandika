# 62teknologi senior backend test Farhandika

# Stacks
1. Golang
2. PostgreSQL
3. Docker

# How to run
1. Go to `files->secrets`
2. Create `secrets.config.json`
3. Input secrets value
4. Run `go mod tidy` or `go mod vendor` or `make init`
5. Run `make services-up`
6. If you want to run a specific service, run `SVCNAME={SERVICE_NAME} make services-up`
7. If database is not connected, please restart the http service
7. You can check all the commands available on `Makefile`
8. If you want to check the api docs, go to `files->docs`

# Maintainer
18220015@std.stei.itb.ac.id (Farhandika Zahrir Mufti Guenia)