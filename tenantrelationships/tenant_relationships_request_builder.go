package tenantrelationships

import (
    i2a0d4734956a5e963566aae111af73e279c3d5379a451b3d2c44cd13b01a47a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/managedtenants"
    i764e92d08cb9718848f3e0ba8dbe891a44a52ede8e667e46244db3f6c20da77e "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadmincustomers"
    i83ea38918d7b5bf5d9afba525ad5cb3ea378a623d3e00cf8109b420010f0c284 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadminrelationships"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i794f862ecc921099af0c35b23c015b58bfdfa8b986158639b5bbfb2aacaea17e "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadmincustomers/item"
    ia7f4fac9af2719d5ab9e4339a8b2f55b843d5b618e4117c5460ccbfca8c2f196 "github.com/microsoftgraph/msgraph-beta-sdk-go/tenantrelationships/delegatedadminrelationships/item"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
)

// TenantRelationshipsRequestBuilder provides operations to manage the tenantRelationship singleton.
type TenantRelationshipsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TenantRelationshipsRequestBuilderGetOptions options for Get
type TenantRelationshipsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TenantRelationshipsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TenantRelationshipsRequestBuilderGetQueryParameters get tenantRelationships
type TenantRelationshipsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TenantRelationshipsRequestBuilderPatchOptions options for Patch
type TenantRelationshipsRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TenantRelationshipable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewTenantRelationshipsRequestBuilderInternal instantiates a new TenantRelationshipsRequestBuilder and sets the default values.
func NewTenantRelationshipsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TenantRelationshipsRequestBuilder) {
    m := &TenantRelationshipsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/tenantRelationships{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTenantRelationshipsRequestBuilder instantiates a new TenantRelationshipsRequestBuilder and sets the default values.
func NewTenantRelationshipsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TenantRelationshipsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantRelationshipsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get tenantRelationships
func (m *TenantRelationshipsRequestBuilder) CreateGetRequestInformation(options *TenantRelationshipsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update tenantRelationships
func (m *TenantRelationshipsRequestBuilder) CreatePatchRequestInformation(options *TenantRelationshipsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *TenantRelationshipsRequestBuilder) DelegatedAdminCustomers()(*i764e92d08cb9718848f3e0ba8dbe891a44a52ede8e667e46244db3f6c20da77e.DelegatedAdminCustomersRequestBuilder) {
    return i764e92d08cb9718848f3e0ba8dbe891a44a52ede8e667e46244db3f6c20da77e.NewDelegatedAdminCustomersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DelegatedAdminCustomersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.delegatedAdminCustomers.item collection
func (m *TenantRelationshipsRequestBuilder) DelegatedAdminCustomersById(id string)(*i794f862ecc921099af0c35b23c015b58bfdfa8b986158639b5bbfb2aacaea17e.DelegatedAdminCustomerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedAdminCustomer_id"] = id
    }
    return i794f862ecc921099af0c35b23c015b58bfdfa8b986158639b5bbfb2aacaea17e.NewDelegatedAdminCustomerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TenantRelationshipsRequestBuilder) DelegatedAdminRelationships()(*i83ea38918d7b5bf5d9afba525ad5cb3ea378a623d3e00cf8109b420010f0c284.DelegatedAdminRelationshipsRequestBuilder) {
    return i83ea38918d7b5bf5d9afba525ad5cb3ea378a623d3e00cf8109b420010f0c284.NewDelegatedAdminRelationshipsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DelegatedAdminRelationshipsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.tenantRelationships.delegatedAdminRelationships.item collection
func (m *TenantRelationshipsRequestBuilder) DelegatedAdminRelationshipsById(id string)(*ia7f4fac9af2719d5ab9e4339a8b2f55b843d5b618e4117c5460ccbfca8c2f196.DelegatedAdminRelationshipItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["delegatedAdminRelationship_id"] = id
    }
    return ia7f4fac9af2719d5ab9e4339a8b2f55b843d5b618e4117c5460ccbfca8c2f196.NewDelegatedAdminRelationshipItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get tenantRelationships
func (m *TenantRelationshipsRequestBuilder) Get(options *TenantRelationshipsRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TenantRelationshipable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateTenantRelationshipFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TenantRelationshipable), nil
}
func (m *TenantRelationshipsRequestBuilder) ManagedTenants()(*i2a0d4734956a5e963566aae111af73e279c3d5379a451b3d2c44cd13b01a47a0.ManagedTenantsRequestBuilder) {
    return i2a0d4734956a5e963566aae111af73e279c3d5379a451b3d2c44cd13b01a47a0.NewManagedTenantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update tenantRelationships
func (m *TenantRelationshipsRequestBuilder) Patch(options *TenantRelationshipsRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
