$action = $args[0]

If($action -eq 'run')
{
    Write-Output 'app is starting...'
    $env:APP_CHAR_BED_DEVELOP=1 & go run .\src\main.go
}
elseIf($action -eq 'build'){
    Write-Output 'app is packing...'
    go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-w -s" .\src\main.go
}
else{
    Write-Output "Please enter the correct command："
    Write-Output "                      run   : run the app"
    Write-Output "                      build : pack the app"
}


