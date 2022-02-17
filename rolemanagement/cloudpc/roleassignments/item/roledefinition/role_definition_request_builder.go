package roledefinition

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i44ebb3756e32798ca2c6bc715afaa56295c6e5bb59d461eee46541b9ac7d0755 "github.com/microsoftgraph/msgraph-beta-sdk-go/rolemanagement/cloudpc/roleassignments/item/roledefinition/ref"
)

// RoleDefinitionRequestBuilder builds and executes requests for operations under \roleManagement\cloudPC\roleAssignments\{unifiedRoleAssignmentMultiple-id}\roleDefinition
type RoleDefinitionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RoleDefinitionRequestBuilderGetOptions options for Get
type RoleDefinitionRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RoleDefinitionRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RoleDefinitionRequestBuilderGetQueryParameters specifies the roleDefinition that the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment.  Supports $filter (eq operator on id, isBuiltIn, and displayName, and startsWith operator on displayName)  and $expand.
type RoleDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// NewRoleDefinitionRequestBuilderInternal instantiates a new RoleDefinitionRequestBuilder and sets the default values.
func NewRoleDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleDefinitionRequestBuilder) {
    m := &RoleDefinitionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/cloudPC/roleAssignments/{unifiedRoleAssignmentMultiple_id}/roleDefinition{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleDefinitionRequestBuilder instantiates a new RoleDefinitionRequestBuilder and sets the default values.
func NewRoleDefinitionRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation specifies the roleDefinition that the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment.  Supports $filter (eq operator on id, isBuiltIn, and displayName, and startsWith operator on displayName)  and $expand.
func (m *RoleDefinitionRequestBuilder) CreateGetRequestInformation(options *RoleDefinitionRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get specifies the roleDefinition that the assignment is for. Provided so that callers can get the role definition using $expand at the same time as getting the role assignment.  Supports $filter (eq operator on id, isBuiltIn, and displayName, and startsWith operator on displayName)  and $expand.
func (m *RoleDefinitionRequestBuilder) Get(options *RoleDefinitionRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleDefinition, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUnifiedRoleDefinition() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UnifiedRoleDefinition), nil
}
func (m *RoleDefinitionRequestBuilder) Ref()(*i44ebb3756e32798ca2c6bc715afaa56295c6e5bb59d461eee46541b9ac7d0755.RefRequestBuilder) {
    return i44ebb3756e32798ca2c6bc715afaa56295c6e5bb59d461eee46541b9ac7d0755.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
