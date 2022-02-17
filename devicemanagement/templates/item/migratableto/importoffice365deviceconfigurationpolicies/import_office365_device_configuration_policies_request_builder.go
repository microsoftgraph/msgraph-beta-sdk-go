package importoffice365deviceconfigurationpolicies

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ImportOffice365DeviceConfigurationPoliciesRequestBuilder builds and executes requests for operations under \deviceManagement\templates\{deviceManagementTemplate-id}\migratableTo\microsoft.graph.importOffice365DeviceConfigurationPolicies
type ImportOffice365DeviceConfigurationPoliciesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ImportOffice365DeviceConfigurationPoliciesRequestBuilderPostOptions options for Post
type ImportOffice365DeviceConfigurationPoliciesRequestBuilderPostOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewImportOffice365DeviceConfigurationPoliciesRequestBuilderInternal instantiates a new ImportOffice365DeviceConfigurationPoliciesRequestBuilder and sets the default values.
func NewImportOffice365DeviceConfigurationPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ImportOffice365DeviceConfigurationPoliciesRequestBuilder) {
    m := &ImportOffice365DeviceConfigurationPoliciesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/templates/{deviceManagementTemplate_id}/migratableTo/microsoft.graph.importOffice365DeviceConfigurationPolicies";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewImportOffice365DeviceConfigurationPoliciesRequestBuilder instantiates a new ImportOffice365DeviceConfigurationPoliciesRequestBuilder and sets the default values.
func NewImportOffice365DeviceConfigurationPoliciesRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ImportOffice365DeviceConfigurationPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImportOffice365DeviceConfigurationPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action importOffice365DeviceConfigurationPolicies
func (m *ImportOffice365DeviceConfigurationPoliciesRequestBuilder) CreatePostRequestInformation(options *ImportOffice365DeviceConfigurationPoliciesRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
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
// Post invoke action importOffice365DeviceConfigurationPolicies
func (m *ImportOffice365DeviceConfigurationPoliciesRequestBuilder) Post(options *ImportOffice365DeviceConfigurationPoliciesRequestBuilderPostOptions)([]ImportOffice365DeviceConfigurationPolicies, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollectionAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewImportOffice365DeviceConfigurationPolicies() }, nil, nil)
    if err != nil {
        return nil, err
    }
    val := make([]ImportOffice365DeviceConfigurationPolicies, len(res))
    for i, v := range res {
        val[i] = *(v.(*ImportOffice365DeviceConfigurationPolicies))
    }
    return val, nil
}
