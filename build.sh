go build -ldflags -H=windowsgui  
tar czf winawake.tgz winawake.exe
rm -rf winawake.exe
