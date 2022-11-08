package officeconfiguration

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ib51775d0103f2848789780fed67e4a5c9a298b92986a1aeabcf8f735e6666d7f "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations"
    i4f3ec96b42f2d1014637f13f95e0e6619fd4e35e09fad4a2b5571a24e4f6d381 "github.com/microsoftgraph/msgraph-beta-sdk-go/officeconfiguration/clientconfigurations/item"
)

// OfficeConfigurationRequestBuilder provides operations to manage the officeConfiguration singleton.
type OfficeConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// OfficeConfigurationRequestBuilderGetQueryParameters get officeConfiguration
type OfficeConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// OfficeConfigurationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OfficeConfigurationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *OfficeConfigurationRequestBuilderGetQueryParameters
}
// OfficeConfigurationRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OfficeConfigurationRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClientConfigurations provides operations to manage the clientConfigurations property of the microsoft.graph.officeConfiguration entity.
func (m *OfficeConfigurationRequestBuilder) ClientConfigurations()(*ib51775d0103f2848789780fed67e4a5c9a298b92986a1aeabcf8f735e6666d7f.ClientConfigurationsRequestBuilder) {
    return ib51775d0103f2848789780fed67e4a5c9a298b92986a1aeabcf8f735e6666d7f.NewClientConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClientConfigurationsById provides operations to manage the clientConfigurations property of the microsoft.graph.officeConfiguration entity.
func (m *OfficeConfigurationRequestBuilder) ClientConfigurationsById(id string)(*i4f3ec96b42f2d1014637f13f95e0e6619fd4e35e09fad4a2b5571a24e4f6d381.OfficeClientConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["officeClientConfiguration%2Did"] = id
    }
    return i4f3ec96b42f2d1014637f13f95e0e6619fd4e35e09fad4a2b5571a24e4f6d381.NewOfficeClientConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewOfficeConfigurationRequestBuilderInternal instantiates a new OfficeConfigurationRequestBuilder and sets the default values.
func NewOfficeConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OfficeConfigurationRequestBuilder) {
    m := &OfficeConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/officeConfiguration{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewOfficeConfigurationRequestBuilder instantiates a new OfficeConfigurationRequestBuilder and sets the default values.
func NewOfficeConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*OfficeConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewOfficeConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get officeConfiguration
func (m *OfficeConfigurationRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *OfficeConfigurationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update officeConfiguration
func (m *OfficeConfigurationRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OfficeConfigurationable, requestConfiguration *OfficeConfigurationRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get officeConfiguration
func (m *OfficeConfigurationRequestBuilder) Get(ctx context.Context, requestConfiguration *OfficeConfigurationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OfficeConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOfficeConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OfficeConfigurationable), nil
}
// Patch update officeConfiguration
func (m *OfficeConfigurationRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OfficeConfigurationable, requestConfiguration *OfficeConfigurationRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OfficeConfigurationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOfficeConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OfficeConfigurationable), nil
}
