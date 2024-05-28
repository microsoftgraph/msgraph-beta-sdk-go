package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder provides operations to manage the managementTemplateCollections property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilderGetQueryParameters get managementTemplateCollections from tenantRelationships
type ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilderGetQueryParameters
}
// ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilderInternal instantiates a new ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder and sets the default values.
func NewManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder) {
    m := &ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managementTemplateCollections/{managementTemplateCollection%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder instantiates a new ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder and sets the default values.
func NewManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property managementTemplateCollections for tenantRelationships
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get managementTemplateCollections from tenantRelationships
// returns a ManagementTemplateCollectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateCollectionable, error) {
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
// ManagementTemplates provides operations to manage the managementTemplates property of the microsoft.graph.managedTenants.managementTemplateCollection entity.
// returns a *ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder when successful
func (m *ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder) ManagementTemplates()(*ManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilder) {
    return NewManagedtenantsManagementtemplatecollectionsItemManagementtemplatesManagementTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property managementTemplateCollections in tenantRelationships
// returns a ManagementTemplateCollectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder) Patch(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateCollectionable, requestConfiguration *ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilderPatchRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateCollectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property managementTemplateCollections for tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get managementTemplateCollections from tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property managementTemplateCollections in tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagementTemplateCollectionable, requestConfiguration *ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder when successful
func (m *ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder) {
    return NewManagedtenantsManagementtemplatecollectionsManagementTemplateCollectionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
