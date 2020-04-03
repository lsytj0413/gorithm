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

package p1394

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type p1394TestSuite struct {
	suite.Suite
}

type result struct {
	arg1   []int
	target int
}

var values = []result{
	{
		arg1:   []int{2, 2, 3, 4},
		target: 2,
	},
	{
		arg1:   []int{1, 2, 2, 3, 3, 3},
		target: 3,
	},
	{
		arg1:   []int{2, 2, 2, 3, 3},
		target: -1,
	},
	{
		arg1:   []int{5},
		target: -1,
	},
	{
		arg1:   []int{7, 7, 7, 7, 7, 7, 7},
		target: 7,
	},
}

func (s *p1394TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target, findLucky(v.arg1))
	}
}

func TestP1394TestSuite(t *testing.T) {
	s := &p1394TestSuite{}
	suite.Run(t, s)
}
