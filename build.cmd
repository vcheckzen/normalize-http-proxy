@echo off

if exist dist (
    del /s /q dist\*
)

set CGO_ENABLED=0
set GOARCH=amd64

set GOOS=windows
call :compile .exe

set GOOS=linux
call :compile

set GOARCH=arm64
call :compile

set GOOS=darwin
call :compile

set TOOLCHAIN=C:\Users\LOGI\AppData\Local\Android\Sdk\ndk\android-ndk-r21b\toolchains\llvm\prebuilt\windows-x86_64
set CC=%TOOLCHAIN%\bin\armv7a-linux-androideabi16-clang.cmd
set CGO_ENABLED=1
set GOARCH=arm
set GOOS=android
call :compile

exit

:compile
go build -v -o dist\np_%GOOS%_%GOARCH%%1 -trimpath -ldflags "-s -w -buildid="
exit /b
