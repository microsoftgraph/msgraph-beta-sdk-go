package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MultitenantorganizationMultiTenantOrganizationRequestBuilder provides operations to manage the multiTenantOrganization property of the microsoft.graph.tenantRelationship entity.
type MultitenantorganizationMultiTenantOrganizationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MultitenantorganizationMultiTenantOrganizationRequestBuilderGetQueryParameters get properties of the multitenant organization.
type MultitenantorganizationMultiTenantOrganizationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MultitenantorganizationMultiTenantOrganizationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MultitenantorganizationMultiTenantOrganizationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MultitenantorganizationMultiTenantOrganizationRequestBuilderGetQueryParameters
}
// MultitenantorganizationMultiTenantOrganizationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MultitenantorganizationMultiTenantOrganizationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMultitenantorganizationMultiTenantOrganizationRequestBuilderInternal instantiates a new MultitenantorganizationMultiTenantOrganizationRequestBuilder and sets the default values.
func NewMultitenantorganizationMultiTenantOrganizationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MultitenantorganizationMultiTenantOrganizationRequestBuilder) {
    m := &MultitenantorganizationMultiTenantOrganizationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/multiTenantOrganization{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMultitenantorganizationMultiTenantOrganizationRequestBuilder instantiates a new MultitenantorganizationMultiTenantOrganizationRequestBuilder and sets the default values.
func NewMultitenantorganizationMultiTenantOrganizationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MultitenantorganizationMultiTenantOrganizationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMultitenantorganizationMultiTenantOrganizationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get properties of the multitenant organization.
// returns a MultiTenantOrganizationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/multitenantorganization-get?view=graph-rest-beta
func (m *MultitenantorganizationMultiTenantOrganizationRequestBuilder) Get(ctx context.Context, requestConfiguration *MultitenantorganizationMultiTenantOrganizationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MultiTenantOrganizationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMultiTenantOrganizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MultiTenantOrganizationable), nil
}
// JoinRequest provides operations to manage the joinRequest property of the microsoft.graph.multiTenantOrganization entity.
// returns a *MultitenantorganizationJoinrequestJoinRequestRequestBuilder when successful
func (m *MultitenantorganizationMultiTenantOrganizationRequestBuilder) JoinRequest()(*MultitenantorganizationJoinrequestJoinRequestRequestBuilder) {
    return NewMultitenantorganizationJoinrequestJoinRequestRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a multi-tenant organization.
// returns a MultiTenantOrganizationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/multitenantorganization-update?view=graph-rest-beta
func (m *MultitenantorganizationMultiTenantOrganizationRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MultiTenantOrganizationable, requestConfiguration *MultitenantorganizationMultiTenantOrganizationRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MultiTenantOrganizationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMultiTenantOrganizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MultiTenantOrganizationable), nil
}
// Tenants provides operations to manage the tenants property of the microsoft.graph.multiTenantOrganization entity.
// returns a *MultitenantorganizationTenantsRequestBuilder when successful
func (m *MultitenantorganizationMultiTenantOrganizationRequestBuilder) Tenants()(*MultitenantorganizationTenantsRequestBuilder) {
    return NewMultitenantorganizationTenantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get properties of the multitenant organization.
// returns a *RequestInformation when successful
func (m *MultitenantorganizationMultiTenantOrganizationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MultitenantorganizationMultiTenantOrganizationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a multi-tenant organization.
// returns a *RequestInformation when successful
func (m *MultitenantorganizationMultiTenantOrganizationRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MultiTenantOrganizationable, requestConfiguration *MultitenantorganizationMultiTenantOrganizationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MultitenantorganizationMultiTenantOrganizationRequestBuilder when successful
func (m *MultitenantorganizationMultiTenantOrganizationRequestBuilder) WithUrl(rawUrl string)(*MultitenantorganizationMultiTenantOrganizationRequestBuilder) {
    return NewMultitenantorganizationMultiTenantOrganizationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
