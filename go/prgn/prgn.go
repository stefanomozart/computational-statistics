/* --------------------------------------------------------------------------------------------------
 * prgn.go (https://github.com/stefanomozart/computational-statistics/go/prgn)
 * copywrite 2019 - Stefano Mozart (stefanomozart@ieee.org)
 *
 * This work is distributed under the GNU General Public Licence
 * (https://www.gnu.org/licenses/gpl-3.0.en.html)
 *------------------------------------------------------------------------------------------------*/

package prgn

// PRGN defines the basic interface for all generators in the package
type PRGN interface {
	Seed(s int) error
	Sample(n int) ([]int, error)
}

// LCG is an implementation of the Linear Congruential Generator
type LCG struct {
	seed int
	m    int
	a    int
	c    int
}

func newLCG(seed int) PRGN {
	return &LCG{
		seed: seed,
		m:    8,
		a:    5,
		c:    1,
	}
}

// Seed restarts teh PRGN
func (cg *LCG) Seed(seed int) error {
	cg.seed = seed
	return nil
}

func (cg *LCG) sampleOne() int {
	return ((cg.a * cg.seed) + cg.c) % cg.m
}

// Sample runs the PRGN in order to return n samples
func (cg *LCG) Sample(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = cg.sampleOne()
	}
	return s
}
