@echo off
rem jalankan dalam mode admin

if not exist testservice.exe (
    echo "File testservice.exe tidak ditemukan"
    echo "Build dulu project dengan output testservice.exe"
    goto :exit
)

sc create test-service-golang binpath= "%CD%\testservice.exe" start= auto DisplayName= "Test Service Golang"
sc description test-service-golang "Testing windows service dengan bahasa golang"
net start test-service-golang
sc query test-service-golang

:exit