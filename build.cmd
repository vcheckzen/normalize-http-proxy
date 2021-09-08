@echo off

if exist dist (
    del /s /q dist\*
)

set CGO_ENABLED=0
set GOARCH=amd64

set GOOS=windows
call :compile windows_amd64.exe

set GOOS=linux
call :compile linux_amd64

set GOARCH=arm64
call :compile linux_arm64

set GOOS=darwin
call :compile darwin_amd64

set TOOLCHAIN=C:\Users\LOGI\AppData\Local\Android\Sdk\ndk\android-ndk-r21b\toolchains\llvm\prebuilt\windows-x86_64
set CC=%TOOLCHAIN%\bin\armv7a-linux-androideabi16-clang.cmd
set CGO_ENABLED=1
set GOARCH=arm
set GOOS=android
call :compile android_arm

exit

:compile
go build -v -o dist\np_%1 -trimpath -ldflags "-s -w -buildid="
exit /b
