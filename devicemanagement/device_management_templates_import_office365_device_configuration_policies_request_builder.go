package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder provides operations to call the importOffice365DeviceConfigurationPolicies method.
type DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilderInternal instantiates a new ImportOffice365DeviceConfigurationPoliciesRequestBuilder and sets the default values.
func NewDeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder) {
    m := &DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder{
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
// NewDeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder instantiates a new ImportOffice365DeviceConfigurationPoliciesRequestBuilder and sets the default values.
func NewDeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action importOffice365DeviceConfigurationPolicies
func (m *DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder) Post(ctx context.Context, requestConfiguration *DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplatesImportOffice365DeviceConfigurationPoliciesResponseable), nil
}
