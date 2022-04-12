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
// ImportOffice365DeviceConfigurationPoliciesRequestBuilderPostOptions options for Post
type ImportOffice365DeviceConfigurationPoliciesRequestBuilderPostOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
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
func (m *ImportOffice365DeviceConfigurationPoliciesRequestBuilder) CreatePostRequestInformation(options *ImportOffice365DeviceConfigurationPoliciesRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Post invoke action importOffice365DeviceConfigurationPolicies
func (m *ImportOffice365DeviceConfigurationPoliciesRequestBuilder) Post(options *ImportOffice365DeviceConfigurationPoliciesRequestBuilderPostOptions)(ImportOffice365DeviceConfigurationPoliciesResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateImportOffice365DeviceConfigurationPoliciesResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(ImportOffice365DeviceConfigurationPoliciesResponseable), nil
}
