package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// FindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilder provides operations to call the findTenantInformationByTenantId method.
type FindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilderInternal instantiates a new FindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilder and sets the default values.
func NewFindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, tenantId *string)(*FindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilder) {
    m := &FindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/findTenantInformationByTenantId(tenantId='{tenantId}')", pathParameters),
    }
    if tenantId != nil {
        m.BaseRequestBuilder.PathParameters["tenantId"] = *tenantId
    }
    return m
}
// NewFindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilder instantiates a new FindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilder and sets the default values.
func NewFindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get given a tenant ID, search for a tenant and read its tenantInformation. You can use this API to validate tenant information and use their tenantId to configure cross-tenant cross-tenant access settings between you and the tenant.
// returns a TenantInformationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tenantrelationship-findtenantinformationbytenantid?view=graph-rest-beta
func (m *FindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilder) Get(ctx context.Context, requestConfiguration *FindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantInformationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTenantInformationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TenantInformationable), nil
}
// ToGetRequestInformation given a tenant ID, search for a tenant and read its tenantInformation. You can use this API to validate tenant information and use their tenantId to configure cross-tenant cross-tenant access settings between you and the tenant.
// returns a *RequestInformation when successful
func (m *FindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilder when successful
func (m *FindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilder) WithUrl(rawUrl string)(*FindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilder) {
    return NewFindtenantinformationbytenantidwithtenantidFindTenantInformationByTenantIdWithTenantIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
