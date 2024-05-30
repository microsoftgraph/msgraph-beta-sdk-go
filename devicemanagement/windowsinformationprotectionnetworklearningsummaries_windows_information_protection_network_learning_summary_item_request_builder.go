package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder provides operations to manage the windowsInformationProtectionNetworkLearningSummaries property of the microsoft.graph.deviceManagement entity.
type WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderGetQueryParameters the windows information protection network learning summaries.
type WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderGetQueryParameters
}
// WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewWindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderInternal instantiates a new WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder and sets the default values.
func NewWindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) {
    m := &WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsInformationProtectionNetworkLearningSummaries/{windowsInformationProtectionNetworkLearningSummary%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder instantiates a new WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder and sets the default values.
func NewWindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property windowsInformationProtectionNetworkLearningSummaries for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the windows information protection network learning summaries.
// returns a WindowsInformationProtectionNetworkLearningSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionNetworkLearningSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsInformationProtectionNetworkLearningSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionNetworkLearningSummaryable), nil
}
// Patch update the navigation property windowsInformationProtectionNetworkLearningSummaries in deviceManagement
// returns a WindowsInformationProtectionNetworkLearningSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionNetworkLearningSummaryable, requestConfiguration *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionNetworkLearningSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsInformationProtectionNetworkLearningSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionNetworkLearningSummaryable), nil
}
// ToDeleteRequestInformation delete navigation property windowsInformationProtectionNetworkLearningSummaries for deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the windows information protection network learning summaries.
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property windowsInformationProtectionNetworkLearningSummaries in deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsInformationProtectionNetworkLearningSummaryable, requestConfiguration *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder when successful
func (m *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) WithUrl(rawUrl string)(*WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder) {
    return NewWindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummaryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
