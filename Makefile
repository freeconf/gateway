export YANGPATH=$(abspath ./yang)

test :
	go test .

# run on host in emulation mode
run:
	go run ./cmd/fc-gateway/main.go -startup ./cmd/fc-gateway/startup.json

# build web source into package
.PHONY: web
web : ./web/dist/index.html
./web/dist/index.html : $(wildcard web/src/*.js) web/index.html $(wildcard web/images/*.*)
	cd web; \
		npx parcel build --public-url /web/ ./index.html
	touch ./web/dist/index.html

clean:
	-rm -rf ./bartend ./web/dist docs/api.md

# REST API docs that host in github
docs : docs/api.md
docs/api.md : ./yang/fc-gateway.yang
	go run github.com/freeconf/yang/cmd/fc-yang \
		doc -module fc-gateway -title 'FreeCONF Gateway' -f md > $@

# update the local copy of the yang model files from freeconf. Only has to
# be run when freeconf dep is updated
update-yang :
	go run github.com/freeconf/yang/cmd/fc-yang get -dir ./yang
