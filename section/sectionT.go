// Section T:
//    | |
//    | | Plate 1
//    | |
// --------- Plate 2
// ----*----
//     (0,0) - center

package section

type sectionT struct {
	plate1, plate2 sectionPlate
}

func (s sectionT) check() error {
	err := s.plate1.check()
	if err != nil {
		return err
	}
	err = s.plate2.check()
	if err != nil {
		return err
	}
	err = s.convert().check()
	return err
}

func (s sectionT) convert() sectionRectanglePart {
	parts := make([]rectanglePart, 0)
	// plate 1
	parts = append(parts, rectanglePart{
		xCenter: 0,
		zCenter: s.plate2.t + s.plate1.h/2.0,
		height:  s.plate1.h,
		width:   s.plate1.t,
	})
	// plate 2
	parts = append(parts, rectanglePart{
		xCenter: 0,
		zCenter: s.plate2.t / 2.0,
		height:  s.plate2.t,
		width:   s.plate2.h,
	})
	return sectionRectanglePart{parts: parts}
}
