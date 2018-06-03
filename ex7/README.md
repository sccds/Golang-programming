# Word Count


### Goal
1. Get the English word count result from multiple files 
2. Output words in decending order
3. Support concurrent

### Steps

#### Environemnt
```
cd $HOME
mkdir -p golang/src/wordcount
export GOPATH=$HOME/golang
cd $GOPATH/src/wordcount
```

#### Main Program
Create `wordcount.go` under `$GOPATH/src/wordcount`

Create `wordfreq.go` under `$GOPATH`

#### Build and Run
```
go build wordfreq.go
./wordfreq shakespeare.txt | head -n 6
```



