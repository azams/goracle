solc ../contracts/Price.sol --abi | tail -n1 > contract.abi
solc --abi --bin ../contracts/Price.sol -o ../build/ --overwrite
abigen --pkg=PricePkg --out=../PricePkg/PricePkg.go --abi=../build/Price.abi
