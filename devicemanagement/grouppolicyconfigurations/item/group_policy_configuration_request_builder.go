package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i16db673b6039e5f54f2159476d6177eaa1e9bfedc38455788aab318c22dffc2e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/updatedefinitionvalues"
    i2573cd7f18c0bef6763b0cd5d5daf2b302474f359820e0e9a5b01b4456663a1b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/definitionvalues"
    ib82d0e0e2e13d1b253232a5b5c89e5d5704adda9d4f7e9e1e53d75f5f967101d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/assignments"
    if10892f25c880cff6f351597e09e418fb5b4358d3bea03afb26c9d1e6f52d6e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/assign"
    i0cb6dc720cdfefbc227fbe58a3f3d6fae0bd97db3acd773de3b5dc3e12dd5792 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/assignments/item"
    id5d6fab6089b28b24032ceb492d4726a5f5b3dd8719e632a9f878e2fd9a9c9e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/definitionvalues/item"
)

// Builds and executes requests for operations under \deviceManagement\groupPolicyConfigurations\{groupPolicyConfiguration-id}
type GroupPolicyConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type GroupPolicyConfigurationRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type GroupPolicyConfigurationRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GroupPolicyConfigurationRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The group policy configurations created by this account.
type GroupPolicyConfigurationRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type GroupPolicyConfigurationRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyConfiguration;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *GroupPolicyConfigurationRequestBuilder) Assign()(*if10892f25c880cff6f351597e09e418fb5b4358d3bea03afb26c9d1e6f52d6e1.AssignRequestBuilder) {
    return if10892f25c880cff6f351597e09e418fb5b4358d3bea03afb26c9d1e6f52d6e1.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupPolicyConfigurationRequestBuilder) Assignments()(*ib82d0e0e2e13d1b253232a5b5c89e5d5704adda9d4f7e9e1e53d75f5f967101d.AssignmentsRequestBuilder) {
    return ib82d0e0e2e13d1b253232a5b5c89e5d5704adda9d4f7e9e1e53d75f5f967101d.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.deviceManagement.groupPolicyConfigurations.item.assignments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *GroupPolicyConfigurationRequestBuilder) AssignmentsById(id string)(*i0cb6dc720cdfefbc227fbe58a3f3d6fae0bd97db3acd773de3b5dc3e12dd5792.GroupPolicyConfigurationAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyConfigurationAssignment_id"] = id
    }
    return i0cb6dc720cdfefbc227fbe58a3f3d6fae0bd97db3acd773de3b5dc3e12dd5792.NewGroupPolicyConfigurationAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new GroupPolicyConfigurationRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewGroupPolicyConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyConfigurationRequestBuilder) {
    m := &GroupPolicyConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyConfigurations/{groupPolicyConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new GroupPolicyConfigurationRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewGroupPolicyConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// The group policy configurations created by this account.
// Parameters:
//  - options : Options for the request
func (m *GroupPolicyConfigurationRequestBuilder) CreateDeleteRequestInformation(options *GroupPolicyConfigurationRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The group policy configurations created by this account.
// Parameters:
//  - options : Options for the request
func (m *GroupPolicyConfigurationRequestBuilder) CreateGetRequestInformation(options *GroupPolicyConfigurationRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// The group policy configurations created by this account.
// Parameters:
//  - options : Options for the request
func (m *GroupPolicyConfigurationRequestBuilder) CreatePatchRequestInformation(options *GroupPolicyConfigurationRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *GroupPolicyConfigurationRequestBuilder) DefinitionValues()(*i2573cd7f18c0bef6763b0cd5d5daf2b302474f359820e0e9a5b01b4456663a1b.DefinitionValuesRequestBuilder) {
    return i2573cd7f18c0bef6763b0cd5d5daf2b302474f359820e0e9a5b01b4456663a1b.NewDefinitionValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.deviceManagement.groupPolicyConfigurations.item.definitionValues.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *GroupPolicyConfigurationRequestBuilder) DefinitionValuesById(id string)(*id5d6fab6089b28b24032ceb492d4726a5f5b3dd8719e632a9f878e2fd9a9c9e8.GroupPolicyDefinitionValueRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyDefinitionValue_id"] = id
    }
    return id5d6fab6089b28b24032ceb492d4726a5f5b3dd8719e632a9f878e2fd9a9c9e8.NewGroupPolicyDefinitionValueRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The group policy configurations created by this account.
// Parameters:
//  - options : Options for the request
func (m *GroupPolicyConfigurationRequestBuilder) Delete(options *GroupPolicyConfigurationRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// The group policy configurations created by this account.
// Parameters:
//  - options : Options for the request
func (m *GroupPolicyConfigurationRequestBuilder) Get(options *GroupPolicyConfigurationRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyConfiguration, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewGroupPolicyConfiguration() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyConfiguration), nil
}
// The group policy configurations created by this account.
// Parameters:
//  - options : Options for the request
func (m *GroupPolicyConfigurationRequestBuilder) Patch(options *GroupPolicyConfigurationRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *GroupPolicyConfigurationRequestBuilder) UpdateDefinitionValues()(*i16db673b6039e5f54f2159476d6177eaa1e9bfedc38455788aab318c22dffc2e.UpdateDefinitionValuesRequestBuilder) {
    return i16db673b6039e5f54f2159476d6177eaa1e9bfedc38455788aab318c22dffc2e.NewUpdateDefinitionValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
