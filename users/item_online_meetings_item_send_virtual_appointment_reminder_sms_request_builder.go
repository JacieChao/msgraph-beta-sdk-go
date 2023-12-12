package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilder provides operations to call the sendVirtualAppointmentReminderSms method.
type ItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilderInternal instantiates a new SendVirtualAppointmentReminderSmsRequestBuilder and sets the default values.
func NewItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilder) {
    m := &ItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onlineMeetings/{onlineMeeting%2Did}/sendVirtualAppointmentReminderSms", pathParameters),
    }
    return m
}
// NewItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilder instantiates a new SendVirtualAppointmentReminderSmsRequestBuilder and sets the default values.
func NewItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post send an SMS reminder to external attendees for a Teams Virtual Appointment. This feature requires Teams Premium and attendees must have a valid United States phone number to receive SMS notifications.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/virtualappointment-sendvirtualappointmentremindersms?view=graph-rest-1.0
func (m *ItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilder) Post(ctx context.Context, body ItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsPostRequestBodyable, requestConfiguration *ItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation send an SMS reminder to external attendees for a Teams Virtual Appointment. This feature requires Teams Premium and attendees must have a valid United States phone number to receive SMS notifications.
func (m *ItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsPostRequestBodyable, requestConfiguration *ItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilder) WithUrl(rawUrl string)(*ItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilder) {
    return NewItemOnlineMeetingsItemSendVirtualAppointmentReminderSmsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
