package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i039e9ad1e0aeeb56dd607902862da4d97e74a08c6c1345dbc1e47e578403eb62 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroledefinitions/item/rolesetting"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i923c7ce7c1de619778c9eaada0c78ec5cb3873ccafc4089803ef40514954ad47 "github.com/microsoftgraph/msgraph-beta-sdk-go/governanceroledefinitions/item/resource"
)

// GovernanceRoleDefinitionRequestBuilder builds and executes requests for operations under \governanceRoleDefinitions\{governanceRoleDefinition-id}
type GovernanceRoleDefinitionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GovernanceRoleDefinitionRequestBuilderDeleteOptions options for Delete
type GovernanceRoleDefinitionRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GovernanceRoleDefinitionRequestBuilderGetOptions options for Get
type GovernanceRoleDefinitionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GovernanceRoleDefinitionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GovernanceRoleDefinitionRequestBuilderGetQueryParameters get entity from governanceRoleDefinitions by key
type GovernanceRoleDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// GovernanceRoleDefinitionRequestBuilderPatchOptions options for Patch
type GovernanceRoleDefinitionRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleDefinition;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGovernanceRoleDefinitionRequestBuilderInternal instantiates a new GovernanceRoleDefinitionRequestBuilder and sets the default values.
func NewGovernanceRoleDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GovernanceRoleDefinitionRequestBuilder) {
    m := &GovernanceRoleDefinitionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/governanceRoleDefinitions/{governanceRoleDefinition_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGovernanceRoleDefinitionRequestBuilder instantiates a new GovernanceRoleDefinitionRequestBuilder and sets the default values.
func NewGovernanceRoleDefinitionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GovernanceRoleDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGovernanceRoleDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from governanceRoleDefinitions
func (m *GovernanceRoleDefinitionRequestBuilder) CreateDeleteRequestInformation(options *GovernanceRoleDefinitionRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from governanceRoleDefinitions by key
func (m *GovernanceRoleDefinitionRequestBuilder) CreateGetRequestInformation(options *GovernanceRoleDefinitionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in governanceRoleDefinitions
func (m *GovernanceRoleDefinitionRequestBuilder) CreatePatchRequestInformation(options *GovernanceRoleDefinitionRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from governanceRoleDefinitions
func (m *GovernanceRoleDefinitionRequestBuilder) Delete(options *GovernanceRoleDefinitionRequestBuilderDeleteOptions)(error) {
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
// Get get entity from governanceRoleDefinitions by key
func (m *GovernanceRoleDefinitionRequestBuilder) Get(options *GovernanceRoleDefinitionRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleDefinition, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGovernanceRoleDefinition() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GovernanceRoleDefinition), nil
}
// Patch update entity in governanceRoleDefinitions
func (m *GovernanceRoleDefinitionRequestBuilder) Patch(options *GovernanceRoleDefinitionRequestBuilderPatchOptions)(error) {
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
func (m *GovernanceRoleDefinitionRequestBuilder) Resource()(*i923c7ce7c1de619778c9eaada0c78ec5cb3873ccafc4089803ef40514954ad47.ResourceRequestBuilder) {
    return i923c7ce7c1de619778c9eaada0c78ec5cb3873ccafc4089803ef40514954ad47.NewResourceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GovernanceRoleDefinitionRequestBuilder) RoleSetting()(*i039e9ad1e0aeeb56dd607902862da4d97e74a08c6c1345dbc1e47e578403eb62.RoleSettingRequestBuilder) {
    return i039e9ad1e0aeeb56dd607902862da4d97e74a08c6c1345dbc1e47e578403eb62.NewRoleSettingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
