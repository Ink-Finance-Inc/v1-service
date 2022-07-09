for item in $(ls ./abi/*.json);do
    fileName=${item##*/}
    contractName=${fileName%%.*}
    echo "generate $contractName"
    mkdir -p "./chain/$contractName"
    abigen --abi "$item" --type "$contractName" --pkg "$contractName" --out "./chain/$contractName/$contractName.go" --lang go
done;

