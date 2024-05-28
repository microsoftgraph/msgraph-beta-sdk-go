package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilder provides operations to call the addLanguageFiles method.
type GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilderInternal instantiates a new GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilder and sets the default values.
func NewGrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilder) {
    m := &GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyUploadedDefinitionFiles/{groupPolicyUploadedDefinitionFile%2Did}/addLanguageFiles", pathParameters),
    }
    return m
}
// NewGrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilder instantiates a new GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilder and sets the default values.
func NewGrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action addLanguageFiles
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilder) Post(ctx context.Context, body GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesPostRequestBodyable, requestConfiguration *GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action addLanguageFiles
// returns a *RequestInformation when successful
func (m *GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesPostRequestBodyable, requestConfiguration *GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilder when successful
func (m *GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilder) WithUrl(rawUrl string)(*GrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilder) {
    return NewGrouppolicyuploadeddefinitionfilesItemAddlanguagefilesAddLanguageFilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
