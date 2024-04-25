package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PermissionsManagementScheduledPermissionsRequestsRequestBuilder provides operations to manage the scheduledPermissionsRequests property of the microsoft.graph.permissionsManagement entity.
type PermissionsManagementScheduledPermissionsRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PermissionsManagementScheduledPermissionsRequestsRequestBuilderGetQueryParameters represents a permissions request that Permissions Management uses to manage permissions for an identity on resources in the authorization system. This request can be granted, rejected or canceled by identities in Permissions Management.
type PermissionsManagementScheduledPermissionsRequestsRequestBuilderGetQueryParameters struct {
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
// PermissionsManagementScheduledPermissionsRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsManagementScheduledPermissionsRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionsManagementScheduledPermissionsRequestsRequestBuilderGetQueryParameters
}
// PermissionsManagementScheduledPermissionsRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsManagementScheduledPermissionsRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPermissionsManagementScheduledPermissionsRequestsRequestBuilderInternal instantiates a new PermissionsManagementScheduledPermissionsRequestsRequestBuilder and sets the default values.
func NewPermissionsManagementScheduledPermissionsRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsManagementScheduledPermissionsRequestsRequestBuilder) {
    m := &PermissionsManagementScheduledPermissionsRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/permissionsManagement/scheduledPermissionsRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewPermissionsManagementScheduledPermissionsRequestsRequestBuilder instantiates a new PermissionsManagementScheduledPermissionsRequestsRequestBuilder and sets the default values.
func NewPermissionsManagementScheduledPermissionsRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsManagementScheduledPermissionsRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionsManagementScheduledPermissionsRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *PermissionsManagementScheduledPermissionsRequestsFilterByCurrentUserWithOnRequestBuilder when successful
func (m *PermissionsManagementScheduledPermissionsRequestsRequestBuilder) FilterByCurrentUserWithOn(on *string)(*PermissionsManagementScheduledPermissionsRequestsFilterByCurrentUserWithOnRequestBuilder) {
    return NewPermissionsManagementScheduledPermissionsRequestsFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get represents a permissions request that Permissions Management uses to manage permissions for an identity on resources in the authorization system. This request can be granted, rejected or canceled by identities in Permissions Management.
// returns a ScheduledPermissionsRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsManagementScheduledPermissionsRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionsManagementScheduledPermissionsRequestsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledPermissionsRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateScheduledPermissionsRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledPermissionsRequestCollectionResponseable), nil
}
// Post create new navigation property to scheduledPermissionsRequests for identityGovernance
// returns a ScheduledPermissionsRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PermissionsManagementScheduledPermissionsRequestsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledPermissionsRequestable, requestConfiguration *PermissionsManagementScheduledPermissionsRequestsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledPermissionsRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateScheduledPermissionsRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledPermissionsRequestable), nil
}
// ToGetRequestInformation represents a permissions request that Permissions Management uses to manage permissions for an identity on resources in the authorization system. This request can be granted, rejected or canceled by identities in Permissions Management.
// returns a *RequestInformation when successful
func (m *PermissionsManagementScheduledPermissionsRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionsManagementScheduledPermissionsRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to scheduledPermissionsRequests for identityGovernance
// returns a *RequestInformation when successful
func (m *PermissionsManagementScheduledPermissionsRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ScheduledPermissionsRequestable, requestConfiguration *PermissionsManagementScheduledPermissionsRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *PermissionsManagementScheduledPermissionsRequestsRequestBuilder when successful
func (m *PermissionsManagementScheduledPermissionsRequestsRequestBuilder) WithUrl(rawUrl string)(*PermissionsManagementScheduledPermissionsRequestsRequestBuilder) {
    return NewPermissionsManagementScheduledPermissionsRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
