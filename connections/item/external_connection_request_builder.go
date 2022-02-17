package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i16aec0d979679a78de026ea357af647676e9bd2f7c8534da4da2643e5d6f81b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/operations"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ia229bb0ea4c647efe1ef7f3f3d49c1c41e13acf3b845ba20fa1cc1c6f2730bdd "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/schema"
    ia8a9fb046e54d24538d7c6d25cc69813af40692eee89a9d2eb29dc5570b69c1d "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/quota"
    ic4ecba2a36e9429147361e0927b616e9b14c0047f0c81f4d383857826bf622b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/groups"
    ieb276fe54a3917eca6ba7222449ff2e6a5642f1c6bebeaabc0427903311a92df "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/items"
    i20ab49c75285e4cd8a2b4e4645b09f58717fe4efdaf039c7973589d395d6d3a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/groups/item"
    i38a4dbb9e776d366f238688a9031f7f382abfd310dedd2207973525e4859aebd "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/operations/item"
    ifb639436ebf63dc6cfd09347d97c3adc1ba4e054cb845ef13d79cea6a7aece70 "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/items/item"
)

// ExternalConnectionRequestBuilder builds and executes requests for operations under \connections\{externalConnection-id}
type ExternalConnectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ExternalConnectionRequestBuilderDeleteOptions options for Delete
type ExternalConnectionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ExternalConnectionRequestBuilderGetOptions options for Get
type ExternalConnectionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ExternalConnectionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ExternalConnectionRequestBuilderGetQueryParameters get entity from connections by key
type ExternalConnectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ExternalConnectionRequestBuilderPatchOptions options for Patch
type ExternalConnectionRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExternalConnection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewExternalConnectionRequestBuilderInternal instantiates a new ExternalConnectionRequestBuilder and sets the default values.
func NewExternalConnectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExternalConnectionRequestBuilder) {
    m := &ExternalConnectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/connections/{externalConnection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExternalConnectionRequestBuilder instantiates a new ExternalConnectionRequestBuilder and sets the default values.
func NewExternalConnectionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ExternalConnectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExternalConnectionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from connections
func (m *ExternalConnectionRequestBuilder) CreateDeleteRequestInformation(options *ExternalConnectionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from connections by key
func (m *ExternalConnectionRequestBuilder) CreateGetRequestInformation(options *ExternalConnectionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in connections
func (m *ExternalConnectionRequestBuilder) CreatePatchRequestInformation(options *ExternalConnectionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from connections
func (m *ExternalConnectionRequestBuilder) Delete(options *ExternalConnectionRequestBuilderDeleteOptions)(error) {
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
// Get get entity from connections by key
func (m *ExternalConnectionRequestBuilder) Get(options *ExternalConnectionRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExternalConnection, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewExternalConnection() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ExternalConnection), nil
}
func (m *ExternalConnectionRequestBuilder) Groups()(*ic4ecba2a36e9429147361e0927b616e9b14c0047f0c81f4d383857826bf622b0.GroupsRequestBuilder) {
    return ic4ecba2a36e9429147361e0927b616e9b14c0047f0c81f4d383857826bf622b0.NewGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.connections.item.groups.item collection
func (m *ExternalConnectionRequestBuilder) GroupsById(id string)(*i20ab49c75285e4cd8a2b4e4645b09f58717fe4efdaf039c7973589d395d6d3a7.ExternalGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalGroup_id"] = id
    }
    return i20ab49c75285e4cd8a2b4e4645b09f58717fe4efdaf039c7973589d395d6d3a7.NewExternalGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ExternalConnectionRequestBuilder) Items()(*ieb276fe54a3917eca6ba7222449ff2e6a5642f1c6bebeaabc0427903311a92df.ItemsRequestBuilder) {
    return ieb276fe54a3917eca6ba7222449ff2e6a5642f1c6bebeaabc0427903311a92df.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.connections.item.items.item collection
func (m *ExternalConnectionRequestBuilder) ItemsById(id string)(*ifb639436ebf63dc6cfd09347d97c3adc1ba4e054cb845ef13d79cea6a7aece70.ExternalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["externalItem_id"] = id
    }
    return ifb639436ebf63dc6cfd09347d97c3adc1ba4e054cb845ef13d79cea6a7aece70.NewExternalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ExternalConnectionRequestBuilder) Operations()(*i16aec0d979679a78de026ea357af647676e9bd2f7c8534da4da2643e5d6f81b1.OperationsRequestBuilder) {
    return i16aec0d979679a78de026ea357af647676e9bd2f7c8534da4da2643e5d6f81b1.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.connections.item.operations.item collection
func (m *ExternalConnectionRequestBuilder) OperationsById(id string)(*i38a4dbb9e776d366f238688a9031f7f382abfd310dedd2207973525e4859aebd.ConnectionOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connectionOperation_id"] = id
    }
    return i38a4dbb9e776d366f238688a9031f7f382abfd310dedd2207973525e4859aebd.NewConnectionOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in connections
func (m *ExternalConnectionRequestBuilder) Patch(options *ExternalConnectionRequestBuilderPatchOptions)(error) {
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
func (m *ExternalConnectionRequestBuilder) Quota()(*ia8a9fb046e54d24538d7c6d25cc69813af40692eee89a9d2eb29dc5570b69c1d.QuotaRequestBuilder) {
    return ia8a9fb046e54d24538d7c6d25cc69813af40692eee89a9d2eb29dc5570b69c1d.NewQuotaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ExternalConnectionRequestBuilder) Schema()(*ia229bb0ea4c647efe1ef7f3f3d49c1c41e13acf3b845ba20fa1cc1c6f2730bdd.SchemaRequestBuilder) {
    return ia229bb0ea4c647efe1ef7f3f3d49c1c41e13acf3b845ba20fa1cc1c6f2730bdd.NewSchemaRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
