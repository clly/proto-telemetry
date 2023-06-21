package generators

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/shoenig/test/must"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/compiler/protogen"
)

func Test_Message(t *testing.T) {
	testcases := map[string]struct {
		numMsg   int
		children int
		expected int
	}{
		"OneMessage": {
			numMsg:   1,
			expected: 1,
		},
		"ThreeMessages": {
			numMsg:   3,
			expected: 3,
		},
		"OneMessageOneChild": {
			numMsg:   1,
			children: 1,
			expected: 3,
		},
		"ThreeMessageOneChild": {
			numMsg:   3,
			children: 1,
			expected: 9,
		},
		"ThreeMessageThreeChildren": {
			numMsg:   3,
			children: 2,
			expected: 21,
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			// can I use math to figure out how many messages I get out of this?
			testMessages := testMessageGenerator(t, tc.numMsg, tc.children)

			msgs := []Message{}
			for _, testMsg := range testMessages {
				msg := MessageGenerator(testMsg, &OpentelemetryGenerator{})
				msgs = append(msgs, msg)
			}

			totalMsgs := 0
			for _, msg := range msgs {
				totalMsgs += len(AllMessages(msg))
			}

			require.Equal(t, tc.expected, totalMsgs)
		})
	}
}

func testMessageGenerator(t *testing.T, numMsg int, children int) []*protogen.Message {
	msgs := make([]*protogen.Message, 0, numMsg*children)
	fmt.Println(children)
	for i := 0; i < numMsg; i++ {
		msg := genMsg(t)
		if children > 0 {
			children := testMessageGenerator(t, 2, children-1)
			msg.Messages = children
		}
		msgs = append(msgs, msg)
	}
	return msgs
}

func genMsg(t *testing.T) *protogen.Message {
	b := make([]byte, 128)
	_, err := rand.Read(b)
	must.NoError(t, err)
	name := hex.EncodeToString(b)

	return &protogen.Message{
		GoIdent: protogen.GoIdent{
			GoName:       name,
			GoImportPath: "",
		},
		Fields:     []*protogen.Field{},
		Oneofs:     []*protogen.Oneof{},
		Enums:      []*protogen.Enum{},
		Messages:   []*protogen.Message{},
		Extensions: []*protogen.Field{},
		Location: protogen.Location{
			SourceFile: "",
			Path:       []int32{},
		},
		Comments: protogen.CommentSet{
			LeadingDetached: []protogen.Comments{},
			Leading:         "",
			Trailing:        "",
		},
	}
}
