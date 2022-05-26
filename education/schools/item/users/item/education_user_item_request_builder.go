package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i0769b3ec1aee8601e6d6b9a87117bb7e2e546451df472fb7d5e9cdde6a433211 "github.com/microsoftgraph/msgraph-beta-sdk-go/education/schools/item/users/item/ref"
)

// EducationUserItemRequestBuilder builds and executes requests for operations under \education\schools\{educationSchool-id}\users\{educationUser-id}
type EducationUserItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEducationUserItemRequestBuilderInternal instantiates a new EducationUserItemRequestBuilder and sets the default values.
func NewEducationUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationUserItemRequestBuilder) {
    m := &EducationUserItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/schools/{educationSchool%2Did}/users/{educationUser%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationUserItemRequestBuilder instantiates a new EducationUserItemRequestBuilder and sets the default values.
func NewEducationUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref the ref property
func (m *EducationUserItemRequestBuilder) Ref()(*i0769b3ec1aee8601e6d6b9a87117bb7e2e546451df472fb7d5e9cdde6a433211.RefRequestBuilder) {
    return i0769b3ec1aee8601e6d6b9a87117bb7e2e546451df472fb7d5e9cdde6a433211.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
