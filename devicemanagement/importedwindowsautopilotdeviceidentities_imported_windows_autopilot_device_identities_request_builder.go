package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder provides operations to manage the importedWindowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
type ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetQueryParameters collection of imported Windows autopilot devices.
type ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetQueryParameters struct {
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
// ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetQueryParameters
}
// ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByImportedWindowsAutopilotDeviceIdentityId provides operations to manage the importedWindowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
// returns a *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder when successful
func (m *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) ByImportedWindowsAutopilotDeviceIdentityId(importedWindowsAutopilotDeviceIdentityId string)(*ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if importedWindowsAutopilotDeviceIdentityId != "" {
        urlTplParams["importedWindowsAutopilotDeviceIdentity%2Did"] = importedWindowsAutopilotDeviceIdentityId
    }
    return NewImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentityItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal instantiates a new ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder and sets the default values.
func NewImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    m := &ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/importedWindowsAutopilotDeviceIdentities{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder instantiates a new ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder and sets the default values.
func NewImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ImportedwindowsautopilotdeviceidentitiesCountRequestBuilder when successful
func (m *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) Count()(*ImportedwindowsautopilotdeviceidentitiesCountRequestBuilder) {
    return NewImportedwindowsautopilotdeviceidentitiesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of imported Windows autopilot devices.
// returns a ImportedWindowsAutopilotDeviceIdentityCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) Get(ctx context.Context, requestConfiguration *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedWindowsAutopilotDeviceIdentityCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateImportedWindowsAutopilotDeviceIdentityCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedWindowsAutopilotDeviceIdentityCollectionResponseable), nil
}
// ImportEscaped provides operations to call the import method.
// returns a *ImportedwindowsautopilotdeviceidentitiesImportRequestBuilder when successful
func (m *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) ImportEscaped()(*ImportedwindowsautopilotdeviceidentitiesImportRequestBuilder) {
    return NewImportedwindowsautopilotdeviceidentitiesImportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to importedWindowsAutopilotDeviceIdentities for deviceManagement
// returns a ImportedWindowsAutopilotDeviceIdentityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedWindowsAutopilotDeviceIdentityable, requestConfiguration *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedWindowsAutopilotDeviceIdentityable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateImportedWindowsAutopilotDeviceIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedWindowsAutopilotDeviceIdentityable), nil
}
// ToGetRequestInformation collection of imported Windows autopilot devices.
// returns a *RequestInformation when successful
func (m *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to importedWindowsAutopilotDeviceIdentities for deviceManagement
// returns a *RequestInformation when successful
func (m *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ImportedWindowsAutopilotDeviceIdentityable, requestConfiguration *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder when successful
func (m *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) WithUrl(rawUrl string)(*ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    return NewImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
