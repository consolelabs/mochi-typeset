package typeset

import (
	"testing"

	"github.com/consolelabs/mochi-typeset/common/error/typeset"
)

func Test_KafkaMessage_Validation(t *testing.T) {
	// testcases
	tests := []struct {
		name string
		msg  *KafkaMessage
		want error
	}{
		{
			name: "Test_valid",
			msg: &KafkaMessage{
				Type:   KafkaMessageTypeAudit,
				Topic:  "mochi.api",
				Data:   []byte(`{"id":1,"last_updated_time":"2020-01-01T00:00:00Z","last_updated_block":1,"creation_block":1,"created_time":"2020-01-01T00:00:00Z","address":"0x1234567890","chain_id":1,"type":"mochi.api"}`),
				Sender: "mochi.api",
			},
			want: nil,
		}, {
			name: "Test_Missing_Type",
			want: typeset.ERROR_FAIL_VALIDATION,
			msg: &KafkaMessage{
				Topic:  "mochi.api",
				Data:   []byte(`{"id":1,"last_updated_time":"2020-01-01T00:00:00Z","last_updated_block":1,"creation_block":1,"created_time":"2020-01-01T00:00:00Z","address":"0x1234567890","chain_id":1,"type":"mochi.api"}`),
				Sender: "mochi.api",
			},
		},
	}

	// run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.msg.Validate(); err != tt.want {
				t.Errorf("KafkaMessage.Validate() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}
