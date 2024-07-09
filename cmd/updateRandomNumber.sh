solc ../contracts/Random.sol --abi | tail -n1 > random.abi
solc --abi --bin ../contracts/Random.sol -o ../build/ --overwrite
abigen --pkg=RandomPkg --out=../RandomPkg/RandomPkg.go --abi=../build/RandomNumber.abi
