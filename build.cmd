@echo off

if exist dist (
    del /s /q dist\*
)

set GOARCH=amd64
call :build-desktop

set GOARCH=arm64
call :build-desktop

call :build-mobile

exit 0

:compile
go build -v -o dist\np_%GOOS%_%GOARCH%%1 -trimpath -ldflags "-s -w -buildid="
exit /b

:build-desktop
set CGO_ENABLED=0
set GOOS=windows
call :compile .exe

set GOOS=linux
call :compile

set GOOS=darwin
call :compile
exit /b

:build-mobile
set TOOLCHAIN=C:\Users\LOGI\AppData\Local\Android\Sdk\ndk\android-ndk-r21b\toolchains\llvm\prebuilt\windows-x86_64
set CC=%TOOLCHAIN%\bin\armv7a-linux-androideabi16-clang.cmd
set CGO_ENABLED=1
set GOOS=android
set GOARCH=arm
call :compile
exit /b