package zoom

import (
	"context"
	"fmt"
)

// UpdateMeetingStatusOptions are the options to update a meeting's status
type UpdateMeetingStatusOptions struct {
	MeetingID int                       `json:"-"`
	Action    MeetingStatusUpdateAction `json:"action"`
}

// MeetingStatusUpdateAction is a specific action perform when updating a meeting's status
type MeetingStatusUpdateAction string

// MeetingStatusUpdateActionEnd ends the meeting
const MeetingStatusUpdateActionEnd = "end"

// UpdateMeetingStatusPath - v2 update a meeting's status
const UpdateMeetingStatusPath = "/meetings/%d/status"

// UpdateMeetingStatus calls PUT /meetings/{meetingID}/status
func UpdateMeetingStatus(opts UpdateMeetingStatusOptions) error {
	return defaultClient.UpdateMeetingStatus(opts)
}

// UpdateMeetingStatus calls PUT /meetings/{meetingID}/status
func (c *Client) UpdateMeetingStatus(opts UpdateMeetingStatusOptions) error {
	return c.UpdateMeetingStatusContext(context.Background(), opts)
}

// UpdateMeetingStatusContext calls PUT /meetings/{meetingID}/status
func (c *Client) UpdateMeetingStatusContext(ctx context.Context, opts UpdateMeetingStatusOptions) error {
	return c.requestV2Context(ctx, requestV2Opts{
		Method:         Put,
		Path:           fmt.Sprintf(UpdateMeetingStatusPath, opts.MeetingID),
		DataParameters: &opts,
		HeadResponse:   true,
	})
}
