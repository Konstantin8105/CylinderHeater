package section

// Section T:
//    | |
//    | | Plate 1
//    | |
// --------- Plate 2
// ----*----
//     (0,0) - center
type sectionT struct {
	plate1, plate2 Plate
}

func (s sectionT) check() error {
	err := s.plate1.Check()
	if err != nil {
		return err
	}
	err = s.plate2.Check()
	if err != nil {
		return err
	}
	err = s.convert().Check()
	return err
}

func (s sectionT) convert() RectangleSection {
	var parts []Rectangle
	// plate 1
	parts = append(parts, Rectangle{
		XCenter: 0,
		ZCenter: s.plate2.Thickness + s.plate1.Height/2.0,
		Height:  s.plate1.Height,
		Width:   s.plate1.Thickness,
	})
	// plate 2
	parts = append(parts, Rectangle{
		XCenter: 0,
		ZCenter: s.plate2.Thickness / 2.0,
		Height:  s.plate2.Thickness,
		Width:   s.plate2.Height,
	})
	return RectangleSection{Parts: parts}
}
