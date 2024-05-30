package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TenantRelationshipsRequestBuilder provides operations to manage the tenantRelationship singleton.
type TenantRelationshipsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TenantRelationshipsRequestBuilderGetQueryParameters get tenantRelationships
type TenantRelationshipsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TenantRelationshipsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TenantRelationshipsRequestBuilderGetQueryParameters
}
// TenantRelationshipsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TenantRelationshipsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTenantRelationshipsRequestBuilderInternal instantiates a new TenantRelationshipsRequestBuilder and sets the default values.
func NewTenantRelationshipsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsRequestBuilder) {
    m := &TenantRelationshipsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTenantRelationshipsRequestBuilder instantiates a new TenantRelationshipsRequestBuilder and sets the default values.
func NewTenantRelationshipsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TenantRelationshipsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTenantRelationshipsRequestBuilderInternal(urlParams, requestAdapter)
}
// DelegatedAdminCustomers provides operations to manage the delegatedAdminCustomers property of the microsoft.graph.tenantRelationship entity.
// returns a *DelegatedadmincustomersDelegatedAdminCustomersRequestBuilder when successful
func (m *TenantRelationshipsRequestBuilder) DelegatedAdminCustomers()(*DelegatedadmincustomersDelegatedAdminCustomersRequestBuilder) {
    return NewDelegatedadmincustomersDelegatedAdminCustomersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DelegatedAdminRelationships provides operations to manage the delegatedAdminRelationships property of the microsoft.graph.tenantRelationship entity.
// returns a *DelegatedadminrelationshipsDelegatedAdminRelationshipsRequestBuilder when successful
func (m *TenantRelationshipsRequestBuilder) DelegatedAdminRelationships()(*DelegatedadminrelationshipsDelegatedAdminRelationshipsRequestBuilder) {
    return NewDelegatedadminrelationshipsDelegatedAdminRelationshipsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FindTenantInformationByDomainNameWithDomainName provides operations to call the findTenantInformationByDomainName method.
// returns a *FindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilder when successful
func (m *TenantRelationshipsRequestBuilder) FindTenantInformationByDomainNameWithDomainName(domainName *string)(*FindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilder) {
    return NewFindtenantinformationbydomainnamewithdomainnameFindTenantInformationByDomainNameWithDomainNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, domainName)
}
// FindTenantInformationByTenantIdWithTenantId provides operations to call the findTenantInformationByTenantId method.
// returns a *FindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilder when successful
func (m *TenantRelationshipsRequestBuilder) FindTenantInformationByTenantIdWithTenantId(tenantId *string)(*FindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilder) {
    return NewFindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, tenantId)
}
// Get get tenantRelationships
// returns a TenantRelationshipable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TenantRelationshipsRequestBuilder) Get(ctx context.Context, requestConfiguration *TenantRelationshipsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantRelationshipable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTenantRelationshipFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantRelationshipable), nil
}
// ManagedTenants provides operations to manage the managedTenants property of the microsoft.graph.tenantRelationship entity.
// returns a *ManagedtenantsManagedTenantsRequestBuilder when successful
func (m *TenantRelationshipsRequestBuilder) ManagedTenants()(*ManagedtenantsManagedTenantsRequestBuilder) {
    return NewManagedtenantsManagedTenantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MultiTenantOrganization provides operations to manage the multiTenantOrganization property of the microsoft.graph.tenantRelationship entity.
// returns a *MultitenantorganizationMultiTenantOrganizationRequestBuilder when successful
func (m *TenantRelationshipsRequestBuilder) MultiTenantOrganization()(*MultitenantorganizationMultiTenantOrganizationRequestBuilder) {
    return NewMultitenantorganizationMultiTenantOrganizationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update tenantRelationships
// returns a TenantRelationshipable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TenantRelationshipsRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantRelationshipable, requestConfiguration *TenantRelationshipsRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantRelationshipable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTenantRelationshipFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantRelationshipable), nil
}
// ToGetRequestInformation get tenantRelationships
// returns a *RequestInformation when successful
func (m *TenantRelationshipsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TenantRelationshipsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update tenantRelationships
// returns a *RequestInformation when successful
func (m *TenantRelationshipsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantRelationshipable, requestConfiguration *TenantRelationshipsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TenantRelationshipsRequestBuilder when successful
func (m *TenantRelationshipsRequestBuilder) WithUrl(rawUrl string)(*TenantRelationshipsRequestBuilder) {
    return NewTenantRelationshipsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
