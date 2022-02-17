package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    icffa63881025e30e94ee621d689ad739df1d6c5f54801abebaf0e6571205dd01 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/publishedresources/item/agentgroups/item/publishedresources"
    ie7fe4ccfff052401443978ac091cbae88e1cf14e4ff2071c64d21c0152352ec0 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/publishedresources/item/agentgroups/item/agents"
    i2b3dc5b3beae25c25ad041275e164179676fcb837611ad79911642fba5228ebe "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/publishedresources/item/agentgroups/item/publishedresources/item"
    i5c382280fd562ccfaf5ac4114acc2547213ff1b2a4b04ae1dedc167accbe54b3 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/publishedresources/item/agentgroups/item/agents/item"
)

// OnPremisesAgentGroupRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\publishedResources\{publishedResource-id}\agentGroups\{onPremisesAgentGroup-id}
type OnPremisesAgentGroupRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// OnPremisesAgentGroupRequestBuilderDeleteOptions options for Delete
type OnPremisesAgentGroupRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnPremisesAgentGroupRequestBuilderGetOptions options for Get
type OnPremisesAgentGroupRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *OnPremisesAgentGroupRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// OnPremisesAgentGroupRequestBuilderGetQueryParameters list of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
type OnPremisesAgentGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// OnPremisesAgentGroupRequestBuilderPatchOptions options for Patch
type OnPremisesAgentGroupRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnPremisesAgentGroup;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *OnPremisesAgentGroupRequestBuilder) Agents()(*ie7fe4ccfff052401443978ac091cbae88e1cf14e4ff2071c64d21c0152352ec0.AgentsRequestBuilder) {
    return ie7fe4ccfff052401443978ac091cbae88e1cf14e4ff2071c64d21c0152352ec0.NewAgentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.publishedResources.item.agentGroups.item.agents.item collection
func (m *OnPremisesAgentGroupRequestBuilder) AgentsById(id string)(*i5c382280fd562ccfaf5ac4114acc2547213ff1b2a4b04ae1dedc167accbe54b3.OnPremisesAgentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onPremisesAgent_id"] = id
    }
    return i5c382280fd562ccfaf5ac4114acc2547213ff1b2a4b04ae1dedc167accbe54b3.NewOnPremisesAgentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewOnPremisesAgentGroupRequestBuilderInternal instantiates a new OnPremisesAgentGroupRequestBuilder and sets the default values.
func NewOnPremisesAgentGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnPremisesAgentGroupRequestBuilder) {
    m := &OnPremisesAgentGroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile_id}/publishedResources/{publishedResource_id}/agentGroups/{onPremisesAgentGroup_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOnPremisesAgentGroupRequestBuilder instantiates a new OnPremisesAgentGroupRequestBuilder and sets the default values.
func NewOnPremisesAgentGroupRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnPremisesAgentGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOnPremisesAgentGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation list of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
func (m *OnPremisesAgentGroupRequestBuilder) CreateDeleteRequestInformation(options *OnPremisesAgentGroupRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation list of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
func (m *OnPremisesAgentGroupRequestBuilder) CreateGetRequestInformation(options *OnPremisesAgentGroupRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation list of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
func (m *OnPremisesAgentGroupRequestBuilder) CreatePatchRequestInformation(options *OnPremisesAgentGroupRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete list of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
func (m *OnPremisesAgentGroupRequestBuilder) Delete(options *OnPremisesAgentGroupRequestBuilderDeleteOptions)(error) {
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
// Get list of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
func (m *OnPremisesAgentGroupRequestBuilder) Get(options *OnPremisesAgentGroupRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnPremisesAgentGroup, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewOnPremisesAgentGroup() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.OnPremisesAgentGroup), nil
}
// Patch list of onPremisesAgentGroups that a publishedResource is assigned to. Read-only. Nullable.
func (m *OnPremisesAgentGroupRequestBuilder) Patch(options *OnPremisesAgentGroupRequestBuilderPatchOptions)(error) {
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
func (m *OnPremisesAgentGroupRequestBuilder) PublishedResources()(*icffa63881025e30e94ee621d689ad739df1d6c5f54801abebaf0e6571205dd01.PublishedResourcesRequestBuilder) {
    return icffa63881025e30e94ee621d689ad739df1d6c5f54801abebaf0e6571205dd01.NewPublishedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PublishedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.publishedResources.item.agentGroups.item.publishedResources.item collection
func (m *OnPremisesAgentGroupRequestBuilder) PublishedResourcesById(id string)(*i2b3dc5b3beae25c25ad041275e164179676fcb837611ad79911642fba5228ebe.PublishedResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["publishedResource_id1"] = id
    }
    return i2b3dc5b3beae25c25ad041275e164179676fcb837611ad79911642fba5228ebe.NewPublishedResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
