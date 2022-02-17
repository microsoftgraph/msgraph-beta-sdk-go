package identity

import (
    i11f22886c33eed99dd85e4d242bbf0976ec471e9e441cfa6a58b5864040bd094 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/apiconnectors"
    i5f1ee64209aad8370602f56501934370194f1248216e73d51b8543369c558033 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/continuousaccessevaluationpolicy"
    i6635ff8f2aa1c074739a0c7312e2c31642186b9aadaa7290a91096bd632e192d "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows"
    i96c0428619bb9ada314a888efda4bc4c3bb5c55513e2cfaa28f956e608f98c45 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/userflows"
    iaa68fab70e5038f618ea9044bb8a9fbc651b67c06a35357c7b18c20d63880c85 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/userflowattributes"
    ic34e66a722321196953b295f6b75e587885dec11e8a616d1a3f940a14b88d50e "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/conditionalaccess"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ie614621c7e356610a3cda4707ec623ea542f5d84eaede8bc6ac868a655447bc5 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows"
    if8e62443e000757279f6963237e640c3a2f6bf54d567fb126a582ea7bac48a91 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/identityproviders"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i056de954ae8d2bd3c37cd62ab74e2a653c76cc631e79df2abe7d5423b433ad7f "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2xuserflows/item"
    i4064b0ed26ff7ac321cce776d8d90a08dbeca0d51fec9f75ba0ee4621fd09ffa "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/identityproviders/item"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i66a685596e55f332f9dea7615db3653de9c5945d4f7b70187e76b70ed8344761 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/b2cuserflows/item"
    i80b7576b721756e8211e557cb66927cff088fae6e30a6b5990e694f8c8938b4d "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/userflows/item"
    i96fbb19ece75424f3a1291715def826b0b7091a7b004df122cb970e740300385 "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/userflowattributes/item"
    id9b7138b09b41c1638c96829912584afe356128a86a38484a5041a342124094b "github.com/microsoftgraph/msgraph-beta-sdk-go/identity/apiconnectors/item"
)

