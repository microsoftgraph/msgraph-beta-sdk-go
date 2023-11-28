package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilder provides operations to manage the managedTenantTicketingEndpoints property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderGetQueryParameters get managedTenantTicketingEndpoints from tenantRelationships
type ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderGetQueryParameters
}
// ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderInternal instantiates a new ManagedTenantTicketingEndpointItemRequestBuilder and sets the default values.
func NewManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilder) {
    m := &ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managedTenantTicketingEndpoints/{managedTenantTicketingEndpoint%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilder instantiates a new ManagedTenantTicketingEndpointItemRequestBuilder and sets the default values.
func NewManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property managedTenantTicketingEndpoints for tenantRelationships
func (m *ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get managedTenantTicketingEndpoints from tenantRelationships
func (m *ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantTicketingEndpointable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantTicketingEndpointFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantTicketingEndpointable), nil
}
// Patch update the navigation property managedTenantTicketingEndpoints in tenantRelationships
func (m *ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantTicketingEndpointable, requestConfiguration *ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantTicketingEndpointable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantTicketingEndpointFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantTicketingEndpointable), nil
}
// ToDeleteRequestInformation delete navigation property managedTenantTicketingEndpoints for tenantRelationships
func (m *ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get managedTenantTicketingEndpoints from tenantRelationships
func (m *ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property managedTenantTicketingEndpoints in tenantRelationships
func (m *ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantTicketingEndpointable, requestConfiguration *ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilder) {
    return NewManagedTenantsManagedTenantTicketingEndpointsManagedTenantTicketingEndpointItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
