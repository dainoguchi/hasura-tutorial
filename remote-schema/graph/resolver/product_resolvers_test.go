package resolver

import (
	"database/sql"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/DATA-DOG/go-txdb"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"remote-schema/graph/generated"
	"remote-schema/graph/model"
	"remote-schema/pkg/config"
	"remote-schema/pkg/db"
	"remote-schema/test/factories"
	"testing"
)

func OpenDBForTest(t *testing.T) *gorm.DB {
	cfg, err := config.Environ()
	if err != nil {
		t.Fatal(err)
	}

	txdb.Register("txdb", "postgres", cfg.Database.TestURL)
	sqlDB, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		t.Fatal(err)
	}
	gormDB, err := db.InitWithDB(sqlDB)
	if err != nil {
		t.Fatal(err)
	}

	return gormDB
}

func TestQueryResolver_MaskingProducts(t *testing.T) {
	t.Helper()

	// resolver生成
	r := Resolver{
		DB: OpenDBForTest(t),
	}

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
		r.DB.Create(product)
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
