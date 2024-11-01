package types

type Element struct {
	Name         string // Chemical name (e.g., "Hydrogen")
	Symbol       string // Chemical symbol (e.g., "H")
	AtomicNumber int    // Atomic number (e.g., 1)
	Group        int    // Group number in the periodic table (e.g., 1: Alkali metals, 18: Noble Gas, etc)
	Period       int    // Period number in the periodic table
	Category     string // Category (e.g., "Nonmetal", "Alkali metal", etc.)
	Color        string // Visual color used for rendering (e.g., "Light Blue" for noble gases)
	ColNum       int    // Horizontal position in the periodic table
	RowNum       int    // Vertical position in the periodic table
}
