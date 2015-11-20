crawl:
	bin/crawler
	
build:
	go build -o bin/crawler crawl_products.go
	
clean:
	rm json/*
	
all:
	ls json/*.json | head -n 1 | xargs cat | jq --raw-output '.'
	
process:
	ls json/*.json | xargs cat | jq --raw-output '. | ."consolidated-categories" | .categories[] \
	 | .products[] | [.PRODUCT_ID, .SHORT_DESC, .defaultSku.HEX_VALUE_STRING, .defaultSku.formattedPrice, .defaultSku.PRODUCT_SIZE, .url ] | @csv'
	
.PHONY: crawl build clean