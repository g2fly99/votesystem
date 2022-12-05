# votesystem
a simple vote system

# interface
1. look interface.md
2. run

# how to run ？
there are 3 way ro start the program:

## dockerfile 1
1. download Dockerfile
2. get ready mysql and redis
3. redis: 127.0.0.1:6379; password:Aa123456!@#
4. mysql: 127.0.0.1:3306; database:ipcc; root password:  Aa123456!@#
5. after build then run the docker file

## dockerfile 2
1. download Dockerfile and directory:conf
2. get ready mysql and redis
3. change Dockerfile, add ENV: ./conf/:/go/src/votesystem/conf
4. change config
5. after build then run the docker file

## docker-compose
1. download docker-compose.tar
2. then run docker-compose up




