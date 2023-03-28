package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedTenantsTenantGroupsManagedTenantsTenantSearchRequestBuilder provides operations to call the tenantSearch method.
type ManagedTenantsTenantGroupsManagedTenantsTenantSearchRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsTenantGroupsManagedTenantsTenantSearchRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsTenantGroupsManagedTenantsTenantSearchRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedTenantsTenantGroupsManagedTenantsTenantSearchRequestBuilderInternal instantiates a new ManagedTenantsTenantSearchRequestBuilder and sets the default values.
func NewManagedTenantsTenantGroupsManagedTenantsTenantSearchRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsTenantGroupsManagedTenantsTenantSearchRequestBuilder) {
    m := &ManagedTenantsTenantGroupsManagedTenantsTenantSearchRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/tenantGroups/managedTenants.tenantSearch", pathParameters),
    }
    return m
}
// NewManagedTenantsTenantGroupsManagedTenantsTenantSearchRequestBuilder instantiates a new ManagedTenantsTenantSearchRequestBuilder and sets the default values.
func NewManagedTenantsTenantGroupsManagedTenantsTenantSearchRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsTenantGroupsManagedTenantsTenantSearchRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsTenantGroupsManagedTenantsTenantSearchRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action tenantSearch
func (m *ManagedTenantsTenantGroupsManagedTenantsTenantSearchRequestBuilder) Post(ctx context.Context, body ManagedTenantsTenantGroupsManagedTenantsTenantSearchTenantSearchPostRequestBodyable, requestConfiguration *ManagedTenantsTenantGroupsManagedTenantsTenantSearchRequestBuilderPostRequestConfiguration)(ManagedTenantsTenantGroupsManagedTenantsTenantSearchTenantSearchResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagedTenantsTenantGroupsManagedTenantsTenantSearchTenantSearchResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagedTenantsTenantGroupsManagedTenantsTenantSearchTenantSearchResponseable), nil
}
// ToPostRequestInformation invoke action tenantSearch
func (m *ManagedTenantsTenantGroupsManagedTenantsTenantSearchRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedTenantsTenantGroupsManagedTenantsTenantSearchTenantSearchPostRequestBodyable, requestConfiguration *ManagedTenantsTenantGroupsManagedTenantsTenantSearchRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
