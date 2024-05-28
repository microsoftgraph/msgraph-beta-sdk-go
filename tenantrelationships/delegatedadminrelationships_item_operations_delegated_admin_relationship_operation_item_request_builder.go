package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder provides operations to manage the operations property of the microsoft.graph.delegatedAdminRelationship entity.
type DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderGetQueryParameters read the properties of a delegatedAdminRelationshipOperation object.
type DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderGetQueryParameters
}
// DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderInternal instantiates a new DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder and sets the default values.
func NewDelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder) {
    m := &DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/delegatedAdminRelationships/{delegatedAdminRelationship%2Did}/operations/{delegatedAdminRelationshipOperation%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder instantiates a new DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder and sets the default values.
func NewDelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property operations for tenantRelationships
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties of a delegatedAdminRelationshipOperation object.
// returns a DelegatedAdminRelationshipOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/delegatedadminrelationshipoperation-get?view=graph-rest-beta
func (m *DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipOperationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDelegatedAdminRelationshipOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipOperationable), nil
}
// Patch update the navigation property operations in tenantRelationships
// returns a DelegatedAdminRelationshipOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipOperationable, requestConfiguration *DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipOperationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDelegatedAdminRelationshipOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipOperationable), nil
}
// ToDeleteRequestInformation delete navigation property operations for tenantRelationships
// returns a *RequestInformation when successful
func (m *DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties of a delegatedAdminRelationshipOperation object.
// returns a *RequestInformation when successful
func (m *DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property operations in tenantRelationships
// returns a *RequestInformation when successful
func (m *DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DelegatedAdminRelationshipOperationable, requestConfiguration *DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder when successful
func (m *DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder) WithUrl(rawUrl string)(*DelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder) {
    return NewDelegatedadminrelationshipsItemOperationsDelegatedAdminRelationshipOperationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
