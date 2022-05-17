#!/bin/sh
name="srun"
version="v0.1.0"
path="./out"

start=$(date +%s)

# Windows
os="windows"
arch=("386" "amd64" "arm64")
for a in ${arch[@]}; do
    GOOS=$os GOARCH=$a go build -o $path/$name-$os-$a-$version.exe -ldflags -w
done

# Linux
os="linux"
arch=("386" "amd64" "arm" "arm64" "mips" "mipsle" "mips64" "mips64le" "ppc64" "ppc64le" "s390x" "riscv64")
for a in ${arch[@]}; do
    GOOS=$os GOARCH=$a go build -o $path/$name-$os-$a-$version -ldflags -w
done

# macOS
os="darwin"
arch=("amd64" "arm64")
for a in ${arch[@]}; do
    GOOS=$os GOARCH=$a go build -o $path/$name-$os-$a-$version -ldflags -w
done

end=$(date +%s)
dur=$(($end - $start))
echo "Build time: $dur seconds"