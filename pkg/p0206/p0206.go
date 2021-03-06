// Copyright (c) 2018 soren yang
//
// Licensed under the MIT License
// you may not use this file except in complicance with the License.
// You may obtain a copy of the License at
//
//     https://opensource.org/licenses/MIT
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package p0206

import (
	"algorithm101.io/gorithm/pkg/utils"
)

// ListNode for singly-linked list
type ListNode = utils.ListNode

func reverseList(head *ListNode) *ListNode {
	var ret *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = ret
		ret = head
		head = tmp
	}

	return ret
}
