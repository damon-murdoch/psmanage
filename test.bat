cd C:/users/sirsc/repos/psmanage
go build psmanage.go
move psmanage.exe bin
cd bin
psmanage.exe export teams_test.sd teams_test
psmanage.exe import teams_test teams_test.sd