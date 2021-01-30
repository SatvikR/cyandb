@echo OFF

set args=-i
set bin=.\bin
set go=go

if "%1" == "clean" (
	Rmdir /S/Q %bin%
) else (
	%go% build %args% -o %bin%\cyan.exe cyan.go
)
