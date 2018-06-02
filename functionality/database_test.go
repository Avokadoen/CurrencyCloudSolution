package functionality

import "testing"
/*

func TestMongoDBInfo_AddDailyFix(t *testing.T) {

}
*/

func TestMongoDBInfo_GetLocalFixer(t *testing.T) {
	_, got := MongoDB.GetLocalFixer(0)
	if got != true{
		t.Error("Could not retrieve latest fixerdata")
	}

}