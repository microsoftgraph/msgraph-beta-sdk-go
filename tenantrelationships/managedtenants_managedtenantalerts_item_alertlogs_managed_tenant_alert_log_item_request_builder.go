package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilder provides operations to manage the alertLogs property of the microsoft.graph.managedTenants.managedTenantAlert entity.
type ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilderGetQueryParameters get alertLogs from tenantRelationships
type ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilderGetQueryParameters
}
// NewManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilderInternal instantiates a new ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilder and sets the default values.
func NewManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilder) {
    m := &ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/managedTenantAlerts/{managedTenantAlert%2Did}/alertLogs/{managedTenantAlertLog%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilder instantiates a new ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilder and sets the default values.
func NewManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get alertLogs from tenantRelationships
// returns a ManagedTenantAlertLogable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertLogable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateManagedTenantAlertLogFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.ManagedTenantAlertLogable), nil
}
// ToGetRequestInformation get alertLogs from tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilder when successful
func (m *ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilder) {
    return NewManagedtenantsManagedtenantalertsItemAlertlogsManagedTenantAlertLogItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
