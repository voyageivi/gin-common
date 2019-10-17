package test

import (
	"fmt"
	"github.com/voyageivi/gin-common/setting"
	"testing"
)

func Test1(t *testing.T) {
	setting.Setup()
	fmt.Println(setting.Config)
	fmt.Println("test")

}
