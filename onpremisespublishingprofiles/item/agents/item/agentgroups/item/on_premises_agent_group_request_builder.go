package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i367fc7ef9476d66e888e14fa6e12431f82d47aa2a5ba37259681d082c71d1533 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agents/item/agentgroups/item/publishedresources"
    ia1f378fe0766dfc5dfea6349a05cfa28f3e62a895660c36f63e44f99148d8843 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agents/item/agentgroups/item/agents"
    i7e840fa7b37fef7aff16c31565cd806fb2481b9e74e26cb7fc98957b305799d2 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agents/item/agentgroups/item/agents/item"
    ifae56f8c6511fee3e2bb5415d36058ab3f529b5629909455a1b4b59c293a0479 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agents/item/agentgroups/item/publishedresources/item"
)

// OnPremisesAgentGroupRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\agents\{onPremisesAgent-id}\agentGroups\{onPremisesAgentGroup-id}
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
// OnPremisesAgentGroupRequestBuilderGetQueryParameters list of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
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
func (m *OnPremisesAgentGroupRequestBuilder) Agents()(*ia1f378fe0766dfc5dfea6349a05cfa28f3e62a895660c36f63e44f99148d8843.AgentsRequestBuilder) {
    return ia1f378fe0766dfc5dfea6349a05cfa28f3e62a895660c36f63e44f99148d8843.NewAgentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.agents.item.agentGroups.item.agents.item collection
func (m *OnPremisesAgentGroupRequestBuilder) AgentsById(id string)(*i7e840fa7b37fef7aff16c31565cd806fb2481b9e74e26cb7fc98957b305799d2.OnPremisesAgentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onPremisesAgent_id1"] = id
    }
    return i7e840fa7b37fef7aff16c31565cd806fb2481b9e74e26cb7fc98957b305799d2.NewOnPremisesAgentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewOnPremisesAgentGroupRequestBuilderInternal instantiates a new OnPremisesAgentGroupRequestBuilder and sets the default values.
func NewOnPremisesAgentGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnPremisesAgentGroupRequestBuilder) {
    m := &OnPremisesAgentGroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile_id}/agents/{onPremisesAgent_id}/agentGroups/{onPremisesAgentGroup_id}{?select,expand}";
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
// CreateDeleteRequestInformation list of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
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
// CreateGetRequestInformation list of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
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
// CreatePatchRequestInformation list of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
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
// Delete list of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
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
// Get list of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
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
// Patch list of onPremisesAgentGroups that an onPremisesAgent is assigned to. Read-only. Nullable.
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
func (m *OnPremisesAgentGroupRequestBuilder) PublishedResources()(*i367fc7ef9476d66e888e14fa6e12431f82d47aa2a5ba37259681d082c71d1533.PublishedResourcesRequestBuilder) {
    return i367fc7ef9476d66e888e14fa6e12431f82d47aa2a5ba37259681d082c71d1533.NewPublishedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PublishedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.agents.item.agentGroups.item.publishedResources.item collection
func (m *OnPremisesAgentGroupRequestBuilder) PublishedResourcesById(id string)(*ifae56f8c6511fee3e2bb5415d36058ab3f529b5629909455a1b4b59c293a0479.PublishedResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["publishedResource_id"] = id
    }
    return ifae56f8c6511fee3e2bb5415d36058ab3f529b5629909455a1b4b59c293a0479.NewPublishedResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
