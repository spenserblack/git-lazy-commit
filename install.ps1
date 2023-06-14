#!/usr/bin/env pwsh
# Downloads and installs from GitHub Releases

if ($Env:PROCESSOR_ARCHITECTURE -eq "x86") {
    $arch = "386"
} else {
    $arch = "amd64"
}

Invoke-WebRequest -OutFile "$Env:ProgramFiles\Git\usr\bin\git-lazy-commit.exe" -Uri "https://github.com/spenserblack/git-lazy-commit/releases/latest/download/git-lazy-commit-windows-$arch.exe"
