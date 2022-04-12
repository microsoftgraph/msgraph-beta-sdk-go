package windowsupdateforbusinessconfiguration

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i72409d3ebb35c3c3f9e273be2486ddf7083aaee0ddce9f46762e2af73feea1a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/windowsupdateforbusinessconfiguration/extendfeatureupdatespause"
    ibd6910a46bd92304bd98f3581e14d46160c5def88b743996b87f376e41604cac "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/windowsupdateforbusinessconfiguration/extendqualityupdatespause"
)

// WindowsUpdateForBusinessConfigurationRequestBuilder builds and executes requests for operations under \deviceManagement\deviceConfigurations\{deviceConfiguration-id}\microsoft.graph.windowsUpdateForBusinessConfiguration
type WindowsUpdateForBusinessConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewWindowsUpdateForBusinessConfigurationRequestBuilderInternal instantiates a new WindowsUpdateForBusinessConfigurationRequestBuilder and sets the default values.
func NewWindowsUpdateForBusinessConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdateForBusinessConfigurationRequestBuilder) {
    m := &WindowsUpdateForBusinessConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}/microsoft.graph.windowsUpdateForBusinessConfiguration";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsUpdateForBusinessConfigurationRequestBuilder instantiates a new WindowsUpdateForBusinessConfigurationRequestBuilder and sets the default values.
func NewWindowsUpdateForBusinessConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdateForBusinessConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdateForBusinessConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// ExtendFeatureUpdatesPause the extendFeatureUpdatesPause property
func (m *WindowsUpdateForBusinessConfigurationRequestBuilder) ExtendFeatureUpdatesPause()(*i72409d3ebb35c3c3f9e273be2486ddf7083aaee0ddce9f46762e2af73feea1a2.ExtendFeatureUpdatesPauseRequestBuilder) {
    return i72409d3ebb35c3c3f9e273be2486ddf7083aaee0ddce9f46762e2af73feea1a2.NewExtendFeatureUpdatesPauseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtendQualityUpdatesPause the extendQualityUpdatesPause property
func (m *WindowsUpdateForBusinessConfigurationRequestBuilder) ExtendQualityUpdatesPause()(*ibd6910a46bd92304bd98f3581e14d46160c5def88b743996b87f376e41604cac.ExtendQualityUpdatesPauseRequestBuilder) {
    return ibd6910a46bd92304bd98f3581e14d46160c5def88b743996b87f376e41604cac.NewExtendQualityUpdatesPauseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
