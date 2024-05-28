package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder provides operations to manage the matchedDevices property of the microsoft.graph.windowsUpdates.applicableContent entity.
type WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilderGetQueryParameters collection of devices and recommendations for applicable catalog content.
type WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilderGetQueryParameters struct {
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
// WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilderGetQueryParameters
}
// WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByApplicableContentDeviceMatchDeviceId provides operations to manage the matchedDevices property of the microsoft.graph.windowsUpdates.applicableContent entity.
// returns a *WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesApplicableContentDeviceMatchDeviceItemRequestBuilder when successful
func (m *WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder) ByApplicableContentDeviceMatchDeviceId(applicableContentDeviceMatchDeviceId string)(*WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesApplicableContentDeviceMatchDeviceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if applicableContentDeviceMatchDeviceId != "" {
        urlTplParams["applicableContentDeviceMatch%2DdeviceId"] = applicableContentDeviceMatchDeviceId
    }
    return NewWindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesApplicableContentDeviceMatchDeviceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilderInternal instantiates a new WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder) {
    m := &WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/deploymentAudiences/{deploymentAudience%2Did}/applicableContent/{applicableContent%2DcatalogEntryId}/matchedDevices{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewWindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder instantiates a new WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder and sets the default values.
func NewWindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesCountRequestBuilder when successful
func (m *WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder) Count()(*WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesCountRequestBuilder) {
    return NewWindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get collection of devices and recommendations for applicable catalog content.
// returns a ApplicableContentDeviceMatchCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentDeviceMatchCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateApplicableContentDeviceMatchCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentDeviceMatchCollectionResponseable), nil
}
// Post create new navigation property to matchedDevices for admin
// returns a ApplicableContentDeviceMatchable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder) Post(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentDeviceMatchable, requestConfiguration *WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilderPostRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentDeviceMatchable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateApplicableContentDeviceMatchFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentDeviceMatchable), nil
}
// ToGetRequestInformation collection of devices and recommendations for applicable catalog content.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to matchedDevices for admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentDeviceMatchable, requestConfiguration *WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder when successful
func (m *WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder) {
    return NewWindowsUpdatesDeploymentaudiencesItemApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
