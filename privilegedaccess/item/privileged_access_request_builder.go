package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
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

// Builds and executes requests for operations under \privilegedAccess\{privilegedAccess-id}
type PrivilegedAccessRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type PrivilegedAccessRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type PrivilegedAccessRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PrivilegedAccessRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get entity from privilegedAccess by key
type PrivilegedAccessRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type PrivilegedAccessRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedAccess;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new PrivilegedAccessRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPrivilegedAccessRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrivilegedAccessRequestBuilder) {
    m := &PrivilegedAccessRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedAccess/{privilegedAccess_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new PrivilegedAccessRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPrivilegedAccessRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PrivilegedAccessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedAccessRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete entity from privilegedAccess
// Parameters:
//  - options : Options for the request
func (m *PrivilegedAccessRequestBuilder) CreateDeleteRequestInformation(options *PrivilegedAccessRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get entity from privilegedAccess by key
// Parameters:
//  - options : Options for the request
func (m *PrivilegedAccessRequestBuilder) CreateGetRequestInformation(options *PrivilegedAccessRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Update entity in privilegedAccess
// Parameters:
//  - options : Options for the request
func (m *PrivilegedAccessRequestBuilder) CreatePatchRequestInformation(options *PrivilegedAccessRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete entity from privilegedAccess
// Parameters:
//  - options : Options for the request
func (m *PrivilegedAccessRequestBuilder) Delete(options *PrivilegedAccessRequestBuilderDeleteOptions)(error) {
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
// Get entity from privilegedAccess by key
// Parameters:
//  - options : Options for the request
func (m *PrivilegedAccessRequestBuilder) Get(options *PrivilegedAccessRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedAccess, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPrivilegedAccess() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.PrivilegedAccess), nil
}
// Update entity in privilegedAccess
// Parameters:
//  - options : Options for the request
func (m *PrivilegedAccessRequestBuilder) Patch(options *PrivilegedAccessRequestBuilderPatchOptions)(error) {
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
func (m *PrivilegedAccessRequestBuilder) Resources()(*i9a0b2cd1cee2ff8b6db6a31960b35b297ea8a82ac1c9d7147b11db2225e14166.ResourcesRequestBuilder) {
    return i9a0b2cd1cee2ff8b6db6a31960b35b297ea8a82ac1c9d7147b11db2225e14166.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.privilegedAccess.item.resources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PrivilegedAccessRequestBuilder) ResourcesById(id string)(*ifef34655781055c0461c93a60eb94832b88d5e7127511a01d91517b0dbb2056f.GovernanceResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceResource_id"] = id
    }
    return ifef34655781055c0461c93a60eb94832b88d5e7127511a01d91517b0dbb2056f.NewGovernanceResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrivilegedAccessRequestBuilder) RoleAssignmentRequests()(*i13776d5c18dfe5cb7926395f0d7b87e53ddf81ba000513408fbd99e65db8c5fb.RoleAssignmentRequestsRequestBuilder) {
    return i13776d5c18dfe5cb7926395f0d7b87e53ddf81ba000513408fbd99e65db8c5fb.NewRoleAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.privilegedAccess.item.roleAssignmentRequests.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PrivilegedAccessRequestBuilder) RoleAssignmentRequestsById(id string)(*i055f2c9d21e7d744d266e629aae4492646c4497616eefa08e3bdc13deae816bc.GovernanceRoleAssignmentRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleAssignmentRequest_id"] = id
    }
    return i055f2c9d21e7d744d266e629aae4492646c4497616eefa08e3bdc13deae816bc.NewGovernanceRoleAssignmentRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrivilegedAccessRequestBuilder) RoleAssignments()(*i6517906c52d298f87f2e9d9ac7e9044b06b89fb46f521c3da9b9d1f087671c44.RoleAssignmentsRequestBuilder) {
    return i6517906c52d298f87f2e9d9ac7e9044b06b89fb46f521c3da9b9d1f087671c44.NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.privilegedAccess.item.roleAssignments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PrivilegedAccessRequestBuilder) RoleAssignmentsById(id string)(*i669c13d28c7fa4ec2373a44ba75602285da1019797e39c430f2ede5ee0850356.GovernanceRoleAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleAssignment_id"] = id
    }
    return i669c13d28c7fa4ec2373a44ba75602285da1019797e39c430f2ede5ee0850356.NewGovernanceRoleAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrivilegedAccessRequestBuilder) RoleDefinitions()(*ie0282f7e2e54b40fe4176464d126008c5f839e06d60480f482df4629ee34909c.RoleDefinitionsRequestBuilder) {
    return ie0282f7e2e54b40fe4176464d126008c5f839e06d60480f482df4629ee34909c.NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.privilegedAccess.item.roleDefinitions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PrivilegedAccessRequestBuilder) RoleDefinitionsById(id string)(*i93a22192c0fdf52ece01982a56d4c7f6efcce6d81281c6a74a3a9d90200baf9d.GovernanceRoleDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleDefinition_id"] = id
    }
    return i93a22192c0fdf52ece01982a56d4c7f6efcce6d81281c6a74a3a9d90200baf9d.NewGovernanceRoleDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PrivilegedAccessRequestBuilder) RoleSettings()(*i83aba682e003297032480b3ecf836e1584c1518b5430625638ccefc15ac6bc8b.RoleSettingsRequestBuilder) {
    return i83aba682e003297032480b3ecf836e1584c1518b5430625638ccefc15ac6bc8b.NewRoleSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go.privilegedAccess.item.roleSettings.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PrivilegedAccessRequestBuilder) RoleSettingsById(id string)(*id87892c36f8f2494356ab0b890f05c975ca897b7061b4af81b617506bdd996cf.GovernanceRoleSettingRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["governanceRoleSetting_id"] = id
    }
    return id87892c36f8f2494356ab0b890f05c975ca897b7061b4af81b617506bdd996cf.NewGovernanceRoleSettingRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
