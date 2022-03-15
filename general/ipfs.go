package general

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	config "github.com/ipfs/go-ipfs-config"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	"github.com/ipfs/go-ipfs/plugin/loader"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	icore "github.com/ipfs/interface-go-ipfs-core"
)

//Part 1 - Getting an IPFS node running: To get get a node running as an ephemeral node (that will cease to exist when the run ends)

//Prepare and set up the plugins

func setupPlugins(externalPluginsPath string) error {
	// Load any external plugins if available on externalPluginsPath
	plugins, err := loader.NewPluginLoader(filepath.Join(externalPluginsPath, "plugins"))
	if err != nil {
		return fmt.Errorf("error loading plugins: %s", err)
	}

	// Load preloaded and external plugins
	if err := plugins.Initialize(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	if err := plugins.Inject(); err != nil {
		return fmt.Errorf("error initializing plugins: %s", err)
	}

	return nil
}

//Create an IPFS repo

func createTempRepo(ctx context.Context) (string, error) {
	repoPath, err := ioutil.TempDir("", "ipfs-shell")
	if err != nil {
		return "", fmt.Errorf("failed to get temp dir: %s", err)
	}

	// Create a config with default options and a 2048 bit key
	cfg, err := config.Init(ioutil.Discard, 2048)
	if err != nil {
		return "", err
	}

	// Create the repo with the config
	err = fsrepo.Init(repoPath, cfg)
	if err != nil {
		return "", fmt.Errorf("failed to init ephemeral node: %s", err)
	}

	return repoPath, nil
}

//Construct the IPFS node instance itself

func spawnEphemeral(ctx context.Context) (icore.CoreAPI, error) {
	if err := setupPlugins(""); err != nil {
		return nil, err
	}

	// Create a Temporary Repo
	repoPath, err := createTempRepo(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create temp repo: %s", err)
	}

	// Spawning an ephemeral IPFS node
	return createNode(ctx, repoPath)
}

func createNode(ctx context.Context, repoPath string) (icore.CoreAPI, error) {
	//Creates an IPFS node and returns its coreAPI
	// Open the repo
	repo, err := fsrepo.Open(repoPath)
	if err != nil {
		return nil, err
	}

	// Construct the node

	nodeOptions := &core.BuildCfg{
		Online:  true,
		Routing: libp2p.DHTOption, // This option sets the node to be a full DHT node (both fetching and storing DHT Records)
		// Routing: libp2p.DHTClientOption, // This option sets the node to be a client DHT node (only fetching records)
		Repo: repo,
	}

	node, err := core.NewNode(ctx, nodeOptions)
	if err != nil {
		return nil, err
	}

	// Attach the Core API to the constructed node
	return coreapi.NewCoreAPI(node)
}

//Part 2 - Adding a file and a directory to IPFS

//Prepare the file to be added to IPFS
func getUnixfsFile(path string) (files.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	st, err := file.Stat()
	if err != nil {
		return nil, err
	}

	f, err := files.NewReaderPathFile(path, file, st)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func getUnixfsNode(path string) (files.Node, error) {
	st, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	f, err := files.NewSerialFile(path, false, st)
	if err != nil {
		return nil, err
	}

	return f, nil
}

//Add a file to IPFS -- main()

//Part 3 - Getting the file and directory you added back

//Get the file back -- main()

//Write the file to local filesystem -- main()

// func main() {
// 	fmt.Println("-- Getting an IPFS node running -- ")

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	fmt.Println("Spawning node on a temporary repo")
// 	ipfs, err := spawnEphemeral(ctx)
// 	if err != nil {
// 		panic(fmt.Errorf("failed to spawn ephemeral node: %s", err))
// 	}

// 	fmt.Println("IPFS node is running")

// 	fmt.Println("\n-- Adding and getting back files & directories --")

// 	inputBasePath := "./ipfs/"
// 	inputPathFile := inputBasePath + "test.txt"
// 	//inputPathDirectory := inputBasePath + "test-dir"

// 	someFile, err := getUnixfsNode(inputPathFile)
// 	if err != nil {
// 		panic(fmt.Errorf("Could not get File: %s", err))
// 	}

// 	cidFile, err := ipfs.Unixfs().Add(ctx, someFile)
// 	if err != nil {
// 		panic(fmt.Errorf("Could not add File: %s", err))
// 	}

// 	fmt.Printf("Added file to IPFS with CID %s\n", cidFile.String())

// 	outputBasePath := "./ipfs/download/"
// 	outputPathFile := outputBasePath + strings.Split(cidFile.String(), "/")[2]
// 	//outputPathDirectory := outputBasePath + strings.Split(cidDirectory.String(), "/")[2]

// 	rootNodeFile, err := ipfs.Unixfs().Get(ctx, cidFile)
// 	if err != nil {
// 		panic(fmt.Errorf("Could not get file with CID: %s", err))
// 	}

// 	err = files.WriteTo(rootNodeFile, outputPathFile)
// 	if err != nil {
// 		panic(fmt.Errorf("Could not write out the fetched CID: %s", err))
// 	}

// 	fmt.Printf("Got file back from IPFS (IPFS path: %s) and wrote it to %s\n", cidFile.String(), outputPathFile)
// }
