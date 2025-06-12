package service

//go:generate powershell -Command "if (Test-Path mocks) { Remove-Item -Recurse -Force mocks }; New-Item -ItemType Directory -Path mocks"
//go:generate minimock -i UserService -o ./mocks/ -s "_minimock.go"
