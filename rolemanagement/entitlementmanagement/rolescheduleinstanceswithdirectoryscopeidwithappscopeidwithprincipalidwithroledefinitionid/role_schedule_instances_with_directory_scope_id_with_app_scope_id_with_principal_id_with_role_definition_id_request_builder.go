package rolescheduleinstanceswithdirectoryscopeidwithappscopeidwithprincipalidwithroledefinitionid

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilder builds and executes requests for operations under \roleManagement\entitlementManagement\microsoft.graph.roleScheduleInstances(directoryScopeId='{directoryScopeId}',appScopeId='{appScopeId}',principalId='{principalId}',roleDefinitionId='{roleDefinitionId}')
type RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilderGetOptions options for Get
type RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewRoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilderInternal instantiates a new RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilder and sets the default values.
func NewRoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, directoryScopeId *string, appScopeId *string, principalId *string, roleDefinitionId *string)(*RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilder) {
    m := &RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/microsoft.graph.roleScheduleInstances(directoryScopeId='{directoryScopeId}',appScopeId='{appScopeId}',principalId='{principalId}',roleDefinitionId='{roleDefinitionId}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if directoryScopeId != nil {
        urlTplParams["directoryScopeId"] = *directoryScopeId
    }
    if appScopeId != nil {
        urlTplParams["appScopeId"] = *appScopeId
    }
    if principalId != nil {
        urlTplParams["principalId"] = *principalId
    }
    if roleDefinitionId != nil {
        urlTplParams["roleDefinitionId"] = *roleDefinitionId
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilder instantiates a new RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilder and sets the default values.
func NewRoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil, nil)
}
// CreateGetRequestInformation invoke function roleScheduleInstances
func (m *RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilder) CreateGetRequestInformation(options *RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
// Get invoke function roleScheduleInstances
func (m *RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilder) Get(options *RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionIdRequestBuilderGetOptions)([]RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionId, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollectionAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionId() }, nil, nil)
    if err != nil {
        return nil, err
    }
    val := make([]RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionId, len(res))
    for i, v := range res {
        val[i] = *(v.(*RoleScheduleInstancesWithDirectoryScopeIdWithAppScopeIdWithPrincipalIdWithRoleDefinitionId))
    }
    return val, nil
}
