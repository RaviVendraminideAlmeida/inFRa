package actions

import (
	"cli/util"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func registryHandler(path, regName string) error {
	err := util.TouchFile(path, fmt.Sprintf("%s.reg", regName))
	return err
}

// func handleDBCreation(imgPath string, ctx *cli.Context) {

// 	containerName := ctx.String("container-name")

// 	dockerContext := context.Background()
// 	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer cli.Close()

// 	reader, err := cli.ImagePull(dockerContext, imgPath, image.PullOptions{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer reader.Close()

// 	io.Copy(os.Stdout, reader)

// 	fmt.Println(containerName)

// 	resp, err := cli.ContainerCreate(
// 		dockerContext, &container.Config{
// 			Image: "alpine",
// 			Cmd:   []string{"echo", "hello world"},
// 			Tty:   false,
// 		},
// 		nil,
// 		nil,
// 		nil,
// 		containerName,
// 	)

// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := cli.ContainerStart(dockerContext, resp.ID, container.StartOptions{}); err != nil {
// 		panic(err)
// 	}

// 	statusCh, errCh := cli.ContainerWait(dockerContext, resp.ID, container.WaitConditionNotRunning)
// 	select {
// 	case err := <-errCh:
// 		if err != nil {
// 			panic(err)
// 		}
// 	case <-statusCh:
// 	}

// 	out, err := cli.ContainerLogs(dockerContext, resp.ID, container.LogsOptions{ShowStdout: true})
// 	if err != nil {
// 		panic(err)
// 	}

// 	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
// }

func CreateAction(ctx *cli.Context) error {
	path := ctx.Args().First()

	err := os.Mkdir(path, 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	err = util.TouchFile(path, "docker-compose.yaml")
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	registryHandler(path, ctx.String("registry"))

	params := util.HandleDBParameters()

	fmt.Printf("%+v\n", params)

	return err
}
