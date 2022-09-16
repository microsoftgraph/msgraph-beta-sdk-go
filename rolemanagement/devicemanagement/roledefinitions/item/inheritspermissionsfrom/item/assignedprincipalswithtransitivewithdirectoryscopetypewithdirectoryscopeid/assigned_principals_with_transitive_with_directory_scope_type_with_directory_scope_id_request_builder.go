package assignedprincipalswithtransitivewithdirectoryscopetypewithdirectoryscopeid

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilder provides operations to call the assignedPrincipals method.
type AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilderGetQueryParameters invoke function assignedPrincipals
type AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilderGetQueryParameters struct {
    // Usage: directoryScopeId='{directoryScopeId}'
    DirectoryScopeId *string
    // Usage: directoryScopeType='{directoryScopeType}'
    DirectoryScopeType *string
    // Usage: transitive={transitive}
    Transitive *bool
}
// AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilderGetQueryParameters
}
// NewAssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilderInternal instantiates a new AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilder and sets the default values.
func NewAssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilder) {
    m := &AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/deviceManagement/roleDefinitions/{unifiedRoleDefinition%2Did}/inheritsPermissionsFrom/{unifiedRoleDefinition%2Did1}/microsoft.graph.assignedPrincipals(transitive={transitive},directoryScopeType='{directoryScopeType}',directoryScopeId='{directoryScopeId}'){?transitive*,directoryScopeType*,directoryScopeId*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilder instantiates a new AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilder and sets the default values.
func NewAssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function assignedPrincipals
func (m *AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function assignedPrincipals
func (m *AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function assignedPrincipals
func (m *AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilder) Get(ctx context.Context, requestConfiguration *AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdRequestBuilderGetRequestConfiguration)(AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, CreateAssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(AssignedPrincipalsWithTransitiveWithDirectoryScopeTypeWithDirectoryScopeIdResponseable), nil
}
