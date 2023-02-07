package users

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder provides operations to call the exportDeviceAndAppManagementData method.
type ItemMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal instantiates a new MicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder and sets the default values.
func NewItemMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, skip *int32, top *int32)(*ItemMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    m := &ItemMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/microsoft.graph.exportDeviceAndAppManagementData(skip={skip},top={top})";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if skip != nil {
        urlTplParams["skip"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*skip), 10)
    }
    if top != nil {
        urlTplParams["top"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*top), 10)
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder instantiates a new MicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder and sets the default values.
func NewItemMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function exportDeviceAndAppManagementData
func (m *ItemMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementDataable, error) {
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
func (m *ItemMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemMicrosoftGraphExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
