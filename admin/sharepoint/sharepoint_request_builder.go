package sharepoint

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i0d83e930c557b824d39150e028a80d9f5a1fe75558f698c157c6c0dd930bcc83 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/tenantadmin"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ibcab7e3e425756e682048e2727794901f7a025dc933c4bf401cac95a8ffb618b "github.com/microsoftgraph/msgraph-beta-sdk-go/admin/sharepoint/settings"
)

// SharepointRequestBuilder provides operations to manage the sharepoint property of the microsoft.graph.admin entity.
type SharepointRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SharepointRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SharepointRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SharepointRequestBuilderGetQueryParameters a container for administrative resources to manage tenant-level settings for SharePoint and OneDrive.
type SharepointRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SharepointRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SharepointRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SharepointRequestBuilderGetQueryParameters
}
// SharepointRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SharepointRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSharepointRequestBuilderInternal instantiates a new SharepointRequestBuilder and sets the default values.
func NewSharepointRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharepointRequestBuilder) {
    m := &SharepointRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/admin/sharepoint{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSharepointRequestBuilder instantiates a new SharepointRequestBuilder and sets the default values.
func NewSharepointRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharepointRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSharepointRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property sharepoint for admin
func (m *SharepointRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property sharepoint for admin
func (m *SharepointRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *SharepointRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation a container for administrative resources to manage tenant-level settings for SharePoint and OneDrive.
func (m *SharepointRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration a container for administrative resources to manage tenant-level settings for SharePoint and OneDrive.
func (m *SharepointRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SharepointRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property sharepoint in admin
func (m *SharepointRequestBuilder) CreatePatchRequestInformation(body i0d83e930c557b824d39150e028a80d9f5a1fe75558f698c157c6c0dd930bcc83.Sharepointable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property sharepoint in admin
func (m *SharepointRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body i0d83e930c557b824d39150e028a80d9f5a1fe75558f698c157c6c0dd930bcc83.Sharepointable, requestConfiguration *SharepointRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property sharepoint for admin
func (m *SharepointRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property sharepoint for admin
func (m *SharepointRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *SharepointRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a container for administrative resources to manage tenant-level settings for SharePoint and OneDrive.
func (m *SharepointRequestBuilder) Get()(i0d83e930c557b824d39150e028a80d9f5a1fe75558f698c157c6c0dd930bcc83.Sharepointable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler a container for administrative resources to manage tenant-level settings for SharePoint and OneDrive.
func (m *SharepointRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *SharepointRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(i0d83e930c557b824d39150e028a80d9f5a1fe75558f698c157c6c0dd930bcc83.Sharepointable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i0d83e930c557b824d39150e028a80d9f5a1fe75558f698c157c6c0dd930bcc83.CreateSharepointFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i0d83e930c557b824d39150e028a80d9f5a1fe75558f698c157c6c0dd930bcc83.Sharepointable), nil
}
// Patch update the navigation property sharepoint in admin
func (m *SharepointRequestBuilder) Patch(body i0d83e930c557b824d39150e028a80d9f5a1fe75558f698c157c6c0dd930bcc83.Sharepointable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property sharepoint in admin
func (m *SharepointRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body i0d83e930c557b824d39150e028a80d9f5a1fe75558f698c157c6c0dd930bcc83.Sharepointable, requestConfiguration *SharepointRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Settings the settings property
func (m *SharepointRequestBuilder) Settings()(*ibcab7e3e425756e682048e2727794901f7a025dc933c4bf401cac95a8ffb618b.SettingsRequestBuilder) {
    return ibcab7e3e425756e682048e2727794901f7a025dc933c4bf401cac95a8ffb618b.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
