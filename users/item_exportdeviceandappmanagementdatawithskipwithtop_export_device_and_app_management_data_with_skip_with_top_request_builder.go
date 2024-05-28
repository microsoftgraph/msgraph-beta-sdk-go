package users

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder provides operations to call the exportDeviceAndAppManagementData method.
type ItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal instantiates a new ItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder and sets the default values.
func NewItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, skip *int32, top *int32)(*ItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    m := &ItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/exportDeviceAndAppManagementData(skip={skip},top={top})", pathParameters),
    }
    if skip != nil {
        m.BaseRequestBuilder.PathParameters["skip"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*skip), 10)
    }
    if top != nil {
        m.BaseRequestBuilder.PathParameters["top"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*top), 10)
    }
    return m
}
// NewItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder instantiates a new ItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder and sets the default values.
func NewItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function exportDeviceAndAppManagementData
// returns a DeviceAndAppManagementDataable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementDataable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAndAppManagementDataFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementDataable), nil
}
// ToGetRequestInformation invoke function exportDeviceAndAppManagementData
// returns a *RequestInformation when successful
func (m *ItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder when successful
func (m *ItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) WithUrl(rawUrl string)(*ItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder) {
    return NewItemExportdeviceandappmanagementdatawithskipwithtopExportDeviceAndAppManagementDataWithSkipWithTopRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
