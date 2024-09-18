.PHONY: sm-solc-all
sm-solc-all:
	@bash ./_scripts/solc-all.sh

.PHONY: sm-abigen-all
sm-abigen-all:
	@bash ./_scripts/abigen-all.sh

.PHONY: sm-clean
sm-clean:
	rm -f contracts/abi/*.abi
	rm -f contracts/*.go

.PHONY: sm-generate
sm-generate: sm-clean sm-solc-all sm-abigen-all
