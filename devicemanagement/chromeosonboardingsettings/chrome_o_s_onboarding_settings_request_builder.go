package chromeosonboardingsettings

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i03210ad32b09351152234b83ec305c89c9e57ede5b483a4b11073f1d5356f1ce "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/chromeosonboardingsettings/count"
    icd47eb9dea6f9de87c7b2bffac772b07ab4e53b5bd39f33b7ea3886e5e11f322 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/chromeosonboardingsettings/disconnect"
    ie99e2dac95e812895dc09a981511e5bcfcf3d589738bbaf03993925b5ede2ada "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/chromeosonboardingsettings/connect"
)

// ChromeOSOnboardingSettingsRequestBuilder provides operations to manage the chromeOSOnboardingSettings property of the microsoft.graph.deviceManagement entity.
type ChromeOSOnboardingSettingsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChromeOSOnboardingSettingsRequestBuilderGetOptions options for Get
type ChromeOSOnboardingSettingsRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChromeOSOnboardingSettingsRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ChromeOSOnboardingSettingsRequestBuilderGetQueryParameters collection of ChromeOSOnboardingSettings settings associated with account.
type ChromeOSOnboardingSettingsRequestBuilderGetQueryParameters struct {
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
// ChromeOSOnboardingSettingsRequestBuilderPostOptions options for Post
type ChromeOSOnboardingSettingsRequestBuilderPostOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChromeOSOnboardingSettingsable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Connect the connect property
func (m *ChromeOSOnboardingSettingsRequestBuilder) Connect()(*ie99e2dac95e812895dc09a981511e5bcfcf3d589738bbaf03993925b5ede2ada.ConnectRequestBuilder) {
    return ie99e2dac95e812895dc09a981511e5bcfcf3d589738bbaf03993925b5ede2ada.NewConnectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewChromeOSOnboardingSettingsRequestBuilderInternal instantiates a new ChromeOSOnboardingSettingsRequestBuilder and sets the default values.
func NewChromeOSOnboardingSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChromeOSOnboardingSettingsRequestBuilder) {
    m := &ChromeOSOnboardingSettingsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/chromeOSOnboardingSettings{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChromeOSOnboardingSettingsRequestBuilder instantiates a new ChromeOSOnboardingSettingsRequestBuilder and sets the default values.
func NewChromeOSOnboardingSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChromeOSOnboardingSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChromeOSOnboardingSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *ChromeOSOnboardingSettingsRequestBuilder) Count()(*i03210ad32b09351152234b83ec305c89c9e57ede5b483a4b11073f1d5356f1ce.CountRequestBuilder) {
    return i03210ad32b09351152234b83ec305c89c9e57ede5b483a4b11073f1d5356f1ce.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation collection of ChromeOSOnboardingSettings settings associated with account.
func (m *ChromeOSOnboardingSettingsRequestBuilder) CreateGetRequestInformation(options *ChromeOSOnboardingSettingsRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to chromeOSOnboardingSettings for deviceManagement
func (m *ChromeOSOnboardingSettingsRequestBuilder) CreatePostRequestInformation(options *ChromeOSOnboardingSettingsRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Disconnect the disconnect property
func (m *ChromeOSOnboardingSettingsRequestBuilder) Disconnect()(*icd47eb9dea6f9de87c7b2bffac772b07ab4e53b5bd39f33b7ea3886e5e11f322.DisconnectRequestBuilder) {
    return icd47eb9dea6f9de87c7b2bffac772b07ab4e53b5bd39f33b7ea3886e5e11f322.NewDisconnectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get collection of ChromeOSOnboardingSettings settings associated with account.
func (m *ChromeOSOnboardingSettingsRequestBuilder) Get(options *ChromeOSOnboardingSettingsRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChromeOSOnboardingSettingsCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChromeOSOnboardingSettingsCollectionResponseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChromeOSOnboardingSettingsCollectionResponseable), nil
}
// Post create new navigation property to chromeOSOnboardingSettings for deviceManagement
func (m *ChromeOSOnboardingSettingsRequestBuilder) Post(options *ChromeOSOnboardingSettingsRequestBuilderPostOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChromeOSOnboardingSettingsable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateChromeOSOnboardingSettingsFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ChromeOSOnboardingSettingsable), nil
}
