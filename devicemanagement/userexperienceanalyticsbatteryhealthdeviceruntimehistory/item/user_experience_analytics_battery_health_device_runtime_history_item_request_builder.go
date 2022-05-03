package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder provides operations to manage the userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory property of the microsoft.graph.deviceManagement entity.
type UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderGetQueryParameters user Experience Analytics Battery Health Device Runtime History
type UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderGetQueryParameters
}
// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderInternal instantiates a new UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) {
    m := &UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory/{userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder instantiates a new UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory for deviceManagement
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory for deviceManagement
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration user Experience Analytics Battery Health Device Runtime History
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration user Experience Analytics Battery Health Device Runtime History
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory in deviceManagement
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory in deviceManagement
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory for deviceManagement
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory for deviceManagement
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// GetWithResponseHandler user Experience Analytics Battery Health Device Runtime History
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) GetWithResponseHandler(requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler user Experience Analytics Battery Health Device Runtime History
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) GetWithResponseHandler(requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable), nil
}
// PatchWithResponseHandler update the navigation property userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory in deviceManagement
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory in deviceManagement
func (m *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable, requestConfiguration *UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
