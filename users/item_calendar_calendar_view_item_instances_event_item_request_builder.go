package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.event entity.
type ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCalendarCalendarViewItemInstancesEventItemRequestBuilderGetQueryParameters the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
type ItemCalendarCalendarViewItemInstancesEventItemRequestBuilderGetQueryParameters struct {
    // The end date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T20:00:00-08:00
    EndDateTime *string
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // The start date and time of the time range, represented in ISO 8601 format. For example, 2019-11-08T19:00:00-08:00
    StartDateTime *string
}
// ItemCalendarCalendarViewItemInstancesEventItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCalendarCalendarViewItemInstancesEventItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilderGetQueryParameters
}
// Accept provides operations to call the accept method.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) Accept()(*ItemCalendarCalendarViewItemInstancesItemAcceptRequestBuilder) {
    return NewItemCalendarCalendarViewItemInstancesItemAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Attachments provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) Attachments()(*ItemCalendarCalendarViewItemInstancesItemAttachmentsRequestBuilder) {
    return NewItemCalendarCalendarViewItemInstancesItemAttachmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AttachmentsById provides operations to manage the attachments property of the microsoft.graph.event entity.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) AttachmentsById(id string)(*ItemCalendarCalendarViewItemInstancesItemAttachmentsAttachmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["attachment%2Did"] = id
    }
    return NewItemCalendarCalendarViewItemInstancesItemAttachmentsAttachmentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Calendar provides operations to manage the calendar property of the microsoft.graph.event entity.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) Calendar()(*ItemCalendarCalendarViewItemInstancesItemCalendarRequestBuilder) {
    return NewItemCalendarCalendarViewItemInstancesItemCalendarRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) Cancel()(*ItemCalendarCalendarViewItemInstancesItemCancelRequestBuilder) {
    return NewItemCalendarCalendarViewItemInstancesItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCalendarCalendarViewItemInstancesEventItemRequestBuilderInternal instantiates a new EventItemRequestBuilder and sets the default values.
func NewItemCalendarCalendarViewItemInstancesEventItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) {
    m := &ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/calendar/calendarView/{event%2Did}/instances/{event%2Did1}{?startDateTime*,endDateTime*,%24select}", pathParameters),
    }
    return m
}
// NewItemCalendarCalendarViewItemInstancesEventItemRequestBuilder instantiates a new EventItemRequestBuilder and sets the default values.
func NewItemCalendarCalendarViewItemInstancesEventItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCalendarCalendarViewItemInstancesEventItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Decline provides operations to call the decline method.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) Decline()(*ItemCalendarCalendarViewItemInstancesItemDeclineRequestBuilder) {
    return NewItemCalendarCalendarViewItemInstancesItemDeclineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DismissReminder provides operations to call the dismissReminder method.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) DismissReminder()(*ItemCalendarCalendarViewItemInstancesItemDismissReminderRequestBuilder) {
    return NewItemCalendarCalendarViewItemInstancesItemDismissReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExceptionOccurrences provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) ExceptionOccurrences()(*ItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesRequestBuilder) {
    return NewItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExceptionOccurrencesById provides operations to manage the exceptionOccurrences property of the microsoft.graph.event entity.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) ExceptionOccurrencesById(id string)(*ItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did2"] = id
    }
    return NewItemCalendarCalendarViewItemInstancesItemExceptionOccurrencesEventItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Extensions provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) Extensions()(*ItemCalendarCalendarViewItemInstancesItemExtensionsRequestBuilder) {
    return NewItemCalendarCalendarViewItemInstancesItemExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExtensionsById provides operations to manage the extensions property of the microsoft.graph.event entity.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) ExtensionsById(id string)(*ItemCalendarCalendarViewItemInstancesItemExtensionsExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension%2Did"] = id
    }
    return NewItemCalendarCalendarViewItemInstancesItemExtensionsExtensionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// Forward provides operations to call the forward method.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) Forward()(*ItemCalendarCalendarViewItemInstancesItemForwardRequestBuilder) {
    return NewItemCalendarCalendarViewItemInstancesItemForwardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Eventable), nil
}
// MultiValueExtendedProperties provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) MultiValueExtendedProperties()(*ItemCalendarCalendarViewItemInstancesItemMultiValueExtendedPropertiesRequestBuilder) {
    return NewItemCalendarCalendarViewItemInstancesItemMultiValueExtendedPropertiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MultiValueExtendedPropertiesById provides operations to manage the multiValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*ItemCalendarCalendarViewItemInstancesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemCalendarCalendarViewItemInstancesItemMultiValueExtendedPropertiesMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// SingleValueExtendedProperties provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) SingleValueExtendedProperties()(*ItemCalendarCalendarViewItemInstancesItemSingleValueExtendedPropertiesRequestBuilder) {
    return NewItemCalendarCalendarViewItemInstancesItemSingleValueExtendedPropertiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SingleValueExtendedPropertiesById provides operations to manage the singleValueExtendedProperties property of the microsoft.graph.event entity.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ItemCalendarCalendarViewItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return NewItemCalendarCalendarViewItemInstancesItemSingleValueExtendedPropertiesSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// SnoozeReminder provides operations to call the snoozeReminder method.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) SnoozeReminder()(*ItemCalendarCalendarViewItemInstancesItemSnoozeReminderRequestBuilder) {
    return NewItemCalendarCalendarViewItemInstancesItemSnoozeReminderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TentativelyAccept provides operations to call the tentativelyAccept method.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) TentativelyAccept()(*ItemCalendarCalendarViewItemInstancesItemTentativelyAcceptRequestBuilder) {
    return NewItemCalendarCalendarViewItemInstancesItemTentativelyAcceptRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCalendarCalendarViewItemInstancesEventItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
