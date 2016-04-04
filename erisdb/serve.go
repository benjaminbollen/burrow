// The erisdb package contains tendermint-specific services that goes with the
// server.
package erisdb

import (
	"bytes"
	"path"

	"github.com/eris-ltd/eris-db/Godeps/_workspace/src/github.com/tendermint/log15"
	acm "github.com/eris-ltd/eris-db/Godeps/_workspace/src/github.com/tendermint/tendermint/account"
	. "github.com/eris-ltd/eris-db/Godeps/_workspace/src/github.com/tendermint/tendermint/common"
	cfg "github.com/eris-ltd/eris-db/Godeps/_workspace/src/github.com/tendermint/tendermint/config"
	tmcfg "github.com/eris-ltd/eris-db/Godeps/_workspace/src/github.com/tendermint/tendermint/config/tendermint"
	dbm "github.com/eris-ltd/eris-db/Godeps/_workspace/src/github.com/tendermint/tendermint/db"
	"github.com/eris-ltd/eris-db/Godeps/_workspace/src/github.com/tendermint/tendermint/events"
	"github.com/eris-ltd/eris-db/Godeps/_workspace/src/github.com/tendermint/tendermint/node"
	"github.com/eris-ltd/eris-db/Godeps/_workspace/src/github.com/tendermint/tendermint/p2p"

	sm "github.com/eris-ltd/eris-db/Godeps/_workspace/src/github.com/tendermint/tendermint/state"
	stypes "github.com/eris-ltd/eris-db/Godeps/_workspace/src/github.com/tendermint/tendermint/state/types"
	"github.com/eris-ltd/eris-db/Godeps/_workspace/src/github.com/tendermint/tendermint/wire"

	ep "github.com/eris-ltd/eris-db/erisdb/pipe"
	"github.com/eris-ltd/eris-db/server"

	edbapp "github.com/eris-ltd/eris-db/tmsp"
	tmsp "github.com/tendermint/tmsp/server"
)

var log = log15.New("module", "eris/erisdb_server")
var tmConfig cfg.Config

