.PHONY: swagger-ui
swagger-ui:
	git clone \
		-c advice.detachedHead=false \
		--depth 1 --branch v5.30.3 \
		https://github.com/swagger-api/swagger-ui /tmp/swagger-ui
	perl -pi -e 's/https:\/\/petstore.swagger.io\/v2//g' /tmp/swagger-ui/dist/swagger-initializer.js
	ls pkg/* | grep -xv -e "pkg/embed.go" -e "pkg/embed_test.go" | xargs rm
	cp /tmp/swagger-ui/dist/* ./pkg
	rm -rf /tmp/swagger-ui
