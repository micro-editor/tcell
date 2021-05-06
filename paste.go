// Copyright 2020 The TCell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tcell

import (
	"time"
)

// EventPaste is used to store the contents of a bracketed paste.
type EventPaste struct {
	text string
	t    time.Time
}

// When returns the time when this EventPaste was created.
func (ev *EventPaste) When() time.Time {
	return ev.t
}

// Text returns the pasted text.
func (ev *EventPaste) Text() string {
	return ev.text
}

// NewEventPaste returns a new EventPaste.
func NewEventPaste(text string) *EventPaste {
	return &EventPaste{t: time.Now(), text: text}
}