// IdentityRequestBuilder builds and executes requests for operations under \identity
type IdentityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// IdentityRequestBuilderGetOptions options for Get
type IdentityRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *IdentityRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// IdentityRequestBuilderGetQueryParameters get identity
type IdentityRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// IdentityRequestBuilderPatchOptions options for Patch
type IdentityRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentityContainer;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *IdentityRequestBuilder) ApiConnectors()(*i11f22886c33eed99dd85e4d242bbf0976ec471e9e441cfa6a58b5864040bd094.ApiConnectorsRequestBuilder) {
    return i11f22886c33eed99dd85e4d242bbf0976ec471e9e441cfa6a58b5864040bd094.NewApiConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ApiConnectorsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.apiConnectors.item collection
func (m *IdentityRequestBuilder) ApiConnectorsById(id string)(*id9b7138b09b41c1638c96829912584afe356128a86a38484a5041a342124094b.IdentityApiConnectorRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityApiConnector_id"] = id
    }
    return id9b7138b09b41c1638c96829912584afe356128a86a38484a5041a342124094b.NewIdentityApiConnectorRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *IdentityRequestBuilder) B2cUserFlows()(*i6635ff8f2aa1c074739a0c7312e2c31642186b9aadaa7290a91096bd632e192d.B2cUserFlowsRequestBuilder) {
    return i6635ff8f2aa1c074739a0c7312e2c31642186b9aadaa7290a91096bd632e192d.NewB2cUserFlowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// B2cUserFlowsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2cUserFlows.item collection
func (m *IdentityRequestBuilder) B2cUserFlowsById(id string)(*i66a685596e55f332f9dea7615db3653de9c5945d4f7b70187e76b70ed8344761.B2cIdentityUserFlowRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["b2cIdentityUserFlow_id"] = id
    }
    return i66a685596e55f332f9dea7615db3653de9c5945d4f7b70187e76b70ed8344761.NewB2cIdentityUserFlowRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *IdentityRequestBuilder) B2xUserFlows()(*ie614621c7e356610a3cda4707ec623ea542f5d84eaede8bc6ac868a655447bc5.B2xUserFlowsRequestBuilder) {
    return ie614621c7e356610a3cda4707ec623ea542f5d84eaede8bc6ac868a655447bc5.NewB2xUserFlowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// B2xUserFlowsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.b2xUserFlows.item collection
func (m *IdentityRequestBuilder) B2xUserFlowsById(id string)(*i056de954ae8d2bd3c37cd62ab74e2a653c76cc631e79df2abe7d5423b433ad7f.B2xIdentityUserFlowRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["b2xIdentityUserFlow_id"] = id
    }
    return i056de954ae8d2bd3c37cd62ab74e2a653c76cc631e79df2abe7d5423b433ad7f.NewB2xIdentityUserFlowRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *IdentityRequestBuilder) ConditionalAccess()(*ic34e66a722321196953b295f6b75e587885dec11e8a616d1a3f940a14b88d50e.ConditionalAccessRequestBuilder) {
    return ic34e66a722321196953b295f6b75e587885dec11e8a616d1a3f940a14b88d50e.NewConditionalAccessRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewIdentityRequestBuilderInternal instantiates a new IdentityRequestBuilder and sets the default values.
func NewIdentityRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IdentityRequestBuilder) {
    m := &IdentityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityRequestBuilder instantiates a new IdentityRequestBuilder and sets the default values.
func NewIdentityRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IdentityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *IdentityRequestBuilder) ContinuousAccessEvaluationPolicy()(*i5f1ee64209aad8370602f56501934370194f1248216e73d51b8543369c558033.ContinuousAccessEvaluationPolicyRequestBuilder) {
    return i5f1ee64209aad8370602f56501934370194f1248216e73d51b8543369c558033.NewContinuousAccessEvaluationPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get identity
func (m *IdentityRequestBuilder) CreateGetRequestInformation(options *IdentityRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update identity
func (m *IdentityRequestBuilder) CreatePatchRequestInformation(options *IdentityRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get get identity
func (m *IdentityRequestBuilder) Get(options *IdentityRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentityContainer, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewIdentityContainer() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.IdentityContainer), nil
}
func (m *IdentityRequestBuilder) IdentityProviders()(*if8e62443e000757279f6963237e640c3a2f6bf54d567fb126a582ea7bac48a91.IdentityProvidersRequestBuilder) {
    return if8e62443e000757279f6963237e640c3a2f6bf54d567fb126a582ea7bac48a91.NewIdentityProvidersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IdentityProvidersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.identityProviders.item collection
func (m *IdentityRequestBuilder) IdentityProvidersById(id string)(*i4064b0ed26ff7ac321cce776d8d90a08dbeca0d51fec9f75ba0ee4621fd09ffa.IdentityProviderBaseRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityProviderBase_id"] = id
    }
    return i4064b0ed26ff7ac321cce776d8d90a08dbeca0d51fec9f75ba0ee4621fd09ffa.NewIdentityProviderBaseRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update identity
func (m *IdentityRequestBuilder) Patch(options *IdentityRequestBuilderPatchOptions)(error) {
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
func (m *IdentityRequestBuilder) UserFlowAttributes()(*iaa68fab70e5038f618ea9044bb8a9fbc651b67c06a35357c7b18c20d63880c85.UserFlowAttributesRequestBuilder) {
    return iaa68fab70e5038f618ea9044bb8a9fbc651b67c06a35357c7b18c20d63880c85.NewUserFlowAttributesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserFlowAttributesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.userFlowAttributes.item collection
func (m *IdentityRequestBuilder) UserFlowAttributesById(id string)(*i96fbb19ece75424f3a1291715def826b0b7091a7b004df122cb970e740300385.IdentityUserFlowAttributeRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityUserFlowAttribute_id"] = id
    }
    return i96fbb19ece75424f3a1291715def826b0b7091a7b004df122cb970e740300385.NewIdentityUserFlowAttributeRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *IdentityRequestBuilder) UserFlows()(*i96c0428619bb9ada314a888efda4bc4c3bb5c55513e2cfaa28f956e608f98c45.UserFlowsRequestBuilder) {
    return i96c0428619bb9ada314a888efda4bc4c3bb5c55513e2cfaa28f956e608f98c45.NewUserFlowsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserFlowsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.identity.userFlows.item collection
func (m *IdentityRequestBuilder) UserFlowsById(id string)(*i80b7576b721756e8211e557cb66927cff088fae6e30a6b5990e694f8c8938b4d.IdentityUserFlowRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identityUserFlow_id"] = id
    }
    return i80b7576b721756e8211e557cb66927cff088fae6e30a6b5990e694f8c8938b4d.NewIdentityUserFlowRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
