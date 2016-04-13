package ppsremote

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/flow"
	"github.com/TIBCOSoftware/flogo-lib/util"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("flowprovider")

// RemoteFlowProvider is an implementation of FlowProvider service
// that can access flowes via URI
type RemoteFlowProvider struct {
	//todo: switch to LRU cache
	flowCache   map[string]*flow.Definition
	embeddedMgr *util.EmbeddedFlowManager
}

// NewRemoteFlowProvider creates a RemoteFlowProvider
func NewRemoteFlowProvider() *RemoteFlowProvider {

	var service RemoteFlowProvider
	service.flowCache = make(map[string]*flow.Definition)

	return &service
}

// Start implements util.Managed.Start()
func (pps *RemoteFlowProvider) Start() {
	// no-op
}

// Stop implements util.Managed.Stop()
func (pps *RemoteFlowProvider) Stop() {
	// no-op
}

// Init implements services.FlowProviderService.Init()
func (pps *RemoteFlowProvider) Init(settings map[string]string, embeddedFlowMgr *util.EmbeddedFlowManager) {
	pps.embeddedMgr = embeddedFlowMgr
}

// GetFlow implements flow.Provider.GetFlow
func (pps *RemoteFlowProvider) GetFlow(flowURI string) *flow.Definition {

	// todo turn pps.flowCache to real cache
	if flow, ok := pps.flowCache[flowURI]; ok {
		log.Debugf("Accessing cached Flow: %s\n")
		return flow
	}

	log.Debugf("Get Flow: %s\n", flowURI)

	var flowJSON []byte

	if strings.Index(flowURI, "local://") == 0 {

		log.Debugf("Loading Embedded Flow: %s\n", flowURI)
		flowJSON = pps.embeddedMgr.GetEmbeddedFlowJSON(flowURI[8:])

	} else {

		req, err := http.NewRequest("GET", flowURI, nil)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		log.Debug("response Status:", resp.Status)

		if resp.StatusCode >= 300 {
			//not found
			return nil
		}

		body, _ := ioutil.ReadAll(resp.Body)
		flowJSON = body
	}

	if flowJSON != nil {
		var defRep flow.DefinitionRep
		json.Unmarshal(flowJSON, &defRep)

		def := flow.NewDefinition(&defRep)

		pps.flowCache[flowURI] = def

		return def
	}

	log.Debugf("Flow Not Found: %s\n", flowURI)

	return nil
}
