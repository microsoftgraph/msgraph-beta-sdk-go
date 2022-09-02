package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i5024ea90c8fd45ac7dbd6e184330159a989748ef5af2324e47ab26071b412298 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/microsofttunnelsites/item/microsofttunnelservers"
    ib99129148e5b54549dbf1a1827b35b5e072543abc1f750848c53c7a2bfcbbf7d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/microsofttunnelsites/item/requestupgrade"
    ibbd8be26def4f1c33c115243694230b27231cfa01a930726d7cd6e16f6198b0a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/microsofttunnelsites/item/microsofttunnelconfiguration"
    i9350d6dec8de249f17a367f9b40b072f9094b90586b5f3d60925874bf5f54981 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/microsofttunnelsites/item/microsofttunnelservers/item"
)

// MicrosoftTunnelSiteItemRequestBuilder provides operations to manage the microsoftTunnelSites property of the microsoft.graph.deviceManagement entity.
type MicrosoftTunnelSiteItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftTunnelSiteItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftTunnelSiteItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MicrosoftTunnelSiteItemRequestBuilderGetQueryParameters collection of MicrosoftTunnelSite settings associated with account.
type MicrosoftTunnelSiteItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MicrosoftTunnelSiteItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftTunnelSiteItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MicrosoftTunnelSiteItemRequestBuilderGetQueryParameters
}
// MicrosoftTunnelSiteItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftTunnelSiteItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftTunnelSiteItemRequestBuilderInternal instantiates a new MicrosoftTunnelSiteItemRequestBuilder and sets the default values.
func NewMicrosoftTunnelSiteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftTunnelSiteItemRequestBuilder) {
    m := &MicrosoftTunnelSiteItemRequestBuilder{
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
// NewMicrosoftTunnelSiteItemRequestBuilder instantiates a new MicrosoftTunnelSiteItemRequestBuilder and sets the default values.
func NewMicrosoftTunnelSiteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftTunnelSiteItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftTunnelSiteItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property microsoftTunnelSites for deviceManagement
func (m *MicrosoftTunnelSiteItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property microsoftTunnelSites for deviceManagement
func (m *MicrosoftTunnelSiteItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *MicrosoftTunnelSiteItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MicrosoftTunnelSiteItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration collection of MicrosoftTunnelSite settings associated with account.
func (m *MicrosoftTunnelSiteItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *MicrosoftTunnelSiteItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MicrosoftTunnelSiteItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property microsoftTunnelSites in deviceManagement
func (m *MicrosoftTunnelSiteItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable, requestConfiguration *MicrosoftTunnelSiteItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property microsoftTunnelSites for deviceManagement
func (m *MicrosoftTunnelSiteItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MicrosoftTunnelSiteItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
func (m *MicrosoftTunnelSiteItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosoftTunnelSiteItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// MicrosoftTunnelConfiguration the microsoftTunnelConfiguration property
func (m *MicrosoftTunnelSiteItemRequestBuilder) MicrosoftTunnelConfiguration()(*ibbd8be26def4f1c33c115243694230b27231cfa01a930726d7cd6e16f6198b0a.MicrosoftTunnelConfigurationRequestBuilder) {
    return ibbd8be26def4f1c33c115243694230b27231cfa01a930726d7cd6e16f6198b0a.NewMicrosoftTunnelConfigurationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftTunnelServers the microsoftTunnelServers property
func (m *MicrosoftTunnelSiteItemRequestBuilder) MicrosoftTunnelServers()(*i5024ea90c8fd45ac7dbd6e184330159a989748ef5af2324e47ab26071b412298.MicrosoftTunnelServersRequestBuilder) {
    return i5024ea90c8fd45ac7dbd6e184330159a989748ef5af2324e47ab26071b412298.NewMicrosoftTunnelServersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MicrosoftTunnelServersById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.microsoftTunnelSites.item.microsoftTunnelServers.item collection
func (m *MicrosoftTunnelSiteItemRequestBuilder) MicrosoftTunnelServersById(id string)(*i9350d6dec8de249f17a367f9b40b072f9094b90586b5f3d60925874bf5f54981.MicrosoftTunnelServerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["microsoftTunnelServer%2Did"] = id
    }
    return i9350d6dec8de249f17a367f9b40b072f9094b90586b5f3d60925874bf5f54981.NewMicrosoftTunnelServerItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property microsoftTunnelSites in deviceManagement
func (m *MicrosoftTunnelSiteItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MicrosoftTunnelSiteable, requestConfiguration *MicrosoftTunnelSiteItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// RequestUpgrade the requestUpgrade property
func (m *MicrosoftTunnelSiteItemRequestBuilder) RequestUpgrade()(*ib99129148e5b54549dbf1a1827b35b5e072543abc1f750848c53c7a2bfcbbf7d.RequestUpgradeRequestBuilder) {
    return ib99129148e5b54549dbf1a1827b35b5e072543abc1f750848c53c7a2bfcbbf7d.NewRequestUpgradeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
