package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AuthorizationSystemsItemDataCollectionInfoRequestBuilder provides operations to manage the dataCollectionInfo property of the microsoft.graph.authorizationSystem entity.
type AuthorizationSystemsItemDataCollectionInfoRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthorizationSystemsItemDataCollectionInfoRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthorizationSystemsItemDataCollectionInfoRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthorizationSystemsItemDataCollectionInfoRequestBuilderGetQueryParameters defines how and whether Permissions Management collects data from the onboarded authorization system. Supports $filter (eq) as follows:  $filter=dataCollectionInfo/entitlements/permissionsModificationCapability and $filter=dataCollectionInfo/entitlements/status.
type AuthorizationSystemsItemDataCollectionInfoRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthorizationSystemsItemDataCollectionInfoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthorizationSystemsItemDataCollectionInfoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthorizationSystemsItemDataCollectionInfoRequestBuilderGetQueryParameters
}
// AuthorizationSystemsItemDataCollectionInfoRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthorizationSystemsItemDataCollectionInfoRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthorizationSystemsItemDataCollectionInfoRequestBuilderInternal instantiates a new AuthorizationSystemsItemDataCollectionInfoRequestBuilder and sets the default values.
func NewAuthorizationSystemsItemDataCollectionInfoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthorizationSystemsItemDataCollectionInfoRequestBuilder) {
    m := &AuthorizationSystemsItemDataCollectionInfoRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/authorizationSystems/{authorizationSystem%2Did}/dataCollectionInfo{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAuthorizationSystemsItemDataCollectionInfoRequestBuilder instantiates a new AuthorizationSystemsItemDataCollectionInfoRequestBuilder and sets the default values.
func NewAuthorizationSystemsItemDataCollectionInfoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthorizationSystemsItemDataCollectionInfoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthorizationSystemsItemDataCollectionInfoRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property dataCollectionInfo for external
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthorizationSystemsItemDataCollectionInfoRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthorizationSystemsItemDataCollectionInfoRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get defines how and whether Permissions Management collects data from the onboarded authorization system. Supports $filter (eq) as follows:  $filter=dataCollectionInfo/entitlements/permissionsModificationCapability and $filter=dataCollectionInfo/entitlements/status.
// returns a DataCollectionInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthorizationSystemsItemDataCollectionInfoRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthorizationSystemsItemDataCollectionInfoRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataCollectionInfoable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDataCollectionInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataCollectionInfoable), nil
}
// Patch update the navigation property dataCollectionInfo in external
// returns a DataCollectionInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthorizationSystemsItemDataCollectionInfoRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataCollectionInfoable, requestConfiguration *AuthorizationSystemsItemDataCollectionInfoRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataCollectionInfoable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDataCollectionInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataCollectionInfoable), nil
}
// ToDeleteRequestInformation delete navigation property dataCollectionInfo for external
// returns a *RequestInformation when successful
func (m *AuthorizationSystemsItemDataCollectionInfoRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthorizationSystemsItemDataCollectionInfoRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/external/authorizationSystems/{authorizationSystem%2Did}/dataCollectionInfo", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation defines how and whether Permissions Management collects data from the onboarded authorization system. Supports $filter (eq) as follows:  $filter=dataCollectionInfo/entitlements/permissionsModificationCapability and $filter=dataCollectionInfo/entitlements/status.
// returns a *RequestInformation when successful
func (m *AuthorizationSystemsItemDataCollectionInfoRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthorizationSystemsItemDataCollectionInfoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property dataCollectionInfo in external
// returns a *RequestInformation when successful
func (m *AuthorizationSystemsItemDataCollectionInfoRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DataCollectionInfoable, requestConfiguration *AuthorizationSystemsItemDataCollectionInfoRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/external/authorizationSystems/{authorizationSystem%2Did}/dataCollectionInfo", m.BaseRequestBuilder.PathParameters)
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
// returns a *AuthorizationSystemsItemDataCollectionInfoRequestBuilder when successful
func (m *AuthorizationSystemsItemDataCollectionInfoRequestBuilder) WithUrl(rawUrl string)(*AuthorizationSystemsItemDataCollectionInfoRequestBuilder) {
    return NewAuthorizationSystemsItemDataCollectionInfoRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
