crawl:
	bin/crawler
	
build:
	go build -o bin/crawler crawl_products.go
	
	
.PHONY: crawl build 