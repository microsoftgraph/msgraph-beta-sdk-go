package trustframework

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// KeySetsItemUploadSecretRequestBuilder provides operations to call the uploadSecret method.
type KeySetsItemUploadSecretRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// KeySetsItemUploadSecretRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type KeySetsItemUploadSecretRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewKeySetsItemUploadSecretRequestBuilderInternal instantiates a new UploadSecretRequestBuilder and sets the default values.
func NewKeySetsItemUploadSecretRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeySetsItemUploadSecretRequestBuilder) {
    m := &KeySetsItemUploadSecretRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/trustFramework/keySets/{trustFrameworkKeySet%2Did}/uploadSecret", pathParameters),
    }
    return m
}
// NewKeySetsItemUploadSecretRequestBuilder instantiates a new UploadSecretRequestBuilder and sets the default values.
func NewKeySetsItemUploadSecretRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*KeySetsItemUploadSecretRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewKeySetsItemUploadSecretRequestBuilderInternal(urlParams, requestAdapter)
}
// Post upload a plain text secret to a trustFrameworkKeyset. Examples of secrets are application secrets in Azure Active Directory, Google, Facebook, or any other identity provider. his method returns trustFrameworkKey. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/trustframeworkkeyset-uploadsecret?view=graph-rest-1.0
func (m *KeySetsItemUploadSecretRequestBuilder) Post(ctx context.Context, body KeySetsItemUploadSecretPostRequestBodyable, requestConfiguration *KeySetsItemUploadSecretRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTrustFrameworkKeyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrustFrameworkKeyable), nil
}
// ToPostRequestInformation upload a plain text secret to a trustFrameworkKeyset. Examples of secrets are application secrets in Azure Active Directory, Google, Facebook, or any other identity provider. his method returns trustFrameworkKey. This API is available in the following national cloud deployments.
func (m *KeySetsItemUploadSecretRequestBuilder) ToPostRequestInformation(ctx context.Context, body KeySetsItemUploadSecretPostRequestBodyable, requestConfiguration *KeySetsItemUploadSecretRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *KeySetsItemUploadSecretRequestBuilder) WithUrl(rawUrl string)(*KeySetsItemUploadSecretRequestBuilder) {
    return NewKeySetsItemUploadSecretRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
