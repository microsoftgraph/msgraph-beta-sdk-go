package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphExportDeviceAndAppManagementDataExportDeviceAndAppManagementDataRequestBuilder provides operations to call the exportDeviceAndAppManagementData method.
type ItemMicrosoftGraphExportDeviceAndAppManagementDataExportDeviceAndAppManagementDataRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphExportDeviceAndAppManagementDataExportDeviceAndAppManagementDataRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphExportDeviceAndAppManagementDataExportDeviceAndAppManagementDataRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMicrosoftGraphExportDeviceAndAppManagementDataExportDeviceAndAppManagementDataRequestBuilderInternal instantiates a new ExportDeviceAndAppManagementDataRequestBuilder and sets the default values.
func NewItemMicrosoftGraphExportDeviceAndAppManagementDataExportDeviceAndAppManagementDataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphExportDeviceAndAppManagementDataExportDeviceAndAppManagementDataRequestBuilder) {
    m := &ItemMicrosoftGraphExportDeviceAndAppManagementDataExportDeviceAndAppManagementDataRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/microsoft.graph.exportDeviceAndAppManagementData()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphExportDeviceAndAppManagementDataExportDeviceAndAppManagementDataRequestBuilder instantiates a new ExportDeviceAndAppManagementDataRequestBuilder and sets the default values.
func NewItemMicrosoftGraphExportDeviceAndAppManagementDataExportDeviceAndAppManagementDataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphExportDeviceAndAppManagementDataExportDeviceAndAppManagementDataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphExportDeviceAndAppManagementDataExportDeviceAndAppManagementDataRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function exportDeviceAndAppManagementData
func (m *ItemMicrosoftGraphExportDeviceAndAppManagementDataExportDeviceAndAppManagementDataRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMicrosoftGraphExportDeviceAndAppManagementDataExportDeviceAndAppManagementDataRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementDataable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAndAppManagementDataFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementDataable), nil
}
// ToGetRequestInformation invoke function exportDeviceAndAppManagementData
func (m *ItemMicrosoftGraphExportDeviceAndAppManagementDataExportDeviceAndAppManagementDataRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMicrosoftGraphExportDeviceAndAppManagementDataExportDeviceAndAppManagementDataRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
