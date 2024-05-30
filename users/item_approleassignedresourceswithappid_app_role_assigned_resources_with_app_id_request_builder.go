package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilder provides operations to manage the appRoleAssignedResources property of the microsoft.graph.user entity.
type ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilderGetQueryParameters get appRoleAssignedResources from users
type ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilderGetQueryParameters
}
// NewItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilderInternal instantiates a new ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilder and sets the default values.
func NewItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, appId *string)(*ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilder) {
    m := &ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/appRoleAssignedResources(appId='{appId}'){?%24expand,%24select}", pathParameters),
    }
    if appId != nil {
        m.BaseRequestBuilder.PathParameters["appId"] = *appId
    }
    return m
}
// NewItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilder instantiates a new ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilder and sets the default values.
func NewItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get appRoleAssignedResources from users
// returns a ServicePrincipalable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable), nil
}
// ToGetRequestInformation get appRoleAssignedResources from users
// returns a *RequestInformation when successful
func (m *ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilder when successful
func (m *ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilder) WithUrl(rawUrl string)(*ItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilder) {
    return NewItemApproleassignedresourceswithappidAppRoleAssignedResourcesWithAppIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
