build-server:
	cd server && go build && mv pocket-trader-backend ../

run-server:
	./pocket-trader-backend

clean:
	rm pocket-trader-backend


.PHONY: build-server run-server clean