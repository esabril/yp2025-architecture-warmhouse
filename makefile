swagger-generate:
	$(RM) $(PWD)/docs/api/dist/api.yaml && \
	docker run --rm \
	-u ${shell id -u}:$(shell id -g) \
	-v $(PWD)/docs/api:/code \
 	-v /etc/passwd:/etc/passwd \
 	--platform=linux/amd64 \
 	--interactive \
 	jeanberu/swagger-cli swagger-cli bundle code/api.yaml --outfile code/dist/api.yaml --type=yaml

diagrams-generate:
	  $(RM) diagrams/*/*.png && plantuml diagrams/*/*.puml