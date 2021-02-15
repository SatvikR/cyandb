@echo OFF

set args=-i
set bin=.\bin
set go=go

if "%1" == "clean" (
	Rmdir /S/Q %bin%
) else (
	%go% build %args% -o %bin%\cyandb.exe cyan.go
)
