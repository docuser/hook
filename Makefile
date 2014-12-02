default:
	docker build -t hook .
	docker run --rm -it -p 80:80 hook
