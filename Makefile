run:
	rm -R go-diagrams -f
	go run cmd/api/main.go
	cd go-diagrams && dot -Tpng app.dot > app.png
