# Go-Playground
Go

## Initializing Package
go mod init <package_name>

## Multi-Module Development
### Initialize workspace
go work init
### Initialize modules/packages
mkdir <package_name1> && cd <package_name1> && go mod init <package_name1>
mkdir <package_name2> && cd <package_name2> && go mod init <package_name2>
### Add moddules to workspace
go work init <package_name1> <package_name2>
### Or Use new module
go work use ./<package_name3>
