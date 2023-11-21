package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilder provides operations to manage the managementTemplateCollections property of the microsoft.graph.managedTenants.managementTemplate entity.
type ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilderGetQueryParameters get managementTemplateCollections from tenantRelationships
type ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilderGetQueryParameters
}
// ByManagementTemplateCollectionId provides operations to manage the managementTemplateCollections property of the microsoft.graph.managedTenants.managementTemplate entity.
func (m *ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilder) ByManagementTemplateCollectionId(managementTemplateCollectionId string)(*ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managementTemplateCollectionId != "" {
        urlTplParams["managementTemplateCollection%2Did"] = managementTemplateCollectionId
    }
    return NewManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilderInternal instantiates a new ManagementTemplateCollectionsRequestBuilder and sets the default values.
func NewManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilder) {
    m := &ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managementTemplates/{managementTemplate%2Did}/managementTemplateCollections{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilder instantiates a new ManagementTemplateCollectionsRequestBuilder and sets the default values.
func NewManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilder) Count()(*ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsCountRequestBuilder) {
    return NewManagedTenantsManagementTemplatesItemManagementTemplateCollectionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get managementTemplateCollections from tenantRelationships
func (m *ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateCollectionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateCollectionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateCollectionCollectionResponseable), nil
}
// ToGetRequestInformation get managementTemplateCollections from tenantRelationships
func (m *ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilder) {
    return NewManagedTenantsManagementTemplatesItemManagementTemplateCollectionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
