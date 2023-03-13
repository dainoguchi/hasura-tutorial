package factories

import (
	"fmt"
	"remote-schema/graph/model"
	"testing"
)

func TestProductFactory(t *testing.T) {
	for i := 0; i < 3; i++ {
		product := ProductFactory.MustCreate().(*model.Product)
		fmt.Println("ID:", product.ID, " Name:", product.Name)
	}
}
