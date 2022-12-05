package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder provides operations to manage the microsoftTunnelSites property of the microsoft.graph.deviceManagement entity.
type DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderGetQueryParameters collection of MicrosoftTunnelSite settings associated with account.
type DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderGetQueryParameters
}
// DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderInternal instantiates a new MicrosoftTunnelSiteItemRequestBuilder and sets the default values.
func NewDeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) {
    m := &DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/microsoftTunnelSites/{microsoftTunnelSite%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder instantiates a new MicrosoftTunnelSiteItemRequestBuilder and sets the default values.
func NewDeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property microsoftTunnelSites for deviceManagement
func (m *DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation collection of MicrosoftTunnelSite settings associated with account.
func (m *DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property microsoftTunnelSites in deviceManagement
func (m *DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable, requestConfiguration *DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property microsoftTunnelSites for deviceManagement
func (m *DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get collection of MicrosoftTunnelSite settings associated with account.
func (m *DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelSiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable), nil
}
// MicrosoftTunnelConfiguration provides operations to manage the microsoftTunnelConfiguration property of the microsoft.graph.microsoftTunnelSite entity.
func (m *DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) MicrosoftTunnelConfiguration()(*DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelConfigurationRequestBuilder) {
    return NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftTunnelServers provides operations to manage the microsoftTunnelServers property of the microsoft.graph.microsoftTunnelSite entity.
func (m *DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) MicrosoftTunnelServers()(*DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilder) {
    return NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftTunnelServersById provides operations to manage the microsoftTunnelServers property of the microsoft.graph.microsoftTunnelSite entity.
func (m *DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) MicrosoftTunnelServersById(id string)(*DeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftTunnelServer%2Did"] = id
    }
    return NewDeviceManagementMicrosoftTunnelSitesItemMicrosoftTunnelServersMicrosoftTunnelServerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property microsoftTunnelSites in deviceManagement
func (m *DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable, requestConfiguration *DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMicrosoftTunnelSiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable), nil
}
// RequestUpgrade provides operations to call the requestUpgrade method.
func (m *DeviceManagementMicrosoftTunnelSitesMicrosoftTunnelSiteItemRequestBuilder) RequestUpgrade()(*DeviceManagementMicrosoftTunnelSitesItemRequestUpgradeRequestBuilder) {
    return NewDeviceManagementMicrosoftTunnelSitesItemRequestUpgradeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
