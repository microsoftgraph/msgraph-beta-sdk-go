package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i5eebf6a3c08792a822f37ae918c3ae9ce1a890dccb9fa349cfa742f76d96a4aa "github.com/microsoftgraph/msgraph-beta-sdk-go/education/schools/item/classes/item/ref"
)

// EducationClassItemRequestBuilder builds and executes requests for operations under \education\schools\{educationSchool-id}\classes\{educationClass-id}
type EducationClassItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEducationClassItemRequestBuilderInternal instantiates a new EducationClassItemRequestBuilder and sets the default values.
func NewEducationClassItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationClassItemRequestBuilder) {
    m := &EducationClassItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/schools/{educationSchool%2Did}/classes/{educationClass%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationClassItemRequestBuilder instantiates a new EducationClassItemRequestBuilder and sets the default values.
func NewEducationClassItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationClassItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationClassItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of educationRoot entities.
func (m *EducationClassItemRequestBuilder) Ref()(*i5eebf6a3c08792a822f37ae918c3ae9ce1a890dccb9fa349cfa742f76d96a4aa.RefRequestBuilder) {
    return i5eebf6a3c08792a822f37ae918c3ae9ce1a890dccb9fa349cfa742f76d96a4aa.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
