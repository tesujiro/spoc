all:
	go vet .
	go build .

search:
	./cli search artist つのだ
	./cli search album Secret
	./cli search playlist harakami

profile:
	./cli get profile

playlist:
	./cli get playlists

test: all
	./test.sh ; echo RESULT=$$?
proxy:
	./reverse-proxy/reverse-proxy ./reverse-proxy/cache.gob &
	sleep 0.1
	curl -X GET "http://localhost:8080/load"

save-cache:
	curl -X GET "http://localhost:8080/save"


load-cache:
	curl -X GET "http://localhost:8080/load"


