package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/pagination"
)

type EventSubscription struct {
	// The event subscription identifier.
	ID string `json:"id,required"`
	// The time the event subscription was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// This indicates if we'll send notifications to this subscription.
	Status EventSubscriptionStatus `json:"status,required"`
	// If specified, this subscription will only receive webhooks for Events with the
	// specified `category`.
	SelectedEventCategory EventSubscriptionSelectedEventCategory `json:"selected_event_category,required,nullable"`
	// The webhook url where we'll send notifications.
	URL string `json:"url,required"`
	// The key that will be used to sign webhooks.
	SharedSecret string `json:"shared_secret,required"`
	// A constant representing the object's type. For this resource it will always be
	// `event_subscription`.
	Type EventSubscriptionType `json:"type,required"`
	JSON EventSubscriptionJSON
}

type EventSubscriptionJSON struct {
	ID                    pjson.Metadata
	CreatedAt             pjson.Metadata
	Status                pjson.Metadata
	SelectedEventCategory pjson.Metadata
	URL                   pjson.Metadata
	SharedSecret          pjson.Metadata
	Type                  pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EventSubscription using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EventSubscription) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type EventSubscriptionStatus string

const (
	EventSubscriptionStatusActive            EventSubscriptionStatus = "active"
	EventSubscriptionStatusDisabled          EventSubscriptionStatus = "disabled"
	EventSubscriptionStatusDeleted           EventSubscriptionStatus = "deleted"
	EventSubscriptionStatusRequiresAttention EventSubscriptionStatus = "requires_attention"
)

type EventSubscriptionSelectedEventCategory string

