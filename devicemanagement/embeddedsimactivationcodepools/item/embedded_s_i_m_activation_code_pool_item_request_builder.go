package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i14f99bca7a957a331b5e7cb13fd8a56e803d7b84d2e6c38be90eb5551714c94e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/embeddedsimactivationcodepools/item/assignments"
    i1a8c6326f7b74495a2d4a1da0167deac80e12606e08ca1320fb159025e628c81 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/embeddedsimactivationcodepools/item/assign"
    ic4dac275164f8d53c7fa85048cdeb7907d2f7d42ec422190e6843e16ab03bfd3 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/embeddedsimactivationcodepools/item/devicestates"
    i4a0e4526cd89dcbf6b464d731ad1079beb451ea477a00769d03119affe5e8e45 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/embeddedsimactivationcodepools/item/assignments/item"
    if044a617b4d30ce4a26c463e72b94403997b338b8aeafd10eb8d535111e4cc01 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/embeddedsimactivationcodepools/item/devicestates/item"
)

// EmbeddedSIMActivationCodePoolItemRequestBuilder provides operations to manage the embeddedSIMActivationCodePools property of the microsoft.graph.deviceManagement entity.
type EmbeddedSIMActivationCodePoolItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EmbeddedSIMActivationCodePoolItemRequestBuilderDeleteOptions options for Delete
type EmbeddedSIMActivationCodePoolItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EmbeddedSIMActivationCodePoolItemRequestBuilderGetOptions options for Get
type EmbeddedSIMActivationCodePoolItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EmbeddedSIMActivationCodePoolItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EmbeddedSIMActivationCodePoolItemRequestBuilderGetQueryParameters the embedded SIM activation code pools created by this account.
type EmbeddedSIMActivationCodePoolItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EmbeddedSIMActivationCodePoolItemRequestBuilderPatchOptions options for Patch
type EmbeddedSIMActivationCodePoolItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EmbeddedSIMActivationCodePoolable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) Assign()(*i1a8c6326f7b74495a2d4a1da0167deac80e12606e08ca1320fb159025e628c81.AssignRequestBuilder) {
    return i1a8c6326f7b74495a2d4a1da0167deac80e12606e08ca1320fb159025e628c81.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) Assignments()(*i14f99bca7a957a331b5e7cb13fd8a56e803d7b84d2e6c38be90eb5551714c94e.AssignmentsRequestBuilder) {
    return i14f99bca7a957a331b5e7cb13fd8a56e803d7b84d2e6c38be90eb5551714c94e.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.embeddedSIMActivationCodePools.item.assignments.item collection
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) AssignmentsById(id string)(*i4a0e4526cd89dcbf6b464d731ad1079beb451ea477a00769d03119affe5e8e45.EmbeddedSIMActivationCodePoolAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["embeddedSIMActivationCodePoolAssignment_id"] = id
    }
    return i4a0e4526cd89dcbf6b464d731ad1079beb451ea477a00769d03119affe5e8e45.NewEmbeddedSIMActivationCodePoolAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEmbeddedSIMActivationCodePoolItemRequestBuilderInternal instantiates a new EmbeddedSIMActivationCodePoolItemRequestBuilder and sets the default values.
func NewEmbeddedSIMActivationCodePoolItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EmbeddedSIMActivationCodePoolItemRequestBuilder) {
    m := &EmbeddedSIMActivationCodePoolItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/embeddedSIMActivationCodePools/{embeddedSIMActivationCodePool_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEmbeddedSIMActivationCodePoolItemRequestBuilder instantiates a new EmbeddedSIMActivationCodePoolItemRequestBuilder and sets the default values.
func NewEmbeddedSIMActivationCodePoolItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EmbeddedSIMActivationCodePoolItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEmbeddedSIMActivationCodePoolItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property embeddedSIMActivationCodePools for deviceManagement
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) CreateDeleteRequestInformation(options *EmbeddedSIMActivationCodePoolItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the embedded SIM activation code pools created by this account.
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) CreateGetRequestInformation(options *EmbeddedSIMActivationCodePoolItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property embeddedSIMActivationCodePools in deviceManagement
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) CreatePatchRequestInformation(options *EmbeddedSIMActivationCodePoolItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property embeddedSIMActivationCodePools for deviceManagement
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) Delete(options *EmbeddedSIMActivationCodePoolItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) DeviceStates()(*ic4dac275164f8d53c7fa85048cdeb7907d2f7d42ec422190e6843e16ab03bfd3.DeviceStatesRequestBuilder) {
    return ic4dac275164f8d53c7fa85048cdeb7907d2f7d42ec422190e6843e16ab03bfd3.NewDeviceStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.embeddedSIMActivationCodePools.item.deviceStates.item collection
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) DeviceStatesById(id string)(*if044a617b4d30ce4a26c463e72b94403997b338b8aeafd10eb8d535111e4cc01.EmbeddedSIMDeviceStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["embeddedSIMDeviceState_id"] = id
    }
    return if044a617b4d30ce4a26c463e72b94403997b338b8aeafd10eb8d535111e4cc01.NewEmbeddedSIMDeviceStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the embedded SIM activation code pools created by this account.
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) Get(options *EmbeddedSIMActivationCodePoolItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EmbeddedSIMActivationCodePoolable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateEmbeddedSIMActivationCodePoolFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.EmbeddedSIMActivationCodePoolable), nil
}
// Patch update the navigation property embeddedSIMActivationCodePools in deviceManagement
func (m *EmbeddedSIMActivationCodePoolItemRequestBuilder) Patch(options *EmbeddedSIMActivationCodePoolItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
