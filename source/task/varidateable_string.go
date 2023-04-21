package task

import (
	"fmt"
)

// 長さ制限付き文字列.
type varidateableString struct {
	minLength int
	maxLength int
	value     string
}

func (p *varidateableString) MinLength() int { return p.minLength }
func (p *varidateableString) MaxLength() int { return p.maxLength }
func (p *varidateableString) Value() string  { return p.value }

func (p *varidateableString) Set(value string) error {

	if len(value) < p.minLength {
		return fmt.Errorf("value too short")
	}

	if p.maxLength < len(value) {
		return fmt.Errorf("value too long")
	}

	p.value = value
	return nil
}
