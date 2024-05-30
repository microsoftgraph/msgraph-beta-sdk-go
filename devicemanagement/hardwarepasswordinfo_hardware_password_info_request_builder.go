package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// HardwarepasswordinfoHardwarePasswordInfoRequestBuilder provides operations to manage the hardwarePasswordInfo property of the microsoft.graph.deviceManagement entity.
type HardwarepasswordinfoHardwarePasswordInfoRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HardwarepasswordinfoHardwarePasswordInfoRequestBuilderGetQueryParameters the hardware password info for this account.
type HardwarepasswordinfoHardwarePasswordInfoRequestBuilderGetQueryParameters struct {
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
// HardwarepasswordinfoHardwarePasswordInfoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HardwarepasswordinfoHardwarePasswordInfoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *HardwarepasswordinfoHardwarePasswordInfoRequestBuilderGetQueryParameters
}
// HardwarepasswordinfoHardwarePasswordInfoRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HardwarepasswordinfoHardwarePasswordInfoRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByHardwarePasswordInfoId provides operations to manage the hardwarePasswordInfo property of the microsoft.graph.deviceManagement entity.
// returns a *HardwarepasswordinfoHardwarePasswordInfoItemRequestBuilder when successful
func (m *HardwarepasswordinfoHardwarePasswordInfoRequestBuilder) ByHardwarePasswordInfoId(hardwarePasswordInfoId string)(*HardwarepasswordinfoHardwarePasswordInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if hardwarePasswordInfoId != "" {
        urlTplParams["hardwarePasswordInfo%2Did"] = hardwarePasswordInfoId
    }
    return NewHardwarepasswordinfoHardwarePasswordInfoItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewHardwarepasswordinfoHardwarePasswordInfoRequestBuilderInternal instantiates a new HardwarepasswordinfoHardwarePasswordInfoRequestBuilder and sets the default values.
func NewHardwarepasswordinfoHardwarePasswordInfoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HardwarepasswordinfoHardwarePasswordInfoRequestBuilder) {
    m := &HardwarepasswordinfoHardwarePasswordInfoRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/hardwarePasswordInfo{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewHardwarepasswordinfoHardwarePasswordInfoRequestBuilder instantiates a new HardwarepasswordinfoHardwarePasswordInfoRequestBuilder and sets the default values.
func NewHardwarepasswordinfoHardwarePasswordInfoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HardwarepasswordinfoHardwarePasswordInfoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHardwarepasswordinfoHardwarePasswordInfoRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *HardwarepasswordinfoCountRequestBuilder when successful
func (m *HardwarepasswordinfoHardwarePasswordInfoRequestBuilder) Count()(*HardwarepasswordinfoCountRequestBuilder) {
    return NewHardwarepasswordinfoCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the hardware password info for this account.
// returns a HardwarePasswordInfoCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwarepasswordinfoHardwarePasswordInfoRequestBuilder) Get(ctx context.Context, requestConfiguration *HardwarepasswordinfoHardwarePasswordInfoRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwarePasswordInfoCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateHardwarePasswordInfoCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwarePasswordInfoCollectionResponseable), nil
}
// Post create new navigation property to hardwarePasswordInfo for deviceManagement
// returns a HardwarePasswordInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *HardwarepasswordinfoHardwarePasswordInfoRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwarePasswordInfoable, requestConfiguration *HardwarepasswordinfoHardwarePasswordInfoRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwarePasswordInfoable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation the hardware password info for this account.
// returns a *RequestInformation when successful
func (m *HardwarepasswordinfoHardwarePasswordInfoRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HardwarepasswordinfoHardwarePasswordInfoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to hardwarePasswordInfo for deviceManagement
// returns a *RequestInformation when successful
func (m *HardwarepasswordinfoHardwarePasswordInfoRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.HardwarePasswordInfoable, requestConfiguration *HardwarepasswordinfoHardwarePasswordInfoRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *HardwarepasswordinfoHardwarePasswordInfoRequestBuilder when successful
func (m *HardwarepasswordinfoHardwarePasswordInfoRequestBuilder) WithUrl(rawUrl string)(*HardwarepasswordinfoHardwarePasswordInfoRequestBuilder) {
    return NewHardwarepasswordinfoHardwarePasswordInfoRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
