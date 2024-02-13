package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CloudPCRoleAssignmentsItemAppScopesRequestBuilder provides operations to manage the appScopes property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
type CloudPCRoleAssignmentsItemAppScopesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CloudPCRoleAssignmentsItemAppScopesRequestBuilderGetQueryParameters read-only collection with details of the app specific scopes when the assignment scopes are app specific. Containment entity. Read-only.
type CloudPCRoleAssignmentsItemAppScopesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// CloudPCRoleAssignmentsItemAppScopesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudPCRoleAssignmentsItemAppScopesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CloudPCRoleAssignmentsItemAppScopesRequestBuilderGetQueryParameters
}
// CloudPCRoleAssignmentsItemAppScopesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudPCRoleAssignmentsItemAppScopesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAppScopeId provides operations to manage the appScopes property of the microsoft.graph.unifiedRoleAssignmentMultiple entity.
// returns a *CloudPCRoleAssignmentsItemAppScopesAppScopeItemRequestBuilder when successful
func (m *CloudPCRoleAssignmentsItemAppScopesRequestBuilder) ByAppScopeId(appScopeId string)(*CloudPCRoleAssignmentsItemAppScopesAppScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if appScopeId != "" {
        urlTplParams["appScope%2Did"] = appScopeId
    }
    return NewCloudPCRoleAssignmentsItemAppScopesAppScopeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCloudPCRoleAssignmentsItemAppScopesRequestBuilderInternal instantiates a new CloudPCRoleAssignmentsItemAppScopesRequestBuilder and sets the default values.
func NewCloudPCRoleAssignmentsItemAppScopesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudPCRoleAssignmentsItemAppScopesRequestBuilder) {
    m := &CloudPCRoleAssignmentsItemAppScopesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/cloudPC/roleAssignments/{unifiedRoleAssignmentMultiple%2Did}/appScopes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCloudPCRoleAssignmentsItemAppScopesRequestBuilder instantiates a new CloudPCRoleAssignmentsItemAppScopesRequestBuilder and sets the default values.
func NewCloudPCRoleAssignmentsItemAppScopesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudPCRoleAssignmentsItemAppScopesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudPCRoleAssignmentsItemAppScopesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CloudPCRoleAssignmentsItemAppScopesCountRequestBuilder when successful
func (m *CloudPCRoleAssignmentsItemAppScopesRequestBuilder) Count()(*CloudPCRoleAssignmentsItemAppScopesCountRequestBuilder) {
    return NewCloudPCRoleAssignmentsItemAppScopesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only collection with details of the app specific scopes when the assignment scopes are app specific. Containment entity. Read-only.
// returns a AppScopeCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudPCRoleAssignmentsItemAppScopesRequestBuilder) Get(ctx context.Context, requestConfiguration *CloudPCRoleAssignmentsItemAppScopesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppScopeCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppScopeCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppScopeCollectionResponseable), nil
}
// Post create new navigation property to appScopes for roleManagement
// returns a AppScopeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudPCRoleAssignmentsItemAppScopesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppScopeable, requestConfiguration *CloudPCRoleAssignmentsItemAppScopesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppScopeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppScopeable), nil
}
// ToGetRequestInformation read-only collection with details of the app specific scopes when the assignment scopes are app specific. Containment entity. Read-only.
// returns a *RequestInformation when successful
func (m *CloudPCRoleAssignmentsItemAppScopesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CloudPCRoleAssignmentsItemAppScopesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to appScopes for roleManagement
// returns a *RequestInformation when successful
func (m *CloudPCRoleAssignmentsItemAppScopesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppScopeable, requestConfiguration *CloudPCRoleAssignmentsItemAppScopesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/roleManagement/cloudPC/roleAssignments/{unifiedRoleAssignmentMultiple%2Did}/appScopes", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CloudPCRoleAssignmentsItemAppScopesRequestBuilder when successful
func (m *CloudPCRoleAssignmentsItemAppScopesRequestBuilder) WithUrl(rawUrl string)(*CloudPCRoleAssignmentsItemAppScopesRequestBuilder) {
    return NewCloudPCRoleAssignmentsItemAppScopesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
