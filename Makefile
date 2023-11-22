.PHONY: swagger-ui
swagger-ui:
	git clone --depth 1 --branch v5.10.3 https://github.com/swagger-api/swagger-ui /tmp/swagger-ui
	perl -pi -e 's/https:\/\/petstore.swagger.io\/v2//g' /tmp/swagger-ui/dist/swagger-initializer.js
	cp /tmp/swagger-ui/dist/* ./pkg
	rm -rf /tmp/swagger-ui