const (
	EventSubscriptionSelectedEventCategoryAccountCreated                                       EventSubscriptionSelectedEventCategory = "account.created"
	EventSubscriptionSelectedEventCategoryAccountUpdated                                       EventSubscriptionSelectedEventCategory = "account.updated"
	EventSubscriptionSelectedEventCategoryAccountNumberCreated                                 EventSubscriptionSelectedEventCategory = "account_number.created"
	EventSubscriptionSelectedEventCategoryAccountNumberUpdated                                 EventSubscriptionSelectedEventCategory = "account_number.updated"
	EventSubscriptionSelectedEventCategoryAccountStatementCreated                              EventSubscriptionSelectedEventCategory = "account_statement.created"
	EventSubscriptionSelectedEventCategoryAccountTransferCreated                               EventSubscriptionSelectedEventCategory = "account_transfer.created"
	EventSubscriptionSelectedEventCategoryAccountTransferUpdated                               EventSubscriptionSelectedEventCategory = "account_transfer.updated"
	EventSubscriptionSelectedEventCategoryACHPrenotificationCreated                            EventSubscriptionSelectedEventCategory = "ach_prenotification.created"
	EventSubscriptionSelectedEventCategoryACHPrenotificationUpdated                            EventSubscriptionSelectedEventCategory = "ach_prenotification.updated"
	EventSubscriptionSelectedEventCategoryACHTransferCreated                                   EventSubscriptionSelectedEventCategory = "ach_transfer.created"
	EventSubscriptionSelectedEventCategoryACHTransferUpdated                                   EventSubscriptionSelectedEventCategory = "ach_transfer.updated"
	EventSubscriptionSelectedEventCategoryCardCreated                                          EventSubscriptionSelectedEventCategory = "card.created"
	EventSubscriptionSelectedEventCategoryCardUpdated                                          EventSubscriptionSelectedEventCategory = "card.updated"
	EventSubscriptionSelectedEventCategoryCardDisputeCreated                                   EventSubscriptionSelectedEventCategory = "card_dispute.created"
	EventSubscriptionSelectedEventCategoryCardDisputeUpdated                                   EventSubscriptionSelectedEventCategory = "card_dispute.updated"
	EventSubscriptionSelectedEventCategoryCheckDepositCreated                                  EventSubscriptionSelectedEventCategory = "check_deposit.created"
	EventSubscriptionSelectedEventCategoryCheckDepositUpdated                                  EventSubscriptionSelectedEventCategory = "check_deposit.updated"
	EventSubscriptionSelectedEventCategoryCheckTransferCreated                                 EventSubscriptionSelectedEventCategory = "check_transfer.created"
	EventSubscriptionSelectedEventCategoryCheckTransferUpdated                                 EventSubscriptionSelectedEventCategory = "check_transfer.updated"
	EventSubscriptionSelectedEventCategoryDeclinedTransactionCreated                           EventSubscriptionSelectedEventCategory = "declined_transaction.created"
	EventSubscriptionSelectedEventCategoryDigitalWalletTokenCreated                            EventSubscriptionSelectedEventCategory = "digital_wallet_token.created"
	EventSubscriptionSelectedEventCategoryDigitalWalletTokenUpdated                            EventSubscriptionSelectedEventCategory = "digital_wallet_token.updated"
	EventSubscriptionSelectedEventCategoryDocumentCreated                                      EventSubscriptionSelectedEventCategory = "document.created"
	EventSubscriptionSelectedEventCategoryEntityCreated                                        EventSubscriptionSelectedEventCategory = "entity.created"
	EventSubscriptionSelectedEventCategoryEntityUpdated                                        EventSubscriptionSelectedEventCategory = "entity.updated"
	EventSubscriptionSelectedEventCategoryExternalAccountCreated                               EventSubscriptionSelectedEventCategory = "external_account.created"
	EventSubscriptionSelectedEventCategoryFileCreated                                          EventSubscriptionSelectedEventCategory = "file.created"
	EventSubscriptionSelectedEventCategoryGroupUpdated                                         EventSubscriptionSelectedEventCategory = "group.updated"
	EventSubscriptionSelectedEventCategoryGroupHeartbeat                                       EventSubscriptionSelectedEventCategory = "group.heartbeat"
	EventSubscriptionSelectedEventCategoryInboundACHTransferReturnCreated                      EventSubscriptionSelectedEventCategory = "inbound_ach_transfer_return.created"
	EventSubscriptionSelectedEventCategoryInboundACHTransferReturnUpdated                      EventSubscriptionSelectedEventCategory = "inbound_ach_transfer_return.updated"
	EventSubscriptionSelectedEventCategoryInboundWireDrawdownRequestCreated                    EventSubscriptionSelectedEventCategory = "inbound_wire_drawdown_request.created"
	EventSubscriptionSelectedEventCategoryOauthConnectionCreated                               EventSubscriptionSelectedEventCategory = "oauth_connection.created"
	EventSubscriptionSelectedEventCategoryOauthConnectionDeactivated                           EventSubscriptionSelectedEventCategory = "oauth_connection.deactivated"
	EventSubscriptionSelectedEventCategoryPendingTransactionCreated                            EventSubscriptionSelectedEventCategory = "pending_transaction.created"
	EventSubscriptionSelectedEventCategoryPendingTransactionUpdated                            EventSubscriptionSelectedEventCategory = "pending_transaction.updated"
	EventSubscriptionSelectedEventCategoryRealTimeDecisionCardAuthorizationRequested           EventSubscriptionSelectedEventCategory = "real_time_decision.card_authorization_requested"
	EventSubscriptionSelectedEventCategoryRealTimeDecisionDigitalWalletTokenRequested          EventSubscriptionSelectedEventCategory = "real_time_decision.digital_wallet_token_requested"
	EventSubscriptionSelectedEventCategoryRealTimeDecisionDigitalWalletAuthenticationRequested EventSubscriptionSelectedEventCategory = "real_time_decision.digital_wallet_authentication_requested"
	EventSubscriptionSelectedEventCategoryRealTimePaymentsTransferCreated                      EventSubscriptionSelectedEventCategory = "real_time_payments_transfer.created"
	EventSubscriptionSelectedEventCategoryRealTimePaymentsTransferUpdated                      EventSubscriptionSelectedEventCategory = "real_time_payments_transfer.updated"
	EventSubscriptionSelectedEventCategoryRealTimePaymentsRequestForPaymentCreated             EventSubscriptionSelectedEventCategory = "real_time_payments_request_for_payment.created"
	EventSubscriptionSelectedEventCategoryRealTimePaymentsRequestForPaymentUpdated             EventSubscriptionSelectedEventCategory = "real_time_payments_request_for_payment.updated"
	EventSubscriptionSelectedEventCategoryTransactionCreated                                   EventSubscriptionSelectedEventCategory = "transaction.created"
	EventSubscriptionSelectedEventCategoryWireDrawdownRequestCreated                           EventSubscriptionSelectedEventCategory = "wire_drawdown_request.created"
	EventSubscriptionSelectedEventCategoryWireDrawdownRequestUpdated                           EventSubscriptionSelectedEventCategory = "wire_drawdown_request.updated"
	EventSubscriptionSelectedEventCategoryWireTransferCreated                                  EventSubscriptionSelectedEventCategory = "wire_transfer.created"
	EventSubscriptionSelectedEventCategoryWireTransferUpdated                                  EventSubscriptionSelectedEventCategory = "wire_transfer.updated"
)

type EventSubscriptionType string

const (
	EventSubscriptionTypeEventSubscription EventSubscriptionType = "event_subscription"
)

type EventSubscriptionList struct {
	// The contents of the list.
	Data []EventSubscription `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       EventSubscriptionListJSON
}

type EventSubscriptionListJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EventSubscriptionList using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EventSubscriptionList) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type EventSubscriptionsPage struct {
	*pagination.Page[EventSubscription]
}

func (r *EventSubscriptionsPage) EventSubscription() *EventSubscription {
	return r.Current()
}

func (r *EventSubscriptionsPage) NextPage() (*EventSubscriptionsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &EventSubscriptionsPage{page}, nil
	}
}
