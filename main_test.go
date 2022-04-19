package main
import (
	"testing"
)
 
//猫会叫
func TestCat(t *testing.T)  {
	saying := Cat()
	if saying != "Miao~~" {
		t.Errorf("Cat saying %s",saying)
	}
}