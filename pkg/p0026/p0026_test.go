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

package p0026

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type result struct {
	arg1    []int
	target  int //nolint
	target2 []int
}

var values = []result{
	{
		arg1:    []int{1, 1, 2},
		target:  2,
		target2: []int{1, 2},
	},
	{
		arg1:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		target:  5,
		target2: []int{0, 1, 2, 3, 4},
	},
}

type p0026TestSuite struct {
	suite.Suite
}

func (s *p0026TestSuite) Test() {
	for _, v := range values {
		s.Equal(v.target2, v.arg1[:removeDuplicates(v.arg1)])
	}
}

func TestP0026TestSuite(t *testing.T) {
	s := &p0026TestSuite{}
	suite.Run(t, s)
}
