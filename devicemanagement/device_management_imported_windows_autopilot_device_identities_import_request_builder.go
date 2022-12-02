package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportRequestBuilder provides operations to call the import method.
type DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportRequestBuilderInternal instantiates a new ImportRequestBuilder and sets the default values.
func NewDeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportRequestBuilder) {
    m := &DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/importedWindowsAutopilotDeviceIdentities/microsoft.graph.import";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportRequestBuilder instantiates a new ImportRequestBuilder and sets the default values.
func NewDeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action import
func (m *DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportPostRequestBodyable, requestConfiguration *DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action import
func (m *DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportPostRequestBodyable, requestConfiguration *DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementImportedWindowsAutopilotDeviceIdentitiesImportResponseable), nil
}
