package postgres

func ConvertIntSliceToArrayString(intSlice []int) string {
	s := "'{"
	for i, n := range intSlice {
		s := s + string(n)

		if i < len(intSlice) - 1 {
			s = s + ","
		}
	}
	s = s + "}'"

	return s
}