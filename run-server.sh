docker-compose stop && docker-compose up -d --build 
go build cmd/server/main.go 
source env.txt 
export $(cut -d= -f1 env.txt)
nohup ./main &
