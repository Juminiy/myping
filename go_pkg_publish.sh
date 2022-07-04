#!/bin/bash
go mod tidy 
go test ./utils/ip_domain_format_test.go
go test ./signal/icmp_test.go

VERSION='v1.0.5'
git add .
git commit -m "myping: publish for $VERSION"
git tag $VERSION
git push origin $VERSION

GOPROXY=proxy.golang.org go list -m github.com/Juminiy/myping/@$VERSION