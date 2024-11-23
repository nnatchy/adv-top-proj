.PHONY: run build down clean info logs restart ps rebuild generate-go-bindings vote docker-build docker-run

info:
	docker compose ps

run:
	docker compose up

down:
	docker compose down

build:
	docker compose build

clean:
	docker compose down --rmi all --volumes --remove-orphans

logs:
	docker compose logs -f

restart:
	docker compose restart

ps:
	docker compose ps

rebuild: clean docker-build run

# Generate Go bindings from the ABI and bytecode
generate-go-bindings:
	abigen --bin=internal/infrastructure/ethereum/solidity/Voting.bin \
	       --abi=internal/infrastructure/ethereum/solidity/Voting.abi \
	       --pkg=solidity \
	       --out=internal/infrastructure/ethereum/solidity/simple_storage.go

vote:
	abigen --bin=internal/infrastructure/ethereum/solidity/voting/Voting.bin \
	       --abi=internal/infrastructure/ethereum/solidity/voting/Voting.abi \
	       --pkg=voting \
	       --out=internal/infrastructure/ethereum/solidity/voting/voting.go

docker-build:
	docker build -t natsith/adv-backend .

docker-run:
	docker run -p 8010:8010 natsith/adv-backend
