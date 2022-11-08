package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i9cff64a5976355358c2613175db8a7ae8b770a08fa88dcaa7c5a293f1f0a8f8d "github.com/microsoftgraph/msgraph-beta-sdk-go/education/users/item/assignments/item/categories/item/ref"
)

// EducationCategoryItemRequestBuilder builds and executes requests for operations under \education\users\{educationUser-id}\assignments\{educationAssignment-id}\categories\{educationCategory-id}
type EducationCategoryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewEducationCategoryItemRequestBuilderInternal instantiates a new EducationCategoryItemRequestBuilder and sets the default values.
func NewEducationCategoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationCategoryItemRequestBuilder) {
    m := &EducationCategoryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/users/{educationUser%2Did}/assignments/{educationAssignment%2Did}/categories/{educationCategory%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationCategoryItemRequestBuilder instantiates a new EducationCategoryItemRequestBuilder and sets the default values.
func NewEducationCategoryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationCategoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationCategoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of educationRoot entities.
func (m *EducationCategoryItemRequestBuilder) Ref()(*i9cff64a5976355358c2613175db8a7ae8b770a08fa88dcaa7c5a293f1f0a8f8d.RefRequestBuilder) {
    return i9cff64a5976355358c2613175db8a7ae8b770a08fa88dcaa7c5a293f1f0a8f8d.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
