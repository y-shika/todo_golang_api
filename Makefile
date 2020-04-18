LOCAL_ENV_VARS := $(shell cat .env.local)

# TODO: fixtureにもローカルなのかHerokuなのかなど種類があるはずで、いちいち環境変数を`.env.local`で変更するのも
#       スマートではないため、envファイルを増やしたり、Makefileものちに種類で分ける
.PHONY: fixture
fixture:
	$(LOCAL_ENV_VARS) go run cmd/dev/fixture/main.go
