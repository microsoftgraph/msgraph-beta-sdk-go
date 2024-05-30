package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder provides operations to manage the delegatedPermissionClassifications property of the microsoft.graph.servicePrincipal entity.
type ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderGetQueryParameters the permission classifications for delegated permissions exposed by the app that this service principal represents. Supports $expand.
type ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderGetQueryParameters
}
// ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderInternal instantiates a new ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder and sets the default values.
func NewItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder) {
    m := &ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/delegatedPermissionClassifications/{delegatedPermissionClassification%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder instantiates a new ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder and sets the default values.
func NewItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a delegatedPermissionClassification which had previously been set for a delegated permission.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-delete-delegatedpermissionclassifications?view=graph-rest-beta
func (m *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the permission classifications for delegated permissions exposed by the app that this service principal represents. Supports $expand.
// returns a DelegatedPermissionClassificationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedPermissionClassificationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDelegatedPermissionClassificationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedPermissionClassificationable), nil
}
// Patch update the navigation property delegatedPermissionClassifications in servicePrincipals
// returns a DelegatedPermissionClassificationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedPermissionClassificationable, requestConfiguration *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedPermissionClassificationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDelegatedPermissionClassificationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedPermissionClassificationable), nil
}
// ToDeleteRequestInformation deletes a delegatedPermissionClassification which had previously been set for a delegated permission.
// returns a *RequestInformation when successful
func (m *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the permission classifications for delegated permissions exposed by the app that this service principal represents. Supports $expand.
// returns a *RequestInformation when successful
func (m *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property delegatedPermissionClassifications in servicePrincipals
// returns a *RequestInformation when successful
func (m *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedPermissionClassificationable, requestConfiguration *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder when successful
func (m *ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder) WithUrl(rawUrl string)(*ItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder) {
    return NewItemDelegatedpermissionclassificationsDelegatedPermissionClassificationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
