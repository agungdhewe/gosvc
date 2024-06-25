@echo off
rem jalankan pada mode admin

net stop test-service-golang
sc delete test-service-golang
