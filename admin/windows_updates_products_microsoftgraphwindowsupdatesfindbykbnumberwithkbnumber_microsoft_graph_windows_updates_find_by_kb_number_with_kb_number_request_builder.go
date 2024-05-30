package admin

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilder provides operations to call the findByKbNumber method.
type WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilderGetQueryParameters invoke function findByKbNumber
type WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilderGetQueryParameters struct {
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
// WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilderGetQueryParameters
}
// NewWindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilderInternal instantiates a new WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilder and sets the default values.
func NewWindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, kbNumber *int32)(*WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilder) {
    m := &WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/products/microsoft.graph.windowsUpdates.findByKbNumber(kbNumber={kbNumber}){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if kbNumber != nil {
        m.BaseRequestBuilder.PathParameters["kbNumber"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*kbNumber), 10)
    }
    return m
}
// NewWindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilder instantiates a new WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilder and sets the default values.
func NewWindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function findByKbNumber
// Deprecated: This method is obsolete. Use GetAsFindByKbNumberWithKbNumberGetResponse instead.
// returns a WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberFindByKbNumberWithKbNumberResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilderGetRequestConfiguration)(WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberFindByKbNumberWithKbNumberResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateWindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberFindByKbNumberWithKbNumberResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberFindByKbNumberWithKbNumberResponseable), nil
}
// GetAsFindByKbNumberWithKbNumberGetResponse invoke function findByKbNumber
// returns a WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberFindByKbNumberWithKbNumberGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilder) GetAsFindByKbNumberWithKbNumberGetResponse(ctx context.Context, requestConfiguration *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilderGetRequestConfiguration)(WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberFindByKbNumberWithKbNumberGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateWindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberFindByKbNumberWithKbNumberGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberFindByKbNumberWithKbNumberGetResponseable), nil
}
// ToGetRequestInformation invoke function findByKbNumber
// returns a *RequestInformation when successful
func (m *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilder when successful
func (m *WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilder) {
    return NewWindowsUpdatesProductsMicrosoftgraphwindowsupdatesfindbykbnumberwithkbnumberMicrosoftGraphWindowsUpdatesFindByKbNumberWithKbNumberRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
