package service

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	cid "github.com/ipfs/go-cid"
	config "github.com/ipfs/go-ipfs-config"
	files "github.com/ipfs/go-ipfs-files"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	"github.com/ipfs/go-ipfs/core/node/libp2p"
	"github.com/ipfs/go-ipfs/plugin/loader"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	icore "github.com/ipfs/interface-go-ipfs-core"
	icorePath "github.com/ipfs/interface-go-ipfs-core/path"
)

type ipfsService struct {
	ctx    context.Context
	cancel context.CancelFunc
	ipfs   icore.CoreAPI
}

type IPFSService interface {
	UploadFile(string) (string, error)
	DownloadFile(hash string) (string, error)
	Destroy()
}

func NewIPFSService() IPFSService {
	ctx, cancel := context.WithCancel(context.Background())
	ipfs, err := spawnEphemeral(ctx)
	if err != nil {
		panic(fmt.Errorf("failed to spawn ephemeral node: %s", err))
	}

	return &ipfsService{
		ctx:    ctx,
		cancel: cancel,
		ipfs:   ipfs,
	}
}

func (service *ipfsService) UploadFile(path string) (string, error) {
	someFile, err := getUnixfsNode(path)
	if err != nil {
		return "", err
	}

	cidFile, err := service.ipfs.Unixfs().Add(service.ctx, someFile)
	if err != nil {
		return "", err
	}
	// multibase.EncoderByName()
	// cidFile.Cid().Encode()

	return strings.Split(cidFile.String(), "/")[2], nil
}

func (service *ipfsService) DownloadFile(hash string) (string, error) {
	c, err := cid.Decode(hash)
	if err != nil {
		return "", fmt.Errorf("failed to download: %v", err)
	}
	cidFile := icorePath.IpfsPath(c)
	// outputBasePath, err := ioutil.TempDir("output", "file_")
	// if err != nil {
	// 	return fmt.Errorf("could not create output dir : %v", err)
	// }
	// // outputBasePath := "/output/"
	// outputPathFile := outputBasePath + strings.Split(cidFile.String(), "/")[2]
	// rootNodeFile, err := service.ipfs.Unixfs().Get(service.ctx, cidFile)
	// if err != nil {
	// 	return fmt.Errorf("failed to download: %v", err)
	// }
	// err = files.WriteTo(rootNodeFile, outputPathFile)
	// if err != nil {
	// 	return fmt.Errorf("Could not write out the fetched : %v", err)
	// }
	f, err := ioutil.TempFile("output", "file_")
	defer f.Close()
	if err != nil {
		return "", fmt.Errorf("could not create output file : %v", err)
	}
	rootNodeFile, err := service.ipfs.Unixfs().Get(service.ctx, cidFile)
	if err != nil {
		return "", fmt.Errorf("failed to download: %v", err)
	}
	switch nd := rootNodeFile.(type) {
	case files.File:
		_, err := io.Copy(f, nd)
		if err != nil {
			return "", fmt.Errorf("failed to copy: %v", err)
		}
	default:
		return "", fmt.Errorf("file type %T at %q is not supported", nd, f.Name())
	}

	return f.Name(), nil
}

func (service *ipfsService) Destroy() {
	service.cancel()
}

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
