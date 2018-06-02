package functionality

// Sources:
// http://cs-guy.com/blog/2015/01/test-main/

import (
	"testing"
	"os"
)



func TestMain(m *testing.M){
	Setup()
	code := m.Run()
	os.Exit(code)
}

func Setup(){
	MongoDB = GetADB()
	MongoDB.Init()
}

