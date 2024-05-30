package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder provides operations to manage the windowsProtectionStates property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilderGetQueryParameters get a list of the windowsProtectionState objects and their properties.
type ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilderGetQueryParameters struct {
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
// ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilderGetQueryParameters
}
// ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWindowsProtectionStateId provides operations to manage the windowsProtectionStates property of the microsoft.graph.managedTenants.managedTenant entity.
// returns a *ManagedtenantsWindowsprotectionstatesWindowsProtectionStateItemRequestBuilder when successful
func (m *ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder) ByWindowsProtectionStateId(windowsProtectionStateId string)(*ManagedtenantsWindowsprotectionstatesWindowsProtectionStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if windowsProtectionStateId != "" {
        urlTplParams["windowsProtectionState%2Did"] = windowsProtectionStateId
    }
    return NewManagedtenantsWindowsprotectionstatesWindowsProtectionStateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilderInternal instantiates a new ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder and sets the default values.
func NewManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder) {
    m := &ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/windowsProtectionStates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder instantiates a new ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder and sets the default values.
func NewManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedtenantsWindowsprotectionstatesCountRequestBuilder when successful
func (m *ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder) Count()(*ManagedtenantsWindowsprotectionstatesCountRequestBuilder) {
    return NewManagedtenantsWindowsprotectionstatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the windowsProtectionState objects and their properties.
// returns a WindowsProtectionStateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/managedtenants-managedtenant-list-windowsprotectionstates?view=graph-rest-beta
func (m *ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.WindowsProtectionStateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateWindowsProtectionStateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.WindowsProtectionStateCollectionResponseable), nil
}
// Post create new navigation property to windowsProtectionStates for tenantRelationships
// returns a WindowsProtectionStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder) Post(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.WindowsProtectionStateable, requestConfiguration *ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.WindowsProtectionStateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateWindowsProtectionStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.WindowsProtectionStateable), nil
}
// ToGetRequestInformation get a list of the windowsProtectionState objects and their properties.
// returns a *RequestInformation when successful
func (m *ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to windowsProtectionStates for tenantRelationships
// returns a *RequestInformation when successful
func (m *ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.WindowsProtectionStateable, requestConfiguration *ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder when successful
func (m *ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder) WithUrl(rawUrl string)(*ManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder) {
    return NewManagedtenantsWindowsprotectionstatesWindowsProtectionStatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
