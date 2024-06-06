package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	repoMap = make(map[string]string)
)

func main() {
	// 检查是否提供了镜像名称
	if len(os.Args) < 2 {
		fmt.Println("Usage: dq <image-name>")
		return
	}

	// 获取输入的镜像名称和标签
	originalImage := os.Args[1]
	subStr := strings.Split(originalImage, "/")[0]

	var newImage = ""
	if repoMap[subStr] == "" {
		newImage = "docker.m.daocloud.io/" + originalImage
	} else {
		newImage = strings.Replace(originalImage, subStr, repoMap[subStr], 1)
	}
	fmt.Println(newImage)

	// 拉取镜像
	fmt.Printf("Pulling image from China mirror: %s\n", newImage)
	if err := execCmd("docker", "pull", newImage); err != nil {
		fmt.Printf("Failed to pull image: %v\n", err)
		return
	}

	// 重命名镜像标签
	fmt.Printf("Tagging image with original name: %s\n", originalImage)
	if err := execCmd("docker", "tag", newImage, originalImage); err != nil {
		fmt.Printf("Failed to tag image: %v\n", err)
		return
	}

	// 可选：删除中间镜像
	fmt.Printf("Removing intermediate image: %s\n", newImage)
	if err := execCmd("docker", "rmi", newImage); err != nil {
		fmt.Printf("Failed to remove intermediate image: %v\n", err)
		return
	}

	fmt.Println("Done. Image", originalImage, "is ready.")
}

func init() {
	repoMap["ghcr.io"] = "ghcr.m.daocloud.io"
	repoMap["gcr.io"] = "gcr.m.daocloud.io"
	repoMap["k8s.gcr.io"] = "k8s-gcr.m.daocloud.io"
	repoMap["registry.k8s.io"] = "k8s.m.daocloud.io"
	repoMap["quay.io"] = "quay.m.daocloud.io"
	repoMap["mcr.microsoft.com"] = "mcr.m.daocloud.io"
}

func execCmd(name string, arg ...string) error {
	var cmd *exec.Cmd

	cmd = exec.Command(name, arg...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Errorf(err.Error())
		return err
	}
	return nil
}
