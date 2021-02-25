@echo OFF

set args=-i
set bin=.\bin
set go=go

if "%1" == "clean" (
	Rmdir /S/Q %bin%
) else if "%1" == "install" (
	%go% install .\cmd\cyand\ %args%
	%go% install .\cmd\cyansh\ %args%
) else (
	%go% build .\cmd\cyand\main.go %args% -o %bin%\cyand.exe
	%go% build .\cmd\cyansh\main.go %args% -o %bin%\cyansh.exe
)
