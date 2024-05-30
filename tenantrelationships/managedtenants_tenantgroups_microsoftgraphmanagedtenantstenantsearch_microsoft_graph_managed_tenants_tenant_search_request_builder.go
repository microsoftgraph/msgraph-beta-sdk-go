package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilder provides operations to call the tenantSearch method.
type ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilderInternal instantiates a new ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilder and sets the default values.
func NewManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilder) {
    m := &ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/tenantGroups/microsoft.graph.managedTenants.tenantSearch", pathParameters),
    }
    return m
}
// NewManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilder instantiates a new ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilder and sets the default values.
func NewManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilderInternal(urlParams, requestAdapter)
}
// Post searches for the specified managed tenants across tenant groups.
// Deprecated: This method is obsolete. Use PostAsTenantSearchPostResponse instead.
// returns a ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchTenantSearchResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilder) Post(ctx context.Context, body ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchTenantSearchPostRequestBodyable, requestConfiguration *ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilderPostRequestConfiguration)(ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchTenantSearchResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchTenantSearchResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchTenantSearchResponseable), nil
}
// PostAsTenantSearchPostResponse searches for the specified managed tenants across tenant groups.
// returns a ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchTenantSearchPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilder) PostAsTenantSearchPostResponse(ctx context.Context, body ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchTenantSearchPostRequestBodyable, requestConfiguration *ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilderPostRequestConfiguration)(ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchTenantSearchPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchTenantSearchPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchTenantSearchPostResponseable), nil
}
// ToPostRequestInformation searches for the specified managed tenants across tenant groups.
// returns a *RequestInformation when successful
func (m *ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilder) ToPostRequestInformation(ctx context.Context, body ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchTenantSearchPostRequestBodyable, requestConfiguration *ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilder when successful
func (m *ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilder) {
    return NewManagedtenantsTenantgroupsMicrosoftgraphmanagedtenantstenantsearchMicrosoftGraphManagedTenantsTenantSearchRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
