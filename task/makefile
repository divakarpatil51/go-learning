build-app:
	go build task.go 

list: build-app
	./task list $(ARGS)

add: build-app
	./task add $(ARGS)

delete: build-app
	./task delete $(ARGS)

update-status: build-app
	./task update-status $(ARGS)
