package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i015798be98f5354e44d25cd3fc8c761a04c2de377021e2334f9f574c0b9d9d50 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/connectorgroups"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4c54046d8ca87c2d4d86246e027b33b58926116766af2be55ea00fe765476f06 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agents"
    i4f4a2aad6de6e818a6b03ab06593dcd209fd37f696b126badf1ccf2cbb487d98 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agentgroups"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i86c4967a04a253d001f5fd0f5c1dd90bb7ff7141cd43dab8638d3f8ddaff0d58 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/connectors"
    if08fda1a25bbc0d180dcba2e00721696a53ac148a9d4209b397df9ba1382eb6f "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/publishedresources"
    i04768df19176575e14a7cad44b0f47526794459b68db53ab71668e3192948132 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/publishedresources/item"
    i4e613208e03a8604c73cff8fea0c650341897927012dab6031821bf1b1ae2305 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/connectorgroups/item"
    i50192dbecdd2d205018f1c9f0d43e5faa1231c597c1e3d9879fb3bfcceb8f7bf "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agents/item"
    i7be5eb505c9819935bb7ecd0914ce542aff26ebb696b731e5bb2ef462b429317 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agentgroups/item"
    ia6c25f4c749e8beb990221d7c4363a8d4383190388b2db9197c94a520cdfffde "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/connectors/item"
)

// OnPremisesPublishingProfileRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}
type OnPremisesPublishingProfileRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnPremisesPublishingProfileRequestBuilderDeleteOptions options for Delete
type OnPremisesPublishingProfileRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnPremisesPublishingProfileRequestBuilderGetOptions options for Get
type OnPremisesPublishingProfileRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnPremisesPublishingProfileRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnPremisesPublishingProfileRequestBuilderGetQueryParameters get entity from onPremisesPublishingProfiles by key
type OnPremisesPublishingProfileRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnPremisesPublishingProfileRequestBuilderPatchOptions options for Patch
type OnPremisesPublishingProfileRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnPremisesPublishingProfile;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *OnPremisesPublishingProfileRequestBuilder) AgentGroups()(*i4f4a2aad6de6e818a6b03ab06593dcd209fd37f696b126badf1ccf2cbb487d98.AgentGroupsRequestBuilder) {
    return i4f4a2aad6de6e818a6b03ab06593dcd209fd37f696b126badf1ccf2cbb487d98.NewAgentGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgentGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.agentGroups.item collection
func (m *OnPremisesPublishingProfileRequestBuilder) AgentGroupsById(id string)(*i7be5eb505c9819935bb7ecd0914ce542aff26ebb696b731e5bb2ef462b429317.OnPremisesAgentGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onPremisesAgentGroup_id"] = id
    }
    return i7be5eb505c9819935bb7ecd0914ce542aff26ebb696b731e5bb2ef462b429317.NewOnPremisesAgentGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnPremisesPublishingProfileRequestBuilder) Agents()(*i4c54046d8ca87c2d4d86246e027b33b58926116766af2be55ea00fe765476f06.AgentsRequestBuilder) {
    return i4c54046d8ca87c2d4d86246e027b33b58926116766af2be55ea00fe765476f06.NewAgentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.agents.item collection
func (m *OnPremisesPublishingProfileRequestBuilder) AgentsById(id string)(*i50192dbecdd2d205018f1c9f0d43e5faa1231c597c1e3d9879fb3bfcceb8f7bf.OnPremisesAgentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onPremisesAgent_id"] = id
    }
    return i50192dbecdd2d205018f1c9f0d43e5faa1231c597c1e3d9879fb3bfcceb8f7bf.NewOnPremisesAgentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnPremisesPublishingProfileRequestBuilder) ConnectorGroups()(*i015798be98f5354e44d25cd3fc8c761a04c2de377021e2334f9f574c0b9d9d50.ConnectorGroupsRequestBuilder) {
    return i015798be98f5354e44d25cd3fc8c761a04c2de377021e2334f9f574c0b9d9d50.NewConnectorGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectorGroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.connectorGroups.item collection
func (m *OnPremisesPublishingProfileRequestBuilder) ConnectorGroupsById(id string)(*i4e613208e03a8604c73cff8fea0c650341897927012dab6031821bf1b1ae2305.ConnectorGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connectorGroup_id"] = id
    }
    return i4e613208e03a8604c73cff8fea0c650341897927012dab6031821bf1b1ae2305.NewConnectorGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *OnPremisesPublishingProfileRequestBuilder) Connectors()(*i86c4967a04a253d001f5fd0f5c1dd90bb7ff7141cd43dab8638d3f8ddaff0d58.ConnectorsRequestBuilder) {
    return i86c4967a04a253d001f5fd0f5c1dd90bb7ff7141cd43dab8638d3f8ddaff0d58.NewConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectorsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.connectors.item collection
func (m *OnPremisesPublishingProfileRequestBuilder) ConnectorsById(id string)(*ia6c25f4c749e8beb990221d7c4363a8d4383190388b2db9197c94a520cdfffde.ConnectorRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connector_id"] = id
    }
    return ia6c25f4c749e8beb990221d7c4363a8d4383190388b2db9197c94a520cdfffde.NewConnectorRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewOnPremisesPublishingProfileRequestBuilderInternal instantiates a new OnPremisesPublishingProfileRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnPremisesPublishingProfileRequestBuilder) {
    m := &OnPremisesPublishingProfileRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnPremisesPublishingProfileRequestBuilder instantiates a new OnPremisesPublishingProfileRequestBuilder and sets the default values.
func NewOnPremisesPublishingProfileRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnPremisesPublishingProfileRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnPremisesPublishingProfileRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfileRequestBuilder) CreateDeleteRequestInformation(options *OnPremisesPublishingProfileRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get entity from onPremisesPublishingProfiles by key
func (m *OnPremisesPublishingProfileRequestBuilder) CreateGetRequestInformation(options *OnPremisesPublishingProfileRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update entity in onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfileRequestBuilder) CreatePatchRequestInformation(options *OnPremisesPublishingProfileRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete entity from onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfileRequestBuilder) Delete(options *OnPremisesPublishingProfileRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get get entity from onPremisesPublishingProfiles by key
func (m *OnPremisesPublishingProfileRequestBuilder) Get(options *OnPremisesPublishingProfileRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnPremisesPublishingProfile, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOnPremisesPublishingProfile() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnPremisesPublishingProfile), nil
}
// Patch update entity in onPremisesPublishingProfiles
func (m *OnPremisesPublishingProfileRequestBuilder) Patch(options *OnPremisesPublishingProfileRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *OnPremisesPublishingProfileRequestBuilder) PublishedResources()(*if08fda1a25bbc0d180dcba2e00721696a53ac148a9d4209b397df9ba1382eb6f.PublishedResourcesRequestBuilder) {
    return if08fda1a25bbc0d180dcba2e00721696a53ac148a9d4209b397df9ba1382eb6f.NewPublishedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PublishedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.publishedResources.item collection
func (m *OnPremisesPublishingProfileRequestBuilder) PublishedResourcesById(id string)(*i04768df19176575e14a7cad44b0f47526794459b68db53ab71668e3192948132.PublishedResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["publishedResource_id"] = id
    }
    return i04768df19176575e14a7cad44b0f47526794459b68db53ab71668e3192948132.NewPublishedResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
