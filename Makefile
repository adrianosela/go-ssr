
run: vite
	go run main.go

vite:
	cd web && yarn install && yarn build

clean:
	rm -rf ./web/dist
	rm -rf ./web/node_modules
