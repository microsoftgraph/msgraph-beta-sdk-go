package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilder provides operations to call the addLanguageFiles method.
type DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilderInternal instantiates a new AddLanguageFilesRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilder) {
    m := &DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyUploadedDefinitionFiles/{groupPolicyUploadedDefinitionFile%2Did}/microsoft.graph.addLanguageFiles";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilder instantiates a new AddLanguageFilesRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action addLanguageFiles
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesPostRequestBodyable, requestConfiguration *DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action addLanguageFiles
func (m *DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesPostRequestBodyable, requestConfiguration *DeviceManagementGroupPolicyUploadedDefinitionFilesItemAddLanguageFilesRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
