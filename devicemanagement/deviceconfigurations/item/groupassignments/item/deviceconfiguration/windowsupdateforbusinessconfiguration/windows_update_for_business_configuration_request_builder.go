package windowsupdateforbusinessconfiguration

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i0d05c7089523495d2bc819e5303b17d97fbcf19eccd2407d52085c077250d401 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/groupassignments/item/deviceconfiguration/windowsupdateforbusinessconfiguration/extendqualityupdatespause"
    i79760522376675a27c0bb443928a73359202ae76f6aca75d5b740c883d3ff75d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/deviceconfigurations/item/groupassignments/item/deviceconfiguration/windowsupdateforbusinessconfiguration/extendfeatureupdatespause"
)

type WindowsUpdateForBusinessConfigurationRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
func NewWindowsUpdateForBusinessConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsUpdateForBusinessConfigurationRequestBuilder) {
    m := &WindowsUpdateForBusinessConfigurationRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/deviceManagement/deviceConfigurations/{deviceConfiguration_id}/groupAssignments/{deviceConfigurationGroupAssignment_id}/deviceConfiguration/microsoft.graph.windowsUpdateForBusinessConfiguration";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewWindowsUpdateForBusinessConfigurationRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*WindowsUpdateForBusinessConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdateForBusinessConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *WindowsUpdateForBusinessConfigurationRequestBuilder) ExtendFeatureUpdatesPause()(*i79760522376675a27c0bb443928a73359202ae76f6aca75d5b740c883d3ff75d.ExtendFeatureUpdatesPauseRequestBuilder) {
    return i79760522376675a27c0bb443928a73359202ae76f6aca75d5b740c883d3ff75d.NewExtendFeatureUpdatesPauseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *WindowsUpdateForBusinessConfigurationRequestBuilder) ExtendQualityUpdatesPause()(*i0d05c7089523495d2bc819e5303b17d97fbcf19eccd2407d52085c077250d401.ExtendQualityUpdatesPauseRequestBuilder) {
    return i0d05c7089523495d2bc819e5303b17d97fbcf19eccd2407d52085c077250d401.NewExtendQualityUpdatesPauseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
