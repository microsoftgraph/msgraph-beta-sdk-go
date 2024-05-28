package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder provides operations to manage the resourceScope property of the microsoft.graph.unifiedRbacResourceAction entity.
type EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderGetQueryParameters get resourceScope from roleManagement
type EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderGetQueryParameters
}
// EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderInternal instantiates a new EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder and sets the default values.
func NewEnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder) {
    m := &EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/enterpriseApps/{rbacApplication%2Did}/resourceNamespaces/{unifiedRbacResourceNamespace%2Did}/resourceActions/{unifiedRbacResourceAction%2Did}/resourceScope{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder instantiates a new EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder and sets the default values.
func NewEnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property resourceScope for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder) Delete(ctx context.Context, requestConfiguration *EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get resourceScope from roleManagement
// returns a UnifiedRbacResourceScopeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder) Get(ctx context.Context, requestConfiguration *EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceScopeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRbacResourceScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceScopeable), nil
}
// Patch update the navigation property resourceScope in roleManagement
// returns a UnifiedRbacResourceScopeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceScopeable, requestConfiguration *EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceScopeable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUnifiedRbacResourceScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceScopeable), nil
}
// ToDeleteRequestInformation delete navigation property resourceScope for roleManagement
// returns a *RequestInformation when successful
func (m *EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get resourceScope from roleManagement
// returns a *RequestInformation when successful
func (m *EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property resourceScope in roleManagement
// returns a *RequestInformation when successful
func (m *EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UnifiedRbacResourceScopeable, requestConfiguration *EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder when successful
func (m *EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder) WithUrl(rawUrl string)(*EnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder) {
    return NewEnterpriseappsItemResourcenamespacesItemResourceactionsItemResourcescopeResourceScopeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
