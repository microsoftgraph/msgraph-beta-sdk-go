package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i1786fc8f851599ebda69e9280c6b1118c2d2be8b0a0440447bbca330fa027c3b "github.com/microsoftgraph/msgraph-beta-sdk-go/contracts/item/restore"
    i2552b8bdbc7e213a9a6eae17e0a7d61fb51f1e4c5000eabd09038e0e3bb9acaa "github.com/microsoftgraph/msgraph-beta-sdk-go/contracts/item/getmemberobjects"
    i40801acc1059e0e77fbe039c879fa1cf08529d5f64814562f3ec79dde1c5a4fc "github.com/microsoftgraph/msgraph-beta-sdk-go/contracts/item/checkmembergroups"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ib7db9dde30bf0fa242f3a4321173cb16756ab77e1673add802f16c530ffa67b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/contracts/item/getmembergroups"
    idde5435a6a0561007c340ade46139ae9f81c57f894f4e9aae890f86aeecbef7d "github.com/microsoftgraph/msgraph-beta-sdk-go/contracts/item/checkmemberobjects"
)

// ContractItemRequestBuilder provides operations to manage the collection of contract entities.
type ContractItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContractItemRequestBuilderDeleteOptions options for Delete
type ContractItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContractItemRequestBuilderGetOptions options for Get
type ContractItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContractItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContractItemRequestBuilderGetQueryParameters get entity from contracts by key
type ContractItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ContractItemRequestBuilderPatchOptions options for Patch
type ContractItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Contractable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ContractItemRequestBuilder) CheckMemberGroups()(*i40801acc1059e0e77fbe039c879fa1cf08529d5f64814562f3ec79dde1c5a4fc.CheckMemberGroupsRequestBuilder) {
    return i40801acc1059e0e77fbe039c879fa1cf08529d5f64814562f3ec79dde1c5a4fc.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContractItemRequestBuilder) CheckMemberObjects()(*idde5435a6a0561007c340ade46139ae9f81c57f894f4e9aae890f86aeecbef7d.CheckMemberObjectsRequestBuilder) {
    return idde5435a6a0561007c340ade46139ae9f81c57f894f4e9aae890f86aeecbef7d.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewContractItemRequestBuilderInternal instantiates a new ContractItemRequestBuilder and sets the default values.
func NewContractItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContractItemRequestBuilder) {
    m := &ContractItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contracts/{contract_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContractItemRequestBuilder instantiates a new ContractItemRequestBuilder and sets the default values.
func NewContractItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContractItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContractItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from contracts
func (m *ContractItemRequestBuilder) CreateDeleteRequestInformation(options *ContractItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from contracts by key
func (m *ContractItemRequestBuilder) CreateGetRequestInformation(options *ContractItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in contracts
func (m *ContractItemRequestBuilder) CreatePatchRequestInformation(options *ContractItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from contracts
func (m *ContractItemRequestBuilder) Delete(options *ContractItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get entity from contracts by key
func (m *ContractItemRequestBuilder) Get(options *ContractItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Contractable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateContractFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Contractable), nil
}
func (m *ContractItemRequestBuilder) GetMemberGroups()(*ib7db9dde30bf0fa242f3a4321173cb16756ab77e1673add802f16c530ffa67b4.GetMemberGroupsRequestBuilder) {
    return ib7db9dde30bf0fa242f3a4321173cb16756ab77e1673add802f16c530ffa67b4.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContractItemRequestBuilder) GetMemberObjects()(*i2552b8bdbc7e213a9a6eae17e0a7d61fb51f1e4c5000eabd09038e0e3bb9acaa.GetMemberObjectsRequestBuilder) {
    return i2552b8bdbc7e213a9a6eae17e0a7d61fb51f1e4c5000eabd09038e0e3bb9acaa.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in contracts
func (m *ContractItemRequestBuilder) Patch(options *ContractItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContractItemRequestBuilder) Restore()(*i1786fc8f851599ebda69e9280c6b1118c2d2be8b0a0440447bbca330fa027c3b.RestoreRequestBuilder) {
    return i1786fc8f851599ebda69e9280c6b1118c2d2be8b0a0440447bbca330fa027c3b.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
