package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i16db673b6039e5f54f2159476d6177eaa1e9bfedc38455788aab318c22dffc2e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/updatedefinitionvalues"
    i2573cd7f18c0bef6763b0cd5d5daf2b302474f359820e0e9a5b01b4456663a1b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/definitionvalues"
    ib82d0e0e2e13d1b253232a5b5c89e5d5704adda9d4f7e9e1e53d75f5f967101d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/assignments"
    if10892f25c880cff6f351597e09e418fb5b4358d3bea03afb26c9d1e6f52d6e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/assign"
    i0cb6dc720cdfefbc227fbe58a3f3d6fae0bd97db3acd773de3b5dc3e12dd5792 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/assignments/item"
    id5d6fab6089b28b24032ceb492d4726a5f5b3dd8719e632a9f878e2fd9a9c9e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/definitionvalues/item"
)

// GroupPolicyConfigurationItemRequestBuilder provides operations to manage the groupPolicyConfigurations property of the microsoft.graph.deviceManagement entity.
type GroupPolicyConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GroupPolicyConfigurationItemRequestBuilderDeleteOptions options for Delete
type GroupPolicyConfigurationItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupPolicyConfigurationItemRequestBuilderGetOptions options for Get
type GroupPolicyConfigurationItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GroupPolicyConfigurationItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupPolicyConfigurationItemRequestBuilderGetQueryParameters the group policy configurations created by this account.
type GroupPolicyConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// GroupPolicyConfigurationItemRequestBuilderPatchOptions options for Patch
type GroupPolicyConfigurationItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyConfigurationable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *GroupPolicyConfigurationItemRequestBuilder) Assign()(*if10892f25c880cff6f351597e09e418fb5b4358d3bea03afb26c9d1e6f52d6e1.AssignRequestBuilder) {
    return if10892f25c880cff6f351597e09e418fb5b4358d3bea03afb26c9d1e6f52d6e1.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *GroupPolicyConfigurationItemRequestBuilder) Assignments()(*ib82d0e0e2e13d1b253232a5b5c89e5d5704adda9d4f7e9e1e53d75f5f967101d.AssignmentsRequestBuilder) {
    return ib82d0e0e2e13d1b253232a5b5c89e5d5704adda9d4f7e9e1e53d75f5f967101d.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.groupPolicyConfigurations.item.assignments.item collection
func (m *GroupPolicyConfigurationItemRequestBuilder) AssignmentsById(id string)(*i0cb6dc720cdfefbc227fbe58a3f3d6fae0bd97db3acd773de3b5dc3e12dd5792.GroupPolicyConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyConfigurationAssignment_id"] = id
    }
    return i0cb6dc720cdfefbc227fbe58a3f3d6fae0bd97db3acd773de3b5dc3e12dd5792.NewGroupPolicyConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewGroupPolicyConfigurationItemRequestBuilderInternal instantiates a new GroupPolicyConfigurationItemRequestBuilder and sets the default values.
func NewGroupPolicyConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyConfigurationItemRequestBuilder) {
    m := &GroupPolicyConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyConfigurations/{groupPolicyConfiguration_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupPolicyConfigurationItemRequestBuilder instantiates a new GroupPolicyConfigurationItemRequestBuilder and sets the default values.
func NewGroupPolicyConfigurationItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property groupPolicyConfigurations for deviceManagement
func (m *GroupPolicyConfigurationItemRequestBuilder) CreateDeleteRequestInformation(options *GroupPolicyConfigurationItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the group policy configurations created by this account.
func (m *GroupPolicyConfigurationItemRequestBuilder) CreateGetRequestInformation(options *GroupPolicyConfigurationItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property groupPolicyConfigurations in deviceManagement
func (m *GroupPolicyConfigurationItemRequestBuilder) CreatePatchRequestInformation(options *GroupPolicyConfigurationItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *GroupPolicyConfigurationItemRequestBuilder) DefinitionValues()(*i2573cd7f18c0bef6763b0cd5d5daf2b302474f359820e0e9a5b01b4456663a1b.DefinitionValuesRequestBuilder) {
    return i2573cd7f18c0bef6763b0cd5d5daf2b302474f359820e0e9a5b01b4456663a1b.NewDefinitionValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefinitionValuesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.groupPolicyConfigurations.item.definitionValues.item collection
func (m *GroupPolicyConfigurationItemRequestBuilder) DefinitionValuesById(id string)(*id5d6fab6089b28b24032ceb492d4726a5f5b3dd8719e632a9f878e2fd9a9c9e8.GroupPolicyDefinitionValueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyDefinitionValue_id"] = id
    }
    return id5d6fab6089b28b24032ceb492d4726a5f5b3dd8719e632a9f878e2fd9a9c9e8.NewGroupPolicyDefinitionValueItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property groupPolicyConfigurations for deviceManagement
func (m *GroupPolicyConfigurationItemRequestBuilder) Delete(options *GroupPolicyConfigurationItemRequestBuilderDeleteOptions)(error) {
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
// Get the group policy configurations created by this account.
func (m *GroupPolicyConfigurationItemRequestBuilder) Get(options *GroupPolicyConfigurationItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateGroupPolicyConfigurationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyConfigurationable), nil
}
// Patch update the navigation property groupPolicyConfigurations in deviceManagement
func (m *GroupPolicyConfigurationItemRequestBuilder) Patch(options *GroupPolicyConfigurationItemRequestBuilderPatchOptions)(error) {
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
func (m *GroupPolicyConfigurationItemRequestBuilder) UpdateDefinitionValues()(*i16db673b6039e5f54f2159476d6177eaa1e9bfedc38455788aab318c22dffc2e.UpdateDefinitionValuesRequestBuilder) {
    return i16db673b6039e5f54f2159476d6177eaa1e9bfedc38455788aab318c22dffc2e.NewUpdateDefinitionValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
