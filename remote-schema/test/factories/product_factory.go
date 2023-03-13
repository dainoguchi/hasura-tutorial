package factories

import (
	"fmt"
	"github.com/bluele/factory-go/factory"
	"remote-schema/graph/model"
)

var ProductFactory = factory.NewFactory(
	&model.Product{},
).SeqInt("ID", func(n int) (interface{}, error) {
	return n, nil
}).Attr("Name", func(args factory.Args) (interface{}, error) {
	product := args.Instance().(*model.Product)
	return fmt.Sprintf("product-%d", product.ID), nil
})
