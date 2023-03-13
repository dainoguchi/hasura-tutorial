package resolver

import (
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/assert"
	"remote-schema/graph/generated"
	"remote-schema/graph/model"
	"remote-schema/test/factories"
	"testing"
)

func TestQueryResolver_MaskingProducts(t *testing.T) {
	t.Helper()

	// resolver生成
	r := Resolver{}

	// clientの作成
	c := client.New(handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &r},
		),
	))

	// テストデータ投入
	for i := 0; i < 3; i++ {
		product := factories.ProductFactory.MustCreate().(*model.Product)

		// resolverに入れたいんだけど...
		r.products = append(r.products, product)
	}

	var resp struct {
		Masking_products []struct {
			Id   int
			Name string
		}
	}

	q := `{ masking_products { id, name } } `

	// クエリ実行
	c.MustPost(q, &resp)

	// assertion
	assert.Equal(t, 3, len(resp.Masking_products))
	for _, v := range resp.Masking_products {
		assert.Equal(t, "xxxxxxxxx", v.Name)
	}
}
