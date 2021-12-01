package settingdefinitions

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i2e88e9256a985641916975c2d6ae64908d2d6f42848ec4c05a5565a2f6a6107c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/reusablepolicysettings/item/referencingconfigurationpolicies/item/settings/item/settingdefinitions/ref"
)

// SettingDefinitionsRequestBuilder builds and executes requests for operations under \deviceManagement\reusablePolicySettings\{deviceManagementReusablePolicySetting-id}\referencingConfigurationPolicies\{deviceManagementConfigurationPolicy-id}\settings\{deviceManagementConfigurationSetting-id}\settingDefinitions
type SettingDefinitionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SettingDefinitionsRequestBuilderGetOptions options for Get
type SettingDefinitionsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SettingDefinitionsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SettingDefinitionsRequestBuilderGetQueryParameters list of related Setting Definitions. This property is read-only.
type SettingDefinitionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// NewSettingDefinitionsRequestBuilderInternal instantiates a new SettingDefinitionsRequestBuilder and sets the default values.
func NewSettingDefinitionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SettingDefinitionsRequestBuilder) {
    m := &SettingDefinitionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reusablePolicySettings/{deviceManagementReusablePolicySetting_id}/referencingConfigurationPolicies/{deviceManagementConfigurationPolicy_id}/settings/{deviceManagementConfigurationSetting_id}/settingDefinitions{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSettingDefinitionsRequestBuilder instantiates a new SettingDefinitionsRequestBuilder and sets the default values.
func NewSettingDefinitionsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SettingDefinitionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSettingDefinitionsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation list of related Setting Definitions. This property is read-only.
func (m *SettingDefinitionsRequestBuilder) CreateGetRequestInformation(options *SettingDefinitionsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get list of related Setting Definitions. This property is read-only.
func (m *SettingDefinitionsRequestBuilder) Get(options *SettingDefinitionsRequestBuilderGetOptions)(*SettingDefinitionsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSettingDefinitionsResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*SettingDefinitionsResponse), nil
}
func (m *SettingDefinitionsRequestBuilder) Ref()(*i2e88e9256a985641916975c2d6ae64908d2d6f42848ec4c05a5565a2f6a6107c.RefRequestBuilder) {
    return i2e88e9256a985641916975c2d6ae64908d2d6f42848ec4c05a5565a2f6a6107c.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