// This function returns a properly configured ErisDb server process with a running
// tendermint node attached to it. To start listening for incoming requests, call
// 'Start()' on the process. Make sure to register any start event listeners before
// that.
func ServeErisDB(workDir string) (*server.ServeProcess, error) {
	log.Info("ErisDB Serve initializing.")
	errEns := EnsureDir(workDir)

	if errEns != nil {
		return nil, errEns
	}

	var sConf *server.ServerConfig

	sConfPath := path.Join(workDir, "server_conf.toml")
	if !FileExists(sConfPath) {
		log.Info("No server configuration, using default.")
		log.Info("Writing to: " + sConfPath)
		sConf = server.DefaultServerConfig()
		errW := server.WriteServerConfig(sConfPath, sConf)
		if errW != nil {
			panic(errW)
		}
	} else {
		var errRSC error
		sConf, errRSC = server.ReadServerConfig(sConfPath)
		if errRSC != nil {
			log.Error("Server config file error.", "error", errRSC.Error())
		}
	}

	// Get tendermint configuration
	tmConfig = tmcfg.GetConfig(workDir)
	tmConfig.Set("version", node.Version)
	cfg.ApplyConfig(tmConfig) // Notify modules of new config

	// load the priv validator
	privVal := types.LoadPrivValidator(tmConfig.GetString("priv_validator_file"))

	// possibly set the signer to use eris-keys
	signerCfg := tmConfig.GetString("signer")
	if !(signerCfg == "default" || signerCfg == "") {
		spl := strings.Split(signerCfg, ":")
		if len(spl) != 2 {
			panic(`"signer" field in config.toml should be "default" or "ERIS_KEYS_HOST:ERIS_KEYS_PORT"`)
		}

		// TODO: if a privKey is found in the privVal, warn the user that we want to import it to eris-keys and delete it

		host := signerCfg
		addr := hex.EncodeToString(privVal.Address)
		signer := NewErisSigner(host, addr)
		privVal.SetSigner(signer)
	}

	// Set the node up.
	// nodeRd := make(chan struct{})
	// nd := node.NewNode()

	// Load the application state
	// The app state used to be managed by tendermint node,
	// but is now managed by ErisDB.
	// The tendermint core only stores the blockchain (history of txs)
	stateDB := dbm.GetDB("state")
	state := sm.LoadState(stateDB)
	var genDoc *stypes.GenesisDoc
	if state == nil {
		genDoc, state = sm.MakeGenesisStateFromFile(stateDB, config.GetString("genesis_file"))
		state.Save()
		// write the gendoc to db
		buf, n, err := new(bytes.Buffer), new(int64), new(error)
		wire.WriteJSON(genDoc, buf, n, err)
		stateDB.Set(stypes.GenDocKey, buf.Bytes())
		if *err != nil {
			Exit(Fmt("Unable to write gendoc to db: %v", err))
		}
	} else {
		genDocBytes := stateDB.Get(stypes.GenDocKey)
		err := new(error)
		wire.ReadJSONPtr(&genDoc, genDocBytes, err)
		if *err != nil {
			Exit(Fmt("Unable to read gendoc from db: %v", err))
		}
	}
	// add the chainid to the global config
	config.Set("chain_id", state.ChainID)

	evsw := events.NewEventSwitch()
	app := edbapp.NewErisDBApp(state, evsw)

	// Start the tmsp listener for state update commands
	go func() {
		// TODO config
		_, err := tmsp.StartListener("tcp://0.0.0.0:46658", app)
		if err != nil {
			// TODO: play nice
			Exit(err.Error())
		}
	}()

	// Load supporting objects.
	pipe := ep.NewPipe(app)
	codec := &TCodec{}
	evtSubs := NewEventSubscriptions(pipe.Events())
	// The services.
	tmwss := NewErisDbWsService(codec, pipe)
	//
	tmjs := NewErisDbJsonService(codec, pipe, evtSubs)
	// The servers.
	jsonServer := NewJsonRpcServer(tmjs)
	restServer := NewRestServer(codec, pipe, evtSubs)
	wsServer := server.NewWebSocketServer(sConf.WebSocket.MaxWebSocketSessions, tmwss)
	// Create a server process.
	proc := server.NewServeProcess(sConf, jsonServer, restServer, wsServer)

	//stopChan := proc.StopEventChannel()
	//go startNode(nd, nodeRd, stopChan)
	//<-nodeRd
	return proc, nil
}

// Private. Create a new node.
func startNode(nd *node.Node, ready chan struct{}, shutDown <-chan struct{}) {
	laddr := tmConfig.GetString("node_laddr")
	if laddr != "" {
		l := p2p.NewDefaultListener("tcp", laddr)
		nd.AddListener(l)
	}

	nd.Start()

	// If seedNode is provided by config, dial out.

	if len(tmConfig.GetString("seeds")) > 0 {
		nd.DialSeed()
	}

	if len(tmConfig.GetString("rpc_laddr")) > 0 {
		nd.StartRPC()
	}
	ready <- struct{}{}
	// Block until everything is shut down.
	<-shutDown
	nd.Stop()
}

type ErisSigner struct {
	Host       string
	Address    string
	SessionKey string
}

func NewErisSigner(host, address string) *ErisSigner {
	if !strings.HasPrefix(host, "http://") {
		host = fmt.Sprintf("http://%s", host)
	}
	return &ErisSigner{Host: host, Address: address}
}

func (es *ErisSigner) Sign(msg []byte) acm.SignatureEd25519 {
	msgHex := hex.EncodeToString(msg)
	sig, err := core.Sign(msgHex, es.Address, es.Host)
	if err != nil {
		panic(err)
	}
	return acm.SignatureEd25519(sig)
}
