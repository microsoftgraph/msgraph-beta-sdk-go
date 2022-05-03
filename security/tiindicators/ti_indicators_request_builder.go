package tiindicators

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i16222b62e050bff07e47b4d54715cd2b8242e600345080f1bf21dd94b9cbf876 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/tiindicators/count"
    i5fd228a8811b3ee0e602ce99917149a6ef997b242fa74579dc28eb98aeff4610 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/tiindicators/updatetiindicators"
    i84faad3a80813a60a94e3416bfa220426811428a0f17f7a125a8a8389a71931a "github.com/microsoftgraph/msgraph-beta-sdk-go/security/tiindicators/deletetiindicatorsbyexternalid"
    id9554bf1b52f05e494c0e5962b41bb42f06710f09ff2286b921260ce9ef0c431 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/tiindicators/submittiindicators"
    ie38c9a1118b781a8cc486e08009768123690676d39af853ebdf06098c5110f09 "github.com/microsoftgraph/msgraph-beta-sdk-go/security/tiindicators/deletetiindicators"
)

// TiIndicatorsRequestBuilder provides operations to manage the tiIndicators property of the microsoft.graph.security entity.
type TiIndicatorsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TiIndicatorsRequestBuilderGetQueryParameters get tiIndicators from security
type TiIndicatorsRequestBuilderGetQueryParameters struct {
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
// TiIndicatorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TiIndicatorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TiIndicatorsRequestBuilderGetQueryParameters
}
// TiIndicatorsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TiIndicatorsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTiIndicatorsRequestBuilderInternal instantiates a new TiIndicatorsRequestBuilder and sets the default values.
func NewTiIndicatorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TiIndicatorsRequestBuilder) {
    m := &TiIndicatorsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/tiIndicators{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTiIndicatorsRequestBuilder instantiates a new TiIndicatorsRequestBuilder and sets the default values.
func NewTiIndicatorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TiIndicatorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTiIndicatorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the count property
func (m *TiIndicatorsRequestBuilder) Count()(*i16222b62e050bff07e47b4d54715cd2b8242e600345080f1bf21dd94b9cbf876.CountRequestBuilder) {
    return i16222b62e050bff07e47b4d54715cd2b8242e600345080f1bf21dd94b9cbf876.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get tiIndicators from security
func (m *TiIndicatorsRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get tiIndicators from security
func (m *TiIndicatorsRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *TiIndicatorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePostRequestInformation create new navigation property to tiIndicators for security
func (m *TiIndicatorsRequestBuilder) CreatePostRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TiIndicatorable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration create new navigation property to tiIndicators for security
func (m *TiIndicatorsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TiIndicatorable, requestConfiguration *TiIndicatorsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DeleteTiIndicators the deleteTiIndicators property
func (m *TiIndicatorsRequestBuilder) DeleteTiIndicators()(*ie38c9a1118b781a8cc486e08009768123690676d39af853ebdf06098c5110f09.DeleteTiIndicatorsRequestBuilder) {
    return ie38c9a1118b781a8cc486e08009768123690676d39af853ebdf06098c5110f09.NewDeleteTiIndicatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeleteTiIndicatorsByExternalId the deleteTiIndicatorsByExternalId property
func (m *TiIndicatorsRequestBuilder) DeleteTiIndicatorsByExternalId()(*i84faad3a80813a60a94e3416bfa220426811428a0f17f7a125a8a8389a71931a.DeleteTiIndicatorsByExternalIdRequestBuilder) {
    return i84faad3a80813a60a94e3416bfa220426811428a0f17f7a125a8a8389a71931a.NewDeleteTiIndicatorsByExternalIdRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get tiIndicators from security
func (m *TiIndicatorsRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TiIndicatorCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get tiIndicators from security
func (m *TiIndicatorsRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *TiIndicatorsRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TiIndicatorCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTiIndicatorCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TiIndicatorCollectionResponseable), nil
}
// Post create new navigation property to tiIndicators for security
func (m *TiIndicatorsRequestBuilder) Post(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TiIndicatorable)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TiIndicatorable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler create new navigation property to tiIndicators for security
func (m *TiIndicatorsRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TiIndicatorable, requestConfiguration *TiIndicatorsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TiIndicatorable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTiIndicatorFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TiIndicatorable), nil
}
// SubmitTiIndicators the submitTiIndicators property
func (m *TiIndicatorsRequestBuilder) SubmitTiIndicators()(*id9554bf1b52f05e494c0e5962b41bb42f06710f09ff2286b921260ce9ef0c431.SubmitTiIndicatorsRequestBuilder) {
    return id9554bf1b52f05e494c0e5962b41bb42f06710f09ff2286b921260ce9ef0c431.NewSubmitTiIndicatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateTiIndicators the updateTiIndicators property
func (m *TiIndicatorsRequestBuilder) UpdateTiIndicators()(*i5fd228a8811b3ee0e602ce99917149a6ef997b242fa74579dc28eb98aeff4610.UpdateTiIndicatorsRequestBuilder) {
    return i5fd228a8811b3ee0e602ce99917149a6ef997b242fa74579dc28eb98aeff4610.NewUpdateTiIndicatorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
