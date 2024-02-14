package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder provides operations to manage the tenantsCustomizedInformation property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderGetQueryParameters read the properties and relationships of a tenantCustomizedInformation object.
type ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderGetQueryParameters
}
// ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderInternal instantiates a new ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder and sets the default values.
func NewManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) {
    m := &ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/tenantsCustomizedInformation/{tenantCustomizedInformation%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder instantiates a new ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder and sets the default values.
func NewManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property tenantsCustomizedInformation for tenantRelationships
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a tenantCustomizedInformation object.
// returns a TenantCustomizedInformationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/managedtenants-tenantcustomizedinformation-get?view=graph-rest-1.0
func (m *ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantCustomizedInformationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateTenantCustomizedInformationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantCustomizedInformationable), nil
}
// Patch update the properties of a tenantCustomizedInformation object.
// returns a TenantCustomizedInformationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/managedtenants-tenantcustomizedinformation-update?view=graph-rest-1.0
func (m *ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantCustomizedInformationable, requestConfiguration *ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantCustomizedInformationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateTenantCustomizedInformationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantCustomizedInformationable), nil
}
// ToDeleteRequestInformation delete navigation property tenantsCustomizedInformation for tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/tenantRelationships/managedTenants/tenantsCustomizedInformation/{tenantCustomizedInformation%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a tenantCustomizedInformation object.
// returns a *RequestInformation when successful
func (m *ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a tenantCustomizedInformation object.
// returns a *RequestInformation when successful
func (m *ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.TenantCustomizedInformationable, requestConfiguration *ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/tenantRelationships/managedTenants/tenantsCustomizedInformation/{tenantCustomizedInformation%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder when successful
func (m *ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder) {
    return NewManagedTenantsTenantsCustomizedInformationTenantCustomizedInformationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
