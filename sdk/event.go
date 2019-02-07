package sdk

import (
	"time"

	"github.com/sqreen/go-agent/agent"
)

// EventPropertyMap is the type used to represent extra custom event properties.
//
//	props := sdk.EventPropertyMap{
//		"key1": "value1",
//		"key2": "value2",
//	}
//	sdk.FromContext(ctx).TrackEvent("my.event").WithProperties(props)
//
type EventPropertyMap map[string]string

// HTTPRequestEvent is a SDK event. Its methods allow request handlers to add
// options further specifying the event, such as a unique user identifier, extra
// properties, etc.
type HTTPRequestEvent struct {
	impl *agent.HTTPRequestEvent
}

// WithTimestamp adds a custom timestamp to the event. By default, the timestamp
// is set to `time.Now()` value at the time of the call to the event creation.
//
//	sdk.FromContext(ctx).TrackEvent("my.event").WithTimestamp(yourTimestamp)
//
func (e *HTTPRequestEvent) WithTimestamp(t time.Time) *HTTPRequestEvent {
	if e == nil {
		return nil
	}
	e.impl.WithTimestamp(t)
	return e
}

// WithProperties adds custom properties to the event.
//
//	props := sdk.EventPropertyMap{
//		"key1": "value1",
//		"key2": "value2",
//	}
//	sdk.FromContext(ctx).TrackEvent("my.event").WithProperties(prop)
//
func (e *HTTPRequestEvent) WithProperties(p EventPropertyMap) *HTTPRequestEvent {
	if e == nil {
		return nil
	}
	e.impl.WithProperties(agent.EventPropertyMap(p))
	return e
}

// WithUserIdentifier associates the given user identifier map `id` to the
// event.
//
//	uid := sdk.EventUserIdentifierMap{"uid": "my-uid"}
//	sdk.FromContext(ctx).Identify(uid)
//
func (e *HTTPRequestEvent) WithUserIdentifiers(id EventUserIdentifiersMap) *HTTPRequestEvent {
	if e == nil || len(id) == 0 {
		return nil
	}
	e.impl.WithUserIdentifier(agent.EventUserIdentifiersMap(id))
	return e
}

// UserHTTPRequestEvent is a SDK event. Its methods allow request handlers to
// add options further specifying the event, such as a unique user identifier,
// extra properties, etc.
type UserHTTPRequestEvent struct {
	impl *HTTPRequestEvent
}

// WithTimestamp adds a custom timestamp to the event. By default, the timestamp
// is set to `time.Now()` value at the time of the call to the event creation.
//
//	sdk.FromContext(ctx).TrackEvent("my.event").WithTimestamp(yourTimestamp)
//
func (e *UserHTTPRequestEvent) WithTimestamp(t time.Time) *UserHTTPRequestEvent {
	if e == nil {
		return nil
	}
	e.impl.WithTimestamp(t)
	return e
}

// WithProperties adds custom properties to the event.
//
//	props := sdk.EventPropertyMap{
//		"key1": "value1",
//		"key2": "value2",
//	}
//	sdk.FromContext(ctx).TrackEvent("my.event").WithProperties(prop)
//
func (e *UserHTTPRequestEvent) WithProperties(p EventPropertyMap) *UserHTTPRequestEvent {
	if e == nil {
		return nil
	}
	e.impl.WithProperties(p)
	return e
}