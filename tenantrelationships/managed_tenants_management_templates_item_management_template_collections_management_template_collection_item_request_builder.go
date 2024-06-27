package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilder provides operations to manage the managementTemplateCollections property of the microsoft.graph.managedTenants.managementTemplate entity.
type ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilderGetQueryParameters get managementTemplateCollections from tenantRelationships
type ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilderGetQueryParameters
}
// NewManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilderInternal instantiates a new ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilder and sets the default values.
func NewManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilder) {
    m := &ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managementTemplates/{managementTemplate%2Did}/managementTemplateCollections/{managementTemplateCollection%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilder instantiates a new ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilder and sets the default values.
func NewManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get managementTemplateCollections from tenantRelationships
// returns a ManagementTemplateCollectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateCollectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagementTemplateCollectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateCollectionable), nil
}
// ToGetRequestInformation get managementTemplateCollections from tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilder when successful
func (m *ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilder) WithUrl(rawUrl string)(*ManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilder) {
    return NewManagedTenantsManagementTemplatesItemManagementTemplateCollectionsManagementTemplateCollectionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
