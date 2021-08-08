package siswa

func Grade(Nilai int) string {
	if Nilai <= 40 {
		return "C"
	} else if Nilai > 40 && Nilai <= 70 {
		return "B"
	} else {
		return "A"
	}
}
