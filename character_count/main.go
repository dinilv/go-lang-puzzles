package main
import (
    "fmt"
    "sync"
    "strings"
)

var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var wg sync.WaitGroup // waits for  diffrent go-routines to complete

type SyncCharacterMap struct {
    // synched operation against singleton
    access *sync.Mutex
    characterCount map[string]int
}

func (s SyncCharacterMap) Add(char string,count int) bool{
    success:=false
    s.access.Lock()
    fmt.Println("In Lock Sync")
    s.characterCount[char] = count
    fmt.Println("In Lock Write")
    success=true
    s.access.Unlock()
    return success
}

func (m SyncCharacterMap) Iterator() {
  fmt.Println("Length",len(m.characterCount))
  for _,v:= range letters {
    fmt.Println("Letter:-",string(v),"Count:-",m.characterCount[string(v)])
  }
}


func main(){
  var text="The quick brown fox jumps over the lazy dog"
  var characterCheck =SyncCharacterMap{
    access: &sync.Mutex{},
    characterCount: make(map[string]int),
}

  for _,char:= range letters{
      wg.Add(1)
      go countChar(text,string(char) ,&characterCheck )
  }
  wg.Wait()
  characterCheck.Iterator()
}

//count each character against text
func countChar(text string,char string,characterCheck  *SyncCharacterMap){
  for {
  result:=characterCheck.Add(char, strings.Count(text, char))
  if result{
    wg.Done()
    break
  }
}
}
