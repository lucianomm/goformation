// Code generated by "go generate". Please don't change this file directly.

package serverless

import (
	"bytes"
	"encoding/json"
	"io"
	"sort"

	"github.com/awslabs/goformation/v7/cloudformation/utils"
)

// Function_Properties is a helper struct that can hold either a S3Event, SNSEvent, SQSEvent, KinesisEvent, DynamoDBEvent, ApiEvent, ScheduleEvent, CloudWatchEventEvent, CloudWatchLogsEvent, IoTRuleEvent, AlexaSkillEvent, or EventBridgeRuleEvent value
type Function_Properties struct {
	S3Event              *Function_S3Event
	SNSEvent             *Function_SNSEvent
	SQSEvent             *Function_SQSEvent
	KinesisEvent         *Function_KinesisEvent
	DynamoDBEvent        *Function_DynamoDBEvent
	ApiEvent             *Function_ApiEvent
	ScheduleEvent        *Function_ScheduleEvent
	CloudWatchEventEvent *Function_CloudWatchEventEvent
	CloudWatchLogsEvent  *Function_CloudWatchLogsEvent
	IoTRuleEvent         *Function_IoTRuleEvent
	AlexaSkillEvent      *Function_AlexaSkillEvent
	EventBridgeRuleEvent *Function_EventBridgeRuleEvent
}

func (r Function_Properties) value() interface{} {
	ret := []interface{}{}

	if r.S3Event != nil {
		ret = append(ret, *r.S3Event)
	}

	if r.SNSEvent != nil {
		ret = append(ret, *r.SNSEvent)
	}

	if r.SQSEvent != nil {
		ret = append(ret, *r.SQSEvent)
	}

	if r.KinesisEvent != nil {
		ret = append(ret, *r.KinesisEvent)
	}

	if r.DynamoDBEvent != nil {
		ret = append(ret, *r.DynamoDBEvent)
	}

	if r.ApiEvent != nil {
		ret = append(ret, *r.ApiEvent)
	}

	if r.ScheduleEvent != nil {
		ret = append(ret, *r.ScheduleEvent)
	}

	if r.CloudWatchEventEvent != nil {
		ret = append(ret, *r.CloudWatchEventEvent)
	}

	if r.CloudWatchLogsEvent != nil {
		ret = append(ret, *r.CloudWatchLogsEvent)
	}

	if r.IoTRuleEvent != nil {
		ret = append(ret, *r.IoTRuleEvent)
	}

	if r.AlexaSkillEvent != nil {
		ret = append(ret, *r.AlexaSkillEvent)
	}

	if r.EventBridgeRuleEvent != nil {
		ret = append(ret, *r.EventBridgeRuleEvent)
	}

	sort.Sort(utils.ByJSONLength(ret)) // Heuristic to select best attribute
	if len(ret) > 0 {
		return ret[0]
	}

	return nil
}

func (r Function_Properties) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.value())
}

// Hook into the marshaller
func (r *Function_Properties) UnmarshalJSON(b []byte) error {

	// Unmarshal into interface{} to check it's type
	var typecheck interface{}
	if err := json.Unmarshal(b, &typecheck); err != nil {
		return err
	}

	switch val := typecheck.(type) {

	case map[string]interface{}:
		val = val // This ensures val is used to stop an error

		reader := bytes.NewReader(b)
		decoder := json.NewDecoder(reader)
		decoder.DisallowUnknownFields()
		reader.Seek(0, io.SeekStart)

		if err := decoder.Decode(&r.S3Event); err != nil {
			r.S3Event = nil
		}
		reader.Seek(0, io.SeekStart)

		if err := decoder.Decode(&r.SNSEvent); err != nil {
			r.SNSEvent = nil
		}
		reader.Seek(0, io.SeekStart)

		if err := decoder.Decode(&r.SQSEvent); err != nil {
			r.SQSEvent = nil
		}
		reader.Seek(0, io.SeekStart)

		if err := decoder.Decode(&r.KinesisEvent); err != nil {
			r.KinesisEvent = nil
		}
		reader.Seek(0, io.SeekStart)

		if err := decoder.Decode(&r.DynamoDBEvent); err != nil {
			r.DynamoDBEvent = nil
		}
		reader.Seek(0, io.SeekStart)

		if err := decoder.Decode(&r.ApiEvent); err != nil {
			r.ApiEvent = nil
		}
		reader.Seek(0, io.SeekStart)

		if err := decoder.Decode(&r.ScheduleEvent); err != nil {
			r.ScheduleEvent = nil
		}
		reader.Seek(0, io.SeekStart)

		if err := decoder.Decode(&r.CloudWatchEventEvent); err != nil {
			r.CloudWatchEventEvent = nil
		}
		reader.Seek(0, io.SeekStart)

		if err := decoder.Decode(&r.CloudWatchLogsEvent); err != nil {
			r.CloudWatchLogsEvent = nil
		}
		reader.Seek(0, io.SeekStart)

		if err := decoder.Decode(&r.IoTRuleEvent); err != nil {
			r.IoTRuleEvent = nil
		}
		reader.Seek(0, io.SeekStart)

		if err := decoder.Decode(&r.AlexaSkillEvent); err != nil {
			r.AlexaSkillEvent = nil
		}
		reader.Seek(0, io.SeekStart)

		if err := decoder.Decode(&r.EventBridgeRuleEvent); err != nil {
			r.EventBridgeRuleEvent = nil
		}
		reader.Seek(0, io.SeekStart)

	case []interface{}:

	}

	return nil
}
