package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i13776d5c18dfe5cb7926395f0d7b87e53ddf81ba000513408fbd99e65db8c5fb "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/roleassignmentrequests"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i6517906c52d298f87f2e9d9ac7e9044b06b89fb46f521c3da9b9d1f087671c44 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/roleassignments"
    i83aba682e003297032480b3ecf836e1584c1518b5430625638ccefc15ac6bc8b "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/rolesettings"
    i9a0b2cd1cee2ff8b6db6a31960b35b297ea8a82ac1c9d7147b11db2225e14166 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/resources"
    ie0282f7e2e54b40fe4176464d126008c5f839e06d60480f482df4629ee34909c "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/roledefinitions"
    i055f2c9d21e7d744d266e629aae4492646c4497616eefa08e3bdc13deae816bc "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/roleassignmentrequests/item"
    i669c13d28c7fa4ec2373a44ba75602285da1019797e39c430f2ede5ee0850356 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/roleassignments/item"
    i93a22192c0fdf52ece01982a56d4c7f6efcce6d81281c6a74a3a9d90200baf9d "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/roledefinitions/item"
    id87892c36f8f2494356ab0b890f05c975ca897b7061b4af81b617506bdd996cf "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/rolesettings/item"
    ifef34655781055c0461c93a60eb94832b88d5e7127511a01d91517b0dbb2056f "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedaccess/item/resources/item"
)

// PrivilegedAccessItemRequestBuilder provides operations to manage the collection of privilegedAccess entities.
type PrivilegedAccessItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PrivilegedAccessItemRequestBuilderDeleteOptions options for Delete
type PrivilegedAccessItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrivilegedAccessItemRequestBuilderGetOptions options for Get
type PrivilegedAccessItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PrivilegedAccessItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PrivilegedAccessItemRequestBuilderGetQueryParameters get entity from privilegedAccess by key
type PrivilegedAccessItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PrivilegedAccessItemRequestBuilderPatchOptions options for Patch
type PrivilegedAccessItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedAccessable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewPrivilegedAccessItemRequestBuilderInternal instantiates a new PrivilegedAccessItemRequestBuilder and sets the default values.
func NewPrivilegedAccessItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrivilegedAccessItemRequestBuilder) {
    m := &PrivilegedAccessItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedAccess/{privilegedAccess_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrivilegedAccessItemRequestBuilder instantiates a new PrivilegedAccessItemRequestBuilder and sets the default values.
func NewPrivilegedAccessItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrivilegedAccessItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from privilegedAccess
func (m *PrivilegedAccessItemRequestBuilder) CreateDeleteRequestInformation(options *PrivilegedAccessItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from privilegedAccess by key
func (m *PrivilegedAccessItemRequestBuilder) CreateGetRequestInformation(options *PrivilegedAccessItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in privilegedAccess
func (m *PrivilegedAccessItemRequestBuilder) CreatePatchRequestInformation(options *PrivilegedAccessItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from privilegedAccess
func (m *PrivilegedAccessItemRequestBuilder) Delete(options *PrivilegedAccessItemRequestBuilderDeleteOptions)(error) {
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
// Get get entity from privilegedAccess by key
func (m *PrivilegedAccessItemRequestBuilder) Get(options *PrivilegedAccessItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedAccessable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreatePrivilegedAccessFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedAccessable), nil
}
// Patch update entity in privilegedAccess
func (m *PrivilegedAccessItemRequestBuilder) Patch(options *PrivilegedAccessItemRequestBuilderPatchOptions)(error) {
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
func (m *PrivilegedAccessItemRequestBuilder) Resources()(*i9a0b2cd1cee2ff8b6db6a31960b35b297ea8a82ac1c9d7147b11db2225e14166.ResourcesRequestBuilder) {
    return i9a0b2cd1cee2ff8b6db6a31960b35b297ea8a82ac1c9d7147b11db2225e14166.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedAccess.item.resources.item collection
func (m *PrivilegedAccessItemRequestBuilder) ResourcesById(id string)(*ifef34655781055c0461c93a60eb94832b88d5e7127511a01d91517b0dbb2056f.GovernanceResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceResource_id"] = id
    }
    return ifef34655781055c0461c93a60eb94832b88d5e7127511a01d91517b0dbb2056f.NewGovernanceResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrivilegedAccessItemRequestBuilder) RoleAssignmentRequests()(*i13776d5c18dfe5cb7926395f0d7b87e53ddf81ba000513408fbd99e65db8c5fb.RoleAssignmentRequestsRequestBuilder) {
    return i13776d5c18dfe5cb7926395f0d7b87e53ddf81ba000513408fbd99e65db8c5fb.NewRoleAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedAccess.item.roleAssignmentRequests.item collection
func (m *PrivilegedAccessItemRequestBuilder) RoleAssignmentRequestsById(id string)(*i055f2c9d21e7d744d266e629aae4492646c4497616eefa08e3bdc13deae816bc.GovernanceRoleAssignmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleAssignmentRequest_id"] = id
    }
    return i055f2c9d21e7d744d266e629aae4492646c4497616eefa08e3bdc13deae816bc.NewGovernanceRoleAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrivilegedAccessItemRequestBuilder) RoleAssignments()(*i6517906c52d298f87f2e9d9ac7e9044b06b89fb46f521c3da9b9d1f087671c44.RoleAssignmentsRequestBuilder) {
    return i6517906c52d298f87f2e9d9ac7e9044b06b89fb46f521c3da9b9d1f087671c44.NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedAccess.item.roleAssignments.item collection
func (m *PrivilegedAccessItemRequestBuilder) RoleAssignmentsById(id string)(*i669c13d28c7fa4ec2373a44ba75602285da1019797e39c430f2ede5ee0850356.GovernanceRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleAssignment_id"] = id
    }
    return i669c13d28c7fa4ec2373a44ba75602285da1019797e39c430f2ede5ee0850356.NewGovernanceRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrivilegedAccessItemRequestBuilder) RoleDefinitions()(*ie0282f7e2e54b40fe4176464d126008c5f839e06d60480f482df4629ee34909c.RoleDefinitionsRequestBuilder) {
    return ie0282f7e2e54b40fe4176464d126008c5f839e06d60480f482df4629ee34909c.NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedAccess.item.roleDefinitions.item collection
func (m *PrivilegedAccessItemRequestBuilder) RoleDefinitionsById(id string)(*i93a22192c0fdf52ece01982a56d4c7f6efcce6d81281c6a74a3a9d90200baf9d.GovernanceRoleDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleDefinition_id"] = id
    }
    return i93a22192c0fdf52ece01982a56d4c7f6efcce6d81281c6a74a3a9d90200baf9d.NewGovernanceRoleDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrivilegedAccessItemRequestBuilder) RoleSettings()(*i83aba682e003297032480b3ecf836e1584c1518b5430625638ccefc15ac6bc8b.RoleSettingsRequestBuilder) {
    return i83aba682e003297032480b3ecf836e1584c1518b5430625638ccefc15ac6bc8b.NewRoleSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleSettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedAccess.item.roleSettings.item collection
func (m *PrivilegedAccessItemRequestBuilder) RoleSettingsById(id string)(*id87892c36f8f2494356ab0b890f05c975ca897b7061b4af81b617506bdd996cf.GovernanceRoleSettingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleSetting_id"] = id
    }
    return id87892c36f8f2494356ab0b890f05c975ca897b7061b4af81b617506bdd996cf.NewGovernanceRoleSettingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
