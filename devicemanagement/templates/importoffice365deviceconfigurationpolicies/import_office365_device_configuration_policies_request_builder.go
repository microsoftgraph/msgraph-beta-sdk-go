package importoffice365deviceconfigurationpolicies

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ImportOffice365DeviceConfigurationPoliciesRequestBuilder provides operations to call the importOffice365DeviceConfigurationPolicies method.
type ImportOffice365DeviceConfigurationPoliciesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewImportOffice365DeviceConfigurationPoliciesRequestBuilderInternal instantiates a new ImportOffice365DeviceConfigurationPoliciesRequestBuilder and sets the default values.
func NewImportOffice365DeviceConfigurationPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportOffice365DeviceConfigurationPoliciesRequestBuilder) {
    m := &ImportOffice365DeviceConfigurationPoliciesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/templates/microsoft.graph.importOffice365DeviceConfigurationPolicies";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewImportOffice365DeviceConfigurationPoliciesRequestBuilder instantiates a new ImportOffice365DeviceConfigurationPoliciesRequestBuilder and sets the default values.
func NewImportOffice365DeviceConfigurationPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportOffice365DeviceConfigurationPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImportOffice365DeviceConfigurationPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action importOffice365DeviceConfigurationPolicies
func (m *ImportOffice365DeviceConfigurationPoliciesRequestBuilder) CreatePostRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action importOffice365DeviceConfigurationPolicies
func (m *ImportOffice365DeviceConfigurationPoliciesRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *ImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action importOffice365DeviceConfigurationPolicies
func (m *ImportOffice365DeviceConfigurationPoliciesRequestBuilder) Post()(ImportOffice365DeviceConfigurationPoliciesResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action importOffice365DeviceConfigurationPolicies
func (m *ImportOffice365DeviceConfigurationPoliciesRequestBuilder) PostWithRequestConfigurationAndResponseHandler(requestConfiguration *ImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ImportOffice365DeviceConfigurationPoliciesResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateImportOffice365DeviceConfigurationPoliciesResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(ImportOffice365DeviceConfigurationPoliciesResponseable), nil
}
