package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder provides operations to manage the hardwarePasswordInfo property of the microsoft.graph.deviceManagement entity.
type HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilderGetQueryParameters intune will provide customer the ability to configure BIOS configuration settings on the enrolled Windows 10 and Windows 11 Microsoft Entra joined devices. Starting from June, 2024, customers should start using hardwarePasswordDetail resource type - Microsoft Graph beta | Microsoft Learn. HardwarePasswordInfo will be marked as deprecated with Intune Release 2409
type HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilderGetQueryParameters
}
// HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewHardwarePasswordInfoHardwarePasswordInfoItemRequestBuilderInternal instantiates a new HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder and sets the default values.
func NewHardwarePasswordInfoHardwarePasswordInfoItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder) {
    m := &HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/hardwarePasswordInfo/{hardwarePasswordInfo%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewHardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder instantiates a new HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder and sets the default values.
func NewHardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHardwarePasswordInfoHardwarePasswordInfoItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property hardwarePasswordInfo for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get intune will provide customer the ability to configure BIOS configuration settings on the enrolled Windows 10 and Windows 11 Microsoft Entra joined devices. Starting from June, 2024, customers should start using hardwarePasswordDetail resource type - Microsoft Graph beta | Microsoft Learn. HardwarePasswordInfo will be marked as deprecated with Intune Release 2409
// returns a HardwarePasswordInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder) Get(ctx context.Context, requestConfiguration *HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwarePasswordInfoable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHardwarePasswordInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwarePasswordInfoable), nil
}
// Patch update the navigation property hardwarePasswordInfo in deviceManagement
// returns a HardwarePasswordInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwarePasswordInfoable, requestConfiguration *HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwarePasswordInfoable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHardwarePasswordInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwarePasswordInfoable), nil
}
// ToDeleteRequestInformation delete navigation property hardwarePasswordInfo for deviceManagement
// returns a *RequestInformation when successful
func (m *HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation intune will provide customer the ability to configure BIOS configuration settings on the enrolled Windows 10 and Windows 11 Microsoft Entra joined devices. Starting from June, 2024, customers should start using hardwarePasswordDetail resource type - Microsoft Graph beta | Microsoft Learn. HardwarePasswordInfo will be marked as deprecated with Intune Release 2409
// returns a *RequestInformation when successful
func (m *HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property hardwarePasswordInfo in deviceManagement
// returns a *RequestInformation when successful
func (m *HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwarePasswordInfoable, requestConfiguration *HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder when successful
func (m *HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder) WithUrl(rawUrl string)(*HardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder) {
    return NewHardwarePasswordInfoHardwarePasswordInfoItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
