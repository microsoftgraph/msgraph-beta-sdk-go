package listitem

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0ba613bc86957dec7613cd9e46e09b8d7a454f90858c274160c7905afe11ee85 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/listitem/fields"
    i176af9c67c2a6f13954c79a5ef04601c3bbd7228362b2769e298f57a1e628d54 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/listitem/driveitem"
    i682982f55902c7b37f180eb10b3a4ad9eb1b3dbea1f3a19d5353ff4e575ec0a5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/listitem/createlink"
    i6ae89857ee2689affd2566e1e5fbb65b2221728f3b7cf966f1fea4e3dc09dbcd "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/listitem/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    iacc2b1de9f4a244ca53c8d57642a9fc80c972ac606ee3825112e82bb46ed5a4d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/listitem/analytics"
    id8d67fb5df16d70b327e82786d4905ad12bdcf3127239632d19764b65b876200 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/listitem/activities"
    if27c9f161a3700f9ff3074a79ef80cbb82ad4babf4ea2c82e141de65a57cc198 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/listitem/versions"
    i65b54269cd09d2c9f4762673baab02ed2d5b0a7ad6f9370129144443f864a555 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/listitem/activities/item"
    if19a367596d4a0b94e88c86e5e1c913bc12f813e4bd0b18909551f197ce4a6a0 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/following/item/listitem/versions/item"
)

// ListItemRequestBuilder provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
type ListItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ListItemRequestBuilderDeleteOptions options for Delete
type ListItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListItemRequestBuilderGetOptions options for Get
type ListItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ListItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListItemRequestBuilderGetQueryParameters for drives in SharePoint, the associated document library list item. Read-only. Nullable.
type ListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ListItemRequestBuilderPatchOptions options for Patch
type ListItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItemable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ListItemRequestBuilder) Activities()(*id8d67fb5df16d70b327e82786d4905ad12bdcf3127239632d19764b65b876200.ActivitiesRequestBuilder) {
    return id8d67fb5df16d70b327e82786d4905ad12bdcf3127239632d19764b65b876200.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.following.item.listItem.activities.item collection
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*i65b54269cd09d2c9f4762673baab02ed2d5b0a7ad6f9370129144443f864a555.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i65b54269cd09d2c9f4762673baab02ed2d5b0a7ad6f9370129144443f864a555.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Analytics()(*iacc2b1de9f4a244ca53c8d57642a9fc80c972ac606ee3825112e82bb46ed5a4d.AnalyticsRequestBuilder) {
    return iacc2b1de9f4a244ca53c8d57642a9fc80c972ac606ee3825112e82bb46ed5a4d.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemRequestBuilderInternal instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/drives/{drive_id}/following/{driveItem_id}/listItem{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListItemRequestBuilder instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property listItem for users
func (m *ListItemRequestBuilder) CreateDeleteRequestInformation(options *ListItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation for drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *ListItemRequestBuilder) CreateGetRequestInformation(options *ListItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ListItemRequestBuilder) CreateLink()(*i682982f55902c7b37f180eb10b3a4ad9eb1b3dbea1f3a19d5353ff4e575ec0a5.CreateLinkRequestBuilder) {
    return i682982f55902c7b37f180eb10b3a4ad9eb1b3dbea1f3a19d5353ff4e575ec0a5.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property listItem in users
func (m *ListItemRequestBuilder) CreatePatchRequestInformation(options *ListItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property listItem for users
func (m *ListItemRequestBuilder) Delete(options *ListItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemRequestBuilder) DriveItem()(*i176af9c67c2a6f13954c79a5ef04601c3bbd7228362b2769e298f57a1e628d54.DriveItemRequestBuilder) {
    return i176af9c67c2a6f13954c79a5ef04601c3bbd7228362b2769e298f57a1e628d54.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*i0ba613bc86957dec7613cd9e46e09b8d7a454f90858c274160c7905afe11ee85.FieldsRequestBuilder) {
    return i0ba613bc86957dec7613cd9e46e09b8d7a454f90858c274160c7905afe11ee85.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get for drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *ListItemRequestBuilder) Get(options *ListItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateListItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItemable), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, interval *string, endDateTime *string)(*i6ae89857ee2689affd2566e1e5fbb65b2221728f3b7cf966f1fea4e3dc09dbcd.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i6ae89857ee2689affd2566e1e5fbb65b2221728f3b7cf966f1fea4e3dc09dbcd.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, interval, endDateTime);
}
// Patch update the navigation property listItem in users
func (m *ListItemRequestBuilder) Patch(options *ListItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemRequestBuilder) Versions()(*if27c9f161a3700f9ff3074a79ef80cbb82ad4babf4ea2c82e141de65a57cc198.VersionsRequestBuilder) {
    return if27c9f161a3700f9ff3074a79ef80cbb82ad4babf4ea2c82e141de65a57cc198.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.following.item.listItem.versions.item collection
func (m *ListItemRequestBuilder) VersionsById(id string)(*if19a367596d4a0b94e88c86e5e1c913bc12f813e4bd0b18909551f197ce4a6a0.ListItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return if19a367596d4a0b94e88c86e5e1c913bc12f813e4bd0b18909551f197ce4a6a0.NewListItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
