package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder provides operations to manage the resourceNamespaces property of the microsoft.graph.unifiedRbacApplication entity.
type ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetQueryParameters resource that represents a collection of related actions.
type ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetQueryParameters
}
// ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderInternal instantiates a new ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder and sets the default values.
func NewExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) {
    m := &ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/exchange/resourceNamespaces/{unifiedRbacResourceNamespace%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder instantiates a new ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder and sets the default values.
func NewExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property resourceNamespaces for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get resource that represents a collection of related actions.
// returns a UnifiedRbacResourceNamespaceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable), nil
}
// ImportResourceActions provides operations to call the importResourceActions method.
// returns a *ExchangeResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilder when successful
func (m *ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) ImportResourceActions()(*ExchangeResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilder) {
    return NewExchangeResourcenamespacesItemImportresourceactionsImportResourceActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property resourceNamespaces in roleManagement
// returns a UnifiedRbacResourceNamespaceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, requestConfiguration *ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRbacResourceNamespaceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable), nil
}
// ResourceActions provides operations to manage the resourceActions property of the microsoft.graph.unifiedRbacResourceNamespace entity.
// returns a *ExchangeResourcenamespacesItemResourceactionsResourceActionsRequestBuilder when successful
func (m *ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) ResourceActions()(*ExchangeResourcenamespacesItemResourceactionsResourceActionsRequestBuilder) {
    return NewExchangeResourcenamespacesItemResourceactionsResourceActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property resourceNamespaces for roleManagement
// returns a *RequestInformation when successful
func (m *ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation resource that represents a collection of related actions.
// returns a *RequestInformation when successful
func (m *ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property resourceNamespaces in roleManagement
// returns a *RequestInformation when successful
func (m *ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceNamespaceable, requestConfiguration *ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder when successful
func (m *ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) WithUrl(rawUrl string)(*ExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder) {
    return NewExchangeResourcenamespacesUnifiedRbacResourceNamespaceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
