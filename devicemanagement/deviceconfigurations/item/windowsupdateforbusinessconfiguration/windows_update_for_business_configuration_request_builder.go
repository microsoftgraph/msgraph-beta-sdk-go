package windowsupdateforbusinessconfiguration

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i72409d3ebb35c3c3f9e273be2486ddf7083aaee0ddce9f46762e2af73feea1a2 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/windowsupdateforbusinessconfiguration/extendfeatureupdatespause"
    ibd6910a46bd92304bd98f3581e14d46160c5def88b743996b87f376e41604cac "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/windowsupdateforbusinessconfiguration/extendqualityupdatespause"
)

// WindowsUpdateForBusinessConfigurationRequestBuilder builds and executes requests for operations under \deviceManagement\deviceConfigurations\{deviceConfiguration-id}\microsoft.graph.windowsUpdateForBusinessConfiguration
type WindowsUpdateForBusinessConfigurationRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NewWindowsUpdateForBusinessConfigurationRequestBuilderInternal instantiates a new WindowsUpdateForBusinessConfigurationRequestBuilder and sets the default values.
func NewWindowsUpdateForBusinessConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsUpdateForBusinessConfigurationRequestBuilder) {
    m := &WindowsUpdateForBusinessConfigurationRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration_id}/microsoft.graph.windowsUpdateForBusinessConfiguration";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewWindowsUpdateForBusinessConfigurationRequestBuilder instantiates a new WindowsUpdateForBusinessConfigurationRequestBuilder and sets the default values.
func NewWindowsUpdateForBusinessConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsUpdateForBusinessConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdateForBusinessConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WindowsUpdateForBusinessConfigurationRequestBuilder) ExtendFeatureUpdatesPause()(*i72409d3ebb35c3c3f9e273be2486ddf7083aaee0ddce9f46762e2af73feea1a2.ExtendFeatureUpdatesPauseRequestBuilder) {
    return i72409d3ebb35c3c3f9e273be2486ddf7083aaee0ddce9f46762e2af73feea1a2.NewExtendFeatureUpdatesPauseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsUpdateForBusinessConfigurationRequestBuilder) ExtendQualityUpdatesPause()(*ibd6910a46bd92304bd98f3581e14d46160c5def88b743996b87f376e41604cac.ExtendQualityUpdatesPauseRequestBuilder) {
    return ibd6910a46bd92304bd98f3581e14d46160c5def88b743996b87f376e41604cac.NewExtendQualityUpdatesPauseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
