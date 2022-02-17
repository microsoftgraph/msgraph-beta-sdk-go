package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i1786fc8f851599ebda69e9280c6b1118c2d2be8b0a0440447bbca330fa027c3b "github.com/microsoftgraph/msgraph-beta-sdk-go/contracts/item/restore"
    i2552b8bdbc7e213a9a6eae17e0a7d61fb51f1e4c5000eabd09038e0e3bb9acaa "github.com/microsoftgraph/msgraph-beta-sdk-go/contracts/item/getmemberobjects"
    i40801acc1059e0e77fbe039c879fa1cf08529d5f64814562f3ec79dde1c5a4fc "github.com/microsoftgraph/msgraph-beta-sdk-go/contracts/item/checkmembergroups"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ib7db9dde30bf0fa242f3a4321173cb16756ab77e1673add802f16c530ffa67b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/contracts/item/getmembergroups"
    idde5435a6a0561007c340ade46139ae9f81c57f894f4e9aae890f86aeecbef7d "github.com/microsoftgraph/msgraph-beta-sdk-go/contracts/item/checkmemberobjects"
)

// ContractRequestBuilder builds and executes requests for operations under \contracts\{contract-id}
type ContractRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContractRequestBuilderDeleteOptions options for Delete
type ContractRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContractRequestBuilderGetOptions options for Get
type ContractRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContractRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContractRequestBuilderGetQueryParameters get entity from contracts by key
type ContractRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ContractRequestBuilderPatchOptions options for Patch
type ContractRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Contract;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ContractRequestBuilder) CheckMemberGroups()(*i40801acc1059e0e77fbe039c879fa1cf08529d5f64814562f3ec79dde1c5a4fc.CheckMemberGroupsRequestBuilder) {
    return i40801acc1059e0e77fbe039c879fa1cf08529d5f64814562f3ec79dde1c5a4fc.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContractRequestBuilder) CheckMemberObjects()(*idde5435a6a0561007c340ade46139ae9f81c57f894f4e9aae890f86aeecbef7d.CheckMemberObjectsRequestBuilder) {
    return idde5435a6a0561007c340ade46139ae9f81c57f894f4e9aae890f86aeecbef7d.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewContractRequestBuilderInternal instantiates a new ContractRequestBuilder and sets the default values.
func NewContractRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContractRequestBuilder) {
    m := &ContractRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/contracts/{contract_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContractRequestBuilder instantiates a new ContractRequestBuilder and sets the default values.
func NewContractRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContractRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContractRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from contracts
func (m *ContractRequestBuilder) CreateDeleteRequestInformation(options *ContractRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ContractRequestBuilder) CreateGetRequestInformation(options *ContractRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ContractRequestBuilder) CreatePatchRequestInformation(options *ContractRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ContractRequestBuilder) Delete(options *ContractRequestBuilderDeleteOptions)(error) {
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
// Get get entity from contracts by key
func (m *ContractRequestBuilder) Get(options *ContractRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Contract, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewContract() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Contract), nil
}
func (m *ContractRequestBuilder) GetMemberGroups()(*ib7db9dde30bf0fa242f3a4321173cb16756ab77e1673add802f16c530ffa67b4.GetMemberGroupsRequestBuilder) {
    return ib7db9dde30bf0fa242f3a4321173cb16756ab77e1673add802f16c530ffa67b4.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContractRequestBuilder) GetMemberObjects()(*i2552b8bdbc7e213a9a6eae17e0a7d61fb51f1e4c5000eabd09038e0e3bb9acaa.GetMemberObjectsRequestBuilder) {
    return i2552b8bdbc7e213a9a6eae17e0a7d61fb51f1e4c5000eabd09038e0e3bb9acaa.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in contracts
func (m *ContractRequestBuilder) Patch(options *ContractRequestBuilderPatchOptions)(error) {
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
func (m *ContractRequestBuilder) Restore()(*i1786fc8f851599ebda69e9280c6b1118c2d2be8b0a0440447bbca330fa027c3b.RestoreRequestBuilder) {
    return i1786fc8f851599ebda69e9280c6b1118c2d2be8b0a0440447bbca330fa027c3b.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
