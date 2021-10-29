package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i12f6208fa635d862ebcbb2efc3fc2ad520e8b11af30ce47c1b7aae81087c3fc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/cancel"
    i24648f3b72cb6913a9588b8f30f65a5861a94ceec8a08eafb72d128c878c149c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances"
    i51f82170e9d7e02f39adef69c5ad83f6c68a8391a942267ee25404779c09df02 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/multivalueextendedproperties"
    i689722744aa2c7a040ce1eea267991c6a8df229fedaf72403c3cd9fd504dc586 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/forward"
    i6a81bfdd259fd587f1bd56d9b047bb56371d6496c53c194845624ba2a958a171 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/snoozereminder"
    i6ccf1ed39358cfe85d5d5e1ed7cc04329f170b3143984349ea66a828df15091d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/attachments"
    i7ea3353b3d2a5a1c8706154adf849e0629e8535acc87cb16c1b741d012b14551 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/calendar"
    i990c9d56acce6df7a509f9f537bf1c3651710e2d54e9785ef179a00a350d7624 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/singlevalueextendedproperties"
    i9f03e8930612bb38b0ee316af407e0390228dbd9c1ce9d5b920322626d4a7ccb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/dismissreminder"
    ia519bb57124ce78ce804d4c8fde3bad9e5de6768daf6ccd8b683b306570d66b9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences"
    ib35ba17b740b364bef12429dc54a7fc3e5647944b9c94a261b2afb41557bdedb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/tentativelyaccept"
    id0b046c29ca78c81c5fc1cdab2b1e6aad652b41020d22dd887163e374abd4714 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/accept"
    idf4c71a10741a39450889d693f4932da9fbbd7fd71fa3872d803f21687c7eeb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/decline"
    ieb846c3f403e62d9249e7d035caf7a3aae7a2c964e81219a2fdbf7417e1f6fbb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/extensions"
    i2ed60d96750ee240e18770a7ac42e716a5f27688d6f8e7757c055b8d4564d81c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/exceptionoccurrences/item"
    i3ab7e30c83492e2917ff34bdeecc0fca059eeb741d94af5ebf45410534e47d9f "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/attachments/item"
    i7a7c70eae23ab3508021e2f69db1d970a3bd21e3036bbb44f0a1291701a6d4da "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/singlevalueextendedproperties/item"
    ia1b7d92a7d10b8e6e611cc5be140d3a64be2800b960d7d8781caf1650e2f0c08 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/instances/item"
    ib2a745f0d686f44f5e53733974718efc67fcdb195c1f77b50aa122d3f058cb4e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/multivalueextendedproperties/item"
    id96b14704d4aef544449833c6800c5d3620733e0995ae628d5280e53eed51574 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/calendars/item/calendarview/item/extensions/item"
)

type EventRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type EventRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    EndDateTime *string;
    Expand []string;
    Select_escaped []string;
    StartDateTime *string;
}
func (m *EventRequestBuilder) Accept()(*id0b046c29ca78c81c5fc1cdab2b1e6aad652b41020d22dd887163e374abd4714.AcceptRequestBuilder) {
    return id0b046c29ca78c81c5fc1cdab2b1e6aad652b41020d22dd887163e374abd4714.NewAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Attachments()(*i6ccf1ed39358cfe85d5d5e1ed7cc04329f170b3143984349ea66a828df15091d.AttachmentsRequestBuilder) {
    return i6ccf1ed39358cfe85d5d5e1ed7cc04329f170b3143984349ea66a828df15091d.NewAttachmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) AttachmentsById(id string)(*i3ab7e30c83492e2917ff34bdeecc0fca059eeb741d94af5ebf45410534e47d9f.AttachmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment_id"] = id
    }
    return i3ab7e30c83492e2917ff34bdeecc0fca059eeb741d94af5ebf45410534e47d9f.NewAttachmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Calendar()(*i7ea3353b3d2a5a1c8706154adf849e0629e8535acc87cb16c1b741d012b14551.CalendarRequestBuilder) {
    return i7ea3353b3d2a5a1c8706154adf849e0629e8535acc87cb16c1b741d012b14551.NewCalendarRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) Cancel()(*i12f6208fa635d862ebcbb2efc3fc2ad520e8b11af30ce47c1b7aae81087c3fc4.CancelRequestBuilder) {
    return i12f6208fa635d862ebcbb2efc3fc2ad520e8b11af30ce47c1b7aae81087c3fc4.NewCancelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func NewEventRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EventRequestBuilder) {
    m := &EventRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/me/calendars/{calendar_id}/calendarView/{event_id}{?startDateTime,endDateTime,select,expand}";
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
func (m *EventRequestBuilder) Decline()(*idf4c71a10741a39450889d693f4932da9fbbd7fd71fa3872d803f21687c7eeb0.DeclineRequestBuilder) {
    return idf4c71a10741a39450889d693f4932da9fbbd7fd71fa3872d803f21687c7eeb0.NewDeclineRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) DismissReminder()(*i9f03e8930612bb38b0ee316af407e0390228dbd9c1ce9d5b920322626d4a7ccb.DismissReminderRequestBuilder) {
    return i9f03e8930612bb38b0ee316af407e0390228dbd9c1ce9d5b920322626d4a7ccb.NewDismissReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrences()(*ia519bb57124ce78ce804d4c8fde3bad9e5de6768daf6ccd8b683b306570d66b9.ExceptionOccurrencesRequestBuilder) {
    return ia519bb57124ce78ce804d4c8fde3bad9e5de6768daf6ccd8b683b306570d66b9.NewExceptionOccurrencesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExceptionOccurrencesById(id string)(*i2ed60d96750ee240e18770a7ac42e716a5f27688d6f8e7757c055b8d4564d81c.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i2ed60d96750ee240e18770a7ac42e716a5f27688d6f8e7757c055b8d4564d81c.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Extensions()(*ieb846c3f403e62d9249e7d035caf7a3aae7a2c964e81219a2fdbf7417e1f6fbb.ExtensionsRequestBuilder) {
    return ieb846c3f403e62d9249e7d035caf7a3aae7a2c964e81219a2fdbf7417e1f6fbb.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) ExtensionsById(id string)(*id96b14704d4aef544449833c6800c5d3620733e0995ae628d5280e53eed51574.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return id96b14704d4aef544449833c6800c5d3620733e0995ae628d5280e53eed51574.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) Forward()(*i689722744aa2c7a040ce1eea267991c6a8df229fedaf72403c3cd9fd504dc586.ForwardRequestBuilder) {
    return i689722744aa2c7a040ce1eea267991c6a8df229fedaf72403c3cd9fd504dc586.NewForwardRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *EventRequestBuilder) Instances()(*i24648f3b72cb6913a9588b8f30f65a5861a94ceec8a08eafb72d128c878c149c.InstancesRequestBuilder) {
    return i24648f3b72cb6913a9588b8f30f65a5861a94ceec8a08eafb72d128c878c149c.NewInstancesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) InstancesById(id string)(*ia1b7d92a7d10b8e6e611cc5be140d3a64be2800b960d7d8781caf1650e2f0c08.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ia1b7d92a7d10b8e6e611cc5be140d3a64be2800b960d7d8781caf1650e2f0c08.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedProperties()(*i51f82170e9d7e02f39adef69c5ad83f6c68a8391a942267ee25404779c09df02.MultiValueExtendedPropertiesRequestBuilder) {
    return i51f82170e9d7e02f39adef69c5ad83f6c68a8391a942267ee25404779c09df02.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ib2a745f0d686f44f5e53733974718efc67fcdb195c1f77b50aa122d3f058cb4e.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return ib2a745f0d686f44f5e53733974718efc67fcdb195c1f77b50aa122d3f058cb4e.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *EventRequestBuilder) SingleValueExtendedProperties()(*i990c9d56acce6df7a509f9f537bf1c3651710e2d54e9785ef179a00a350d7624.SingleValueExtendedPropertiesRequestBuilder) {
    return i990c9d56acce6df7a509f9f537bf1c3651710e2d54e9785ef179a00a350d7624.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i7a7c70eae23ab3508021e2f69db1d970a3bd21e3036bbb44f0a1291701a6d4da.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i7a7c70eae23ab3508021e2f69db1d970a3bd21e3036bbb44f0a1291701a6d4da.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EventRequestBuilder) SnoozeReminder()(*i6a81bfdd259fd587f1bd56d9b047bb56371d6496c53c194845624ba2a958a171.SnoozeReminderRequestBuilder) {
    return i6a81bfdd259fd587f1bd56d9b047bb56371d6496c53c194845624ba2a958a171.NewSnoozeReminderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EventRequestBuilder) TentativelyAccept()(*ib35ba17b740b364bef12429dc54a7fc3e5647944b9c94a261b2afb41557bdedb.TentativelyAcceptRequestBuilder) {
    return ib35ba17b740b364bef12429dc54a7fc3e5647944b9c94a261b2afb41557bdedb.NewTentativelyAcceptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}