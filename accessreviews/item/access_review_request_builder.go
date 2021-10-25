package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i104e81d59a18f4dfbeba211da9c23eb66a86bb3c0eb75b4a45999a45d1eb3e70 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/mydecisions"
    i1154d187957e7be461e0c1c2dc6e93905e63254a5e74becc376782f3efc12e92 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/reviewers"
    i21b8a2b755a37ce23f8670bdd4bc3877676e01bbed32cf8af1cfae68717a6edf "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/sendreminder"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i54ebdeadabb68caf8ef19ad8db4d5cf5f5f6e9dfb31a235af6595d6ea1c98ce0 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/stop"
    i82bb32bf1dad608f8fc8a4d162c2fad29e17008845c88f3f1939fca8639df33f "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/resetdecisions"
    i925eb082010f3d3144e37628501079bf5f355da9fe37d9a4fefdf1a66160bec6 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/applydecisions"
    ia35d22e7f0ea8d709ef1663feab561faae5adc5bcc2afe304389ef4bae030021 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/instances"
    ib639ae13590f5066540e44d9644ba6417e9c8d059dd2ed39c9d2f8a77498b365 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/decisions"
    i674dbad02771f1ef76ff9b64c61a2d6be676fe415bb8516739597b979e68b51c "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/decisions/item"
    i705becb6109e853394ffaa8ec9b9d2ca9bbc3f2d29dd2d667c6d76b0ba676d1c "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/instances/item"
    i9a1acb52dae713764d01984787eec0e51b4aaa96f162a6a5f7316ada71c7758b "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/reviewers/item"
    ia4fff597adb182643edeeb0e0ef392b36f38b279436254fb6004441a7cc386e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/accessreviews/item/mydecisions/item"
)

type AccessReviewRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type AccessReviewRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *AccessReviewRequestBuilder) ApplyDecisions()(*i925eb082010f3d3144e37628501079bf5f355da9fe37d9a4fefdf1a66160bec6.ApplyDecisionsRequestBuilder) {
    return i925eb082010f3d3144e37628501079bf5f355da9fe37d9a4fefdf1a66160bec6.NewApplyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewAccessReviewRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewRequestBuilder) {
    m := &AccessReviewRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/accessReviews/{accessReview_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewAccessReviewRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*AccessReviewRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessReviewRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *AccessReviewRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *AccessReviewRequestBuilder) CreateGetRequestInformation(q func (value *AccessReviewRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(AccessReviewRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *AccessReviewRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReview, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *AccessReviewRequestBuilder) Decisions()(*ib639ae13590f5066540e44d9644ba6417e9c8d059dd2ed39c9d2f8a77498b365.DecisionsRequestBuilder) {
    return ib639ae13590f5066540e44d9644ba6417e9c8d059dd2ed39c9d2f8a77498b365.NewDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewRequestBuilder) DecisionsById(id string)(*i674dbad02771f1ef76ff9b64c61a2d6be676fe415bb8516739597b979e68b51c.AccessReviewDecisionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["accessReviewDecision_id"] = id
    }
    return i674dbad02771f1ef76ff9b64c61a2d6be676fe415bb8516739597b979e68b51c.NewAccessReviewDecisionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessReviewRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *AccessReviewRequestBuilder) Get(q func (value *AccessReviewRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReview, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewAccessReview() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReview), nil
}
func (m *AccessReviewRequestBuilder) Instances()(*ia35d22e7f0ea8d709ef1663feab561faae5adc5bcc2afe304389ef4bae030021.InstancesRequestBuilder) {
    return ia35d22e7f0ea8d709ef1663feab561faae5adc5bcc2afe304389ef4bae030021.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewRequestBuilder) InstancesById(id string)(*i705becb6109e853394ffaa8ec9b9d2ca9bbc3f2d29dd2d667c6d76b0ba676d1c.AccessReviewRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["accessReview_id1"] = id
    }
    return i705becb6109e853394ffaa8ec9b9d2ca9bbc3f2d29dd2d667c6d76b0ba676d1c.NewAccessReviewRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessReviewRequestBuilder) MyDecisions()(*i104e81d59a18f4dfbeba211da9c23eb66a86bb3c0eb75b4a45999a45d1eb3e70.MyDecisionsRequestBuilder) {
    return i104e81d59a18f4dfbeba211da9c23eb66a86bb3c0eb75b4a45999a45d1eb3e70.NewMyDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewRequestBuilder) MyDecisionsById(id string)(*ia4fff597adb182643edeeb0e0ef392b36f38b279436254fb6004441a7cc386e4.AccessReviewDecisionRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["accessReviewDecision_id"] = id
    }
    return ia4fff597adb182643edeeb0e0ef392b36f38b279436254fb6004441a7cc386e4.NewAccessReviewDecisionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessReviewRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.AccessReview, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *AccessReviewRequestBuilder) ResetDecisions()(*i82bb32bf1dad608f8fc8a4d162c2fad29e17008845c88f3f1939fca8639df33f.ResetDecisionsRequestBuilder) {
    return i82bb32bf1dad608f8fc8a4d162c2fad29e17008845c88f3f1939fca8639df33f.NewResetDecisionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewRequestBuilder) Reviewers()(*i1154d187957e7be461e0c1c2dc6e93905e63254a5e74becc376782f3efc12e92.ReviewersRequestBuilder) {
    return i1154d187957e7be461e0c1c2dc6e93905e63254a5e74becc376782f3efc12e92.NewReviewersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewRequestBuilder) ReviewersById(id string)(*i9a1acb52dae713764d01984787eec0e51b4aaa96f162a6a5f7316ada71c7758b.AccessReviewReviewerRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["accessReviewReviewer_id"] = id
    }
    return i9a1acb52dae713764d01984787eec0e51b4aaa96f162a6a5f7316ada71c7758b.NewAccessReviewReviewerRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *AccessReviewRequestBuilder) SendReminder()(*i21b8a2b755a37ce23f8670bdd4bc3877676e01bbed32cf8af1cfae68717a6edf.SendReminderRequestBuilder) {
    return i21b8a2b755a37ce23f8670bdd4bc3877676e01bbed32cf8af1cfae68717a6edf.NewSendReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *AccessReviewRequestBuilder) Stop()(*i54ebdeadabb68caf8ef19ad8db4d5cf5f5f6e9dfb31a235af6595d6ea1c98ce0.StopRequestBuilder) {
    return i54ebdeadabb68caf8ef19ad8db4d5cf5f5f6e9dfb31a235af6595d6ea1c98ce0.NewStopRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
