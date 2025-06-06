// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilder provides operations to call the tenantSearch method.
type ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilderInternal instantiates a new ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilder and sets the default values.
func NewManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilder) {
    m := &ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/tenantGroups/microsoft.graph.managedTenants.tenantSearch", pathParameters),
    }
    return m
}
// NewManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilder instantiates a new ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilder and sets the default values.
func NewManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilderInternal(urlParams, requestAdapter)
}
// Post searches for the specified managed tenants across tenant groups.
// Deprecated: This method is obsolete. Use PostAsTenantSearchPostResponse instead.
// returns a ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilder) Post(ctx context.Context, body ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBodyable, requestConfiguration *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilderPostRequestConfiguration)(ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchResponseable), nil
}
// PostAsTenantSearchPostResponse searches for the specified managed tenants across tenant groups.
// returns a ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilder) PostAsTenantSearchPostResponse(ctx context.Context, body ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBodyable, requestConfiguration *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilderPostRequestConfiguration)(ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostResponseable), nil
}
// ToPostRequestInformation searches for the specified managed tenants across tenant groups.
// returns a *RequestInformation when successful
func (m *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchTenantSearchPostRequestBodyable, requestConfiguration *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilder when successful
func (m *ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilder) {
    return NewManagedTenantsTenantGroupsMicrosoftGraphManagedTenantsTenantSearchRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
