package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0b59187fa51909a371f8b9858abba00a04491a0a4774c16f37ed7ad230dd3c80 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/multivalueextendedproperties"
    i181320aa93ef4321014b74609a44e3f5a32133f9580d21d62327fa7bbb5cda54 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/snoozereminder"
    i19dd778a11375dddbb2d2611e805800876e81cba206c523af607ec0876d42638 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/dismissreminder"
    i336acfa1272f48c2fe797109baa0a33d28eb0913b952cd433351e0dca2bfea27 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/calendar"
    i5793f33136eb9238f8d107847a380f870d05d5e35af7d529b95b28bce2609cf5 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/tentativelyaccept"
    i60fcaa43c34fa0b127246ffcfd3c522b1bb552c2093011690608d937556cadcb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/cancel"
    i83b5cc71f24decf0fa568a1fea8c63f1115a6c0cda7a6f7a48cb1cc942e4eb75 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/attachments"
    i8553f6cdfc8767ac4a225c751996e9393fc28a393eea6cf2b2cb3fb7537c37d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/singlevalueextendedproperties"
    ia64b443dfec1fdf0f96b97e8ce3dc6342bfb217cb244a28fb51f4e2fac80098a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/extensions"
    ib22faf4889a4ff940189a07d0634017ab1700b96946ffe4c7859b1c600c95b13 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/accept"
    ic699aa89cf0fa5dce893a4e798799256b1490ee3a0e7945a5380b385a6e1ebf4 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/forward"
    id81837152e19e46f34e39a5f29933b7bf6301a3b447bcdab80bfcac087daa8df "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/decline"
    ieb69cb4984db33666e55e73bcf1987aa77b5daea0acac2114d5c716cc8ee9590 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances"
    if3620ac1eeab8b06c743846feb104c5b5c1c860caa9b835409096fbd72f4dbc2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences"
    i08d2438ecad8b31d14bf1588846d430634768d2d73ce954a9c7459f7128dee29 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/instances/item"
    i0e2c2c7317d5772b90fb64d735c3dd692b4af05af875d54fd676524e724cd512 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/exceptionoccurrences/item"
    i2c60bfae6ea65be12b4d02f34ad3d2bfe71fdc157260307ad234bfb2acc44811 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/multivalueextendedproperties/item"
    i5f7bc65efefc422b0d02ea65c0c11a8e28f4805da2f6d82655cafcb3c05ae1a7 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/singlevalueextendedproperties/item"
    ia9e0f6249aba2f425cdb963bff747c708eaa818659d217e6df3cd92448321a0f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/attachments/item"
    ic6abe1dc87b2423ac4d607f9d80cfadb179e722bf3ebc7a79351f62c697aff2d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/calendars/item/events/item/extensions/item"
)

type EventRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type EventRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escaped []string;
}
func (m *EventRequestBuilder) Accept()(*ib22faf4889a4ff940189a07d0634017ab1700b96946ffe4c7859b1c600c95b13.AcceptRequestBuilder) {
    return ib22faf4889a4ff940189a07d0634017ab1700b96946ffe4c7859b1c600c95b13.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i83b5cc71f24decf0fa568a1fea8c63f1115a6c0cda7a6f7a48cb1cc942e4eb75.AttachmentsRequestBuilder) {
    return i83b5cc71f24decf0fa568a1fea8c63f1115a6c0cda7a6f7a48cb1cc942e4eb75.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) AttachmentsById(id string)(*ia9e0f6249aba2f425cdb963bff747c708eaa818659d217e6df3cd92448321a0f.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return ia9e0f6249aba2f425cdb963bff747c708eaa818659d217e6df3cd92448321a0f.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i336acfa1272f48c2fe797109baa0a33d28eb0913b952cd433351e0dca2bfea27.CalendarRequestBuilder) {
    return i336acfa1272f48c2fe797109baa0a33d28eb0913b952cd433351e0dca2bfea27.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i60fcaa43c34fa0b127246ffcfd3c522b1bb552c2093011690608d937556cadcb.CancelRequestBuilder) {
    return i60fcaa43c34fa0b127246ffcfd3c522b1bb552c2093011690608d937556cadcb.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/users/{user_id}/calendars/{calendar_id}/events/{event_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewEventRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEventRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *EventRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventRequestBuilder) CreateGetRequestInformation(q func (value *EventRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(EventRequestBuilderGetQueryParameters)
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
func (m *EventRequestBuilder) CreatePatchRequestInformation(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EventRequestBuilder) Decline()(*id81837152e19e46f34e39a5f29933b7bf6301a3b447bcdab80bfcac087daa8df.DeclineRequestBuilder) {
    return id81837152e19e46f34e39a5f29933b7bf6301a3b447bcdab80bfcac087daa8df.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EventRequestBuilder) DismissReminder()(*i19dd778a11375dddbb2d2611e805800876e81cba206c523af607ec0876d42638.DismissReminderRequestBuilder) {
    return i19dd778a11375dddbb2d2611e805800876e81cba206c523af607ec0876d42638.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrences()(*if3620ac1eeab8b06c743846feb104c5b5c1c860caa9b835409096fbd72f4dbc2.ExceptionOccurrencesRequestBuilder) {
    return if3620ac1eeab8b06c743846feb104c5b5c1c860caa9b835409096fbd72f4dbc2.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrencesById(id string)(*i0e2c2c7317d5772b90fb64d735c3dd692b4af05af875d54fd676524e724cd512.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i0e2c2c7317d5772b90fb64d735c3dd692b4af05af875d54fd676524e724cd512.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*ia64b443dfec1fdf0f96b97e8ce3dc6342bfb217cb244a28fb51f4e2fac80098a.ExtensionsRequestBuilder) {
    return ia64b443dfec1fdf0f96b97e8ce3dc6342bfb217cb244a28fb51f4e2fac80098a.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExtensionsById(id string)(*ic6abe1dc87b2423ac4d607f9d80cfadb179e722bf3ebc7a79351f62c697aff2d.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ic6abe1dc87b2423ac4d607f9d80cfadb179e722bf3ebc7a79351f62c697aff2d.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*ic699aa89cf0fa5dce893a4e798799256b1490ee3a0e7945a5380b385a6e1ebf4.ForwardRequestBuilder) {
    return ic699aa89cf0fa5dce893a4e798799256b1490ee3a0e7945a5380b385a6e1ebf4.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Get(q func (value *EventRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEvent() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event), nil
}
func (m *EventRequestBuilder) Instances()(*ieb69cb4984db33666e55e73bcf1987aa77b5daea0acac2114d5c716cc8ee9590.InstancesRequestBuilder) {
    return ieb69cb4984db33666e55e73bcf1987aa77b5daea0acac2114d5c716cc8ee9590.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) InstancesById(id string)(*i08d2438ecad8b31d14bf1588846d430634768d2d73ce954a9c7459f7128dee29.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i08d2438ecad8b31d14bf1588846d430634768d2d73ce954a9c7459f7128dee29.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i0b59187fa51909a371f8b9858abba00a04491a0a4774c16f37ed7ad230dd3c80.MultiValueExtendedPropertiesRequestBuilder) {
    return i0b59187fa51909a371f8b9858abba00a04491a0a4774c16f37ed7ad230dd3c80.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i2c60bfae6ea65be12b4d02f34ad3d2bfe71fdc157260307ad234bfb2acc44811.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i2c60bfae6ea65be12b4d02f34ad3d2bfe71fdc157260307ad234bfb2acc44811.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Patch(body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Event, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i8553f6cdfc8767ac4a225c751996e9393fc28a393eea6cf2b2cb3fb7537c37d6.SingleValueExtendedPropertiesRequestBuilder) {
    return i8553f6cdfc8767ac4a225c751996e9393fc28a393eea6cf2b2cb3fb7537c37d6.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i5f7bc65efefc422b0d02ea65c0c11a8e28f4805da2f6d82655cafcb3c05ae1a7.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i5f7bc65efefc422b0d02ea65c0c11a8e28f4805da2f6d82655cafcb3c05ae1a7.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i181320aa93ef4321014b74609a44e3f5a32133f9580d21d62327fa7bbb5cda54.SnoozeReminderRequestBuilder) {
    return i181320aa93ef4321014b74609a44e3f5a32133f9580d21d62327fa7bbb5cda54.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*i5793f33136eb9238f8d107847a380f870d05d5e35af7d529b95b28bce2609cf5.TentativelyAcceptRequestBuilder) {
    return i5793f33136eb9238f8d107847a380f870d05d5e35af7d529b95b28bce2609cf5.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}