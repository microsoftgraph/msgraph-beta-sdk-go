package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DefenderResourceNamespacesItemResourceActionsRequestBuilder provides operations to manage the resourceActions property of the microsoft.graph.unifiedRbacResourceNamespace entity.
type DefenderResourceNamespacesItemResourceActionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DefenderResourceNamespacesItemResourceActionsRequestBuilderGetQueryParameters operations that an authorized principal is allowed to perform.
type DefenderResourceNamespacesItemResourceActionsRequestBuilderGetQueryParameters struct {
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
// DefenderResourceNamespacesItemResourceActionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DefenderResourceNamespacesItemResourceActionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DefenderResourceNamespacesItemResourceActionsRequestBuilderGetQueryParameters
}
// DefenderResourceNamespacesItemResourceActionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DefenderResourceNamespacesItemResourceActionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUnifiedRbacResourceActionId provides operations to manage the resourceActions property of the microsoft.graph.unifiedRbacResourceNamespace entity.
// returns a *DefenderResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilder when successful
func (m *DefenderResourceNamespacesItemResourceActionsRequestBuilder) ByUnifiedRbacResourceActionId(unifiedRbacResourceActionId string)(*DefenderResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if unifiedRbacResourceActionId != "" {
        urlTplParams["unifiedRbacResourceAction%2Did"] = unifiedRbacResourceActionId
    }
    return NewDefenderResourceNamespacesItemResourceActionsUnifiedRbacResourceActionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDefenderResourceNamespacesItemResourceActionsRequestBuilderInternal instantiates a new DefenderResourceNamespacesItemResourceActionsRequestBuilder and sets the default values.
func NewDefenderResourceNamespacesItemResourceActionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DefenderResourceNamespacesItemResourceActionsRequestBuilder) {
    m := &DefenderResourceNamespacesItemResourceActionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/defender/resourceNamespaces/{unifiedRbacResourceNamespace%2Did}/resourceActions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDefenderResourceNamespacesItemResourceActionsRequestBuilder instantiates a new DefenderResourceNamespacesItemResourceActionsRequestBuilder and sets the default values.
func NewDefenderResourceNamespacesItemResourceActionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DefenderResourceNamespacesItemResourceActionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDefenderResourceNamespacesItemResourceActionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DefenderResourceNamespacesItemResourceActionsCountRequestBuilder when successful
func (m *DefenderResourceNamespacesItemResourceActionsRequestBuilder) Count()(*DefenderResourceNamespacesItemResourceActionsCountRequestBuilder) {
    return NewDefenderResourceNamespacesItemResourceActionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get operations that an authorized principal is allowed to perform.
// returns a UnifiedRbacResourceActionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DefenderResourceNamespacesItemResourceActionsRequestBuilder) Get(ctx context.Context, requestConfiguration *DefenderResourceNamespacesItemResourceActionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceActionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRbacResourceActionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceActionCollectionResponseable), nil
}
// Post create new navigation property to resourceActions for roleManagement
// returns a UnifiedRbacResourceActionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DefenderResourceNamespacesItemResourceActionsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceActionable, requestConfiguration *DefenderResourceNamespacesItemResourceActionsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceActionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRbacResourceActionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceActionable), nil
}
// ToGetRequestInformation operations that an authorized principal is allowed to perform.
// returns a *RequestInformation when successful
func (m *DefenderResourceNamespacesItemResourceActionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DefenderResourceNamespacesItemResourceActionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to resourceActions for roleManagement
// returns a *RequestInformation when successful
func (m *DefenderResourceNamespacesItemResourceActionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceActionable, requestConfiguration *DefenderResourceNamespacesItemResourceActionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DefenderResourceNamespacesItemResourceActionsRequestBuilder when successful
func (m *DefenderResourceNamespacesItemResourceActionsRequestBuilder) WithUrl(rawUrl string)(*DefenderResourceNamespacesItemResourceActionsRequestBuilder) {
    return NewDefenderResourceNamespacesItemResourceActionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
