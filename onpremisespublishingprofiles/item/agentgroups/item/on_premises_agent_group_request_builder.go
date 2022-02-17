package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i19ebc08305d8024aeb70498557d19e50b92342834c2372eba852ff8b2c7bc38c "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agentgroups/item/agents"
    id42635c350598b69768841bef34ccba8e937a2ab99cff155b42e58eb7f4991e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agentgroups/item/publishedresources"
    i3a56cac0ec805de8b5e133f140b642c0d2dc07a3691ad68c64879efefe4b9140 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agentgroups/item/publishedresources/item"
    id6036572ca04fed5badd3bec64b71a79379dcd4ed352d04a4ca26efcc8fa4ed6 "github.com/microsoftgraph/msgraph-beta-sdk-go/onpremisespublishingprofiles/item/agentgroups/item/agents/item"
)

// OnPremisesAgentGroupRequestBuilder builds and executes requests for operations under \onPremisesPublishingProfiles\{onPremisesPublishingProfile-id}\agentGroups\{onPremisesAgentGroup-id}
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
// OnPremisesAgentGroupRequestBuilderGetQueryParameters list of existing onPremisesAgentGroup objects. Read-only. Nullable.
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
func (m *OnPremisesAgentGroupRequestBuilder) Agents()(*i19ebc08305d8024aeb70498557d19e50b92342834c2372eba852ff8b2c7bc38c.AgentsRequestBuilder) {
    return i19ebc08305d8024aeb70498557d19e50b92342834c2372eba852ff8b2c7bc38c.NewAgentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AgentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.agentGroups.item.agents.item collection
func (m *OnPremisesAgentGroupRequestBuilder) AgentsById(id string)(*id6036572ca04fed5badd3bec64b71a79379dcd4ed352d04a4ca26efcc8fa4ed6.OnPremisesAgentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onPremisesAgent_id"] = id
    }
    return id6036572ca04fed5badd3bec64b71a79379dcd4ed352d04a4ca26efcc8fa4ed6.NewOnPremisesAgentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewOnPremisesAgentGroupRequestBuilderInternal instantiates a new OnPremisesAgentGroupRequestBuilder and sets the default values.
func NewOnPremisesAgentGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*OnPremisesAgentGroupRequestBuilder) {
    m := &OnPremisesAgentGroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/onPremisesPublishingProfiles/{onPremisesPublishingProfile_id}/agentGroups/{onPremisesAgentGroup_id}{?select,expand}";
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
// CreateDeleteRequestInformation list of existing onPremisesAgentGroup objects. Read-only. Nullable.
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
// CreateGetRequestInformation list of existing onPremisesAgentGroup objects. Read-only. Nullable.
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
// CreatePatchRequestInformation list of existing onPremisesAgentGroup objects. Read-only. Nullable.
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
// Delete list of existing onPremisesAgentGroup objects. Read-only. Nullable.
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
// Get list of existing onPremisesAgentGroup objects. Read-only. Nullable.
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
// Patch list of existing onPremisesAgentGroup objects. Read-only. Nullable.
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
func (m *OnPremisesAgentGroupRequestBuilder) PublishedResources()(*id42635c350598b69768841bef34ccba8e937a2ab99cff155b42e58eb7f4991e4.PublishedResourcesRequestBuilder) {
    return id42635c350598b69768841bef34ccba8e937a2ab99cff155b42e58eb7f4991e4.NewPublishedResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PublishedResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.onPremisesPublishingProfiles.item.agentGroups.item.publishedResources.item collection
func (m *OnPremisesAgentGroupRequestBuilder) PublishedResourcesById(id string)(*i3a56cac0ec805de8b5e133f140b642c0d2dc07a3691ad68c64879efefe4b9140.PublishedResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["publishedResource_id"] = id
    }
    return i3a56cac0ec805de8b5e133f140b642c0d2dc07a3691ad68c64879efefe4b9140.NewPublishedResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
