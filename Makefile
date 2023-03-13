hasura-init:
	hasura init

hasura-console:
	hasura console

hasura-mg-status:
	hasura migrate status --database-name default

hasura-mg-squash: ## 1678351250823_xxxを1678351250823_initに変える(多分認識間違えてる) https://hasura.io/docs/latest/migrations-metadata-seeds/manage-migrations/
	hasura migrate squash --name "init" --from 1678351250823 --database-name default

hasura-mg-create:
	hasura migrate create init --sql-from-file `schema.sql` --database-name display

hasura-mg-apply:
	hasura migrate apply --database-name default

## hasura cloudのstaging環境にmigration
hasura-mg-apply-staging:
	hasura migrate apply --endpoint $(ENDPOINT) --admin-secret $(ADMIN_SECRET) --database-name default

## migrationにより生成されたmetadataをexport
hasura-meta-export:
	hasura metadata export

hasura-meta-apply:
	hasura metadata apply --endpoint $(ENDPOINT) --admin-secret $(ADMIN_SECRET)

create-test-db:
	PG_DATABASE_URL=postgres://hasura:password@db:5432/test docker-compose up -d hasura
