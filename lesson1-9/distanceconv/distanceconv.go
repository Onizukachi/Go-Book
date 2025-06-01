package distanceconv

import "fmt"

type Metr float64
type Ft float64

func (m Metr) String() string { return fmt.Sprintf("%g Metrov", m) }
func (ft Ft) String() string  { return fmt.Sprintf("%g Ft", ft) }
