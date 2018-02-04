package pulse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepareCompressedPulse(t *testing.T) {
	input := "255 2904 1388 771 11346 0 0 0 0100020002020000020002020000020002000202000200020002000200000202000200020000020002000200020002020002000002000200000002000200020002020002000200020034"

	p := PrepareCompressedPulses(input)

	assert.Equal(t, []int{255, 771, 1388, 2904, 11346}, p.Lengths)
	assert.Equal(t, "0300020002020000020002020000020002000202000200020002000200000202000200020000020002000200020002020002000002000200000002000200020002020002000200020014", p.Pulses)
}
