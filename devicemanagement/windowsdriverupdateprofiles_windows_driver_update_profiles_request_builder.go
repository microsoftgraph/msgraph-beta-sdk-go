package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder provides operations to manage the windowsDriverUpdateProfiles property of the microsoft.graph.deviceManagement entity.
type WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilderGetQueryParameters a collection of windows driver update profiles
type WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilderGetQueryParameters struct {
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
// WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilderGetQueryParameters
}
// WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWindowsDriverUpdateProfileId provides operations to manage the windowsDriverUpdateProfiles property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder when successful
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder) ByWindowsDriverUpdateProfileId(windowsDriverUpdateProfileId string)(*WindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if windowsDriverUpdateProfileId != "" {
        urlTplParams["windowsDriverUpdateProfile%2Did"] = windowsDriverUpdateProfileId
    }
    return NewWindowsdriverupdateprofilesWindowsDriverUpdateProfileItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilderInternal instantiates a new WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder and sets the default values.
func NewWindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder) {
    m := &WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/windowsDriverUpdateProfiles{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewWindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder instantiates a new WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder and sets the default values.
func NewWindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *WindowsdriverupdateprofilesCountRequestBuilder when successful
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder) Count()(*WindowsdriverupdateprofilesCountRequestBuilder) {
    return NewWindowsdriverupdateprofilesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a collection of windows driver update profiles
// returns a WindowsDriverUpdateProfileCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDriverUpdateProfileCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileCollectionResponseable), nil
}
// Post create new navigation property to windowsDriverUpdateProfiles for deviceManagement
// returns a WindowsDriverUpdateProfileable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable, requestConfiguration *WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsDriverUpdateProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable), nil
}
// ToGetRequestInformation a collection of windows driver update profiles
// returns a *RequestInformation when successful
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to windowsDriverUpdateProfiles for deviceManagement
// returns a *RequestInformation when successful
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsDriverUpdateProfileable, requestConfiguration *WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder when successful
func (m *WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder) WithUrl(rawUrl string)(*WindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder) {
    return NewWindowsdriverupdateprofilesWindowsDriverUpdateProfilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
